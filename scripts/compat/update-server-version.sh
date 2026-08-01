#!/usr/bin/env bash
set -euo pipefail

VERSION_FILE="${1:-VERSION}"
CURRENT_FILE="${2:-docs/compat/rustdesk-current.json}"

if ! command -v jq >/dev/null 2>&1; then
  echo "jq is required" >&2
  exit 2
fi
if [ ! -f "$VERSION_FILE" ]; then
  echo "Version file not found: $VERSION_FILE" >&2
  exit 1
fi
if [ ! -f "$CURRENT_FILE" ]; then
  echo "Manifest not found: $CURRENT_FILE" >&2
  exit 1
fi

server_version="$(tr -d '[:space:]' < "$VERSION_FILE")"
client_version="$(jq -r '.client.version' "$CURRENT_FILE")"
if ! [[ "$server_version" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "Invalid server version: $server_version" >&2
  exit 1
fi
if [ -z "$client_version" ] || [ "$client_version" = "null" ]; then
  echo "Client compatibility version is missing" >&2
  exit 1
fi

tmp="$(mktemp)"
trap 'rm -f "$tmp"' EXIT
jq \
  --arg server_version "$server_version" \
  --arg sysinfo "rustdesk-api-server-pro-compat-client-${client_version}-server-${server_version}-latest" \
  '.server.compat_server_version = $server_version
   | .sysinfo_version = $sysinfo' \
  "$CURRENT_FILE" > "$tmp"
mv "$tmp" "$CURRENT_FILE"
trap - EXIT

echo "Updated compatibility manifest to server version $server_version"
