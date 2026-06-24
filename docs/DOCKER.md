# Docker 安装与配置参考（详细版）

本文档提供当前项目的 Docker / Docker Compose 安装、配置、默认参数、升级和常见排查方式。

当前推荐部署形态是**一体化容器**：同一个 `rustdesk-api-server-pro` 容器同时提供管理后台前端、管理后台接口和 RustDesk 客户端 API。

```text
rustdesk-api-server-pro
├── /          管理后台前端静态页面
├── /admin/*   管理后台接口
└── /api/*     RustDesk 客户端接口
```

如果你之前单独部署过 `rustdesk-web` / nginx 前端容器，升级到当前镜像后建议删除旧前端容器，避免旧版 `dist` 覆盖新版登录页，导致第三方登录按钮不显示或页面行为不一致。

## 安装命令示例（可直接执行）

### 方式一：`docker compose`（推荐）

```bash
mkdir -p /opt/rustdesk-api-server-pro/data
cd /opt/rustdesk-api-server-pro

# 下载示例配置（下载后请修改 signKey、端口、数据库等）
curl -L -o server.yaml https://raw.githubusercontent.com/liyan-lucky/rustdesk-api-server-pro/main/backend/server.yaml

cat > docker-compose.yaml <<'YAML'
services:
  rustdesk-api-server-pro:
    container_name: rustdesk-api-server-pro
    image: ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest
    environment:
      - "ADMIN_USER=admin"
      - "ADMIN_PASS=ChangeMe123!"
    volumes:
      - ./server.yaml:/app/server.yaml
      - ./data:/app/data
    network_mode: host
    restart: unless-stopped
    labels:
      name: "RustDesk API Server Pro"
      desc: "RustDesk API 增强服务端：管理后台前端、后端 API、第三方登录、SQLite 数据持久化"
      ports: "12345/tcp"
YAML

docker compose up -d
docker compose logs -f rustdesk-api-server-pro
```

### 方式二：`docker run`（host 网络）

```bash
mkdir -p /opt/rustdesk-api-server-pro/data
cd /opt/rustdesk-api-server-pro

curl -L -o server.yaml https://raw.githubusercontent.com/liyan-lucky/rustdesk-api-server-pro/main/backend/server.yaml

docker run -d \
  --name rustdesk-api-server-pro \
  --restart unless-stopped \
  --network host \
  --label "name=RustDesk API Server Pro" \
  --label "desc=RustDesk API 增强服务端：管理后台前端、后端 API、第三方登录、SQLite 数据持久化" \
  --label "ports=12345/tcp" \
  -e ADMIN_USER=admin \
  -e ADMIN_PASS='ChangeMe123!' \
  -v $(pwd)/server.yaml:/app/server.yaml \
  -v $(pwd)/data:/app/data \
  ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest

docker logs -f rustdesk-api-server-pro
```

### 方式三：OpenWrt / x86 软路由一键对齐更新

OpenWrt 场景推荐阅读：

- [OpenWrt 一体化部署与对齐更新指南](./OPENWRT_ONE_CONTAINER.md)

可直接执行：

```bash
curl -L -o /tmp/update-openwrt-one-container.sh \
  https://raw.githubusercontent.com/liyan-lucky/rustdesk-api-server-pro/main/docker/update-openwrt-one-container.sh

sh /tmp/update-openwrt-one-container.sh
```

脚本默认会：

- 使用 `/mnt/docker/rustdesk-api` 作为数据目录
- 备份旧数据到 `/mnt/docker/backup/rustdesk-api`
- 使用 `--network host`
- 删除旧 `rustdesk-web` 容器，避免旧前端干扰
- 启动一体化 `rustdesk-api-server-pro` 容器
- 添加中文 `label` 和 `ports=12345/tcp` 标记

## 常用命令

```bash
# compose 升级
docker compose pull
docker compose up -d

# compose 重启
docker compose restart rustdesk-api-server-pro

# 查看运行状态
docker compose ps

# 进入容器
docker exec -it rustdesk-api-server-pro sh

# 手动执行数据库同步（通常启动脚本已自动执行）
docker exec -it rustdesk-api-server-pro rustdesk-api-server-pro sync
```

## 1. 容器镜像与启动行为

容器启动时会执行以下流程：

1. 建立可执行文件软链接（如需要）
2. 创建 `/app/data`
3. 将挂载的 `/app/server.yaml` 复制到 `/app/data/server.yaml`
4. 切换工作目录到 `/app/data`
5. 执行 `rustdesk-api-server-pro sync`
6. 若首次启动且设置了 `ADMIN_USER` / `ADMIN_PASS`，自动创建管理员
7. 执行 `rustdesk-api-server-pro start`

程序读取配置文件的位置是当前工作目录下的 `server.yaml`。容器内进程实际在 `/app/data` 下运行，所以最终生效的配置文件是 `/app/data/server.yaml`。启动脚本会把你挂载到 `/app/server.yaml` 的配置复制到 `/app/data/server.yaml`。

## 2. 默认参数与目录

Dockerfile 中的默认项：

- 镜像内工作目录：`/app`
- 默认声明端口：`EXPOSE 8080`
- 启动命令：`sh /app/start.sh`

注意：`EXPOSE 8080` 只是镜像声明，不代表程序一定监听 8080。程序实际监听端口由 `server.yaml` 的 `httpConfig.port` 决定。

示例 `backend/server.yaml` 常见默认值：

- `httpConfig.port: ":12345"`
- `httpConfig.staticdir: "/app/dist"`
- `db.driver: "sqlite"`
- `db.dsn: "./server.db"`

建议持久化 `/app/data`，其中包含：

- `server.db`：SQLite 数据库
- `server.yaml`：实际生效配置
- `.init.lock`：自动创建管理员后的标记
- `record_uploads/`：录制上传文件

## 3. 配置文件最小示例

```yaml
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
  port: ":12345"

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
```

重点：

- `signKey` 必须修改，并且升级时保持固定。
- `httpConfig.staticdir` 在官方镜像内应为 `/app/dist`。
- SQLite 相对路径会落在 `/app/data`，因为进程工作目录是 `/app/data`。

## 4. 环境变量说明

当前镜像明确支持：

- `ADMIN_USER`：首次启动自动创建管理员用户名（可选）
- `ADMIN_PASS`：首次启动自动创建管理员密码（可选）

生效条件：

- 两者都非空
- `/app/data/.init.lock` 不存在

注意：首次创建后会写入 `.init.lock`。后续即使修改环境变量，也不会再次自动创建管理员，除非你删除 `.init.lock` 并自行确认风险。

## 5. 访问方式与端口说明

假设 `server.yaml` 配置 `httpConfig.port: ":12345"`：

- 管理后台页面：`http://<host>:12345/`
- 客户端 API：`http://<host>:12345/api/...`
- 管理后台接口：`http://<host>:12345/admin/...`

更多细节可结合：

- [端口与访问路径说明](./PORTS.md)

## 6. 第三方登录与反向代理

新版推荐使用 `oauth.providers` 配置 GitHub / Google / OIDC。回调路径为：

```text
/admin/auth/oauth/<provider>/callback
```

如果使用外部 Nginx / Caddy / OpenWrt 反代，请统一反代到一体化服务，例如：

```text
127.0.0.1:12345
```

不要再让 `/` 指向旧 `rustdesk-web` 的 dist。

反代建议传递：

```nginx
proxy_set_header Host $host;
proxy_set_header X-Forwarded-Host $host;
proxy_set_header X-Forwarded-Proto $scheme;
proxy_set_header X-Real-IP $remote_addr;
```

用于检查第三方登录配置是否生效：

```bash
curl http://127.0.0.1:12345/admin/auth/oauth/providers
curl "http://127.0.0.1:12345/admin/auth/oauth/url?provider=github"
```

## 7. 升级命令

### Compose 升级

```bash
docker compose pull
docker compose up -d
docker compose logs -f rustdesk-api-server-pro
```

### docker run 升级

```bash
docker pull ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest
docker rm -f rustdesk-api-server-pro
# 然后重新执行 docker run，保持原有 -v / -e / label / 网络参数
```

当前启动脚本会自动执行 `sync`，升级后仍建议查看日志确认数据库同步成功。

## 8. 常见问题

### 8.1 修改了 `server.yaml` 但配置没生效

排查：

1. 确认挂载的是 `/app/server.yaml`
2. 重启容器
3. 进入容器检查 `/app/data/server.yaml` 内容是否已更新

### 8.2 首页打不开，但 `/api/*` 正常

通常是 `httpConfig.staticdir` 配置错误。Docker 环境下应指向：

```text
/app/dist
```

### 8.3 第三方登录按钮不显示

常见原因：

- 仍然访问旧 `rustdesk-web` 容器里的旧前端
- `oauth.providers` 没有启用任何 provider
- 当前访问入口和 OAuth `redirectUrl` 不一致

建议：删除旧 `rustdesk-web` 容器，只保留一体化 `rustdesk-api-server-pro` 容器，再访问 `http://<host>:12345/`。

### 8.4 数据丢失

通常是没有持久化 `/app/data`。请确认：

```text
-v ./data:/app/data
```

或等价挂载已经配置。

### 8.5 录制上传失败

请确认：

- `/app/data/record_uploads` 可写
- 容器日志中无文件写入报错
- 宿主机磁盘空间充足

## 9. GHCR `denied` 排查

如果执行下面命令报错：

```bash
docker pull ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest
# Error response from daemon: Head ... denied
```

通常是以下原因之一：

- GHCR 包未发布
- GHCR 包是私有，匿名拉取会被拒绝
- `latest` 标签不存在
- 未登录 GHCR（私有包场景）

排查建议：

1. 在仓库 `Actions` 中手动运行 Docker build 工作流，确认推送成功。
2. 到 GitHub 仓库右侧 `Packages` 打开 `rustdesk-api-server-pro`，将包可见性改为 `Public`。
3. 若暂时不公开，先执行 `docker login ghcr.io`，使用 GitHub 用户名 + PAT（需 `read:packages` 权限）。
4. 检查镜像标签是否存在。

## 10. 生产部署建议

最小上线清单：

1. 修改并固定 `signKey`
2. 配置正确的 `httpConfig.port`
3. 持久化 `/app/data`
4. 首次启动创建管理员
5. 删除旧前端容器，避免旧 dist 干扰
6. 如使用反代，统一反代到一体化服务
7. 用最新版 RustDesk 客户端做冒烟测试：登录、地址簿、设备列表、分组面板、审计、录制上传
