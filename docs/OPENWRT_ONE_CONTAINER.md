# OpenWrt 一体化部署与对齐更新指南

本文档面向 OpenWrt x86 / 软路由 Docker 环境，目标是把旧的“后端 API 容器 + 单独 rustdesk-web 前端容器”对齐为当前推荐的一体化部署。

## 1. 当前推荐架构

当前镜像 `ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest` 已内置：

- Go 后端 API
- 管理后台接口 `/admin/*`
- RustDesk 客户端 API `/api/*`
- 管理后台前端静态文件 `/app/dist`
- 启动时数据库同步 `sync`

因此推荐只保留一个容器：

```text
rustdesk-api-server-pro
├── /          管理后台前端
├── /admin/*   管理后台接口
└── /api/*     RustDesk 客户端接口
```

旧的 `rustdesk-web` nginx 容器不再是必需组件。继续保留旧前端容器时，可能出现后端已经升级、但登录页仍是旧版，导致第三方登录按钮不显示或页面行为不一致。

## 2. OpenWrt 推荐目录

```text
/mnt/docker/rustdesk-api
├── server.yaml        # 宿主机主配置文件
└── data/              # 持久化数据目录
    ├── server.db      # SQLite 数据库
    ├── server.yaml    # 容器启动时由 /app/server.yaml 复制而来
    ├── .init.lock     # 首次管理员初始化标记
    └── record_uploads/
```

重要说明：

- 宿主机应主要维护 `/mnt/docker/rustdesk-api/server.yaml`。
- 容器启动脚本会将 `/app/server.yaml` 复制到 `/app/data/server.yaml`。
- 不建议只修改 `/mnt/docker/rustdesk-api/data/server.yaml`，因为下次容器启动可能会被覆盖。

## 3. 一键对齐更新

仓库已提供脚本：

```bash
sh docker/update-openwrt-one-container.sh
```

脚本默认行为：

- 使用 host 网络模式
- 使用 `/mnt/docker/rustdesk-api` 作为数据目录
- 启动前备份现有数据到 `/mnt/docker/backup/rustdesk-api`
- 拉取 `ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest`
- 删除旧 `rustdesk-api-server-pro` 容器并重新创建
- 默认删除旧 `rustdesk-web` 容器，避免旧 dist 干扰
- 添加中文 label 与端口 label
- 保留 `/mnt/docker/rustdesk-api/data` 中的数据库和运行数据

直接执行：

```bash
curl -L -o /tmp/update-openwrt-one-container.sh \
  https://raw.githubusercontent.com/liyan-lucky/rustdesk-api-server-pro/main/docker/update-openwrt-one-container.sh

sh /tmp/update-openwrt-one-container.sh
```

如果要指定端口或目录：

```bash
PORT=12345 \
BASE_DIR=/mnt/docker/rustdesk-api \
BACKUP_DIR=/mnt/docker/backup/rustdesk-api \
sh /tmp/update-openwrt-one-container.sh
```

如果暂时不想删除旧 `rustdesk-web` 容器：

```bash
REMOVE_LEGACY_WEB=false sh /tmp/update-openwrt-one-container.sh
```

但这种模式需要确认外部访问入口没有继续指向旧前端。

## 4. 手动 docker run 命令

```bash
mkdir -p /mnt/docker/rustdesk-api/data

# 首次使用时准备 server.yaml，可从仓库示例复制后修改 signKey
curl -L -o /mnt/docker/rustdesk-api/server.yaml \
  https://raw.githubusercontent.com/liyan-lucky/rustdesk-api-server-pro/main/backend/server.yaml

docker rm -f rustdesk-api-server-pro 2>/dev/null || true
docker rm -f rustdesk-web 2>/dev/null || true

docker pull ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest

docker run -d \
  --name rustdesk-api-server-pro \
  --restart unless-stopped \
  --network host \
  --label "name=RustDesk API Server Pro" \
  --label "desc=RustDesk API 增强服务端：管理后台前端、后端 API、第三方登录、SQLite 数据持久化" \
  --label "ports=12345/tcp" \
  -e ADMIN_USER=admin \
  -e ADMIN_PASS='ChangeMe123!' \
  -v /mnt/docker/rustdesk-api/server.yaml:/app/server.yaml \
  -v /mnt/docker/rustdesk-api/data:/app/data \
  ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest
```

说明：

- `ADMIN_USER` / `ADMIN_PASS` 只在 `/app/data/.init.lock` 不存在时自动创建管理员。
- 老数据目录存在时，修改这两个环境变量不会重置管理员密码。
- 程序实际监听端口以 `server.yaml` 的 `httpConfig.port` 为准。

## 5. 关键配置

`/mnt/docker/rustdesk-api/server.yaml` 至少应确认：

```yaml
signKey: "请改成固定随机字符串，不要每次升级变动"

httpConfig:
  printRequestLog: false
  staticdir: "/app/dist"
  port: ":12345"
```

重点：

- `signKey` 必须固定，不能每次升级重新生成。
- `staticdir` 在容器内应为 `/app/dist`。
- 端口改动时，脚本 label 中的 `PORT` 也建议同步。

## 6. 第三方登录对齐

新版推荐使用 `oauth.providers`，例如 GitHub：

```yaml
oauth:
  providers:
    - type: "github"
      name: "github"
      displayName: "GitHub"
      enabled: true
      clientId: "YOUR_GITHUB_CLIENT_ID"
      clientSecret: "YOUR_GITHUB_CLIENT_SECRET"
      redirectUrl: "http://your-host:12345/admin/auth/oauth/github/callback"
      scopes: ["read:user", "user:email"]
      bindByEmail: true
      autoCreateAdmin: true
      stateTtlSeconds: 180
      ticketTtlSeconds: 180
      successRedirect: "/#/login"
      failureRedirect: "/#/login"
```

检查接口：

```bash
curl http://127.0.0.1:12345/admin/auth/oauth/providers
curl "http://127.0.0.1:12345/admin/auth/oauth/url?provider=github"
```

如果返回 providers 非空，并且 URL 接口包含授权地址，说明后端配置已生效。

## 7. 外部反向代理建议

如果使用 Nginx / Caddy / OpenWrt 反代，只反代到一体化服务：

```text
127.0.0.1:12345
```

不要再单独把 `/` 指到旧 `rustdesk-web` 的 dist。

反代时建议传递：

```nginx
proxy_set_header Host $host;
proxy_set_header X-Forwarded-Host $host;
proxy_set_header X-Forwarded-Proto $scheme;
proxy_set_header X-Real-IP $remote_addr;
```

这样第三方登录生成回调地址时能识别外部入口。

## 8. 升级后检查

```bash
docker ps --filter name=rustdesk-api-server-pro
docker logs --tail=120 rustdesk-api-server-pro
curl http://127.0.0.1:12345/admin/auth/oauth/providers
```

浏览器访问：

```text
http://<OpenWrt-IP>:12345/
```

确认：

- 登录页是新版页面
- 第三方登录按钮按配置显示
- `/api/*` 路径正常
- `/admin/*` 路径正常
- 日志中没有数据库字段不存在或 sync 失败

## 9. 回滚

脚本执行前会备份数据，例如：

```text
/mnt/docker/backup/rustdesk-api/rustdesk-api-20260625-120000.tar.gz
```

回滚示例：

```bash
docker rm -f rustdesk-api-server-pro

cd /mnt/docker
mv rustdesk-api rustdesk-api.failed.$(date +%Y%m%d-%H%M%S)
tar -xzf /mnt/docker/backup/rustdesk-api/你的备份文件.tar.gz -C /mnt/docker

sh /tmp/update-openwrt-one-container.sh
```

如果只是镜像问题，也可以指定旧镜像标签：

```bash
IMAGE=ghcr.io/liyan-lucky/rustdesk-api-server-pro:<旧标签> sh /tmp/update-openwrt-one-container.sh
```
