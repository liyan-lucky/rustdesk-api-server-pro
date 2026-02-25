# Release Notes（兼容增强版发布说明模板）

本文档用于当前仓库版本发布时的说明模板，可直接复制到 GitHub Releases 页面后按需微调。

## 标题建议（任选其一）

- `兼容增强版发布（贴近最新版 RustDesk 客户端 API）`
- `Compatibility Update Release (RustDesk latest client API compatible)`
- `vX.Y.Z - RustDesk API Compatibility Update`

## 发布摘要（可直接使用）

本次版本为“兼容增强版”发布，重点目标是提升与最新版 RustDesk 客户端常用 API 调用流程的兼容性，并保持部署轻量、主流程可用。

本版本适合作为自建环境下的第三方 API 服务端使用，已补齐大量新版客户端兼容点（地址簿、设备列表、分组面板基础请求、审计、`devices/cli`、`record` 等），同时保留后续结构重构空间。

> 警告：本分支部分兼容性更新由 ChatGPT 生成/辅助完成。请在生产环境使用前自行审查代码并充分测试。

## 本次更新重点（建议保留）

### 1. 新版客户端 API 兼容增强

- 补齐新版 RustDesk 客户端常见 API 调用路径与字段兼容
- 地址簿新旧接口并存兼容（含备注 `note` 字段）
- 增量同步字段兼容（如 `username` / `hostname` / `platform` / `note`）
- 用户/设备列表响应字段补齐（如 `display_name`、`device_group_name`）

### 2. 客户端附加兼容端点

- `POST /api/devices/cli`：由 no-op 升级为最小可用写入
- `POST /api/record`：支持最小落盘协议（`new/part/tail/remove`）
- `POST /api/sysinfo_ver`：稳定返回，减少重复 `sysinfo` 上传
- `GET /api/login-options`：显式返回 `[]`

### 3. 兼容占位端点（避免 404）

- `POST /api/oidc/auth`
- `GET /api/oidc/auth-query`
- `POST /lic/web/api/plugin-sign`

说明：以上端点当前为兼容实现或占位返回，避免客户端报错，但不等同官方 Pro 完整功能。

### 4. 管理后台与文档整理

- 移除赞助/收款码相关界面与文案
- 仓库 README 调整为中文默认显示 + 英文入口
- 新增文档中心（使用说明 / 端口说明 / 问题排查手册）

## 当前兼容范围（发布说明版）

当前版本可覆盖的主要流程：

- 登录 / 退出 / 当前用户信息
- 地址簿读取与写入（含备注）
- 设备列表、用户列表、分组面板基础加载
- 心跳 / `sysinfo` / 审计上报 / 审计备注更新
- `devices/cli` 基础写入
- `record` 最小上传落盘

## 已知限制（建议保留）

以下能力仍非官方 Pro 完整实现：

- OIDC 完整登录流程（当前为兼容返回）
- 插件签名真实服务（`plugin-sign` 为兼容占位）
- 完整分组权限模型（`accessible` 相关接口为兼容模型）

如果你的目标是“最新版客户端主流程可用”，本版本可直接使用；
如果你的目标是“完全替代官方 Pro 高级能力”，建议继续补齐上述功能。

## 升级与发布前注意事项（建议保留）

### 1. 数据库结构同步（必须）

升级后请执行一次数据库同步：

```bash
rustdesk-api-server-pro.exe sync
```

然后重启服务。

### 2. 发布前冒烟测试（建议）

- 登录 / 退出
- 地址簿列表与备注保存
- 设备列表加载
- 分组面板加载（不报错）
- 审计写入与备注更新
- 如启用录制：确认 `record_uploads/` 可写

### 3. 配置确认

- `server.yaml` 中 `httpConfig.port` 与实际对外端口一致
- `httpConfig.staticdir` 指向正确的前端 `dist` 目录
- Docker 场景确认端口映射与卷挂载正确

## 验证状态（你可以按实际修改）

发布前本地检查（示例）：

- 后端：`go build ./...` ✅
- 后端：`go test ./...` ✅
- 前端：`pnpm typecheck` ✅
- 前端：`pnpm build` ✅

## 文档入口（建议保留）

- 使用说明：`docs/USAGE.md`
- 端口说明：`docs/PORTS.md`
- 问题排查：`docs/TROUBLESHOOTING.md`

## 致使用者

本版本重点是“兼容可用”和“轻量部署”，并不声称完全等价官方 Pro。建议在生产环境上线前完成以下动作：

1. 结合自身业务场景做一轮完整功能回归测试
2. 开启请求日志观察 1~2 天
3. 根据日志决定是否继续补齐 OIDC / plugin-sign / 分组权限模型
