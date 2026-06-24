#!/bin/sh
set -eu

# OpenWrt / x86 Docker one-container updater for rustdesk-api-server-pro.
# Default layout follows the user's deployment convention:
#   - host network mode
#   - /mnt/docker persistent data
#   - labels describing Chinese name, description and used ports
#   - backup before replacing containers

BASE_DIR="${BASE_DIR:-/mnt/docker/rustdesk-api}"
BACKUP_DIR="${BACKUP_DIR:-/mnt/docker/backup/rustdesk-api}"
IMAGE="${IMAGE:-ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest}"
CONTAINER="${CONTAINER:-rustdesk-api-server-pro}"
LEGACY_WEB_CONTAINER="${LEGACY_WEB_CONTAINER:-rustdesk-web}"
PORT="${PORT:-12345}"
ADMIN_USER="${ADMIN_USER:-admin}"
ADMIN_PASS="${ADMIN_PASS:-ChangeMe123!}"
REMOVE_LEGACY_WEB="${REMOVE_LEGACY_WEB:-true}"
CLEAN_OLD_IMAGES="${CLEAN_OLD_IMAGES:-true}"

log() {
  printf '\n==> %s\n' "$*"
}

warn() {
  printf '\n[WARN] %s\n' "$*" >&2
}

require_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "缺少命令: $1" >&2
    exit 1
  fi
}

create_default_config() {
  cat > "$BASE_DIR/server.yaml" <<YAML
signKey: "please-change-this-sign-key"
debugMode: false

db:
  driver: "sqlite"
  dsn: "./server.db"
  timeZone: "Asia/Shanghai"
  showSql: false

httpConfig:
  printRequestLog: false
  staticdir: "/app/dist"
  port: ":$PORT"

smtpConfig:
  host: "127.0.0.1"
  port: 1025
  username: ""
  password: ""
  encryption: "none"
  from: "noreply@example.com"

jobsConfig:
  deviceCheckJob:
    duration: 30

oidc:
  enabled: false
  providerName: "oidc"
  issuer: ""
  clientId: ""
  clientSecret: ""
  redirectUrl: ""
  scopes:
    - "openid"
    - "profile"
    - "email"
  bindByEmail: true
  autoCreateAdmin: false
  stateTtlSeconds: 180
  ticketTtlSeconds: 180
  successRedirect: "/login"
  failureRedirect: "/login"
  subjectClaim: "sub"
  emailClaim: "email"
  nameClaim: "name"
  pictureClaim: "picture"
  prompt: ""
  allowedEmailDomains: []

oauth:
  providers: []
YAML
}

require_cmd docker
require_cmd tar
require_cmd date

log "创建目录"
mkdir -p "$BASE_DIR/data" "$BACKUP_DIR"

log "备份当前数据"
TS="$(date +%Y%m%d-%H%M%S)"
BACKUP_FILE="$BACKUP_DIR/rustdesk-api-$TS.tar.gz"
if [ -d "$BASE_DIR" ]; then
  tar -czf "$BACKUP_FILE" -C "$(dirname "$BASE_DIR")" "$(basename "$BASE_DIR")"
  echo "备份完成: $BACKUP_FILE"
else
  warn "未找到 $BASE_DIR，跳过旧数据备份"
fi

log "准备 server.yaml"
if [ ! -f "$BASE_DIR/server.yaml" ]; then
  warn "未找到 $BASE_DIR/server.yaml，已生成最小配置；上线前务必修改 signKey"
  create_default_config
fi

if grep -q 'please-change-this-sign-key' "$BASE_DIR/server.yaml" 2>/dev/null; then
  warn "server.yaml 仍使用默认 signKey，请修改为固定随机字符串后再用于生产环境"
fi

log "拉取最新镜像"
docker pull "$IMAGE"

log "停止并删除旧 API 容器"
docker rm -f "$CONTAINER" 2>/dev/null || true

if [ "$REMOVE_LEGACY_WEB" = "true" ]; then
  log "删除旧 rustdesk-web 前端容器，避免旧 dist 干扰新版登录页"
  docker rm -f "$LEGACY_WEB_CONTAINER" 2>/dev/null || true
else
  warn "已跳过删除 $LEGACY_WEB_CONTAINER；请确认它没有继续提供旧版 dist"
fi

log "启动一体化容器"
docker run -d \
  --name "$CONTAINER" \
  --restart unless-stopped \
  --network host \
  --label "name=RustDesk API Server Pro" \
  --label "desc=RustDesk API 增强服务端：管理后台前端、后端 API、第三方登录、SQLite 数据持久化" \
  --label "ports=$PORT/tcp" \
  -e ADMIN_USER="$ADMIN_USER" \
  -e ADMIN_PASS="$ADMIN_PASS" \
  -v "$BASE_DIR/server.yaml:/app/server.yaml" \
  -v "$BASE_DIR/data:/app/data" \
  "$IMAGE"

log "等待服务启动"
sleep 3

log "容器状态"
docker ps --filter "name=$CONTAINER"

log "最近日志"
docker logs --tail=120 "$CONTAINER" || true

log "本机接口检查"
if command -v curl >/dev/null 2>&1; then
  curl -fsS "http://127.0.0.1:$PORT/admin/auth/oauth/providers" || true
  printf '\n'
else
  warn "系统没有 curl，跳过接口检查"
fi

if [ "$CLEAN_OLD_IMAGES" = "true" ]; then
  log "清理悬空旧镜像"
  docker image prune -f || true
fi

cat <<EOF

更新完成。
访问地址: http://<OpenWrt-IP>:$PORT/
数据目录: $BASE_DIR/data
配置文件: $BASE_DIR/server.yaml
备份文件: $BACKUP_FILE

注意：
1. ADMIN_USER / ADMIN_PASS 仅在 /app/data/.init.lock 不存在时生效。
2. 配置以 $BASE_DIR/server.yaml 为准，容器启动时会复制到 /app/data/server.yaml。
3. 新版已经内置管理后台前端，不需要单独 rustdesk-web 容器。
4. 若使用外部反代，请统一反代到 127.0.0.1:$PORT，并传递 X-Forwarded-Proto / X-Forwarded-Host。
EOF
