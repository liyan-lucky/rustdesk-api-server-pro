#!/bin/sh

set -eu

if [ ! -f /usr/local/bin/rustdesk-api-server-pro ]; then
    ln -s /app/rustdesk-api-server-pro /usr/local/bin/rustdesk-api-server-pro
fi

mkdir -p /app/data

# Make mounted config at /app/server.yaml effective, because the binary reads server.yaml from CWD.
# The process runs in /app/data, so keep /app/data/server.yaml in sync with /app/server.yaml.
if [ -f /app/server.yaml ]; then
    cp /app/server.yaml /app/data/server.yaml
fi

cd /app/data

# Harden first-run Docker defaults. The image contains a sample server.yaml, but the
# server refuses unsafe placeholder signKey values. Generate a stable random key in
# the persisted /app/data/server.yaml so a fresh container can start safely without
# reusing a public example secret.
if [ -f /app/data/server.yaml ]; then
    current_sign_key="$(grep -E '^signKey:' /app/data/server.yaml | head -n1 | sed 's/^signKey:[[:space:]]*//' | tr -d '"' || true)"
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

/app/rustdesk-api-server-pro start
