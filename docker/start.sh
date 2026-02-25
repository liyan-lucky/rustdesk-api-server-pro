#!/bin/sh

if [ ! -f /usr/local/bin/rustdesk-api-server-pro ]; then
    ln -s /app/rustdesk-api-server-pro /usr/local/bin/rustdesk-api-server-pro
fi

mkdir /app/data || true

# Make mounted config at /app/server.yaml effective, because the binary reads server.yaml from CWD.
# The process runs in /app/data, so keep /app/data/server.yaml in sync with /app/server.yaml.
if [ -f /app/server.yaml ]; then
    cp /app/server.yaml /app/data/server.yaml
fi

cd /app/data

#if [ ! -f /app/server.db ]; then # This is not good if one wants to upgrade instance
/app/rustdesk-api-server-pro sync
#fi

if [ ! -f /app/data/.init.lock ] && [ -n "$ADMIN_USER" ] && [ -n "$ADMIN_PASS" ]; then
    /app/rustdesk-api-server-pro user add $ADMIN_USER $ADMIN_PASS --admin
    touch /app/data/.init.lock
fi

/app/rustdesk-api-server-pro start
