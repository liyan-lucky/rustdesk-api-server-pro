#!/usr/bin/env bash
set -euo pipefail

fail() {
  echo "[security-regression] $*" >&2
  exit 1
}

# Generic security primitives must stay available for sensitive comparisons.
grep -q 'func ConstantTimeStringEqual' backend/util/secure.go || fail "constant time string compare helper missing"
grep -q 'subtle.ConstantTimeCompare' backend/util/secure.go || fail "constant time compare implementation missing"
test -f backend/util/secure_test.go || fail "constant time compare tests missing"
grep -q 'TestConstantTimeStringEqual' backend/util/secure_test.go || fail "constant time compare test case missing"

# AuthToken must keep model-layer hashing as a final safety net.
grep -q 'func (m \*AuthToken) BeforeInsert()' backend/app/model/auth_token.go || fail "AuthToken.BeforeInsert hook missing"
grep -q 'func (m \*AuthToken) BeforeUpdate()' backend/app/model/auth_token.go || fail "AuthToken.BeforeUpdate hook missing"
grep -q 'm.TokenHash = util.Sha256Hex(m.Token)' backend/app/model/auth_token.go || fail "AuthToken token hash normalization missing"
grep -q 'm.Token = ""' backend/app/model/auth_token.go || fail "AuthToken plaintext token clearing missing"

# Admin user list must not expose the raw TOTP secret.
if grep -n '"tfa_secret"' backend/app/controller/admin/users.go; then
  fail "admin users API exposes tfa_secret"
fi
grep -q '"has_2fa"' backend/app/controller/admin/users.go || fail "admin users API should expose has_2fa instead of tfa_secret"

# Generated config files must be owner-only readable/writable.
grep -q 'os.WriteFile(yamlFile, bytes, 0600)' backend/config/config.go || fail "server.yaml write permission must be 0600"

# Unsafe static sign keys must be blocked at startup.
grep -q 'IsUnsafeSignKey' backend/app/main.go || fail "startup signKey safety check missing"
grep -q 'CHANGE_ME_TO_A_RANDOM_32_BYTE_SECRET' backend/config/config.go || fail "unsafe signKey placeholder check missing"

# Docker first-run path must not reuse a public sample signKey.
grep -q 'generated_sign_key=' docker/start.sh || fail "Docker first-run signKey generation missing"
grep -q 'CHANGE_ME_TO_A_RANDOM_32_BYTE_SECRET' docker/start.sh || fail "Docker placeholder signKey detection missing"

# CLI entrypoint must return a normal failure exit code instead of panicking.
if grep -n 'panic(err)' backend/main.go; then
  fail "CLI entrypoint must not panic on command errors"
fi
grep -q 'fmt.Fprintln(os.Stderr, err)' backend/main.go || fail "CLI entrypoint must print command errors to stderr"
grep -q 'os.Exit(1)' backend/main.go || fail "CLI entrypoint must exit with code 1 on command errors"

# Scheduler startup must reuse the app DB engine and return errors instead of panicking.
if grep -n 'panic(err)' backend/app/jobs.go; then
  fail "scheduler startup must return errors instead of panicking"
fi
grep -q 'func StartJobs(cfg \*config.ServerConfig, dbEngine \*xorm.Engine) error' backend/app/jobs.go || fail "StartJobs must accept the app DB engine and return an error"
grep -q 'job db engine is nil' backend/app/jobs.go || fail "StartJobs must reject nil DB engine"
if grep -n 'db.NewEngine' backend/app/jobs.go; then
  fail "StartJobs must reuse the app DB engine instead of creating a second engine"
fi
grep -q 'dbEngine, err := db.NewEngine(cfg.Db)' backend/app/main.go || fail "server startup must create the DB engine once"
grep -q 'newApp(cfg, dbEngine)' backend/app/main.go || fail "app initialization must receive the shared DB engine"
grep -q 'if err := StartJobs(cfg, dbEngine); err != nil' backend/app/main.go || fail "server startup must fail when jobs fail"
grep -q 'func newApp(cfg \*config.ServerConfig, dbEngine \*xorm.Engine)' backend/app/main.go || fail "newApp must accept the shared DB engine"
grep -q 'db engine is nil' backend/app/main.go || fail "newApp must reject nil DB engine"
grep -q 'return fmt.Errorf("create scheduler:' backend/app/jobs.go || fail "StartJobs must return scheduler creation errors"
grep -q 's.NewJob' backend/app/jobs.go || fail "device check job creation missing"
grep -q 'return fmt.Errorf("create device check job:' backend/app/jobs.go || fail "StartJobs must return job creation errors"
grep -q 'jobDuration <= 0' backend/app/jobs.go || fail "StartJobs must validate job duration"
if grep -n 'Logger().Fatal' backend/app/main.go; then
  fail "app initialization must return errors instead of fatal-exiting"
fi

# Process helpers must return errors instead of panicking.
if grep -n 'panic(err)' backend/util/process.go; then
  fail "process helper must return errors instead of panicking"
fi
grep -q 'StartProcess(name string, attr \*ProcessAttr) (\*exec.Cmd, error)' backend/util/process.go || fail "StartProcess must return an error"
grep -q 'return nil, err' backend/util/process.go || fail "StartProcess must return start errors"

# RustDesk process lifecycle must avoid unsafe pid permissions and nil process panics.
grep -q 'writePidFile(pidFile string, pid int) error' backend/helper/rustdesk/server.go || fail "rustdesk pid writer helper missing"
grep -q 'os.WriteFile(pidFile, \[\]byte(strconv.Itoa(pid)), 0644)' backend/helper/rustdesk/server.go || fail "rustdesk pid files must use 0644 permissions"
grep -q 'killProcessByPID' backend/helper/rustdesk/server.go || fail "rustdesk process kill helper missing"
grep -q 'isProcessRunning' backend/helper/rustdesk/server.go || fail "rustdesk process status helper missing"
grep -q 'proc == nil' backend/helper/rustdesk/server.go || fail "rustdesk process helpers must guard nil process values"
grep -q 'os.Remove(hbbrPidFile)' backend/helper/rustdesk/server.go || fail "rustdesk startup failure/stop must clean hbbr pid file"
grep -q 'os.Remove(hbbsPidFile)' backend/helper/rustdesk/server.go || fail "rustdesk stop must clean hbbs pid file"
if grep -n 'os.ModePerm' backend/helper/rustdesk/server.go; then
  fail "rustdesk server helper must not use world-writable permissions"
fi

# RustDesk key output must not reveal or read the private key by default.
grep -q 'show-private' backend/cmd/rustdesk.go || fail "rustdesk private key output must require an explicit flag"
grep -q 'rustdesk.PublicKey()' backend/cmd/rustdesk.go || fail "rustdesk keys command should read public key by default"
grep -q 'rustdesk.PrivateKey()' backend/cmd/rustdesk.go || fail "rustdesk keys command should read private key only when explicitly requested"
grep -q 'func PublicKey() string' backend/helper/rustdesk/server.go || fail "rustdesk public key helper missing"
grep -q 'func PrivateKey() string' backend/helper/rustdesk/server.go || fail "rustdesk private key helper missing"
if grep -n 'public, private := rustdesk.Keys()' backend/cmd/rustdesk.go; then
  fail "rustdesk keys command must not read the private key by default"
fi
grep -q 'private key: hidden' backend/cmd/rustdesk.go || fail "rustdesk keys command should tell users private key is hidden"

# HTTP helpers must reject error statuses, bound response sizes, and write downloads safely.
grep -q 'maxHTTPStringSize' backend/util/http.go || fail "HTTP string response size limit missing"
grep -q 'maxDownloadSize' backend/util/http.go || fail "download size limit missing"
grep -q 'resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices' backend/util/http.go || fail "HTTP status validation missing"
grep -q 'io.LimitReader(resp.Body, maxHTTPStringSize+1)' backend/util/http.go || fail "HTTP string response limit reader missing"
grep -q 'io.LimitReader(resp.Body, maxDownloadSize+1)' backend/util/http.go || fail "download limit reader missing"
grep -q 'os.CreateTemp' backend/util/http.go || fail "download must write to a temporary file first"
grep -q 'tmpFile.Chmod(0644)' backend/util/http.go || fail "downloaded file permission must be 0644"
grep -q 'os.Rename(tmpName, filename)' backend/util/http.go || fail "download must atomically rename completed temporary file"
grep -q 'Timeout: 60 \* time.Second' backend/util/http.go || fail "HTTP client timeout missing"
test -f backend/util/http_test.go || fail "HTTP helper security tests missing"
grep -q 'TestHttpGetStringRejectsErrorStatus' backend/util/http_test.go || fail "HTTP error status test missing"
grep -q 'TestDownloadFileRejectsErrorStatusAndLeavesNoFile' backend/util/http_test.go || fail "download error cleanup test missing"
grep -q 'TestDownloadFileUsesSafePermissionsAndTruncatesOldFile' backend/util/http_test.go || fail "download permission/truncation test missing"
if grep -n 'OpenFile(filename,.*os.ModePerm\|os.Create(filename)' backend/util/http.go; then
  fail "download must not write directly to final path with unsafe permissions"
fi

# 404 handling must not dump request headers, body, or query strings to logs.
if grep -n 'GetBody()\|Request().Header\|fmt.Println\|RequestURI' backend/app/main.go; then
  fail "404 handler must not log raw request headers, body, or query strings"
fi
grep -q 'Logger().Infof("(404)' backend/app/main.go || fail "404 handler should keep sanitized method/path logging"
grep -q 'Request().URL.Path' backend/app/main.go || fail "404 handler must log URL.Path instead of RequestURI"

# Request logger must not dump raw request headers, body, or query strings, even in debug mode.
if grep -n 'GetBody()\|Request().Header\|fmt.Println\|RequestURI' backend/app/middleware/request_logger.go; then
  fail "request logger must not log raw request headers, body, or query strings"
fi
grep -q 'Logger().Infof("▶ %s:%s"' backend/app/middleware/request_logger.go || fail "request logger should keep sanitized method/path logging"
grep -q 'Request().URL.Path' backend/app/middleware/request_logger.go || fail "request logger must log URL.Path instead of RequestURI"
grep -q 'RequestLogger(cfg.DebugMode)' backend/app/main.go || fail "request logger must use startup config instead of reloading config per request"

# Helper packages must return errors to callers rather than terminating the process.
if grep -n 'os.Exit' backend/helper/github/github.go; then
  fail "github helper must return errors instead of exiting the process"
fi
grep -q 'GetLatestRelease(repo string) (\*Release, error)' backend/helper/github/github.go || fail "github latest release helper should return an error"
grep -q 'GetReleaseByTag(repo, tag string) (\*Release, error)' backend/helper/github/github.go || fail "github release by tag helper should return an error"
grep -q 'GetReleases(repo string) (\*\[\]Release, error)' backend/helper/github/github.go || fail "github releases helper should return an error"

# Zip extraction must not allow ZipSlip path traversal or unsafe permission defaults.
grep -q 'safeZipDestination' backend/util/file.go || fail "zip extraction path validation missing"
grep -q 'filepath.IsAbs(name)' backend/util/file.go || fail "zip extraction must reject absolute paths"
grep -q 'filepath.Rel(cleanDst, target)' backend/util/file.go || fail "zip extraction must use filepath.Rel for containment checks"
grep -q 'strings.HasPrefix(rel, ".."+string(os.PathSeparator))' backend/util/file.go || fail "zip extraction must reject parent-directory escapes"
grep -q 'os.ModeSymlink' backend/util/file.go || fail "zip extraction must reject symlinks"
grep -q 'mode.Perm() & 0755' backend/util/file.go || fail "zip extraction must strip group/world write bits"
test -f backend/util/file_test.go || fail "zip extraction security tests missing"
grep -q 'TestSafeZipDestinationRejectsTraversal' backend/util/file_test.go || fail "zip traversal rejection test missing"
grep -q 'TestSafeZipDestinationRejectsAbsolutePath' backend/util/file_test.go || fail "zip absolute path rejection test missing"
grep -q 'TestSafeModesRemoveGroupAndWorldWrite' backend/util/file_test.go || fail "zip permission hardening test missing"
if grep -n 'MkdirAll(.*os.ModePerm' backend/util/file.go; then
  fail "zip extraction must not create world-writable directories"
fi

# OAuth/OIDC callback URL generation must not read X-Forwarded-Host from request headers.
if grep -n 'GetHeader("X-Forwarded-Host")\|GetHeader(\x27X-Forwarded-Host\x27)' backend/app/controller/admin/auth.go; then
  fail "OAuth/OIDC base URL must not read X-Forwarded-Host"
fi
grep -q 'oidc.redirectUrl or oauth.providers\[\].redirectUrl' backend/app/controller/admin/auth.go || fail "OAuth/OIDC redirectUrl operator guidance missing"
if grep -n 'withQuery(redirectTo, "oidc_error", err.Error())\|withQuery(redirectTo, "oauth_error", err.Error())' backend/app/controller/admin/auth.go; then
  fail "OAuth/OIDC redirect errors must not expose raw internal errors"
fi
grep -q 'withQuery(redirectTo, "oidc_error", "auth_failed")' backend/app/controller/admin/auth.go || fail "OIDC redirect error must use sanitized code"
grep -q 'withQuery(redirectTo, "oauth_error", "auth_failed")' backend/app/controller/admin/auth.go || fail "OAuth redirect error must use sanitized code"

# Generic OAuth provider must not rely on plaintext AuthToken writes or timing-sensitive state signature checks.
if awk '
  /authToken := &model\.AuthToken\{/ { in_auth_token=1 }
  in_auth_token && /Token:[[:space:]]*token/ { print; found=1 }
  in_auth_token && /^[[:space:]]*}/ { in_auth_token=0 }
  END { exit found ? 0 : 1 }
' backend/internal/service/oauth_provider_service.go; then
  fail "OAuth provider AuthToken must write TokenHash instead of plaintext Token"
fi
if grep -n 'string(rawSignature)[[:space:]]*!=[[:space:]]*expectedSignature' backend/internal/service/oauth_provider_service.go; then
  fail "OAuth provider state signature comparison must be constant time"
fi
grep -q 'TokenHash:[[:space:]]*util.Sha256Hex(token)' backend/internal/service/oauth_provider_service.go || fail "OAuth provider explicit TokenHash write missing"
grep -q 'util.ConstantTimeStringEqual(string(rawSignature), expectedSignature)' backend/internal/service/oauth_provider_service.go || fail "OAuth provider constant-time state comparison missing"

# Generic OAuth id_token fallback must verify signatures and claims before trusting payload data.
grep -q 'JWKSURI.*json:"jwks_uri"' backend/internal/service/oauth_provider_service.go || fail "OAuth provider jwks_uri metadata support missing"
grep -q 'JWKSURI[[:space:]]*string[[:space:]]*`yaml:"jwksUri"`' backend/config/config.go || fail "OAuth provider jwksUri config support missing"
grep -q 'verifyOAuthIDToken' backend/internal/service/oauth_provider_service.go || fail "OAuth provider ID token verification missing"
grep -q 'verifyIDTokenSignature(idToken' backend/internal/service/oauth_provider_service.go || fail "OAuth provider ID token signature verification call missing"
grep -q 'fillClaimsByOAuthIDToken(idToken, expectedIssuer, provider.ClientID, claims)' backend/internal/service/oauth_provider_service.go || fail "OAuth provider ID token claim validation call missing"
grep -q 'validateIDTokenClaims(claims, expectedIssuer, expectedAudience)' backend/internal/service/oauth_provider_service.go || fail "OAuth provider ID token claim validation missing"
grep -q 'oauth userinfo subject mismatch' backend/internal/service/oauth_provider_service.go || fail "OAuth provider userinfo/id_token subject consistency check missing"

# OIDC ID token fallback must verify signature and validate high-value claims before trusting payload data.
grep -q 'verifyIDTokenSignature' backend/internal/service/oidc_auth_service.go || fail "OIDC ID token signature verification call missing"
grep -q 'JWKSURI.*json:"jwks_uri"' backend/internal/service/oidc_auth_service.go || fail "OIDC jwks_uri metadata support missing"
grep -q 'rsa.VerifyPKCS1v15' backend/internal/service/oidc_jwks_verify.go || fail "OIDC RS256 signature verification missing"
grep -q 'unsupported id token alg' backend/internal/service/oidc_jwks_verify.go || fail "OIDC alg allowlist missing"
grep -q 'validateIDTokenClaims' backend/internal/service/oidc_auth_service.go || fail "OIDC ID token claim validation missing"
grep -q 'id token issuer invalid' backend/internal/service/oidc_auth_service.go || fail "OIDC issuer validation missing"
grep -q 'id token audience invalid' backend/internal/service/oidc_auth_service.go || fail "OIDC audience validation missing"
grep -q 'id token expired' backend/internal/service/oidc_auth_service.go || fail "OIDC expiry validation missing"
grep -q 'id token issued-at invalid' backend/internal/service/oidc_auth_service.go || fail "OIDC issued-at validation missing"

# Recording uploads must have a hard size limit and private directory permissions.
grep -q 'maxCompatRecordSize' backend/internal/service/compat_service.go || fail "record upload size limit missing"
grep -q 'os.MkdirAll(dir, 0700)' backend/internal/service/compat_service.go || fail "record upload directory permission must be 0700"

echo "[security-regression] ok"
