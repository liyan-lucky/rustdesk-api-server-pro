# 当前仓库状态

更新时间：2026-08-14

## 项目定位

RustDesk API Server Pro 是兼容 RustDesk 客户端 API 的第三方服务端实现，包含 Go 后端、Vue 3 管理后台及 SQLite/MySQL 数据存储。推荐使用单容器部署，由同一 HTTP 服务提供客户端 API、管理 API 和后台静态页面。

## 版本与运行基线

- 当前版本以根目录 `VERSION` 为唯一事实来源，文档不复制会在 CI 中自动递增的版本号。
- 主分支为 `main`，RustDesk 当前兼容基线见 `docs/compat/rustdesk-current.json`。
- 默认 HTTP 端口为 `12345`，生产环境可由配置或环境变量覆盖。
- 管理后台构建产物由后端内置并从 `/app/dist` 提供，不需要独立 Web 前端容器。

## 登录链路

### 账号密码

- 管理后台、用户门户和客户端 WebAuth 复用同一套账号认证规则。
- “关于与更新”页面提供登录验证码开关；关闭后上述密码登录入口均不校验验证码。
- 客户端 WebAuth 使用官方轮询协议：浏览器只负责完成认证，客户端通过 poll token 获取直接的认证 JSON，不依赖浏览器协议唤醒才能完成登录。

### 客户端 WebAuth

1. 客户端从 `/api/login-options` 获取 WebAuth 入口。
2. 浏览器打开复用后台主题的登录页，携带一次性 poll token。
3. 登录成功后服务端保存一次性认证结果；客户端轮询接口直接返回认证体。
4. 成功页尝试自行关闭；因浏览器安全策略无法关闭时，显示“返回 RustDesk”和“关闭页面”兜底按钮。
5. `rustdesk://config/` 仅用于可选的客户端聚焦，不参与认证结果传递，也不自动触发外部协议确认框。

### OAuth/OIDC

- 管理后台与客户端统一使用 `/admin/auth/oauth/{provider}/callback`。
- state 中的 poll token 区分客户端登录和 Web 后台登录。
- Web 后台使用一次性 `oauth_ticket`/`oidc_ticket` 换取会话；前端只消费一次并在用户信息就绪后显式进入目标页面。
- 客户端回调写入轮询结果，浏览器成功页不承担 token 回传职责。
- 回调失败统一传递 `ERR-22xx`，前端仍兼容旧版符号参数。
- Provider 配置以数据库中的 `oauth.providers` 为准；GitHub、QQ、Google、Microsoft、Gitee、GitLab、WeChat、Apple 已具备协议适配。真实可用性仍取决于 Provider 配置、账户绑定规则和部署环境外网连通性。
- Web 后台回调固定进入公开登录页消费 ticket，原目标独立放入安全 `redirect`；已有会话不会跳过新的第三方 ticket。
- 第三方身份只有存在有效绑定记录时才直接登录；首次登录或解除绑定后必须进入确认页，用户可验证目标本地账户密码后“绑定并登录”，或在 Provider 允许时明确选择“创建普通用户并登录”。管理员不能被自动创建或在客户端流程中绑定。

### Passkey

Passkey/WebAuthn 是独立的可选后台登录能力，不参与 RustDesk 客户端 WebAuth 或 OAuth 回传。相关凭据管理位于个人资料页面；客户端 WebAuth 页面不展示 Passkey 入口。

## 错误与鉴权规范

- 对外业务错误统一为 `ERR-xxxx: Message`；控制器公共错误出口会为遗留的无编码错误补充 `ERR-B010`。
- OAuth 回调页面显示和传递真实 `ERR-22xx` 编码。
- 鉴权中间件的 401/406 响应均使用 `StopWithJSON`。
- 官方客户端明确要求的协议哨兵值（例如 OIDC 轮询尚未完成）保持官方原文，不套用通用错误包装，以免破坏客户端兼容。

## 页面与前端状态

- 登录页采用后台现有主题与组件，不会把完整管理后台直接嵌入客户端 WebAuth 流程。
- 登录弹窗、邮件日志详情和邮件模板编辑弹窗均限制在当前视口内，并支持内部滚动。
- 移动端登录卡片按可用视口宽度收缩，第三方登录按钮在窄屏改为单列；移动端侧栏会透出全局背景并跟随主题毛玻璃参数。
- 表格页面沿用公共自适应表格容器；窄屏操作区允许换行。
- 已移除未注册且会干扰 Elegant Router 类型生成的旧 `views/about/index.vue`，当前使用 `about/version` 和 `about/help` 子页面。
- 普通用户不显示也不能调用邮件模板和邮件日志管理；设备列表仅返回当前账户关联设备。
- 服务器客户端链接配置在半屏宽度下把操作按钮换行到输入框下方，移动端按单列排列。

## 国际化（i18n）状态

当前翻译覆盖率（基准语言 en-US，共 716 个字段）：

| 语言 | 翻译数 | 覆盖率 | 回退率 |
|------|--------|--------|--------|
| zh-CN | 714 | 99.72% | 0.28% |
| fr-FR | 669 | 93.44% | 6.56% |
| de-DE | 665 | 92.88% | 7.12% |
| it-IT | 490 | 68.44% | 31.56% |
| ko-KR | 513 | 71.65% | 28.35% |
| ja-JP | 512 | 71.51% | 28.49% |
| ru-RU | 512 | 71.51% | 28.49% |
| es-ES | 503 | 70.25% | 29.75% |

- ja-jp、ko-kr、ru-ru、es-es 之前因 UTF-8 编码错误导致乱码，已修复并完成核心 UI 元素翻译。
- fr-fr 修复了 `action` 字段未翻译的问题（`Action` → `Opération`）。
- 翻译质量检查配置：`soybean-admin/scripts/i18n-quality.config.json`（严格模式）。
- 翻译基线对比配置：`soybean-admin/scripts/i18n-quality-baseline.json`。

## 性能优化

- 首页背景图片添加预加载标签（`<link rel="preload" href="/login-background.jpg" as="image" />`），减少首页加载延迟。
- 毛玻璃效果初始化时机优化：确保 CSS 变量注入到 DOM 后再计算毛玻璃变量，避免初始化失败。

## 验证基线

本次状态更新对应的本地验证：

- `go test ./...`
- `go vet ./...`
- `npm run typecheck`
- `npm run build`

测试设备上的真实链路验证：

- Web 后台 Gitee OAuth 已在完全退出的浏览器状态下完成 ticket 兑换并进入绑定用户页面。
- RustDesk 客户端退出后选择 WebAuth，浏览器完成密码认证并自动关闭标签页，客户端结束等待状态并显示已登录账户。

Windows 下 ZIP 路径安全检查已同时识别 Unix 根路径、Windows 根路径和盘符；Unix 文件权限断言只在支持该语义的平台执行。

## 已知边界与后续验收

- OAuth Provider 的真实登录需要逐项配置真实应用并验证回调白名单，不能仅凭单元测试宣称全部 Provider 已完成线上验收。
- 浏览器脚本只能关闭由脚本打开的窗口；普通标签页无法自动关闭属于浏览器限制，页面必须保留手动关闭提示。
- 部署后应至少实测：后台密码登录、后台 OAuth、客户端 WebAuth、客户端 OAuth、验证码开启/关闭各一次，并检查容器错误日志中无 token、Secret 或密码泄漏。
