# 新对话续接提示词

复制下面整个代码块到新对话，继续当前项目任务：

```text
继续维护项目 E:\Visual_Studio_Code\07 Rustdesk-api-server-pro。先读取仓库根 README、docs/README.md、docs/CURRENT_STATUS.md、docs/PROJECT_WORKFLOW.md、docs/OAUTH_PROVIDERS.md、docs/TROUBLESHOOTING.md 和本文档，再检查 git/线上/设备实时状态，不要仅依赖本提示词中的快照。

必须遵守：
1. 所有开发、提交、推送都在 main；本地也保持 main。
2. 远端只保留 main 和必须保留的 backup。每次发布完成后将 backup 对齐 main，禁止删除 backup。
3. 不覆盖用户无关改动，不提交数据库、日志、token、OAuth Secret、生产 server.yaml 或真实密码。
4. 每次构建由 CI 自动递增 PATCH 版本。推送后持续监控全部质量流程、单次 Docker 流程、GHCR 产物和设备更新，不要在目标未完成时提前结束。
5. 前端新增 i18n key 必须补齐 app.d.ts 和 9 种语言，并运行严格翻译检查；中文不得回退英文。
6. 后端通讯/OAuth 改动必须有测试，检查错误处理、权限边界、敏感信息、回调重放、网络超时和重启持久化。
7. 错误消息必须带 ERR-xxxx 编码前缀，与帮助页错误码索引匹配；新增 errcode 时同步更新 errcode.go、app.d.ts 和全部语言文件。
8. Apple 私钥 placeholder 不得包含 `-----BEGIN PRIVATE KEY-----` 文本，避免触发合规检查密钥检测。
9. 新增数据模型必须注册到 cmd/sync.go 的 models 列表，否则 sync 命令不会自动建表。
10. 鉴权中间件返回 401/406 时必须使用 StopWithJSON 返回带 ERR-xxxx 编码的 JSON，不能返回纯文本。

当前快照（2026-08-08，使用前实时复核）：
- VERSION：1.2.29
- GHCR：ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest
- 测试设备：ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net
- root 账户当前不接受已有 SSH 密钥，使用 LiYan；该账户属于 docker/Administrators 组。
- 容器：rustdesk-api-server-pro，host 网络，端口 16888
- 更新脚本：/vol1/docker/rustdesk-api-server-pro/update-rustdesk-api.sh

已完成的 OAuth 状态：
- 管理员在"第三方登录"页面直接配置 Provider，server.yaml 仅作兼容/恢复。
- 已完成 Provider：GitHub、QQ、Google、Microsoft、Gitee、GitLab、WeChat、Apple。
- Apple 使用动态 JWT client_secret（ES256 签名），需配置 Team ID、Key ID 和 .p8 私钥；Apple 不支持 PKCE，用户信息仅从 ID Token 获取。
- WeChat 使用 appid 参数、逗号分隔 scope、GET token 请求，不支持 PKCE。
- GitHub Provider 支持 PKCE S256、数据库持久化一次性 state/ticket、已验证邮箱、admin/user 角色隔离、按邮箱绑定和自动创建开关。
- Client Secret 不回传明文；配置接口返回 secretConfigured 和 secretHint。
- OAuth/OIDC 成功/失败目标统一规范化为 /#/ hash 路由。失败固定回 /#/login。
- 错误码索引体系：errcode.go 注册 105+ 个错误码，Message 统一为 PascalCase；前端 parseBackendMessage() 从 ERR-xxxx: Message 提取编码和翻译；帮助页面提供错误码搜索和筛选；错误日志表 ErrorLog 全局记录后端错误。
- RustDesk 客户端兼容协议：/api/login-options 返回 oidc/ 前缀列表，/api/oidc/auth 返回 {code,url}，/api/oidc/auth-query 返回 {body:"json"}。
- 鉴权中间件 401/406 响应已改为 JSON 格式带 ERR-1010 编码。

当前外部阻碍和真实设备证据：
- GitHub Provider 已启用，Client ID/Secret 均已保存，公开 Provider API 正常返回。
- NAS 访问 github.com:443 偶尔超时；api.github.com:443 返回 HTTP 200。
- 不要擅自开启 autoCreateAdmin，这会扩大管理权限。

已完成的功能清单：
- 8 种 OAuth Provider 协议适配（GitHub/QQ/Google/Microsoft/Gitee/GitLab/WeChat/Apple）
- 错误码索引体系（105+ ERR-xxxx 编码 + 前端帮助页面 + 错误日志管理）
- 前端移动端响应式优化（Modal/Drawer/表格/登录页自适应）
- Token 清除按钮和错误日志管理页面
- 登录对话框压缩（320px 宽度、medium 表单、small OAuth 按钮 2 列）
- OAuth 按钮文案缩短（仅显示 Provider 名）
- fr/de/es i18n 翻译补充至 85%+ 阈值
- ErrorLog 模型注册到 sync 命令自动建表
- 鉴权中间件 JSON 化 401/406 响应
- AdminOrUserAuth 中间件：普通用户可查看审计/系统设置（只读），操作类 API 通过 isAdmin 检查限制
- 客户端 OAuth 统一回调：客户端和 admin 共用回调端点，通过 PollToken 区分，客户端回调返回 HTML 页面包含 rustdesk:// URL scheme

开始新对话后的执行顺序：
1. git status、分支、远端分支、VERSION、最近提交和 Actions/GHCR 实时核验。
2. SSH 检查设备版本、容器状态、Provider API 和 github.com/api.github.com 网络；网络不稳定时按间隔重试，不因单次失败下结论。
3. 按用户指示继续开发或维护任务。
4. 本地验证、提交 main、监控 CI、确认版本自增和 GHCR、对齐 backup、更新设备、重启复验。
```
