# 当前仓库状态

更新时间：2026-07-31

## 定位

`rustdesk-api-server-pro` 是面向 RustDesk 客户端的第三方 API 服务端实现，包含 Go 后端和 Vue 管理后台前端。当前推荐部署方式为单容器一体化服务：同一个 HTTP 端口同时提供 RustDesk 客户端 API、管理后台 API 和管理后台静态页面。

## 当前架构

- 后端：`backend/`，Go 1.21.4（CI 使用 1.23.x）。
- 前端：`soybean-admin/`，Vue 3 / Vite / TypeScript / Naive UI。
- 默认数据库：SQLite。
- 可选数据库：MySQL。
- 默认配置文件：`backend/server.yaml`；容器内持久化配置为 `/app/data/server.yaml`。
- 默认端口：`12345/tcp`，可通过 `PORT` 环境变量覆盖。
- Docker 镜像：`ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest`。
- RustDesk 兼容版本：1.4.9。
- 服务端版本：由 VERSION 文件控制，CI 自动递增 PATCH 号。

## 当前能力边界

- RustDesk 客户端主流程 API 兼容增强（兼容 1.4.9）。
- 地址簿、设备列表、用户列表、审计日志和文件传输日志等基础管理能力。
- 我的同步设备：管理员查看所有设备，普通用户通过用户门户查看自己的设备。
- 通讯录：地址簿、地址簿管理、标签管理三个子页面。
- 版本自动递增系统：VERSION 文件为单一事实来源，CI 每次构建自动递增 PATCH 版本号。
- 首页更新日志区域显示服务端版本与构建时间。
- 管理后台前端已内置到镜像，旧 `rustdesk-web` / nginx 前端容器不再是必需组件。
- 第三方登录支持骨架：`oidc`、`google`、`github`，新版推荐使用 `oauth.providers`。
- plugin-sign、部分 OIDC 和高级企业能力仍以兼容占位或主流程兼容为主，不能宣称完整替代官方 Pro。

## 当前目录职责

- `backend/`：Go 后端 API 服务、配置、数据库、命令行和业务逻辑。
- `soybean-admin/`：管理后台前端源码和多语言词条。
- `docker/`：容器启动脚本与 OpenWrt 一体化部署脚本。
- `docs/`：Docker、OpenWrt、端口、排障和项目说明文档。
- `Dockerfile`：多阶段镜像构建文件。
- `docker-compose.yaml`：Compose 示例。

## 当前分支和备份

- `main`：当前主工作分支。
- `backup`：`main` 的快照备份分支。
- `.github/workflows/force-backup-main.yml`：手动输入 `YES` 后，把 `main` 当前提交强制覆盖到 `backup`。

## 部署结论

当前文档和部署脚本应统一推荐单容器一体化部署。升级旧部署时，应停止继续访问旧 `rustdesk-web` 前端容器；后端镜像内置的 `/app/dist` 才是当前管理后台入口。

## 维护要求

1. 修改端口、部署方式、默认配置或 OAuth 行为时，同步更新本文件、根 README 和 `docs/PROJECT_DESCRIPTION.md`。
2. Docker / OpenWrt 命令应保持 host 网络、`/mnt/docker` 数据目录、中文 label 和端口 label 风格。
3. 生产环境必须修改 `signKey`，并在升级时保持固定。
4. 不要在仓库中提交数据库、密钥、token、真实账号密码、OAuth secret 或生产配置。
5. YAML 配置文件中安全相关字段不要加行尾注释（会导致 shell 脚本解析失败）。
6. 启动脚本最后一行必须用 `exec`（确保 Go 进程成为 PID 1）。
7. sed 替换 YAML 值时必须限定段范围（避免误改其他段的同名字段）。
8. 新增前端 i18n key 时必须同步更新 `app.d.ts` 的 TypeScript 类型定义。
9. 修改测试期望值后必须同步提交测试文件。
