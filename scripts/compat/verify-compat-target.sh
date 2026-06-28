#!/usr/bin/env bash
set -euo pipefail

CURRENT_FILE="${1:-docs/compat/rustdesk-current.json}"
SERVICE_FILE="${2:-backend/internal/service/compat_service.go}"

if ! command -v jq >/dev/null 2>&1; then
  echo "jq is required" >&2
  exit 2
fi

if [ ! -f "$CURRENT_FILE" ]; then
  echo "Manifest not found: $CURRENT_FILE" >&2
  exit 1
fi
if [ ! -f "$SERVICE_FILE" ]; then
  echo "Service file not found: $SERVICE_FILE" >&2
  exit 1
fi

manifest_version="$(jq -r '.client.version' "$CURRENT_FILE")"
manifest_release_date="$(jq -r '.client.release_date' "$CURRENT_FILE")"
manifest_server_version="$(jq -r '.server.compat_server_version' "$CURRENT_FILE")"
manifest_sysinfo="$(jq -r '.sysinfo_version' "$CURRENT_FILE")"

service_client_version="$(grep -E '^const CompatClientVersion = ' "$SERVICE_FILE" | sed -E 's/.*"([^"]+)".*/\1/')"
service_release_date="$(grep -E '^const CompatClientReleaseDate = ' "$SERVICE_FILE" | sed -E 's/.*"([^"]+)".*/\1/')"
service_server_version="$(grep -E '^const CompatServerVersion = ' "$SERVICE_FILE" | sed -E 's/.*"([^"]+)".*/\1/')"
service_sysinfo="$(grep -E '^const CompatSysinfoVersion = ' "$SERVICE_FILE" | sed -E 's/.*"([^"]+)".*/\1/')"

check_equal() {
  local name="$1"
  local left="$2"
  local right="$3"
  if [ "$left" != "$right" ]; then
    echo "Mismatch for $name: manifest=$left service=$right" >&2
    exit 1
  fi
}

check_equal client.version "$manifest_version" "$service_client_version"
check_equal client.release_date "$manifest_release_date" "$service_release_date"
check_equal server.compat_server_version "$manifest_server_version" "$service_server_version"
check_equal sysinfo_version "$manifest_sysinfo" "$service_sysinfo"

for endpoint in $(jq -r '.public_probe_endpoints[]' "$CURRENT_FILE"); do
  if ! grep -q "\"$endpoint\"" "$SERVICE_FILE"; then
    echo "Endpoint $endpoint is listed in manifest but missing from service Target()" >&2
    exit 1
  fi
done

echo "Compatibility target manifest matches Go constants"
