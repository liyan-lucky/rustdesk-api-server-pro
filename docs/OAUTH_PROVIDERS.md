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
