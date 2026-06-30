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
