# 第三方登录规范

## 当前支持状态

第三方登录统一使用 `oauth.providers`。GitHub 是第一套完成安全加固和端到端测试的 Provider；QQ 已完成网站应用 OAuth2 协议适配，等待使用已审核应用做真实回调验收；Google 与通用 OIDC 保持兼容。后续 Provider 应复用同一套一次性 state、短期 ticket、账号绑定和审计逻辑。

规划中的国际 Provider 包括 Microsoft、Apple、GitLab；国内 Provider 包括 Gitee、微信开放平台、支付宝、钉钉和飞书。未完成官方协议适配及测试前，不得仅添加一个按钮就宣称支持。

## QQ 网站应用配置

在 QQ 互联创建并审核网站应用后，在后台“第三方登录”中选择 QQ，填写 AppID、AppKey 和审核时登记的回调地址。默认回调路径为 `/admin/auth/oauth/qq/callback`，授权范围使用最小权限 `get_user_info`。

QQ 的 token 接口使用 GET，用户身份需要通过 `openid` 识别，再调用 `get_user_info` 获取昵称和头像。QQ 不提供可信邮箱，因此配置会强制关闭按邮箱绑定和邮箱域名限制。新增 QQ 配置默认使用普通用户角色并允许自动创建普通用户，禁止自动创建管理员；管理员接入应先建立显式 OpenID 绑定。

QQ Connect 当前不提供 PKCE 参数。服务端继续使用数据库持久化、一次性消费且短时有效的 state 防止 CSRF 和回放，回调成功后仅向浏览器返回一次性短期 ticket，不保存第三方 access token。

## GitHub OAuth App 配置

登录管理后台后打开“系统管理 → 第三方登录”，点击“添加登录方式”，填写 GitHub Client ID、Client Secret、回调地址和账户角色，保存后使用“检查必填项”确认配置完整且能够生成授权地址。该检查不访问 GitHub，不能证明 Client ID 或 Client Secret 有效；凭据有效性必须通过一次完整的 GitHub 登录回调确认。Client Secret 是只写字段：保存后编辑框显示 `********` 加末 8 位用于识别，未修改该提示直接保存会保留原密钥；列表和接口不会返回明文。

下面的 `server.yaml` 配置只作为旧版兼容、自动化部署或后台不可用时的恢复方式；日常配置无需直接编辑该文件。页面保存后会原子更新实际配置，并立即用于后续登录请求，无需重启服务。

在 GitHub OAuth App 中设置：

- Homepage URL：后台公开访问地址，例如 `https://desk.example.com/`
- Authorization callback URL：`https://desk.example.com/admin/auth/oauth/github/callback`

兼容用 `server.yaml` 示例：

```yaml
oauth:
  providers:
    - type: "github"
      name: "github"
      displayName: "GitHub"
      enabled: true
      clientId: "YOUR_GITHUB_CLIENT_ID"
      clientSecret: "YOUR_GITHUB_CLIENT_SECRET"
      redirectUrl: "https://desk.example.com/admin/auth/oauth/github/callback"
      scopes: ["read:user", "user:email"]
      accountRole: "admin" # admin 或 user
      bindByEmail: true
      autoCreateAdmin: false
      autoCreateUser: false
      allowedEmailDomains: []
      stateTtlSeconds: 180
      ticketTtlSeconds: 180
      successRedirect: "/#/login"
      failureRedirect: "/#/login"
```

默认不自动创建账号。推荐先在系统中创建同邮箱、同角色的账号，再使用 `bindByEmail: true` 完成首次绑定。若配置 `accountRole: user`，只会绑定或创建普通用户；`accountRole: admin` 只会绑定或创建管理员，不能跨角色匹配。

## 安全与生命周期

- 授权请求使用 OAuth authorization code flow、随机一次性 state 和 PKCE S256。
- state、PKCE verifier 与登录 ticket 持久化到 `oauth_login_session`，只保存浏览器值的 SHA-256，均为一次性并有短有效期，服务重启后流程仍可完成。
- GitHub 私有邮箱通过 `/user/emails` 获取，只接受 `verified=true` 的主邮箱或首个已验证邮箱；按邮箱绑定时没有已验证邮箱会拒绝登录。
- GitHub API 请求使用 Bearer token、官方 JSON media type 和明确 API 版本头；访问令牌不会写入日志或 OAuth 账号表。
- 回调成功后按 `provider + subject + role` 绑定 `oauth_account`，浏览器只能获得一次性 ticket，再换取站内 token。
- `successRedirect` 与 `failureRedirect` 只允许站内路径，拒绝完整外部 URL，避免开放重定向。
- 管理后台使用 hash 路由，普通站内路径会规范化为 `/#/...`。OAuth/OIDC 失败始终回到 `/#/login` 并显示经过白名单映射的错误，不会跳入受保护页面造成提示闪退。
- `clientSecret` 由第三方登录页面写入部署环境的实际配置文件，或由部署密钥管理系统注入，禁止提交仓库；后台接口不会回传明文。

## 验证

```bash
curl https://desk.example.com/admin/auth/oauth/providers
curl "https://desk.example.com/admin/auth/oauth/url?provider=github"
```

第二个响应中的 GitHub 授权 URL 应包含 `state`、`code_challenge` 和 `code_challenge_method=S256`。完整验收还必须分别验证：已有管理员绑定、已有普通用户绑定、私有已验证邮箱、state/ticket 重放拒绝、服务重启后的回调，以及禁用 Provider 后按钮消失。

## 其他国内 Provider 状态

微信开放平台等 Provider 尚未实现。接入前必须重新查阅其当前官方文档并确认：应用审核类型、网站应用资质、公开 HTTPS 回调域名、Client ID/AppID 与 Secret、授权/换 token/userinfo 端点、unionid/openid 身份规则和国内网络可达性。没有完成官方协议、错误处理、安全测试和真实回调验收前，只能标记为“计划”，不得仅增加按钮。

## 客户端第三方登录（/api/oauth/*）

桌面和移动客户端没有浏览器地址栏，无法像 Web 后台那样接收 OAuth 回调重定向。因此客户端第三方登录采用**服务端回调 + 客户端轮询**流程，复用同一套 `oauth.providers` 配置，仅暴露 `accountRole: user` 的 Provider。

### 接口清单

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/oauth/providers` | 列出客户端可用的已启用 Provider（仅 accountRole=user） |
| POST | `/api/oauth/start` | 生成授权 URL 和一次性 poll_token，绑定客户端设备标识 |
| GET | `/api/oauth/{provider}/callback` | Provider 授权后回调服务端，处理后返回 HTML 提示页 |
| POST | `/api/oauth/poll` | 客户端轮询拿一次性 ticket（未就绪返回 ready=false） |
| POST | `/api/oauth/exchange` | 用 ticket 换客户端 access_token（90 天，绑定 rustdesk_id+uuid） |

### 完整流程

1. 客户端调用 `POST /api/oauth/start`，请求体 `{provider, id, uuid, deviceInfo:{os,type,name}}`，服务端生成 state+PKCE+poll_token 并持久化，返回 `{enabled, url, poll_token}`。
2. 客户端用系统浏览器打开返回的 `url`，用户在 Provider 页面授权。
3. Provider 回调 `GET /api/oauth/{provider}/callback?code=xxx&state=xxx`，服务端换 token、获取用户 claims、绑定/创建账号、签发一次性 client ticket，用 poll_token 关联存储，返回 HTML 提示页“已成功登录，请回到客户端继续”。
4. 客户端每隔 2 秒调用 `POST /api/oauth/poll`，请求体 `{poll_token}`，未就绪返回 `{ready:false}`，就绪返回 `{ready:true, ticket}`。
5. 客户端拿到 ticket 后调用 `POST /api/oauth/exchange`，请求体 `{ticket}`，返回 `{access_token, type:"access_token", user:{...}}`，与 `/api/login` 成功响应格式一致。

### 安全机制

- state、poll_token、ticket 均为一次性随机值，仅持久化 SHA-256，有短有效期（默认 180 秒）。
- 客户端 token 签发与 `/api/login` 一致：90 天有效期、`is_admin=false`、绑定 `rustdesk_id`+`uuid`，新 token 签发前作废同设备旧 token。
- 仅 `accountRole: user` 的 Provider 暴露给客户端；`accountRole: admin` 的 Provider 仅供后台使用。
- 回调地址默认为 `{baseURL}/api/oauth/{provider}/callback`，也可在 Provider 配置的 `redirectUrl` 中显式指定。
- 回调成功/失败均返回中文 HTML 提示页，不暴露内部状态；失败页显示白名单错误码（`oauth_account_not_bound`/`oauth_provider_unreachable`/`oauth_state_expired`/`oauth_provider_not_for_client`/`oauth_auth_failed`）。

### Provider 配置要求

客户端 OAuth 复用后台 `oauth.providers` 配置，无需额外配置段。要将某个 Provider 同时用于客户端登录，设置 `accountRole: user` 即可。Provider 的 `redirectUrl` 若显式配置，需同时覆盖后台和客户端回调；未配置时后台用 `/admin/auth/oauth/{provider}/callback`，客户端用 `/api/oauth/{provider}/callback`，由服务端各自推导。

> 注意：同一 Provider 若要同时服务后台和客户端，`redirectUrl` 不能写死单一回调路径，应留空让服务端按请求路径推导，或注册两个 Provider 实例（一个 admin 一个 user）分别配置回调。

## RustDesk 客户端兼容协议（/api/oidc/* + /api/login-options）

RustDesk 桌面/移动客户端（Flutter 版）使用与 `/api/oauth/*` 不同的端点名和响应格式。为兼容官方客户端，服务器在 `/api/oidc/*` 和 `/api/login-options` 上提供了适配层，内部复用同一套 `OAuthProviderService` 逻辑。

### 客户端发现流程

1. 客户端调用 `GET /api/login-options`，服务器返回 `["oidc/github", "oidc/google", ...]`（仅 `accountRole=user` 的 Provider，每项以 `oidc/` 前缀）。
2. 客户端解析前缀，提取 provider 名称，渲染第三方登录按钮。

### 客户端登录流程

| 步骤 | 端点 | 说明 |
| --- | --- | --- |
| 1. 启动 | `POST /api/oidc/auth` | 请求体 `{op, id, uuid, deviceInfo, apiDomain}`，返回 `{code, url}`（`code` 即 poll_token，`url` 即授权 URL） |
| 2. 浏览器授权 | 系统浏览器打开 `url` | 用户在 Provider 页面授权 |
| 3. 回调 | `GET /api/oauth/{provider}/callback` | Provider 回调服务端（与 `/api/oauth/*` 共用），签发 ticket 关联到 poll_token |
| 4. 轮询 | `GET /api/oidc/auth-query?code=&id=&uuid=` | 客户端每秒轮询，未就绪返回 `{body:"{\"error\":\"No authed oidc is found\"}"}`，就绪返回 `{body:"{\"access_token\":\"...\",\"type\":\"access_token\",\"user\":{...}}"}` |

### 与 /api/oauth/* 的对应关系

| 客户端协议 | 内部复用 | 说明 |
| --- | --- | --- |
| `GET /api/login-options` | `ListClientProviders()` | 加 `oidc/` 前缀返回 |
| `POST /api/oidc/auth` | `BuildClientAuthURL()` | `op` → providerName，返回 `code` → pollToken |
| `GET /api/oidc/auth-query` | `PollClientTicket()` + `ExchangeClientTicket()` | `code` → pollToken，一步完成轮询+换 token |
