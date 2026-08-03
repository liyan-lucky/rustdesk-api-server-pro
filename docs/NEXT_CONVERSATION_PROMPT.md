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

当前快照（2026-08-03，使用前实时复核）：
- main/backup/本地：f2633f2669d7a8b94284e5163a173df3348166ae
- VERSION：1.1.54
- GHCR：ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest
- 最近 Docker 发布流程成功：https://github.com/liyan-lucky/rustdesk-api-server-pro/actions/runs/30786562895
- 测试设备：ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net
- root 账户当前不接受已有 SSH 密钥，使用 LiYan；该账户属于 docker/Administrators 组。
- 容器：rustdesk-api-server-pro，host 网络，端口 16888，运行版本已验证为 1.1.54。
- 更新脚本：/vol1/docker/rustdesk-api-server-pro/update-rustdesk-api.sh

已完成的 OAuth 状态：
- 管理员在“第三方登录”页面直接配置 Provider，server.yaml 仅作兼容/恢复。
- GitHub Provider 支持 PKCE S256、数据库持久化一次性 state/ticket、已验证邮箱、admin/user 角色隔离、按邮箱绑定和自动创建开关。
- Client Secret 不回传明文；配置接口返回 secretConfigured 和 secretHint。secretHint 格式为 ******** 加末 8 位，前端原样保存提示时保留旧密钥。
- “检查必填项”只检查启用、Client ID、Secret、回调地址及授权 URL 生成，不宣称凭据有效；真实有效性必须完成 Provider 回调。
- OAuth/OIDC 成功/失败目标统一规范化为 /#/ hash 路由。失败固定回 /#/login，只有成功进入原目标页。
- 安全错误码：oauth_account_not_bound、oauth_provider_unreachable、oauth_state_expired、oauth_auth_failed；9 种语言均有提示。
- 启动脚本已修复后台保存 YAML 后 PORT sed 破坏四空格缩进的问题。
- Docker 工作流已修复同一 SHA 被多个 workflow_run 重复触发并互相取消的问题，只由“发布前检查”启动一次，内部仍等待全部检查。

当前外部阻碍和真实设备证据：
- GitHub Provider 已启用，Client ID/Secret 均已保存，公开 Provider API 正常返回 GitHub。
- NAS 连续多次访问 github.com:443 超时；api.github.com:443 返回 HTTP 200。
- GitHub token 交换地址是 https://github.com/login/oauth/access_token，因此该网络问题会产生 oauth_provider_unreachable。
- 历史 security_audit 还记录过 no bindable oauth account：当前有效管理员没有填写邮箱，同时配置为 bindByEmail=true、autoCreateAdmin=false。网络恢复后必须给目标管理员填写与 GitHub verified email 一致的邮箱，或先取得用户明确授权再开启自动创建管理员。
- 不要擅自开启 autoCreateAdmin，这会扩大管理权限。

用户最新方向：
- 怀疑国内到 GitHub 的代理/网络不稳定，希望考虑使用微信或 QQ 测试。
- 微信、QQ 当前尚未实现，不能只加按钮宣称支持。
- 下一步先联网查阅微信开放平台、QQ 互联的最新官方文档，比较网站应用资质、审核、公开 HTTPS 回调、授权/token/userinfo 端点、openid/unionid 身份模型和测试条件。
- 在没有真实 AppID/Secret、已审核应用或回调域名时，不得伪造“测试成功”。可以先完成通用 Provider 架构设计、配置模型、接口适配器和离线单元测试；需要真实凭据或外部审核时明确报告阻碍。
- 优先选择一个能在现有网页后台场景合法落地且可真实验收的国内 Provider；若微信/QQ均需要用户先完成平台审核，列出精确准备清单并继续完成不依赖凭据的代码与测试。

开始新对话后的执行顺序：
1. git status、分支、远端分支、VERSION、最近提交和 Actions/GHCR 实时核验。
2. SSH 检查设备版本、容器状态、Provider API和 github.com/api.github.com 网络；网络不稳定时按间隔重试，不因单次失败下结论。
3. 阅读官方微信/QQ文档，给出可实施比较和选择；只引用官方来源。
4. 继续完成选定 Provider 的 P0-P2 实现、权限/错误/i18n/导入导出等相关一致性检查。
5. 本地验证、提交 main、监控 CI、确认版本自增和 GHCR、对齐 backup、更新设备、重启复验。
```
