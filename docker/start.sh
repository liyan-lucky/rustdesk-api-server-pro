#!/bin/sh

set -eu

if [ ! -f /usr/local/bin/rustdesk-api-server-pro ]; then
    ln -s /app/rustdesk-api-server-pro /usr/local/bin/rustdesk-api-server-pro
fi

mkdir -p /app/data

# Make mounted config at /app/server.yaml effective, because the binary reads server.yaml from CWD.
# The process runs in /app/data, so keep /app/data/server.yaml in sync with /app/server.yaml.
# Only copy when /app/data/server.yaml does not exist yet, to avoid overwriting user edits.
if [ -f /app/server.yaml ] && [ ! -f /app/data/server.yaml ]; then
    cp /app/server.yaml /app/data/server.yaml
fi

cd /app/data

# Allow PORT env var to override the listening port (e.g. Docker -p 21114:21114 -e PORT=21114).
# The Go binary also reads PORT env var directly, but updating server.yaml ensures
# config display and health probes are consistent.
if [ -n "${PORT:-}" ]; then
    port_val="$PORT"
    case "$port_val" in :*) ;; *) port_val=":$port_val" ;; esac
    if [ -f /app/data/server.yaml ]; then
        sed -i "/^httpConfig:/,/^[^ ]/ s|^\([[:space:]]*port:\).*|  port: \"$port_val\"|" /app/data/server.yaml
    fi
fi

# Harden first-run Docker defaults. The image contains a sample server.yaml, but the
# server refuses unsafe placeholder signKey values. Generate a stable random key in
# the persisted /app/data/server.yaml so a fresh container can start safely without
# reusing a public example secret.
if [ -f /app/data/server.yaml ]; then
    current_sign_key="$(grep -E '^signKey:' /app/data/server.yaml | head -n1 | sed 's/#.*//' | sed 's/^signKey:[[:space:]]*//' | tr -d '"' | tr -d "'" | sed 's/[[:space:]]*$//' || true)"
    case "$current_sign_key" in
        ""|"CHANGE_ME_TO_A_RANDOM_32_BYTE_SECRET"|'sercrethatmaycontainch@r$32chars')
            generated_sign_key="$(head -c 48 /dev/urandom | base64 | tr -d '=+/\n' | cut -c1-48)"
            sed -i "s|^signKey:.*|signKey: \"$generated_sign_key\"|" /app/data/server.yaml
            ;;
    esac
fi

#if [ ! -f /app/server.db ]; then # This is not good if one wants to upgrade instance
/app/rustdesk-api-server-pro sync
#fi

if [ ! -f /app/data/.init.lock ] && [ -n "${ADMIN_USER:-}" ] && [ -n "${ADMIN_PASS:-}" ]; then
    /app/rustdesk-api-server-pro user add "$ADMIN_USER" "$ADMIN_PASS" --admin
    touch /app/data/.init.lock
fi

export APP_VERSION="${APP_VERSION:-latest}"
export BUILD_TIME="${BUILD_TIME:-}"

exec /app/rustdesk-api-server-pro start
