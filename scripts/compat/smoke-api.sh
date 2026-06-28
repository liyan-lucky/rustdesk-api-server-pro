#!/usr/bin/env bash
set -euo pipefail

BASE_URL="${BASE_URL:-${1:-http://127.0.0.1:12345}}"
TOKEN="${TOKEN:-}"

failures=0

request() {
  local method="$1"
  local path="$2"
  local expect="$3"
  local url="${BASE_URL%/}${path}"
  local args=(-fsS -X "$method" "$url")
  if [ -n "$TOKEN" ]; then
    args+=(-H "Authorization: Bearer $TOKEN")
  fi

  echo "==> $method $path"
  if out="$(curl "${args[@]}" 2>&1)"; then
    if [ -n "$expect" ] && ! grep -q "$expect" <<<"$out"; then
      echo "Expected to find '$expect' but got: $out" >&2
      failures=$((failures + 1))
    else
      echo "OK"
    fi
  else
    echo "FAILED: $out" >&2
    failures=$((failures + 1))
  fi
}

# Public compatibility probes. These should not require login.
request GET  /api/health "ok"
request GET  /api/ping "ok"
request GET  /api/status "compat_target"
request POST /api/status "compat_target"
request GET  /api/version "compat_target"
request GET  /api/info "compat_target"
request GET  /api/features "compat_target"
request GET  /api/capabilities "compat_target"
request GET  /api/compat/features "compat_target"
request GET  /api/config "compat_target"
request GET  /api/client-config "compat_target"
request GET  /api/client_config "compat_target"
request GET  /api/server-config "compat_target"
request GET  /api/server_config "compat_target"
request GET  /api/server/info "compat_target"
request GET  /api/compat-target "1.4.8"
request GET  /api/compat/target "1.4.8"
request GET  /api/compat/version "1.4.8"
request GET  /api/sysinfo_ver "1.4.8"
request GET  /api/devices/deploy "NOT_ENABLED"

if [ -n "$TOKEN" ]; then
  request GET  /api/currentUser "data"
  request POST /api/currentUser "data"
  request GET  /api/me "username"
  request GET  /api/devices/current "rustdesk_id"
fi

if [ "$failures" -ne 0 ]; then
  echo "Smoke test failed: $failures failure(s)" >&2
  exit 1
fi

echo "Compatibility smoke test passed"
