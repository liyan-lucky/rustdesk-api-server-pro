# 项目协作与发布规范

本文用于统一仓库结构、开发分支、验证、提交和发布规则。若本文与临时操作习惯冲突，以本文和 GitHub Actions 中的质量门禁为准。

## 1. 分支与工作区

- `main` 是唯一日常开发、集成和发布分支；本地修改、提交、推送均在 `main` 上进行。
- 开始工作前先同步 `origin/main`，结束工作后确认本地 `main` 与远端一致。
- `backup` 是唯一必须保留的非主分支，仅用于 `main` 快照备份，不在该分支开发；需要刷新时使用 `.github/workflows/force-backup-main.yml`。
- 除 `backup` 外，不保留长期功能分支、临时发布分支或版本维护分支。确需隔离实验时，不得将实验分支作为发布来源，最终变更仍需整理回 `main` 并删除实验分支。
- 工作区存在未确认来源的改动时不得覆盖、重置或批量暂存；先核对变更归属，再按文件提交。

## 2. 项目结构与职责

| 路径 | 职责 | 变更后的最低验证 |
| --- | --- | --- |
| `backend/` | Go API、鉴权、兼容层、数据库、任务和静态资源服务 | `go vet ./...`、`go test ./...` |
| `soybean-admin/` | Vue 3 管理后台、用户门户、类型和多语言 | `pnpm ci:frontend` |
| `docs/` | 当前状态、API 矩阵、部署、排障和发布规范 | 核对链接、版本和实现一致性 |
| `docker/`、`Dockerfile`、`docker-compose.yaml` | 镜像、容器启动和一体化部署 | Docker 构建与 smoke |
| `scripts/` | 兼容检查、发布辅助和维护脚本 | 执行对应脚本的最小验证 |
| `.github/workflows/` | 质量门禁、版本递增、GHCR 与 Release | 检查触发条件和最小权限 |
| `VERSION` | 服务端版本单一事实来源 | 不手工制造与自动递增冲突 |

## 3. 当前架构事实

- 后端使用 Go、Iris、Xorm，默认 SQLite，同时支持 MySQL。
- 前端使用 Vue 3、Vite、TypeScript 和 Naive UI，由后端同端口提供构建后的静态文件。
- `/api/*` 服务 RustDesk 客户端，`/admin/*` 服务管理后台，`/user-portal/*` 服务普通用户门户。
- 推荐单容器部署，运行数据和实际配置持久化到 `/app/data`。
- RustDesk 兼容目标、已实现范围和占位能力以 `API_MATRIX.md`、`CURRENT_STATUS.md` 为准，不能把兼容占位描述为官方 Pro 完整实现。

## 4. 修改规则

1. 先从官方客户端源码或公开文档确认真实请求路径、方法和字段形态。
2. 控制器保持薄层，业务逻辑放入 `internal/service`，数据访问放入 `internal/repository`，HTTP 结构放入 `internal/transport/httpdto`。
3. 修复兼容问题必须增加回归测试，至少覆盖正常请求、缺省字段和权限边界。
4. 数据模型变化必须确认 `sync` 可升级 SQLite/MySQL，并同步更新部署和排障文档。
5. API 行为变化必须同步更新 `API_MATRIX.md`、`CURRENT_STATUS.md` 及相关 README/专项文档。
6. 新增数据模型必须注册到 `cmd/sync.go` 的 models 列表，确保 `sync` 命令能自动建表。
6. 禁止提交数据库、日志、录屏、密钥、token、OAuth secret、生产配置或其他运行数据。

## 5. 提交与发布

1. 在本地 `main` 同步远端并完成修改。
2. 运行与变更范围相符的后端、前端、兼容、Docker 或安全检查。
3. 只暂存本次修改文件，使用可读的简短提交说明。
4. 直接推送 `main`；推送后不得绕过失败的 GitHub Actions 门禁手工发布镜像。
5. `main` 的线上质量流程全部成功后，`ghcr-docker.yml` 自动递增 PATCH 版本，同时更新 `VERSION` 和 `docs/compat/rustdesk-current.json`，再把同一版本注入后端、前端和 GHCR 镜像。
6. 发布完成后核对版本号、`latest`、`main`、`sha-*` 标签以及 `VERSION` 自动递增提交；缺少版本号标签不得视为发布完成。
7. 发布完成后将远端 `backup` 对齐到已发布的 `main`；`backup` 只保存快照，不触发正式发布。
8. 任一门禁失败时先修复失败原因并重新推送 `main`，不得用旧提交或本地未验证镜像冒充正式发布。

## 6. 文档事实入口

- 项目当前状态：`CURRENT_STATUS.md`
- 完整项目说明：`PROJECT_DESCRIPTION.md`
- API 实现状态：`API_MATRIX.md`
- 官方 API 兼容计划：`OFFICIAL_API_COMPATIBILITY_PLAN.md`
- 发布流程：`RELEASE_PROCESS.md`
- 故障排查：`TROUBLESHOOTING.md`

## 7. 通讯录与翻译发布门禁

- 地址簿、标签和 peer 的写操作必须分别从管理后台和 RustDesk 客户端 API 做真实增删改查回归，不能只验证接口可达。
- 标签重命名或删除必须同步更新 peer 内的标签引用；新版地址簿标签优先于遗留 `tags` 表，旧客户端兼容响应不得被陈旧数据遮蔽。
- 每个语言文件必须实际注册到 locale、语言选择器、Day.js 和 Naive UI；发布前执行严格键值检查并生成覆盖率报告。
- 新菜单必须明确标注 `R_SUPER`、`R_USER` 或双方共享；隐藏菜单不能替代后端授权，普通用户接口必须按当前用户过滤且不得接受客户端提供的 user_id。
- 角色差异应优先体现在同一业务菜单的可见子页面和数据范围中，不新增“个人工作台”：普通用户仅看到个人资料、自己的通讯录/设备/登录状态，以及只读可复制的服务器配置。
- 更新检查地址由“关于与更新”页面配置，不得硬编码为唯一发布站点；浏览器直连自定义地址，避免服务端更新检查形成 SSRF 入口。
- 通讯录页面必须直接展示当前账户可管理的全部联系人和标签，以“地址簿名称”作为数据列及表头筛选条件，不使用页面级地址簿选择器切换数据集。
- 管理员新增联系人、标签或地址簿时，界面必须要求先选择归属用户，并仅展示该用户的地址簿；普通用户创建时归属只能是自己。管理员为用户创建的地址簿必须记录来源，普通用户不得删除，且限制必须由后端强制执行。
- 具有新增能力的数据管理页面必须同时提供 CSV 批量导入、导出；数据表业务列应在表头提供筛选入口，导出应覆盖当前筛选结果而非仅当前分页。
## 设备更新与版本确认

- 设备更新统一使用 `scripts/deploy/update-rustdesk-api.sh`。脚本保留旧容器，直至新容器健康检查和 `/api/version` 版本匹配成功；失败自动回滚。
- 发布确认必须同时记录 `VERSION`、镜像 OCI version/revision、RepoDigest 和设备 `/api/version` 返回值，不能只检查 `latest` 标签。
