# API 兼容矩阵

本文档用于跟踪 RustDesk 客户端 API、管理后台 API、license/plugin-sign 兼容接口的实现状态。

状态说明：

- `完整`：已实现主要业务逻辑，并经过客户端验证。
- `基础`：主流程可用，但字段或边界场景仍需补齐。
- `占位`：避免客户端 404 或报错，暂未完整实现业务逻辑。
- `待核验`：源码中可能已有实现，但需要用真实客户端重新验证。
- `计划`：尚未实现。

## 1. 客户端公开接口

| 模块 | 方法 | 路径 | 鉴权 | 状态 | 说明 |
| --- | --- | --- | --- | --- | --- |
| 系统 | GET/POST | `/api/health`、`/api/ping` | 否 | 基础 | 健康检查，已纳入 smoke 与兼容命中审计 |
| 系统 | GET/POST | `/api/status`、`/api/version`、`/api/info` | 否 | 基础 | 返回版本与 `compat_target`，已纳入 smoke 与兼容命中审计 |
| 系统 | GET/POST | `/api/features`、`/api/capabilities`、`/api/compat/features` | 否 | 基础 | 返回能力开关，包含 `compat_api_audit` |
| 系统 | GET/POST | `/api/config`、`/api/client-config`、`/api/server-config` | 否 | 基础 | 返回客户端配置形状，已纳入 smoke 与兼容命中审计 |
| 系统 | GET/POST | `/api/compat-target`、`/api/compat/target`、`/api/compat/version` | 否 | 基础 | 返回当前匹配对象，目标 RustDesk 1.4.8 |
| 系统 | GET/POST | `/api/sysinfo_ver` | 否 | 基础 | 返回兼容 sysinfo 版本字符串 |
| 登录 | GET/POST | `/api/login-options` | 否 | 基础 | 返回可用登录方式，已纳入公开 smoke 与兼容命中审计 |
| 登录 | POST | `/api/login` | 否 | 基础/待核验 | 账号登录、token 返回；成功/失败/验证码或 2FA 中间态已写 `security_audit` |
| 审计 | POST | `/api/audit/conn` | 否/待核验 | 基础 | 连接开始、关闭、备注更新 |
| 审计 | POST | `/api/audit/file` | 否/待核验 | 基础 | 文件传输日志 |
| 审计 | POST | `/api/audit/alarm` | 否/待核验 | 基础 | 已落库到 `alarm_audit` |
| 兼容 | * | `/api/*` 公开探测别名 | 否 | 基础 | 公开兼容控制器已写入 `compat_api_audit`，用于发现新版客户端真实探测 |

## 2. 客户端鉴权接口

| 模块 | 方法 | 路径 | 鉴权 | 状态 | 说明 |
| --- | --- | --- | --- | --- | --- |
| 用户 | GET/POST | `/api/user/*` | 是 | 待核验 | 用户信息、当前账号 |
| 退出登录 | GET/POST/DELETE | `/api/logout` | 是 | 基础 | 退出登录成功/失败已写 `security_audit` |
| token 鉴权 | * | `/api/*` 鉴权接口 | 是 | 基础 | 客户端 token 无效、token 对应用户无效已写 `security_audit` |
| 地址簿 | GET/POST | `/api/peers/*` | 是 | 基础 | 老接口地址簿兼容 |
| 地址簿 | GET/POST | `/api/ab/*` | 是 | 基础 | 新地址簿主体 |
| 地址簿标签 | GET/POST | `/api/ab/tags/*` | 是 | 基础 | 标签、颜色、备注兼容 |
| 地址簿设备 | GET/POST | `/api/ab/peers/*` | 是 | 基础 | 地址簿设备条目 |
| 设备组 | GET/POST | `/api/device-group/*` | 是 | 基础/待核验 | 企业设备组兼容 |
| 企业兼容 | * | `/api/*` | 是 | 基础/占位 | 策略、组、企业字段兼容 |
| 鉴权兼容 | * | `/api/*` | 是 | 基础/占位 | token、会话相关兼容 |

## 3. 管理后台接口

| 模块 | 路径 | 状态 | 建议 |
| --- | --- | --- | --- |
| 后台登录 | `/admin/auth/*` | 基础 | 管理员账号登录成功/失败、OIDC/OAuth 回调与 ticket 换 token、后台 token 无效已写 `security_audit` |
| 仪表盘 | `/admin/dashboard/*` | 已有 | 增加审计概览 |
| 用户管理 | `/admin/users/*` | 基础 | 新增、修改、删除用户已写 `operation_audit`；不记录密码和 2FA 密钥明文；空删除列表会返回 `NoUserIds` 并记录失败审计 |
| 会话管理 | `/admin/sessions/*` | 已有 | 增加 token 安全事件 |
| 设备管理 | `/admin/devices/*` | 已有 | 增加设备变更审计 |
| 审计日志 | `/admin/audit/*` | 已有 | 增加高级筛选、导出、报警审计、兼容探测审计视图 |
| 邮件模板 | `/admin/mail-template/*` | 已有 | 增加修改审计 |
| 邮件日志 | `/admin/mail-logs/*` | 已有 | 增加发送失败分析 |

## 4. License / plugin-sign 兼容接口

| 方法 | 路径 | 状态 | 建议 |
| --- | --- | --- | --- |
| POST | `/lic/web/api/plugin-sign` | 基础/占位 | 已纳入 smoke 与 `compat_api_audit`；当前为稳定 JSON 形状 + 消息透传，不声明官方签名等价 |
| * | `/lic/web/api/*` | 待核验 | 建立真实客户端抓包验证样例 |

## 5. 审计覆盖矩阵

| 事件 | 是否已有 | 目标表 | 优先级 |
| --- | --- | --- | --- |
| 远程连接开始 | 是 | `audit` / `connection_audit` | P0 |
| 远程连接关闭 | 是 | `audit` / `connection_audit` | P0 |
| 连接备注修改 | 是 | `audit` | P0 |
| 文件传输 | 是 | `file_transfer` | P0 |
| 客户端报警 | 是 | `alarm_audit` | P1 |
| 客户端登录成功 | 是 | `security_audit` | P0 |
| 客户端登录失败 | 是 | `security_audit` | P0 |
| 客户端验证码/2FA 中间态 | 是 | `security_audit` | P1 |
| 客户端退出登录 | 是 | `security_audit` | P1 |
| 客户端 token 无效 | 是 | `security_audit` | P1 |
| 客户端 token 对应用户无效 | 是 | `security_audit` | P1 |
| 后台登录成功 | 是 | `security_audit` | P0 |
| 后台登录失败 | 是 | `security_audit` | P0 |
| 后台 token 无效 | 是 | `security_audit` | P1 |
| 后台 token 对应管理员无效 | 是 | `security_audit` | P1 |
| OIDC/OAuth 回调成功 | 是 | `security_audit` | P1 |
| OIDC/OAuth 回调失败 | 是 | `security_audit` | P1 |
| OIDC/OAuth ticket 换 token | 是 | `security_audit` | P1 |
| 后台新增用户 | 是 | `operation_audit` | P0 |
| 后台修改用户 | 是 | `operation_audit` | P0 |
| 后台删除用户 | 是 | `operation_audit` | P0 |
| 修改设备 | 待补 | `operation_audit` | P1 |
| 修改地址簿 | 待补 | `operation_audit` | P1 |
| 修改策略 | 待补 | `operation_audit` | P2 |
| 兼容/占位接口命中 | 是 | `compat_api_audit` | P1 |

## 6. 每个接口的记录模板

后续每补齐一个接口，请按以下格式更新本文：

```text
模块：
路径：
方法：
鉴权：是/否
当前状态：完整/基础/占位/计划
请求字段：
响应字段：
是否写审计：
兼容客户端版本：
SQLite 验证：通过/失败
MySQL 验证：通过/失败
备注：
```

## 7. 下一批建议开发任务

1. 后台审计页面增加 `security_audit` / `compat_api_audit` / `operation_audit` 视图，按事件、path、method、is_stub、result、client_version、resource_type 聚合。
2. 新增 `operation_audit` 接入设备管理修改。
3. 新增 `operation_audit` 接入地址簿和策略修改。
4. 建立真实 RustDesk 客户端抓包样例目录，补齐官方接口字段差异。
5. 对 `/lic/web/api/*` 继续做真实客户端验证，避免误以为 plugin-sign 透传等价于官方签名服务。
6. 增强 authenticated smoke：用测试用户登录后验证 `/api/currentUser`、`/api/logout` 和安全审计落库。
