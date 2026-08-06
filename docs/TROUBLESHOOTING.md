# 问题排查手册（常见问题 / 日志定位 / 处理建议）

本文档面向当前“兼容增强版”代码状态，重点覆盖上线与使用过程中的常见问题。

## 1. 排查流程总览（建议顺序）

出现问题时，先按以下顺序排查：

1. 看服务端是否启动成功（端口是否监听）
2. 看配置文件是否生效（`server.yaml`）
3. 看数据库结构是否已同步（`sync`）
4. 看客户端请求路径是否命中兼容端点（日志）
5. 看是否属于“兼容占位实现”的预期限制（OIDC / plugin-sign / 分组权限模型）

## 2. 服务启动失败 / 无法访问页面

### 现象

- 程序启动后立即退出
- 浏览器访问 `http://IP:端口/` 无响应
- 客户端请求 API 超时

### 排查

1. 检查后端启动日志是否有报错（数据库连接、端口占用、配置解析失败）
2. 检查 `backend/server.yaml` 的 `httpConfig.port`
3. 检查端口是否被其他程序占用
4. 检查防火墙/安全组是否放行
5. 若 `/api` 可用但首页不可用，重点检查 `httpConfig.staticdir`

### 常见原因

- `server.yaml` 未放在正确目录
- `httpConfig.port` 填写错误
- `staticdir` 指向不存在目录（导致后台页面 404）
- Docker 端口映射未配置或映射错

### v1.1.16 特有问题：signKey 启动失败

**现象**：容器启动后立即退出，日志显示 `unsafe signKey: set a unique random signKey with at least 32 characters`

**原因**：`server.yaml` 中 signKey 行尾注释导致启动脚本解析失败，无法自动替换默认值。

**处理**：删除容器内 `/app/data/server.yaml`，重新创建容器让新镜像配置生效。

### v1.1.16 特有问题：更新后 Web 页面 404

**现象**：API 接口正常但访问管理后台返回 404

**原因**：`PORT` 环境变量 sed 替换误改了 `smtpConfig.port`（整数→字符串），导致 YAML 解析失败，程序用默认值覆盖配置，默认 `staticdir` 为相对路径 `dist`，在 `/app/data` 下找不到前端文件。

**处理**：删除 `/app/data/server.yaml`，重新创建容器。v1.1.16 已修复 sed 仅作用于 `httpConfig` 段，且解析失败不再覆盖文件。

### v1.1.16 特有问题：容器停止报错

**原因**：旧版 `start.sh` 中 Go 进程不是 PID 1，Shell 不转发 SIGTERM，Docker 超时后 SIGKILL 强杀。

**处理**：v1.1.16 已修复：`exec` 替换 Shell 为 PID 1，Go 进程直接接收信号优雅关闭。

## 2.1 第三方登录不显示、提示闪退或 GitHub 登录失败

### Provider 不显示

检查 `/admin/auth/oauth/providers`。只有 `enabled=true` 且 Client ID/Secret 完整的 Provider 才会返回。后台编辑 Secret 时显示 `********` 加末 8 位属于正常安全提示，不代表配置为空。

### 点击后提示闪一下消失

旧版本在 OAuth 失败时会跳向成功目标页，受保护页面随后再次重定向到登录页，导致错误参数进入嵌套 `redirect`。从 1.1.54 起，OAuth/OIDC 失败固定回到 `/#/login?oauth_error=...`；请升级并强制刷新浏览器缓存。

### `oauth_provider_unreachable`

表示服务端无法访问 Provider，不代表 Secret 未保存。应从容器宿主机分别测试：

```bash
curl -I --connect-timeout 8 https://github.com/
curl -I --connect-timeout 8 https://api.github.com/
```

GitHub OAuth 换 token 必须访问 `https://github.com/login/oauth/access_token`。仅 `api.github.com` 可用仍不能完成登录，需要检查出站防火墙、DNS、透明代理或路由。

### `oauth_account_not_bound`

若启用按邮箱绑定且关闭自动创建，现有同角色账户必须填写与 Provider 已验证邮箱完全一致的邮箱。管理员 Provider 不能绑定普通用户，普通用户 Provider 也不能绑定管理员。

### 国内 Provider 测试

QQ Provider 已实现网站应用 OAuth2 流程。必须使用已审核的 QQ 互联网站应用，并确保配置的回调地址与审核登记值一致；QQ 不提供邮箱，首次测试建议使用普通用户角色和自动创建普通用户。WeChat Provider 已完成协议适配，使用 appid 参数和逗号分隔 scope，不支持 PKCE。Apple Provider 已完成协议适配，使用动态 JWT client_secret，需配置 Team ID、Key ID 和 .p8 私钥。

### 错误码 ERR-xxxx

所有后端错误消息均带 `ERR-xxxx` 编码前缀。前端显示格式为 `[ERR-xxxx] 翻译后消息`。可在"关于 → 错误码帮助"页面搜索编码查看详细说明和解决方案。错误日志表 `ErrorLog` 全局记录后端错误，可在"审计 → 错误日志"页面查看。

## 3. 登录成功但部分页面/接口报错（尤其升级后）

### 现象

- 登录后地址簿打不开
- 设备列表报错
- 审计页异常
- 日志里出现 SQL 字段不存在

### 高概率原因

数据库结构未同步（尤其是新增兼容字段后）。

### 处理

执行一次数据库同步：

```powershell
rustdesk-api-server-pro.exe sync
```

然后重启服务。

说明：

- 本仓库近期兼容补丁包含地址簿字段扩展（如 `peer.note`、`peer.same_server`），未同步会导致运行时报错或字段不生效

## 4. 新版客户端提示接口 404 / 功能打不开

### 现象

- 客户端某个页面打开时报错
- 服务端日志里有 `/api/...` 的 404 记录

### 排查步骤

1. 打开请求日志（`httpConfig.printRequestLog: true`）
2. 观察 404 的具体路径、方法、请求体
3. 判断是否属于：
   - 已支持端点但配置/鉴权问题
   - 兼容占位端点（预期有限功能）
   - 尚未补齐的新端点

### 当前已补齐的常见兼容端点（用于避免 404）

- `/api/oidc/auth`
- `/api/oidc/auth-query`
- `/api/devices/cli`
- `/api/devices/deploy`
- `/api/record`
- `/api/device-group/accessible`
- `/lic/web/api/plugin-sign`

如果仍然是 404，通常说明：

- 代码未更新到当前兼容版本
- 启动的是旧二进制
- 路由未生效（部署目录不一致）

## 5. 客户端反复同步策略或心跳后状态异常

### 现象

- 客户端心跳成功，但持续触发策略同步
- 日志中 `/api/heartbeat` 正常返回，客户端仍认为策略版本变化

### 排查

1. 确认服务端版本已包含心跳 `modified_at` 回显补丁
2. 查看 `/api/heartbeat` 请求体中的 `modified_at`
3. 确认响应 JSON 中的 `modified_at` 与请求值一致
4. 如果未来启用了真实策略分发，再按策略版本更新时间重新核验该行为

## 6. 地址簿字段或共享地址簿不保存 / 不回显

### 现象

- 客户端修改地址簿备注或 `same_server` 后刷新消失
- 共享地址簿能看到 profile，但 peer/tag 列表为空
- 共享地址簿写入失败
- 新版客户端同步字段不完整

### 排查

1. 确认服务端版本已包含地址簿 `note`、`same_server` 和共享地址簿兼容补丁
2. 执行 `sync`
3. 重启服务
4. 查看日志中 `/api/ab/peers`、`/api/ab/tags/*`、`/api/ab/peer/update/*` 请求是否成功
5. 如果是共享地址簿写入，确认当前地址簿 `shared=true` 且 `rule >= 2`；只读共享地址簿只能读取，不能写入

### 说明

当前兼容版本已支持：

- 地址簿 `note` 字段读写
- 新版客户端增量更新字段（`username` / `hostname` / `platform` / `note` 等）
- 新增 peer 时 `same_server` 支持布尔值、`null` 或缺省
- 共享地址簿读取使用 owner 数据；共享写入按 `rule >= 2` 控制

## 7. 分组面板能打开，但看不到完整分组/权限效果

### 现象

- 分组页面不再报错，但内容为空或不完整
- `accessible` 相关数据与官方 Pro 不一致

### 原因

这是当前版本的已知边界：相关接口是“兼容模型”，不是官方 Pro 完整权限实现。

涉及接口：

- `/api/device-group/accessible`
- `/api/users?accessible=...`
- `/api/peers?accessible=...`

### 结论

- 如果你的目标是“客户端页面可正常打开，不报错”：当前行为是可接受的
- 如果你的目标是“完整企业级分组权限”：需要继续开发

## 8. 客户端第三方登录（/api/oauth/*）

### 现象

- 客户端希望通过 GitHub/QQ 等 Provider 登录，但不确定服务端是否支持

### 说明

客户端第三方登录已通过 `/api/oauth/*` 接口族实现，采用**服务端回调 + 客户端轮询**流程，复用后台 `oauth.providers` 中 `accountRole: user` 的 Provider。完整流程见 `docs/OAUTH_PROVIDERS.md`“客户端第三方登录”章节。

### 排查

- 客户端调用 `GET /api/oauth/providers` 无返回：检查后台是否配置了 `accountRole: user` 且 `enabled: true` 的 Provider。
- `POST /api/oauth/start` 返回 `enabled:false`：Provider 未启用或 `accountRole` 不是 `user`。
- 回调页显示 `oauth_state_expired`：state 超时（默认 180 秒），用户授权过慢，调大 `stateTtlSeconds`。
- 回调页显示 `oauth_account_not_bound`：未配置 `bindByEmail` 或 `autoCreateUser`，且无已绑定账号。
- 轮询一直 `ready:false`：回调未完成或失败，检查浏览器是否成功跳转到 Provider 授权页。
- `/api/oidc/auth` 与 `/api/oidc/auth-query` 已实现完整客户端兼容协议，复用 `/api/oauth/*` 的服务端回调+轮询逻辑。

## 8.1 错误日志页面报 ERR-B010: DatabaseError: no such table: error_log

### 原因

`ErrorLog` 模型未注册到 `cmd/sync.go` 的自动建表列表，导致 `sync` 命令不会创建 `error_log` 表。

### 处理

1. 确认服务端版本已包含 `ErrorLog` 注册到 `sync.go`（1.2.13+ 已修复）。
2. 在容器内执行 `rustdesk-api-server-pro sync`。
3. 重启容器。

## 8.2 Token 管理页面操作返回 401 且无 ERR 编码

### 原因

旧版鉴权中间件使用 `StopWithText` 返回纯文本 `Unauthorized`，前端无法解析为 JSON 格式的错误消息。

### 处理

确认服务端版本已包含中间件 JSON 化修复（1.2.13+ 已修复）。401 响应现在返回 `{"code":401,"message":"ERR-1010: Unauthorized"}` 格式。

## 9. 插件签名（plugin-sign）不可用或签名结果不符合预期

### 现象

- 客户端请求 `/lic/web/api/plugin-sign` 成功，但签名流程不通过

### 原因

当前为兼容占位实现，返回结构对齐，但不提供官方真实签名能力。

### 处理建议

- 不使用插件签名：可忽略
- 需要插件签名：实现真实签名服务逻辑（私钥管理、签名算法、校验链路）

## 10. 录制上传（record）报错或没有文件

### 现象

- 客户端录制上传时报错
- 服务端没有录制文件

### 排查

1. 检查服务端是否为包含 `/api/record` 最小落盘实现的版本
2. 检查运行目录下 `record_uploads/` 是否存在、是否可写
3. 检查磁盘空间
4. 查看日志中 `type=new/part/tail/remove` 请求是否正常

### 说明

当前实现是“最小兼容落盘”，用于兼容客户端上传流程：

- 支持 `new / part / tail / remove`
- 未实现完整归档、索引、生命周期管理

## 11. 前端构建失败（开发时）

### 排查命令

```powershell
pnpm typecheck
pnpm build
```

### 当前已知提示（通常不阻塞构建）

- `vite` 提示某些路由文件名含大写（建议项）
- `unocss` 提示部分 icon 加载失败（项目现有提示，构建通常仍成功）

### 若构建仍失败

- 先看是否是本地改动引入的类型错误
- 再看 `src/locales`、页面字段名、接口类型映射是否被误改

## 12. 日志建议（上线后）

建议在排查阶段开启：

```yaml
httpConfig:
  printRequestLog: true
```

重点关注：

- 404 请求路径（判断缺接口还是配置问题）
- 401/鉴权失败（token/session 问题）
- SQL 报错（字段缺失、约束冲突）
- 文件写入报错（`record_uploads` 权限）

定位完成后可按需关闭请求日志，避免日志过大。

## 13. 何时提交缺陷或开展二次开发

建议按下面标准判断：

- “主流程报错、字段丢失、明显 404”：应优先修复（兼容问题）
- “高级功能不完整但已在 README 写明”：属于已知限制（按需开发）
- “部署/配置导致的问题”：优先修复配置和环境，不一定是代码缺陷

如果你要进一步对齐官方 Pro，请优先排期：

1. OIDC 完整流程
2. plugin-sign 真实签名服务
3. 完整分组权限模型
