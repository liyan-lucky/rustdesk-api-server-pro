# RustDesk 官方 API 兼容完善计划

本文档用于指导 `rustdesk-api-server-pro` 后续补齐 RustDesk 官方客户端常用 API 能力，并避免接口修改时破坏现有自托管部署。

> 目标：优先保证客户端主流程稳定可用，再逐步补齐企业管理能力。本文不承诺复刻官方商业服务全部能力，而是把兼容边界、实现优先级和验收方式写清楚。

## 1. 当前项目基础

当前项目采用单端口一体化架构：

- `/api/*`：RustDesk 客户端 API。
- `/admin/*`：管理后台 API。
- `/lic/web/api/*`：license/plugin-sign 兼容入口。
- `/`：管理后台静态前端。

后端路由集中在 `backend/app/route.go` 注册。客户端侧已包含：

- `SystemController`
- `LoginController`
- `AuditController`
- `CompatPublicController`
- `UserController`
- `PeerController`
- `CompatAuthController`
- `DeviceGroupController`
- `EnterpriseCompatController`
- `AddressBookController`
- `AddressBookPeerController`
- `AddressBookTagController`

这说明项目已经具备 API 兼容扩展基础，后续重点应放在接口覆盖矩阵、响应结构兼容、审计字段完整性和后台可视化。

## 2. 兼容原则

### 2.1 不破坏旧客户端

新增字段必须向后兼容：

- 老客户端不识别字段时应能忽略。
- 缺省字段必须有安全默认值。
- 不随意修改已有 JSON 字段名。
- 不随意修改状态码和通用响应结构。

### 2.2 兼容优先级

优先级从高到低：

1. 登录、退出、token 刷新、token 校验。
2. 设备上线、心跳、sysinfo、设备列表。
3. 地址簿、标签、备注、共享地址簿。
4. 连接审计、文件传输审计、报警审计。
5. 用户组、设备组、策略分发。
6. OIDC、OAuth、license、plugin-sign 兼容。
7. 高级企业能力：策略模板、权限继承、细粒度审计、导出归档。

### 2.3 兼容占位必须可观察

对于暂未完整实现的官方接口，不建议只返回空成功。应做到：

- 返回客户端能接受的最小结构。
- 后端记录兼容占位日志。
- 在调试模式下可标记 `compat_stub: true`。
- 管理后台可查看命中的兼容占位接口次数。

## 3. API 覆盖矩阵

建议新增 `docs/API_MATRIX.md` 或在本文维护表格。

| 模块 | 路径示例 | 当前状态 | 建议动作 | 优先级 |
| --- | --- | --- | --- | --- |
| 登录 | `/api/login` | 已有 | 对齐新版字段、失败原因、2FA 流程 | P0 |
| 登出 | `/api/logout` | 已有/待核验 | token 失效写入审计 | P0 |
| 用户信息 | `/api/currentUser` | 待核验 | 返回套餐、权限、组信息兼容字段 | P0 |
| 心跳 | `/api/heartbeat` | 已有/待核验 | 记录在线状态和客户端版本 | P0 |
| sysinfo | `/api/sysinfo` | 已有/待核验 | 扩展 OS、版本、主机名、IP、MAC | P1 |
| 设备列表 | `/api/devices` | 已有/待核验 | 支持分页、搜索、在线过滤 | P1 |
| 地址簿 | `/api/ab`、`/api/peers` | 已有 | 对齐 note、tags、hash、更新时间 | P0 |
| 连接审计 | `/api/audit/conn` | 已有基础 | 增加完整字段、状态、方向、会话时长 | P0 |
| 文件审计 | `/api/audit/file` | 已有基础 | 增加大小、结果、方向、错误信息 | P0 |
| 报警审计 | `/api/audit/alarm` | 兼容占位 | 落库并后台展示 | P1 |
| 录屏 | `/api/record/*` | 已有基础/待核验 | 完整 new/part/tail/remove 状态机 | P1 |
| 设备组 | `/api/device-group/*` | 已有/待核验 | 补齐分组权限和策略关系 | P2 |
| 用户组 | `/api/user-group/*` | 已有模型/待核验 | 补齐成员、权限、策略绑定 | P2 |
| 策略 | `/api/strategy/*` | 兼容基础 | 支持策略下发、默认策略、冲突处理 | P2 |
| OIDC | `/api/oidc/*` | 部分兼容 | 区分后台 OAuth 与客户端 OIDC | P2 |
| plugin-sign | `/lic/web/api/plugin-sign` | 兼容入口 | 配置化签名策略、审计命中 | P2 |

## 4. 后端实现建议

### 4.1 新增兼容层包

建议保持现有 Iris MVC 控制器，但新增内部兼容层：

```text
backend/internal/compat/
├─ official/
│  ├─ request.go
│  ├─ response.go
│  ├─ matrix.go
│  └─ version.go
├─ audit/
│  ├─ normalizer.go
│  └─ mapper.go
└─ stub/
   ├─ recorder.go
   └─ response.go
```

职责：

- `official/request.go`：统一解析官方客户端请求。
- `official/response.go`：统一生成兼容响应。
- `official/version.go`：根据客户端版本返回不同字段。
- `audit/normalizer.go`：把不同客户端版本审计字段归一化。
- `stub/recorder.go`：记录占位接口命中情况。

### 4.2 控制器保持薄层

控制器只做：

1. 读取请求。
2. 调用 service。
3. 返回兼容响应。

不要在控制器中堆积业务逻辑，避免以后官方接口变动时难维护。

### 4.3 统一错误响应

建议定义错误码：

| 错误码 | 含义 |
| --- | --- |
| `0` | 成功 |
| `1001` | 未登录或 token 无效 |
| `1002` | 权限不足 |
| `1003` | 参数缺失 |
| `1004` | 参数格式错误 |
| `1100` | 客户端版本不兼容 |
| `2001` | 设备不存在 |
| `2002` | 用户不存在 |
| `3001` | 审计记录不存在 |
| `9000` | 服务端内部错误 |

## 5. 数据库兼容建议

### 5.1 保留 Xorm 自动同步，但增加迁移记录

当前项目使用 Xorm 自动同步模型，简单方便。建议增加 `schema_migrations` 表记录每次结构升级：

```go
type SchemaMigration struct {
    Id        int       `xorm:"'id' pk autoincr"`
    Version   string    `xorm:"'version' varchar(64) unique"`
    Name      string    `xorm:"'name' varchar(255)"`
    AppliedAt time.Time `xorm:"'applied_at' created"`
}
```

这样可以避免后续字段越来越多时无法判断升级状态。

### 5.2 索引建议

建议为以下字段添加索引：

- `audit.conn_id`
- `audit.session_id`
- `audit.rustdesk_id`
- `audit.peer_id` 或归一化后的对端字段
- `audit.created_at`
- `file_transfer.rustdesk_id`
- `file_transfer.peer_id`
- `file_transfer.created_at`
- `device.rustdesk_id`
- `device.uuid`
- `auth_token.token`

SQLite 和 MySQL 都会受益，尤其是后台日志分页查询。

## 6. 验收清单

每次补齐一个官方接口，都要至少验证：

- 未登录访问是否正确拒绝。
- 已登录访问是否返回客户端可接受结构。
- 老客户端是否不报错。
- 新客户端是否不出现 404。
- 返回字段缺省值是否稳定。
- 管理后台是否能看到相关记录。
- Docker 容器重启后数据是否保留。
- SQLite 和 MySQL 是否都能同步表结构。

## 7. 推荐实施顺序

### 第一阶段：接口矩阵与审计落库

- 建立 API 覆盖矩阵。
- 把 `/api/audit/alarm` 从占位改成落库。
- 扩展连接审计字段。
- 扩展文件传输审计字段。
- 后台增加高级筛选。

### 第二阶段：客户端主流程兼容

- 登录响应对齐。
- currentUser/user info 对齐。
- heartbeat/sysinfo 对齐。
- devices/cli/devices 列表对齐。

### 第三阶段：企业能力兼容

- 用户组。
- 设备组。
- 策略。
- 权限绑定。
- 操作审计。

### 第四阶段：高级能力

- 客户端 OIDC。
- plugin-sign 配置化。
- 录屏完整状态机。
- 审计导出、归档、清理策略。

## 8. 文档要求

每新增或修改一个接口，必须同步更新：

- API 路径。
- 请求字段。
- 响应字段。
- 是否需要 token。
- 是否写入审计。
- 是否为完整实现或兼容占位。
- 客户端版本验证结果。

建议文档统一放在：

```text
docs/
├─ OFFICIAL_API_COMPATIBILITY_PLAN.md
├─ AUDIT_LOG_DESIGN.md
├─ API_MATRIX.md
└─ RELEASE_CHECKLIST.md
```
