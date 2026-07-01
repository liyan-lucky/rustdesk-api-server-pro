# RustDesk API Server Pro 文档索引

## 当前状态入口

1. `CURRENT_STATUS.md`：当前仓库事实、架构、能力边界、分支/备份策略和部署结论。
2. `PROJECT_DESCRIPTION.md`：项目定位、系统组成、接口边界、数据模型和维护重点。
3. `DOCKER.md`：Docker / Docker Compose 安装、配置、升级和排查。
4. `OPENWRT_ONE_CONTAINER.md`：OpenWrt x86 / 软路由一体化部署与对齐更新。
5. `PORTS.md`：端口与访问路径说明。
6. `TROUBLESHOOTING.md`：常见问题与排查。
7. `CLEANUP_HISTORY.md`：清理历史记录说明。

## 当前主事实

- 当前主工作分支：`main`。
- 当前备份分支：`backup`，由 `.github/workflows/force-backup-main.yml` 手动输入 `YES` 后强制覆盖。
- 当前推荐部署方式：单容器一体化服务。
- 当前默认端口：`12345/tcp`，以 `server.yaml` 的 `httpConfig.port` 为准。
- 当前镜像内置管理后台前端，旧 `rustdesk-web` / nginx 前端容器不再是必需组件。
- 当前第三方登录推荐使用 `oauth.providers`，支持 `oidc`、`google`、`github`。

## 文档维护要求

- 部署方式、端口、Docker/OpenWrt 脚本、OAuth 行为、数据库或默认配置变化时，必须同步更新 `CURRENT_STATUS.md`、根 README 和相关专项文档。
- 不要把生产数据库、密钥、token、真实账号密码、OAuth secret 或生产配置提交到仓库。
- OpenWrt / Docker 命令应保持 host 网络、`/mnt/docker` 数据目录、中文 label 和端口 label 风格。
