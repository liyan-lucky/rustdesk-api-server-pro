#!/usr/bin/env bash
set -euo pipefail

fail() {
  echo "[security-regression] $*" >&2
  exit 1
}

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

# OAuth/OIDC callback URL generation must not trust X-Forwarded-Host.
if grep -n 'X-Forwarded-Host' backend/app/controller/admin/auth.go; then
  fail "OAuth/OIDC base URL must not trust X-Forwarded-Host"
fi
grep -q 'oidc.redirectUrl or oauth.providers\[\].redirectUrl' backend/app/controller/admin/auth.go || fail "OAuth/OIDC redirectUrl operator guidance missing"

# OIDC ID token fallback must validate high-value claims before trusting payload data.
grep -q 'validateIDTokenClaims' backend/internal/service/oidc_auth_service.go || fail "OIDC ID token claim validation missing"
grep -q 'id token issuer invalid' backend/internal/service/oidc_auth_service.go || fail "OIDC issuer validation missing"
grep -q 'id token audience invalid' backend/internal/service/oidc_auth_service.go || fail "OIDC audience validation missing"
grep -q 'id token expired' backend/internal/service/oidc_auth_service.go || fail "OIDC expiry validation missing"

# Recording uploads must have a hard size limit and private directory permissions.
grep -q 'maxCompatRecordSize' backend/internal/service/compat_service.go || fail "record upload size limit missing"
grep -q 'os.MkdirAll(dir, 0700)' backend/internal/service/compat_service.go || fail "record upload directory permission must be 0700"

echo "[security-regression] ok"
