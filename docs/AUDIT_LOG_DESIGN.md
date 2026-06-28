# 审计日志增强设计

本文档用于指导 `rustdesk-api-server-pro` 从“基础连接/文件日志”升级为“完整审计中心”。

## 1. 目标

审计系统需要覆盖：

1. 客户端连接审计。
2. 文件传输审计。
3. 报警审计。
4. 后台管理员操作审计。
5. 登录与安全事件审计。
6. 兼容占位接口命中审计。
7. 数据导出、清理、归档。

最终后台应能回答这些问题：

- 谁连接了谁？
- 什么时间连接？
- 连接多久？
- 使用了哪个 RustDesk ID、UUID、IP、客户端版本？
- 是否传输了文件？文件方向、路径、大小、结果是什么？
- 是否触发报警？报警类型和详情是什么？
- 哪个管理员修改了用户、设备、策略、地址簿？
- 哪些官方兼容接口被客户端调用但当前还是占位实现？

## 2. 当前审计基础

当前已有能力：

- `/api/audit/conn`：连接审计基础流程。
- `/api/audit/file`：文件传输审计基础流程。
- `/api/audit/alarm`：当前为兼容占位，接受请求但不落库。
- `audit` 表：保存连接相关字段。
- `file_transfer` 表：保存文件传输相关字段。
- 管理后台已有审计相关页面。

当前不足：

- `audit` 表字段偏少。
- `alarm` 未落库。
- 管理后台操作缺少统一审计。
- 登录失败、token 失效、OAuth 绑定等安全事件未统一记录。
- 兼容占位接口没有统计。
- 缺少保留周期、导出、清理策略。

## 3. 审计分类

建议统一审计类型：

| 类型 | code | 说明 |
| --- | --- | --- |
| 连接审计 | `connection` | 远程连接开始、更新、结束 |
| 文件审计 | `file_transfer` | 文件上传、下载、目录传输 |
| 报警审计 | `alarm` | 客户端上报安全或异常事件 |
| 登录审计 | `login` | 登录成功、失败、退出、token 刷新 |
| 管理操作 | `admin_operation` | 后台用户、设备、策略、配置修改 |
| API 兼容 | `compat_api` | 官方兼容接口调用、占位接口命中 |
| 系统事件 | `system` | 启动、同步、定时任务、配置变更 |

## 4. 连接审计字段设计

建议扩展 `audit` 表或新增 `connection_audit` 表。

如果保持兼容，推荐先扩展现有 `audit` 表：

```go
type Audit struct {
    Id              int       `xorm:"'id' int notnull pk autoincr"`
    ConnId          int       `xorm:"'conn_id' int index"`
    SessionId       string    `xorm:"'session_id' varchar(100) index"`
    RustdeskId      string    `xorm:"'rustdesk_id' varchar(100) index"`
    PeerId          string    `xorm:"'peer_id' varchar(100) index"`
    Uuid            string    `xorm:"'uuid' varchar(255) index"`
    IP              string    `xorm:"'ip' varchar(64)"`
    PeerIP          string    `xorm:"'peer_ip' varchar(64)"`
    Type            int       `xorm:"'type' tinyint"`
    Direction       string    `xorm:"'direction' varchar(32)"`
    Status          string    `xorm:"'status' varchar(32) index"`
    ClientVersion   string    `xorm:"'client_version' varchar(64)"`
    Platform        string    `xorm:"'platform' varchar(64)"`
    Hostname        string    `xorm:"'hostname' varchar(255)"`
    Note            string    `xorm:"'note' varchar(1024)"`
    Raw             string    `xorm:"'raw' text"`
    StartedAt       time.Time `xorm:"'started_at' datetime index"`
    ClosedAt        time.Time `xorm:"'closed_at' datetime index"`
    DurationSeconds int64     `xorm:"'duration_seconds' bigint"`
    CreatedAt       time.Time `xorm:"'created_at' datetime created index"`
    UpdatedAt       time.Time `xorm:"'updated_at' datetime updated"`
}
```

字段说明：

- `ConnId`：客户端连接编号。
- `SessionId`：会话 ID，也可兼容旧字段 `guid`。
- `RustdeskId`：发起方或当前设备 RustDesk ID。
- `PeerId`：对端 RustDesk ID。
- `Direction`：`incoming`、`outgoing`、`unknown`。
- `Status`：`open`、`closed`、`failed`、`unknown`。
- `Raw`：原始请求 JSON，方便兼容新客户端字段。
- `DurationSeconds`：连接关闭时计算。

## 5. 文件传输审计字段设计

建议扩展 `file_transfer` 表：

```go
type FileTransfer struct {
    Id              int       `xorm:"'id' int notnull pk autoincr"`
    RustdeskId      string    `xorm:"'rustdesk_id' varchar(100) index"`
    PeerId          string    `xorm:"'peer_id' varchar(100) index"`
    SessionId       string    `xorm:"'session_id' varchar(100) index"`
    Uuid            string    `xorm:"'uuid' varchar(255) index"`
    Direction       string    `xorm:"'direction' varchar(32) index"`
    Path            string    `xorm:"'path' text"`
    FileName        string    `xorm:"'file_name' varchar(512)"`
    IsFile          bool      `xorm:"'is_file' bool"`
    SizeBytes       int64     `xorm:"'size_bytes' bigint"`
    Result          string    `xorm:"'result' varchar(32) index"`
    ErrorMessage    string    `xorm:"'error_message' varchar(1024)"`
    Type            int       `xorm:"'type' tinyint"`
    Info            string    `xorm:"'info' text"`
    Raw             string    `xorm:"'raw' text"`
    CreatedAt       time.Time `xorm:"'created_at' datetime created index"`
    UpdatedAt       time.Time `xorm:"'updated_at' datetime updated"`
}
```

建议 `Result` 可选值：

- `success`
- `failed`
- `canceled`
- `unknown`

## 6. 报警审计字段设计

新增 `alarm_audit` 表：

```go
type AlarmAudit struct {
    Id            int       `xorm:"'id' int notnull pk autoincr"`
    RustdeskId    string    `xorm:"'rustdesk_id' varchar(100) index"`
    PeerId        string    `xorm:"'peer_id' varchar(100) index"`
    SessionId     string    `xorm:"'session_id' varchar(100) index"`
    AlarmType     string    `xorm:"'alarm_type' varchar(100) index"`
    Severity      string    `xorm:"'severity' varchar(32) index"`
    Message       string    `xorm:"'message' varchar(1024)"`
    Raw           string    `xorm:"'raw' text"`
    CreatedAt     time.Time `xorm:"'created_at' datetime created index"`
}
```

`/api/audit/alarm` 不应再只返回成功，应至少落库：

- 原始 JSON。
- RustDesk ID。
- 对端 ID。
- 会话 ID。
- 报警类型。
- 严重级别。

如果无法识别字段，也要保存 `Raw`。

## 7. 后台操作审计设计

新增 `operation_audit` 表：

```go
type OperationAudit struct {
    Id            int       `xorm:"'id' int notnull pk autoincr"`
    ActorUserId   int       `xorm:"'actor_user_id' int index"`
    ActorUsername string    `xorm:"'actor_username' varchar(100) index"`
    Action        string    `xorm:"'action' varchar(100) index"`
    ResourceType  string    `xorm:"'resource_type' varchar(100) index"`
    ResourceId    string    `xorm:"'resource_id' varchar(100) index"`
    BeforeData    string    `xorm:"'before_data' text"`
    AfterData     string    `xorm:"'after_data' text"`
    IP            string    `xorm:"'ip' varchar(64)"`
    UserAgent     string    `xorm:"'user_agent' varchar(512)"`
    Result        string    `xorm:"'result' varchar(32) index"`
    ErrorMessage  string    `xorm:"'error_message' varchar(1024)"`
    CreatedAt     time.Time `xorm:"'created_at' datetime created index"`
}
```

建议记录的后台操作：

- 新增、修改、删除用户。
- 重置密码。
- 启用、禁用账号。
- 删除或修改设备。
- 修改地址簿。
- 修改设备组、用户组、策略。
- 修改系统配置。
- 修改 SMTP、OAuth、OIDC 配置。
- 导出审计日志。
- 清理审计日志。

## 8. 安全事件审计设计

新增 `security_audit` 表：

```go
type SecurityAudit struct {
    Id            int       `xorm:"'id' int notnull pk autoincr"`
    UserId        int       `xorm:"'user_id' int index"`
    Username      string    `xorm:"'username' varchar(100) index"`
    Event         string    `xorm:"'event' varchar(100) index"`
    IP            string    `xorm:"'ip' varchar(64)"`
    UserAgent     string    `xorm:"'user_agent' varchar(512)"`
    Success       bool      `xorm:"'success' bool index"`
    Reason        string    `xorm:"'reason' varchar(512)"`
    CreatedAt     time.Time `xorm:"'created_at' datetime created index"`
}
```

建议事件：

- `login_success`
- `login_failed`
- `logout`
- `token_refresh`
- `token_invalid`
- `password_changed`
- `totp_enabled`
- `totp_disabled`
- `oauth_bind`
- `oauth_unbind`
- `oauth_login_failed`

## 9. API 兼容占位审计

新增 `compat_api_audit` 表：

```go
type CompatAPIAudit struct {
    Id            int       `xorm:"'id' int notnull pk autoincr"`
    Method        string    `xorm:"'method' varchar(16) index"`
    Path          string    `xorm:"'path' varchar(255) index"`
    ClientVersion string    `xorm:"'client_version' varchar(64)"`
    RustdeskId    string    `xorm:"'rustdesk_id' varchar(100) index"`
    IsStub        bool      `xorm:"'is_stub' bool index"`
    StatusCode    int       `xorm:"'status_code' int"`
    IP            string    `xorm:"'ip' varchar(64)"`
    UserAgent     string    `xorm:"'user_agent' varchar(512)"`
    CreatedAt     time.Time `xorm:"'created_at' datetime created index"`
}
```

用途：

- 统计哪些官方兼容接口最常被调用。
- 找出仍是占位实现的接口。
- 辅助决定下一步开发优先级。

## 10. 管理后台页面建议

建议后台新增“审计中心”菜单：

```text
审计中心
├─ 连接审计
├─ 文件传输
├─ 报警事件
├─ 登录安全
├─ 管理操作
├─ API 兼容命中
└─ 审计设置
```

筛选条件：

- 时间范围。
- RustDesk ID。
- 对端 ID。
- 用户名。
- IP。
- 事件类型。
- 成功/失败。
- 客户端版本。
- 是否占位接口。

导出能力：

- CSV。
- JSON。
- 按时间范围导出。
- 按类型导出。

## 11. 保留和清理策略

建议配置：

```yaml
audit:
  enabled: true
  retainDays: 180
  rawPayload: true
  maxRawBytes: 65535
  exportEnabled: true
  cleanupJob:
    enabled: true
    hour: 3
```

说明：

- `retainDays`：日志保留天数。
- `rawPayload`：是否保存原始请求。
- `maxRawBytes`：限制原始 JSON 大小，防止数据库膨胀。
- `cleanupJob`：每天自动清理过期日志。

## 12. 实施步骤

### 阶段 1：不破坏现有表的字段扩展

- 扩展 `Audit` 和 `FileTransfer` 模型。
- `sync` 自动补字段。
- 保留旧字段和旧查询。

### 阶段 2：新增独立审计表

- `AlarmAudit`
- `OperationAudit`
- `SecurityAudit`
- `CompatAPIAudit`

### 阶段 3：接入写入点

- `/api/audit/alarm` 落库。
- 登录成功/失败落库。
- 后台关键操作落库。
- 兼容占位接口落库。

### 阶段 4：后台页面增强

- 增加审计中心菜单。
- 增加分页查询。
- 增加导出。
- 增加清理策略配置。

### 阶段 5：验收

- RustDesk 客户端连接、断开、传文件后后台能看到完整记录。
- 登录失败能看到安全审计。
- 管理员修改用户能看到操作审计。
- 调用占位接口能看到兼容命中记录。
- SQLite/MySQL 都能完成 `sync`。

## 13. 注意事项

- 不要在日志里保存明文密码、token、OAuth secret。
- 原始请求 `Raw` 需要脱敏。
- 文件路径可能包含隐私信息，导出时要提示。
- 审计日志删除也应该记录一条操作审计。
- 大量日志场景下必须依赖索引和分页，不要一次性全量查询。
