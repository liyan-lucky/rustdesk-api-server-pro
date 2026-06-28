#!/usr/bin/env bash
set -euo pipefail

DETECTED_FILE="${1:-docs/compat/rustdesk-latest-detected.json}"
CURRENT_FILE="${2:-docs/compat/rustdesk-current.json}"
SERVICE_FILE="${3:-backend/internal/service/compat_service.go}"

if ! command -v jq >/dev/null 2>&1; then
  echo "jq is required" >&2
  exit 2
fi

if [ ! -f "$DETECTED_FILE" ]; then
  echo "Detected release file not found: $DETECTED_FILE" >&2
  exit 1
fi

LATEST_TAG="$(jq -r '.latest.tag // empty' "$DETECTED_FILE")"
PUBLISHED_AT="$(jq -r '.latest.published_at // empty' "$DETECTED_FILE")"

if [ -z "$LATEST_TAG" ]; then
  echo "latest.tag is empty in $DETECTED_FILE" >&2
  exit 1
fi

RELEASE_DATE="${PUBLISHED_AT%%T*}"
if [ -z "$RELEASE_DATE" ] || [ "$RELEASE_DATE" = "null" ]; then
  RELEASE_DATE="$(date -u +%Y-%m-%d)"
fi

if [ ! -f "$SERVICE_FILE" ]; then
  echo "Service file not found: $SERVICE_FILE" >&2
  exit 1
fi

python3 - "$SERVICE_FILE" "$LATEST_TAG" "$RELEASE_DATE" <<'PY'
from pathlib import Path
import re
import sys

path = Path(sys.argv[1])
version = sys.argv[2]
release_date = sys.argv[3]
text = path.read_text(encoding="utf-8")
text = re.sub(r'const CompatClientVersion = "[^"]+"', f'const CompatClientVersion = "{version}"', text)
text = re.sub(r'const CompatClientReleaseDate = "[^"]+"', f'const CompatClientReleaseDate = "{release_date}"', text)
text = re.sub(
    r'const CompatSysinfoVersion = "[^"]+"',
    f'const CompatSysinfoVersion = "rustdesk-api-server-pro-compat-client-{version}-server-1.1.15-latest"',
    text,
)
path.write_text(text, encoding="utf-8")
PY

tmp="$(mktemp)"
jq \
  --arg version "$LATEST_TAG" \
  --arg release_date "$RELEASE_DATE" \
  --arg sysinfo "rustdesk-api-server-pro-compat-client-${LATEST_TAG}-server-1.1.15-latest" \
  '.client.version = $version
   | .client.tag = $version
   | .client.release_date = $release_date
   | .sysinfo_version = $sysinfo' \
  "$CURRENT_FILE" > "$tmp"
mv "$tmp" "$CURRENT_FILE"

echo "Updated compatibility target to RustDesk $LATEST_TAG ($RELEASE_DATE)"
