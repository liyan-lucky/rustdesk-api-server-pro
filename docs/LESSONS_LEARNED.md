# 经验教训与技术笔记

本文档记录项目维护过程中形成的经验、偏好、技术实现细节和避免重复踩坑的约定。

## 线上展示文案

- 线上页面、线上包和用户可见文案应使用中文业务描述。
- 不在用户可见区域展示提交信息、提交哈希、提交链接、`chore:`、`feat:`、`fix:` 等开发流转痕迹。
- 更新日志卡片只保留用户能理解的功能说明，例如"支持客户端 1.4.9 主流程""新增手机扫码导入配置"。
- 如果确实需要保留提交链接或提交历史，应放在 Git 仓库历史、开发文档或发布内部记录中，不放入线上页面。
- GitHub Release 不使用自动生成发布说明，避免线上 Release 正文出现提交标题、提交哈希和提交链接；发布说明应手写为中文功能摘要。
- 发布包构建时应注入前端版本标签，避免管理后台更新日志长期显示 `latest`，影响线上版本确认。

## RustDesk 配置导入

- "复制RustDesk模板"和"生成二维码"必须复用同一份文本生成逻辑，避免手机扫码导入与剪贴板导入不一致。
- 二维码仅在前端本地生成，不上传配置内容，避免泄露 ID Server、Relay Server、API Server 和 KEY。
- 修改服务器配置页面后，至少执行前端类型检查和生产构建，确认二维码依赖、自动导入组件和语言包类型均正常。
- 涉及前端按钮和后端兼容版本的更新，部署时必须同时替换后端程序和 `dist` 静态资源；只换其中一边会导致页面仍显示旧版本或新按钮不可见。

## 版本兼容记录

- RustDesk 官方客户端版本更新后，需要同步检查后端 `CompatSysinfoVersion`，并用最新客户端做登录、地址簿、设备列表、审计和扫码导入冒烟验证。
- RustDesk 官方客户端版本更新后，不能只改兼容版本字符串；需要从官方源码梳理新增 API 路径。1.4.7 新增的 `/api/devices/deploy` 在当前项目中应返回 `NOT_ENABLED`，除非后续真正实现 hbbs 显式部署白名单，不能伪造 `OK`。
- 版本字符串只表示已适配和验证的兼容目标，不等于完整实现官方 Pro 所有能力。

## OAuth 统一回调实现

### 问题背景
OAuth provider 只能注册一个回调URL，但 admin 和客户端需要不同的处理：
- Admin: 回调后重定向到管理后台，携带 ticket
- 客户端: 回调后返回HTML页面，包含 `rustdesk://` URL scheme

### 解决方案
客户端和 admin 共用同一个回调URL (`/admin/auth/oauth/{provider}/callback`)，通过 state 中的 `PollToken` 区分：

1. **客户端发起登录**: 调用 `BuildClientAuthURL` 生成授权URL，state 中包含 `PollToken`，返回授权URL和 pollToken 给客户端
2. **Admin 发起登录**: 调用 `BuildAdminAuthURL` 生成授权URL，state 中不包含 `PollToken`，返回授权URL给前端
3. **统一回调处理**: `ConsumeUnifiedCallback` 处理回调，检查 state 中的 `PollToken`。如果有则返回HTML页面包含 `rustdesk://oauth/callback?poll_token=xxx`；如果没有则返回 ticket 和重定向URL

### 浏览器安全提示
当用户点击 `rustdesk://...` 链接时，浏览器会显示安全提示（"此网站想打开 RustDesk"）。这是浏览器的正常行为，提示框显示的是当前页面域名，而不是目标URL scheme。

### 调试信息
在 `renderOAuthCallbackPage` 中添加了HTML注释：`<!-- DEBUG: schemeURL=rustdesk://oauth/callback?poll_token=xxx pollToken=xxx -->`，可以通过查看网页源代码确认 pollToken 是否正确。

## AdminOrUserAuth 中间件

### 问题背景
需要让普通用户可以查看日志审计和系统设置信息，但不能执行操作。

### 解决方案
新建 `AdminOrUserAuth` 中间件，允许管理员和普通用户访问，但操作类API需要检查 `isAdmin`。

### 前端权限控制
- 路由 roles 配置：`['R_SUPER', 'R_USER']`
- 按钮显示控制：`v-if="isAdmin"`
- 按钮禁用控制：`:disabled="!isAdmin"`

## 错误码规范

### ERR1xxx: 通用错误
- ERR1001: 验证码错误
- ERR1002: 用户不存在
- ERR1003: 用户名或密码错误
- ERR1004: Token无效或已过期
- ERR1005: 权限不足

### ERR2xxx: OAuth/OIDC 相关错误
- ERR2001: Provider不存在
- ERR2002: Provider未启用
- ERR2003: 缺少必要参数
- ERR2004: State无效或已过期
- ERR2005: 生成State失败
- ERR2006: 生成Ticket失败
- ERR2007: Ticket无效
- ERR2008: Ticket已过期
- ERR2009: 用户不可用

### ERR22xx: 客户端 OAuth 相关错误
- ERR2203: Provider不支持客户端登录
- ERR2204: PollToken无效
- ERR2205: PollToken已过期
- ERR2206: Ticket不是客户端类型
- ERR2207: 生成PollToken失败
- ERR2208: 账户不可绑定
- ERR2209: Provider不可达
- ERR2210: State已过期
- ERR2212: 未知错误

## 部署流程

### 编译（交叉编译 Linux 版本）
```bash
cd backend
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o rustdesk-api-server-pro
```

### 上传到容器并重启
```bash
# 上传到 NAS /tmp/
scp -P 2828 backend/rustdesk-api-server-pro LiYan@liyan-fnosnas.dynv6.net:/tmp/

# 复制到容器并设置权限
ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net 'docker cp /tmp/rustdesk-api-server-pro rustdesk-api-server-pro:/app/rustdesk-api-server-pro'
ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net 'docker exec rustdesk-api-server-pro chmod 755 /app/rustdesk-api-server-pro'

# 重启容器
ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net 'docker restart rustdesk-api-server-pro'
```

### 查看日志
```bash
ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net 'docker logs rustdesk-api-server-pro --tail 100'
```

### 部署关键教训
- 容器运行的是镜像内 `/app/rustdesk-api-server-pro` 二进制，不是数据目录 `/app/data/` 中的文件。
- 必须交叉编译 Linux 版本（`GOOS=linux GOARCH=amd64 CGO_ENABLED=0`），不能用 Windows PE 格式。
- 上传后必须 `chmod 755` 确保可执行权限。

## 性能优化建议

### 数据库
- 添加索引：`token_hash`, `user_id`, `rustdesk_id`
- 定期清理过期数据：`auth_token`, `oauth_login_session`
- 考虑使用连接池

### 缓存
- OAuth metadata 缓存（已实现）
- 用户信息缓存
- 配置缓存

### 前端
- 路由懒加载
- 组件按需加载
- 静态资源CDN

## 安全建议

### Token安全
- 使用强随机数生成 token
- Token 哈希存储
- 定期清理过期 token
- 防止重放攻击

### OAuth安全
- State 一次性使用
- PKCE 防止授权码拦截
- 回调URL验证
- 防止CSRF攻击

### 数据安全
- 敏感信息加密
- SQL注入防护
- XSS防护
- CSRF防护

## GHCR 镜像发布

- GHCR package 如果已存在且未授权当前仓库 Actions 写入，即使 workflow 配置了 `packages: write` 也会在推送阶段返回 `write_package` 权限拒绝。
- main 分支的 GHCR 发布必须挂在质量门禁后面：同一提交的线上工作流全部成功后才自动推送，任一流程失败都应阻止镜像发布。
- 发布 `v*.*.*` 标签时应自动构建并推送 GHCR `latest`、版本标签（含 `vX.Y.Z` 与 `X.Y.Z`）和 sha 标签，否则 Docker 用户拉取 `latest` 拿不到新后端和新 `dist`。
- Docker 发布默认使用 `GITHUB_TOKEN`，并依赖 GHCR Package settings 里的 `Manage Actions access` 给本仓库写权限。只有在仓库变量 `GHCR_USE_PAT=true` 且 `GHCR_TOKEN` 确认有效时，才切换到 PAT 登录。
- 需要发布 GHCR 镜像时，先在 GitHub Package 设置中授予本仓库写入权限，或配置带 `write:packages` 权限的有效发布令牌，再手动触发工作流并选择推送镜像。

## Docker 启动脚本（v1.1.16 血泪教训）

### signKey 行尾注释导致自动替换失败

- **现象**：容器启动报 `unsafe signKey` 后退出
- **根因**：`server.yaml` 中 `signKey: "CHANGE_ME..." # comment` 的行尾注释导致 `grep | sed | tr` 管道提取的值包含注释文本，不匹配 case 中的任何不安全 key 模式，sed 替换被跳过
- **教训**：**YAML 配置文件中安全相关字段不要加行尾注释**；shell 脚本解析 YAML 时必须先 `sed 's/#.*//'` 去注释再提取值

### start.sh 无条件覆盖用户配置

- **现象**：用户修改的 `server.yaml` 在容器重启后被恢复为默认值
- **根因**：`cp /app/server.yaml /app/data/server.yaml` 每次启动都覆盖
- **教训**：**仅在目标文件不存在时才复制**：`if [ ! -f /app/data/server.yaml ]; then cp ...; fi`

### PORT sed 误改 smtpConfig.port 导致 Web 404

- **现象**：设置 `PORT` 环境变量后，管理后台返回 404
- **根因**：`sed "s|^\([[:space:]]*port:\).*|..."` 匹配了所有 `port:` 行，把 `smtpConfig.port: 1025`（整数）也改成了字符串，YAML 解析失败，`config.go` 用默认值覆盖整个文件，默认 `staticdir: "dist"`（相对路径）在 `/app/data` 下找不到前端文件
- **教训**：
  1. **sed 替换必须限定 YAML 段范围**：`/^httpConfig:/,/^[^ ]/ s|...`
  2. **YAML 解析失败时绝不能覆盖用户配置文件**，只返回内存中的默认值
  3. **默认 `staticdir` 必须用绝对路径** `/app/dist`，不能用相对路径 `dist`

### Shell 不是 PID 1 导致容器停止报错

- **现象**：`docker stop` 后容器 exit code 非零
- **根因**：Shell 作为 PID 1 不转发 SIGTERM 给子进程，Docker 超时后 SIGKILL 强杀
- **教训**：**启动脚本最后一行必须用 `exec`**：`exec /app/rustdesk-api-server-pro start`

### iris.Context 不是 context.Context

- **现象**：编译报错 `cannot convert ctx to type iris.Context`
- **根因**：Iris 框架的 `iris.Context` 是请求上下文接口，不是 Go 标准库的 `context.Context`
- **教训**：**优雅关闭用 goroutine 监听 context + `app.Shutdown()`**，不要把 `context.Context` 传给 Iris 的 Listen 选项

## KEY 自动检测（v1.1.16 新增）

- KEY 优先级：`RUSTDESK_KEY` 环境变量 > `RUSTDESK_HBBS_DIR` 环境变量目录 > CWD > `/app/data` > `/root`
- hbbs 默认输出 `id_ed25519.pub` 到 `/root`，Docker 挂载 `-v /宿主机/hbbs:/root` 即可自动检测
- 后端返回的 source 必须是前端 i18n `sourceType` 中已定义的类型（`env`/`inferred`/`auto`/`empty`），否则前端无法显示
- 新增 i18n key 时必须同步更新 `app.d.ts` 的 TypeScript 类型定义，否则前端 CI 类型检查失败
- 新增 i18n key 时必须同步更新全部 9 种语言文件（en-us、zh-cn、ja-jp、ko-kr、fr-fr、de-de、es-es、ru-ru、it-it），否则 TypeScript 类型检查失败
- errcode Message 必须使用 PascalCase（如 `CaptchaError`），确保可作为 i18n key 使用
- Apple 私钥 placeholder 不得包含 `-----BEGIN PRIVATE KEY-----` 文本，否则合规检查会误判为真实密钥
- 新增数据模型必须注册到 `cmd/sync.go` 的 models 列表，否则 `sync` 命令不会自动建表，运行时报 `no such table` 错误
- 鉴权中间件返回 401/406 时必须使用 `StopWithJSON` 返回 JSON 格式（含 ERR-xxxx 编码），不能使用 `StopWithText` 返回纯文本；前端 `onError` 必须处理 HTTP 错误中的后端 JSON 消息

## CI/CD 相关

- **修改测试期望值后必须同步提交**：只改实现不改测试，CI 必然失败
- **go.mod 声明的 Go 版本应与 CI 一致**：`go 1.21.4` 但 CI 用 `1.23.x` 可能导致行为差异
- **`pnpm install --frozen-lockfile` 要求 lockfile 必须与 package.json 完全匹配**，否则 Docker 构建失败
- **所有修改必须 `git add` + `git commit`**：工作区修改不会自动包含在推送中

## OpenWrt/Clash 代理相关

- Clash fake-ip 模式的 `store_fakeip` 必须开启（设为 `1`），否则重启后缓存丢失，已分配的 fake-ip 映射失效，导致 TLS 握手失败（`wrong version number`）
- `cachesize_dns` 建议设为 `9999`，减少 DNS 缓存未命中导致的连接中断
- `docker pull ghcr.io` 报 `http: server gave HTTP response to HTTPS client` 时，通常是 Clash fake-ip 缓存失效，重启 Clash 或开启持久化即可恢复
