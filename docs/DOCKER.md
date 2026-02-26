# Docker 安装与配置参考（详细版）

本文档提供当前项目的 Docker / Docker Compose 安装、配置、默认参数说明和常见排查方式。

适用对象：

- 希望快速部署并上线使用
- 需要可复制的参考命令
- 需要了解容器内默认目录、端口、环境变量与卷挂载关系

## 快速安装命令（参考原作者写法风格）

下面这部分优先给出“可直接复制”的命令块，说明放在后面章节。

### 方式一：`docker compose`（推荐）

```bash
# 创建目录
mkdir -p /opt/rustdesk-api-server-pro/data
cd /opt/rustdesk-api-server-pro

# 下载示例配置（下载后请修改 signKey、端口、数据库等）
curl -L -o server.yaml https://raw.githubusercontent.com/liyan-lucky/rustdesk-api-server-pro/master/backend/server.yaml

# 创建 docker-compose.yaml
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
YAML

# 启动
docker compose up -d

# 查看日志
docker compose logs -f rustdesk-api-server-pro
```

### 方式二：`docker run`（host 网络）

```bash
mkdir -p /opt/rustdesk-api-server-pro/data
cd /opt/rustdesk-api-server-pro

curl -L -o server.yaml https://raw.githubusercontent.com/liyan-lucky/rustdesk-api-server-pro/master/backend/server.yaml

docker run -d \
  --name rustdesk-api-server-pro \
  --restart unless-stopped \
  --network host \
  -e ADMIN_USER=admin \
  -e ADMIN_PASS='ChangeMe123!' \
  -v $(pwd)/server.yaml:/app/server.yaml \
  -v $(pwd)/data:/app/data \
  ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest

docker logs -f rustdesk-api-server-pro
```

### 常用命令（升级 / 重启 / 进入容器）

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

快速说明：

- 默认示例使用 `host` 网络，实际监听端口以 `server.yaml` 中 `httpConfig.port` 为准（常见 `:12345`）
- 首次启动会自动 `sync`
- 同时设置 `ADMIN_USER` 和 `ADMIN_PASS` 时会自动创建管理员（仅首次）

## 1. 先说结论（推荐部署方式）

推荐优先使用 `docker-compose.yaml`，并至少持久化以下内容：

- `./server.yaml -> /app/server.yaml`
- `./data -> /app/data`

原因：

- `/app/data` 会存放 SQLite 数据库、初始化锁文件、录制上传目录等运行数据
- 容器启动脚本会在启动时自动执行 `sync`
- 若设置了 `ADMIN_USER` + `ADMIN_PASS`，首次启动会自动创建管理员账号

## 2. 容器镜像与启动行为（当前仓库）

### 2.1 镜像入口行为（`docker/start.sh`）

容器启动时会执行以下流程：

1. 建立可执行文件软链接（如需要）
2. 创建 `/app/data`
3. 将挂载的 `/app/server.yaml` 复制到 `/app/data/server.yaml`（保证配置生效）
4. 切换工作目录到 `/app/data`
5. 执行 `rustdesk-api-server-pro sync`
6. 若首次启动且设置了 `ADMIN_USER` / `ADMIN_PASS`，自动创建管理员
7. 执行 `rustdesk-api-server-pro start`

### 2.2 为什么配置要特别说明

程序读取配置文件的位置是“当前工作目录下的 `server.yaml`”。

- 容器内进程实际在 `/app/data` 下运行
- 所以最终生效的配置文件是 `/app/data/server.yaml`

本仓库的启动脚本已经做了兼容处理：会把你挂载到 `/app/server.yaml` 的配置复制到 `/app/data/server.yaml`。

## 3. 默认参数与默认目录（重要）

### 3.1 Dockerfile 中的默认项

- 镜像内工作目录：`/app`
- 默认暴露端口：`8080`（`EXPOSE 8080`）
- 启动命令：`sh /app/start.sh`

注意：

- `EXPOSE 8080` 只是镜像声明，不代表程序一定监听 `8080`
- 程序实际监听端口由 `server.yaml` 的 `httpConfig.port` 决定

### 3.2 示例 `server.yaml` 默认值（仓库内）

当前示例 `backend/server.yaml` 常见默认值：

- `httpConfig.port: ":12345"`
- `httpConfig.staticdir: "/app/dist"`
- `db.driver: "sqlite"`
- `db.dsn: "./server.db"`（在容器内实际路径会是 `/app/data/server.db`，因为工作目录是 `/app/data`）
- `smtpConfig.port: 1025`（连接外部 SMTP 用，不是本程序监听端口）

### 3.3 运行时数据目录（容器内）

建议持久化 `/app/data`，其中会包含（按当前实现）：

- `server.db`（SQLite）
- `server.yaml`（实际生效配置）
- `.init.lock`（自动创建管理员后的标记）
- `record_uploads/`（录制上传文件，若客户端使用录制上传）

## 4. Docker Compose 部署（推荐）

### 4.1 参考配置（当前仓库基础版）

仓库内 `docker-compose.yaml` 已提供基础配置，当前关键点：

- 使用镜像：`ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest`
- `network_mode: host`
- 支持环境变量：
  - `ADMIN_USER`
  - `ADMIN_PASS`
- 建议挂载：
  - `./server.yaml:/app/server.yaml`
  - `./data:/app/data`

### 4.2 Linux 启动命令（推荐）

```bash
docker compose up -d
```

查看日志：

```bash
docker compose logs -f rustdesk-api-server-pro
```

停止：

```bash
docker compose down
```

### 4.3 首次启动建议

1. 先准备 `server.yaml`
2. 设置 `ADMIN_USER` / `ADMIN_PASS`
3. `docker compose up -d`
4. 观察日志，确认 `sync` 成功且管理员创建成功

## 5. `docker run` 参考命令（不使用 Compose）

### 5.1 Host 网络模式（Linux）

适合简单部署，端口由 `server.yaml` 控制：

```bash
docker run -d \
  --name rustdesk-api-server-pro \
  --restart unless-stopped \
  --network host \
  -e ADMIN_USER=admin \
  -e ADMIN_PASS='ChangeMe123!' \
  -v $(pwd)/server.yaml:/app/server.yaml \
  -v $(pwd)/data:/app/data \
  ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest
```

说明：

- `--network host` 下不需要 `-p`
- 程序监听端口以 `server.yaml` 为准（例如 `:12345`）

### 5.2 端口映射模式（更通用）

如果不使用 host 网络，需要端口映射，并保证映射端口与容器内监听端口一致：

假设 `server.yaml` 中为 `:12345`：

```bash
docker run -d \
  --name rustdesk-api-server-pro \
  --restart unless-stopped \
  -p 12345:12345 \
  -e ADMIN_USER=admin \
  -e ADMIN_PASS='ChangeMe123!' \
  -v $(pwd)/server.yaml:/app/server.yaml \
  -v $(pwd)/data:/app/data \
  ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest
```

如果你把 `server.yaml` 改成 `:8080`，则映射改为：

```bash
-p 8080:8080
```

## 6. 从源码自行构建镜像（可选）

在仓库根目录执行：

```bash
docker build -t rustdesk-api-server-pro:local .
```

然后使用 `docker run` 或 `docker compose` 将镜像名替换为：

```bash
rustdesk-api-server-pro:local
```

## 7. 环境变量说明（容器）

当前镜像明确支持：

- `ADMIN_USER`：首次启动自动创建管理员用户名（可选）
- `ADMIN_PASS`：首次启动自动创建管理员密码（可选）

生效条件：

- 两者都非空
- `/app/data/.init.lock` 不存在（即首次初始化）

注意：

- 首次创建后会写入 `.init.lock`
- 后续即使修改环境变量，也不会再次自动创建管理员（除非你删除 `.init.lock` 并自行确认风险）

## 8. `server.yaml` 推荐最小配置（Docker）

下面是一份适合容器部署的参考配置（SQLite）：

```yaml
signKey: "please-change-this-sign-key"
debugMode: false

db:
  driver: "sqlite"
  dsn: "./server.db"
  timeZone: "Asia/Shanghai"
  showSql: false

httpConfig:
  printRequestLog: true
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

说明：

- `staticdir` 在容器内应使用 `/app/dist`
- `db.dsn` 用相对路径时会落在 `/app/data`（因为进程工作目录是 `/app/data`）
- `signKey` 请务必修改

## 9. 访问方式与端口说明（Docker 下）

假设 `server.yaml` 配置 `httpConfig.port: ":12345"`：

- 管理后台页面：`http://<host>:12345/`
- 客户端 API：`http://<host>:12345/api/...`
- 管理后台接口：`http://<host>:12345/admin/...`

更多细节可结合：

- [端口与访问路径说明](./PORTS.md)

## 10. 升级命令（Docker / Compose）

### 10.1 Compose（使用远端镜像）

```bash
docker compose pull
docker compose up -d
```

说明：

- 当前启动脚本会自动执行 `sync`
- 升级后仍建议查看日志确认数据库同步成功

### 10.2 `docker run` 方式升级（示例）

```bash
docker pull ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest
docker rm -f rustdesk-api-server-pro
# 然后重新执行 docker run（保持原有 -v / -e / 端口参数）
```

## 11. 常见问题（Docker 场景）

### 11.1 修改了 `server.yaml` 但配置没生效

排查：

1. 确认挂载的是 `/app/server.yaml`
2. 重启容器（启动脚本会复制到 `/app/data/server.yaml`）
3. 进入容器检查 `/app/data/server.yaml` 内容是否已更新

### 11.2 首页打不开，但 `/api/*` 正常

通常是 `httpConfig.staticdir` 配置错误。

Docker 环境下应指向：

- `/app/dist`

### 11.3 数据丢失（升级后用户/设备/地址簿没了）

通常是没有持久化 `/app/data`。

请确认：

- `-v ./data:/app/data`（或等价挂载）已配置

### 11.4 录制上传失败

请确认：

- `/app/data/record_uploads` 可写（宿主机目录权限正确）
- 容器日志中无文件写入报错

## 12. 上线建议（Docker）

最小上线清单：

1. 修改 `signKey`
2. 配置正确的 `httpConfig.port`
3. 持久化 `/app/data`
4. 首次启动创建管理员（环境变量或手动）
5. 用最新版客户端做冒烟测试（登录、地址簿、设备列表、分组面板、审计、录制上传）

## GHCR `denied` 排查（补充）

如果执行下面命令报错：

```bash
docker pull ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest
# Error response from daemon: Head ... denied
```

通常是以下原因之一：

- GHCR 包未发布（GitHub Actions 的 `Docker build` 工作流尚未手动执行）
- GHCR 包是私有（Private），匿名拉取会被拒绝
- `latest` 标签不存在
- 未登录 GHCR（私有包场景）

排查建议：

1. 先在仓库 `Actions` 中手动运行 `Docker build` 工作流，确认推送成功。
2. 到 GitHub 仓库右侧 `Packages` 打开 `rustdesk-api-server-pro`，将包可见性改为 `Public`。
3. 若暂时不公开，先执行 `docker login ghcr.io`，使用 GitHub 用户名 + PAT（需 `read:packages` 权限）。
4. 检查镜像标签是否存在（确认拉取的是 `latest`）。
