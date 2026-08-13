#!/usr/bin/env bash
set -Eeuo pipefail
IMAGE="${IMAGE:-ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest}"
CONTAINER="${CONTAINER:-rustdesk-api-server-pro}"
API_URL="${API_URL:-}"
EXPECTED_VERSION="${EXPECTED_VERSION:-}"
HBBS_HOST_DIR="${HBBS_HOST_DIR:-/vol1/docker/rustdesk/hbbs}"
HBBS_CONTAINER_DIR="${HBBS_CONTAINER_DIR:-/root}"
BACKUP="${CONTAINER}-rollback"
log(){ printf '[rustdesk-api-update] %s\n' "$*"; }
die(){ log "ERROR: $*" >&2; exit 1; }
command -v docker >/dev/null || die "docker is required"
command -v jq >/dev/null || die "jq is required"
docker inspect "$CONTAINER" >/dev/null 2>&1 || die "container $CONTAINER not found"
work_dir="$(mktemp -d)"; trap 'rm -rf "$work_dir"' EXIT
env_file="$work_dir/container.env"
docker inspect "$CONTAINER" | jq -r '.[0].Config.Env[]' > "$env_file"; chmod 600 "$env_file"
if grep -q '^RUSTDESK_HBBS_DIR=' "$env_file"; then
  sed -i "s|^RUSTDESK_HBBS_DIR=.*|RUSTDESK_HBBS_DIR=$HBBS_CONTAINER_DIR|" "$env_file"
else
  printf 'RUSTDESK_HBBS_DIR=%s\n' "$HBBS_CONTAINER_DIR" >> "$env_file"
fi
mapfile -t mount_args < <(docker inspect "$CONTAINER" | jq -r --arg hbbs "$HBBS_CONTAINER_DIR" '.[0].Mounts[] | select(.Destination != $hbbs) | "-v", (.Source + ":" + .Destination)')
mkdir -p "$HBBS_HOST_DIR"
mount_args+=("-v" "$HBBS_HOST_DIR:$HBBS_CONTAINER_DIR:ro")
mapfile -t port_args < <(docker inspect "$CONTAINER" | jq -r '
  .[0].HostConfig.PortBindings // {} | to_entries[] |
  .key as $container | .value[]? |
  "-p", (((.HostIp // "") | if . == "0.0.0.0" or . == "::" then "" else . + ":" end) + .HostPort + ":" + $container)
')
network="$(docker inspect "$CONTAINER" --format '{{.HostConfig.NetworkMode}}')"
restart="$(docker inspect "$CONTAINER" --format '{{.HostConfig.RestartPolicy.Name}}')"
if [[ -z "$API_URL" ]]; then
  app_port="$(grep -E '^PORT=' "$env_file" | tail -n1 | cut -d= -f2- || true)"
  app_port="${app_port:-12345}"
  host_port="$(docker inspect "$CONTAINER" | jq -r --arg key "${app_port}/tcp" '.[0].HostConfig.PortBindings[$key][0].HostPort // empty')"
  API_URL="http://127.0.0.1:${host_port:-$app_port}"
fi
log "pulling $IMAGE"; docker pull "$IMAGE"
image_version="$(docker image inspect "$IMAGE" --format '{{index .Config.Labels "org.opencontainers.image.version"}}')"
image_revision="$(docker image inspect "$IMAGE" --format '{{index .Config.Labels "org.opencontainers.image.revision"}}')"
image_digest="$(docker image inspect "$IMAGE" --format '{{index .RepoDigests 0}}')"
image_version="${image_version#v}"
[[ -z "$EXPECTED_VERSION" || "$image_version" == "${EXPECTED_VERSION#v}" ]] || die "image version $image_version != expected $EXPECTED_VERSION"
docker rm -f "$BACKUP" >/dev/null 2>&1 || true
docker stop "$CONTAINER" >/dev/null; docker rename "$CONTAINER" "$BACKUP"
run_args=(run -d --name "$CONTAINER" --restart "$restart" --env-file "$env_file")
[[ "$network" == "default" ]] || run_args+=(--network "$network")
run_args+=("${port_args[@]}")
run_args+=("${mount_args[@]}" "$IMAGE")
if ! docker "${run_args[@]}" >/dev/null; then docker rename "$BACKUP" "$CONTAINER"; docker start "$CONTAINER" >/dev/null; die "start failed; rollback restored"; fi
for _ in $(seq 1 30); do
  payload="$(curl -fsS "$API_URL/api/version" 2>/dev/null || true)"
  if [[ "$payload" == *"\"version\":\"$image_version\""* ]]; then
    docker rm -f "$BACKUP" >/dev/null
    log "updated version=$image_version revision=$image_revision digest=$image_digest"
    printf '%s\n' "$payload"; exit 0
  fi
  sleep 2
done
docker rm -f "$CONTAINER" >/dev/null 2>&1 || true; docker rename "$BACKUP" "$CONTAINER"; docker start "$CONTAINER" >/dev/null
die "health/version verification failed; rollback restored"
