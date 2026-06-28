#!/usr/bin/env bash
set -euo pipefail

CURRENT_FILE="${1:-docs/compat/rustdesk-current.json}"
OUT_FILE="${2:-docs/compat/rustdesk-latest-detected.json}"
API_URL="https://api.github.com/repos/rustdesk/rustdesk/releases/latest"

if ! command -v jq >/dev/null 2>&1; then
  echo "jq is required" >&2
  exit 2
fi

mkdir -p "$(dirname "$OUT_FILE")"

LATEST_JSON="$(curl -fsSL "$API_URL")"
LATEST_TAG="$(jq -r '.tag_name // .name // empty' <<<"$LATEST_JSON")"
LATEST_NAME="$(jq -r '.name // .tag_name // empty' <<<"$LATEST_JSON")"
LATEST_PUBLISHED_AT="$(jq -r '.published_at // .created_at // empty' <<<"$LATEST_JSON")"
LATEST_HTML_URL="$(jq -r '.html_url // empty' <<<"$LATEST_JSON")"

CURRENT_VERSION=""
if [ -f "$CURRENT_FILE" ]; then
  CURRENT_VERSION="$(jq -r '.client.version // .client.tag // .tag // empty' "$CURRENT_FILE")"
fi

if [ -z "$LATEST_TAG" ]; then
  echo "Unable to read latest RustDesk release tag" >&2
  exit 1
fi

jq -n \
  --arg source_repo "rustdesk/rustdesk" \
  --arg tag "$LATEST_TAG" \
  --arg name "$LATEST_NAME" \
  --arg published_at "$LATEST_PUBLISHED_AT" \
  --arg html_url "$LATEST_HTML_URL" \
  --arg current_version "$CURRENT_VERSION" \
  --arg checked_at "$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
  '{
    source_repo: $source_repo,
    latest: {
      tag: $tag,
      name: $name,
      published_at: $published_at,
      html_url: $html_url
    },
    current_version: $current_version,
    needs_update: ($tag != $current_version),
    checked_at: $checked_at
  }' > "$OUT_FILE"

NEEDS_UPDATE="$(jq -r '.needs_update' "$OUT_FILE")"

echo "latest_tag=$LATEST_TAG"
echo "current_version=$CURRENT_VERSION"
echo "needs_update=$NEEDS_UPDATE"

if [ -n "${GITHUB_OUTPUT:-}" ]; then
  {
    echo "latest_tag=$LATEST_TAG"
    echo "current_version=$CURRENT_VERSION"
    echo "needs_update=$NEEDS_UPDATE"
    echo "out_file=$OUT_FILE"
  } >> "$GITHUB_OUTPUT"
fi
