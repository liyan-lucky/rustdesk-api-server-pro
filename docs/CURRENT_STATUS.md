# 当前仓库状态

更新时间：2026-08-08

## 定位

`rustdesk-api-server-pro` 是面向 RustDesk 客户端的第三方 API 服务端实现，包含 Go 后端和 Vue 管理后台前端。当前推荐部署方式为单容器一体化服务：同一个 HTTP 端口同时提供 RustDesk 客户端 API、管理后台 API 和管理后台静态页面。

## 当前版本

**v1.2.29** (2026-08-07)，由 `VERSION` 文件控制，CI 自动递增 PATCH 号并同步兼容清单；API、前端、镜像标签使用同一构建版本。

## 当前架构

- 后端：`backend/`，Go 1.21.4（CI 使用 1.23.x）。
- 前端：`soybean-admin/`，Vue 3 / Vite / TypeScript / Naive UI。
- 默认数据库：SQLite。
- 可选数据库：MySQL。
- 默认配置文件：`backend/server.yaml`；容器内持久化配置为 `/app/data/server.yaml`。
- 默认端口：`12345/tcp`，可通过 `PORT` 环境变量覆盖。
- Docker 镜像：`ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest`。
- RustDesk 兼容版本：1.4.9。
- 管理后台前端已内置到镜像，旧 `rustdesk-web` / nginx 前端容器不再是必需组件。

## 当前能力边界

- RustDesk 客户端主流程 API 兼容增强（兼容 1.4.9）。
- 心跳接口回显客户端 `modified_at`，避免未分配策略时触发持续策略重同步。
- 心跳已按设备和设备组下发策略，禁用设备会返回连接断开列表；共享地址簿已实现用户、用户组和 everyone 权限规则。
- 管理前端启动后及每 5 分钟核对后端运行版本，发现版本或构建变化时显示更新提示。
- 地址簿、共享地址簿、设备列表、用户列表、审计日志和文件传输日志等基础管理能力。
- 设备管理：管理员查看所有设备；普通用户在同一菜单中仅查看自己账户已绑定的设备。
- 通讯录：联系人和标签按当前账户全量展示，直接显示"地址簿名称"/归属人并支持表头筛选，不使用页面级地址簿切换器；地址簿管理固定排在通讯录菜单底部。联系人、标签和地址簿管理均支持 CSV 批量导入导出与实际增删改。管理员新增联系人、标签或地址簿时必须先明确选择归属用户，再选择该用户的地址簿；普通用户只能操作自己的数据，且不能删除管理员代建地址簿。客户端 API 已兼容 `/api/ab/get` 与共享 profile 路径别名。
- 标签同步以 `address_book_tag` 为新版唯一事实来源，兼容读取旧 `tags` 表；标签重命名或删除会在同一事务内更新所有 peer 的标签引用，避免管理端与客户端显示不一致。
- 共享地址簿：`shared=true` 可跨用户读取；写入 peer/tag 需要 owner 或共享规则 `rule >= 2`。
- 管理后台语言共 9 种（含此前存在但未注册的意大利语）；i18n 严格检查要求缺失键、额外键和可疑占位符均为 0，覆盖率报告保存在 `soybean-admin/docs/i18n-coverage-report.md`。
- 管理员与普通用户共用业务菜单，不新增"个人工作台"。用户管理对管理员显示用户/会话管理，对普通用户仅显示个人资料；通讯录、设备管理、登录日志按角色自动限制数据范围；服务器配置对普通用户只读并支持复制。所有个人接口均按当前用户 ID 在后端强制过滤。
- "关于与更新"支持修改在线版本检查地址；管理员还可保存、恢复和复制容器更新命令模板，模板中的 `{version}` 自动替换为检查到的最新版本。
- "关于与更新"页面对两类角色开放，显示运行版本、构建时间和兼容版本；在线更新检查地址可修改并保存在浏览器，支持纯语义版本文件或常见 JSON 版本字段，方便镜像发布到其他网站后切换检查源。
- 版本自动递增系统：VERSION 文件为单一事实来源，CI 每次构建自动递增 PATCH 版本号。
- 首页更新日志区域显示服务端版本与构建时间。
- 第三方登录统一使用 `oauth.providers`，管理员优先在"第三方登录"页面配置。GitHub 已完成 authorization code、PKCE S256、持久化一次性 state/ticket、已验证私有邮箱、管理员/普通用户角色绑定及自动创建开关；QQ 已完成网站应用 OAuth2、OpenID 身份与用户资料协议适配，等待真实回调验收。Google、Microsoft、Gitee、GitLab、WeChat、Apple 已完成协议适配和前端配置界面；Apple 使用动态 JWT client_secret（ES256），WeChat 使用 appid 参数和逗号分隔 scope。Client Secret 接口只返回 `********` 加末 8 位的识别提示，不返回明文；未修改提示直接保存会保留原密钥。
- WebAuthn / Passkey 登录已实现：服务器自签名认证，不依赖第三方 OAuth。后端自动推导 RP ID 和 Origins，前端在登录页显示 Passkey 按钮，用户可在个人资料页注册和管理 Passkey 凭据。HTTP 页面通过 `window.open()` 拉起 HTTPS 认证页面完成 WebAuthn 操作（登录/注册），通过 `postMessage` 传回结果。HTTPS 反向代理监听独立端口（`tlsPort`），自签名证书支持 `tlsHosts` 配置多个域名/IP。`GetServerConfig()` 已改为单例缓存，确保 WebAuthn auto-config 跨请求持久化。
- OAuth/OIDC 成功目标和失败目标均已统一为前端 hash 路由；失败回调固定返回登录页，不再先进入受保护页面后闪退。错误码区分账户不可绑定、Provider 网络不可达、state 过期和其他失败。
- OAuth 统一回调：客户端和 admin 共用回调端点 `/admin/auth/oauth/{provider}/callback`，通过 state 中的 `PollToken` 区分；客户端回调返回 HTML 页面包含 `rustdesk://oauth/callback?poll_token=xxx` 链接。
- 错误码索引体系：`errcode.go` 注册 105+ 个错误码，Message 统一为 PascalCase 作为 i18n key；前端 `parseBackendMessage()` 从 `ERR-xxxx: Message` 提取编码和翻译；帮助页面提供错误码搜索和筛选；错误日志表 `ErrorLog` 全局记录后端错误。
- Google 与通用 OIDC 保持兼容。微信、Gitee、Microsoft、GitLab、Apple 已完成协议适配。
- 普通用户查看权限增强：普通用户可查看日志审计和系统设置信息（只读），操作类 API 通过 `isAdmin` 检查限制；前端操作按钮通过 `isAdmin` 控制显示。
- plugin-sign、部分 OIDC 和高级企业能力仍以兼容占位或主流程兼容为主，不能宣称完整替代官方 Pro。

## 最近完成的工作

### v1.2.29 (2026-08-08)

#### 3. WebAuthn / Passkey 登录 ✅

- 后端：添加 `go-webauthn/webauthn v0.17.4` 依赖，新建 `WebauthnCredential` 和 `WebauthnSession` 数据模型，注册到 `cmd/sync.go`
- 后端：新建 `WebauthnService`（`backend/internal/service/webauthn_service.go`），实现注册/登录/列表/删除/重命名完整流程
- 后端：添加公开路由 `GET /auth/webauthn/enabled`、`POST /auth/webauthn/login/begin`、`POST /auth/webauthn/login/finish`
- 后端：添加鉴权路由 `POST /webauthn/register/begin`、`POST /webauthn/register/finish`、`GET /webauthn/credentials`、`DELETE /webauthn/credentials/{id}`、`PUT /webauthn/credentials/{id}`
- 后端：WebAuthn 配置自动推导 RP ID 和 Origins（从请求 Host 和 X-Forwarded-Proto），无需手动配置
- 后端：添加错误码 ERR3101-ERR3110
- 前端：新建 `webauthn.ts` 工具函数（base64url 编解码、WebAuthn 选项转换、凭据序列化）
- 前端：在 `pwd-login.vue` 和 `user-login.vue` 添加 Passkey 登录按钮，自动检测浏览器支持和服务器启用状态
- 前端：在 `auth store` 添加 `loginByPasskey` 方法
- 前端：新建 `passkey-manager.vue` 组件，嵌入 `user/profile` 和 `workspace/profile` 页面，支持注册/列表/删除/重命名
- 前端：添加 9 种语言的 i18n 词条和 `app.d.ts` 类型定义
- 测试状态：已部署到设备，`/admin/auth/webauthn/enabled` 返回 `enabled: true`，登录流程端点验证通过

#### 2. 修复版本提示不停弹出 ✅

- `soybean-admin/src/plugins/app.ts` 增加 localStorage 持久化忽略状态
- 点"稍后"或 X 关闭写入忽略记录，点"立即更新"清除记录并刷新

#### 1. 修复 CI workflow 失败 ✅

- `docs/compat/rustdesk-current.json` 移除 git merge 冲突标记，恢复合法 JSON

### v1.2.28 (2026-08-06)

- OAuth 统一回调基础实现
- 客户端和 admin 共用回调端点

### v1.2.27 (2026-08-05)

- AdminOrUserAuth 中间件基础实现

## 当前目录职责

- `backend/`：Go 后端 API 服务、配置、数据库、命令行和业务逻辑。
- `soybean-admin/`：管理后台前端源码和多语言词条。
- `docker/`：容器启动脚本与 OpenWrt 一体化部署脚本。
- `docs/`：Docker、OpenWrt、端口、排障和项目说明文档。
- `Dockerfile`：多阶段镜像构建文件。
- `docker-compose.yaml`：Compose 示例。

## 当前分支和备份

- `main`：当前主工作分支。本地开发、提交、推送和正式发布统一在 `main` 进行。
- `backup`：`main` 的只读快照备份分支，由 `.github/workflows/force-backup-main.yml` 手动输入 `YES` 后强制覆盖。
- 远端仅额外保留 `backup`，不保留其他长期或临时分支。

## 部署结论

当前文档和部署脚本应统一推荐单容器一体化部署。升级旧部署时，应停止继续访问旧 `rustdesk-web` 前端容器；后端镜像内置的 `/app/dist` 才是当前管理后台入口。

## 当前测试设备状态

- SSH：`ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net`（当前密钥不允许直接以 `root` 登录）。
- 容器：`rustdesk-api-server-pro`，host 网络，API/后台端口 `16888`。
- 更新脚本：`/vol1/docker/rustdesk-api-server-pro/update-rustdesk-api.sh`。
- 运行版本：`1.2.29`，容器状态已验证为 `running`。设备通过 `update-rustdesk-api.sh` 脚本自动拉取最新镜像并重启。
- GitHub/Gitee/Microsoft OAuth 配置已启用且 Secret 已保存。测试设备访问 `github.com:443` 偶尔超时；`api.github.com:443` 返回 200。该网络问题会阻断 token 交换。
- 当前有效管理员没有填写邮箱，且 GitHub 配置为按邮箱绑定、禁止自动创建管理员。网络恢复后需补充邮箱或决定是否开启自动创建。

## 待办事项

### 高优先级
- [ ] 在浏览器中完整测试 WebAuthn 注册和登录流程
- [ ] 确认客户端 OAuth 回调功能正常工作
- [ ] 测试普通用户查看权限是否正确限制操作

### 中优先级
- [ ] 添加更多单元测试覆盖新增功能
- [ ] 优化 OAuth 回调页面的用户体验
- [ ] 完善错误处理和用户提示

### 低优先级
- [ ] 性能优化
- [ ] 代码重构

## 已知问题

### 客户端 OAuth 回调
- **现象**：用户反馈浏览器提示显示 `http://...` 而不是 `rustdesk://...`
- **分析**：可能是浏览器安全提示的正常行为（提示框显示当前页面域名，不是目标 URL scheme）
- **调试**：已添加 HTML 注释显示 schemeURL 和 pollToken
- **下一步**：等待用户测试反馈

## 部署历史

| 版本 | 日期 | 主要变更 | 状态 |
|------|------|----------|------|
| v1.2.29 | 2026-08-08 | WebAuthn/Passkey 登录 + 修复版本提示 + 修复 CI | 已部署 |
| v1.2.29 | 2026-08-07 | 普通用户查看权限 + 客户端 OAuth 统一回调 | 已部署 |
| v1.2.28 | 2026-08-06 | OAuth 统一回调基础实现 | 已部署 |
| v1.2.27 | 2026-08-05 | AdminOrUserAuth 中间件 | 已部署 |

## 关键文件清单

### 后端核心文件
- `backend/app/middleware/auth.go` — 鉴权中间件（AdminAuth、UserAuth、AdminOrUserAuth）
- `backend/app/route.go` — 路由配置
- `backend/app/controller/admin/auth.go` — OAuth 回调处理 + WebAuthn 登录端点
- `backend/app/controller/admin/webauthn.go` — WebAuthn 注册管理控制器
- `backend/internal/service/webauthn_service.go` — WebAuthn 服务
- `backend/app/controller/api/compat_public.go` — 客户端兼容 API
- `backend/internal/service/oauth_provider_service.go` — OAuth 服务
- `backend/internal/errcode/errcode.go` — 错误码定义

### 前端核心文件
- `soybean-admin/src/router/elegant/routes.ts` — 路由配置
- `soybean-admin/src/views/audit/*/` — 审计页面
- `soybean-admin/src/views/system/*/` — 系统设置页面
- `soybean-admin/src/utils/webauthn.ts` — WebAuthn 浏览器端工具函数
- `soybean-admin/src/views/user/profile/modules/passkey-manager.vue` — Passkey 管理组件

### 配置文件
- `backend/server.yaml` — 服务配置
- `VERSION` — 版本号（单一事实来源）

## 维护要求

1. 修改端口、部署方式、默认配置或 OAuth 行为时，同步更新本文件、根 README 和 `docs/PROJECT_DESCRIPTION.md`。
2. Docker / OpenWrt 命令应保持 host 网络、`/mnt/docker` 数据目录、中文 label 和端口 label 风格。
3. 生产环境必须修改 `signKey`，并在升级时保持固定。
4. 不要在仓库中提交数据库、密钥、token、真实账号密码、OAuth secret 或生产配置。
5. YAML 配置文件中安全相关字段不要加行尾注释（会导致 shell 脚本解析失败）。
6. 启动脚本最后一行必须用 `exec`（确保 Go 进程成为 PID 1）。
7. sed 替换 YAML 值时必须限定段范围（避免误改其他段的同名字段）。
8. 新增前端 i18n key 时必须同步更新 `app.d.ts` 的 TypeScript 类型定义和全部 9 种语言文件。
9. 修改测试期望值后必须同步提交测试文件。
10. 所有语言文件必须注册到 locale、语言选择器、Day.js 与 Naive UI 映射；发布前运行 i18n 严格检查和覆盖率报告，禁止遗漏语言或缺失键。
11. 地址簿写入必须同时验证管理端 `/admin/ab/*` 和客户端 `/api/ab/*` 的真实增删改查；标签重命名、删除必须同步更新 peer 标签引用，并验证旧版 `/api/ab/get` 兼容读取结果。
12. 错误消息必须带 `ERR-xxxx` 编码前缀，与帮助页错误码索引匹配；新增 errcode 时同步更新 `errcode.go`、`app.d.ts` 类型定义和全部语言文件。
13. Apple 私钥 placeholder 不得包含 `-----BEGIN PRIVATE KEY-----` 文本，避免触发合规检查密钥检测。
14. 新增数据模型必须注册到 `cmd/sync.go` 的 models 列表，否则 `sync` 命令不会自动建表。
15. 鉴权中间件返回 401/406 时必须使用 `StopWithJSON` 返回带 ERR-xxxx 编码的 JSON，不能返回纯文本；前端 `onError` 必须处理 HTTP 错误中的后端 JSON 消息。
16. 登录对话框宽度 320px，表单 medium，OAuth 按钮 small + 2 列网格；移动端 `w-[calc(100vw-48px)]`。

## 下次开发建议

### 新开对话时应该：
1. 阅读本文件了解当前状态
2. 阅读 `AGENTS.md` 了解开发规则
3. 检查待办事项列表
4. 确认已知问题是否已解决
5. 询问用户下一步要做什么

### 开发流程：
1. 在 `main` 分支开发
2. 编写代码 + 测试
3. 本地验证
4. 部署到设备测试
5. 更新文档
6. 提交代码
7. 更新本文件
