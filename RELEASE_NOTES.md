# Release Notes（兼容增强版发布说明模板）

本文档用于当前仓库版本发布时的说明模板，可直接复制到 GitHub Releases 页面后按需微调。

## 标题建议（任选）

- `兼容增强版发布（贴近最新版 RustDesk 客户端 API）`
- `Compatibility Update Release (RustDesk latest client API compatible)`
- `vX.Y.Z - RustDesk API Compatibility Update`

## 发布摘要（可直接使用）

本次版本为“兼容增强版”，重点目标是提升与最新 RustDesk 客户端常用 API 调用流程的兼容性，并保持部署轻量、主流程可用。

当前版本适合作为自建环境下的第三方 API 服务端使用，已补齐大量新版本客户端兼容点（地址簿、设备列表、分组面板基础请求、审计、`devices/cli`、`record` 等），同时保留后续重构空间。

> <span style="color:#ff4d4f;font-weight:700;">警告：本分支部分兼容性更新由 ChatGPT 生成/辅助完成。请在使用前自行审查代码并充分测试，生产环境务必谨慎。</span>

## 本次更新重点

### 0. 当前发布状态（v1.1.55）

- 第三方登录改为后台页面直接配置，GitHub Provider 支持启用、必填项检查、回调地址复制和角色/绑定策略。
- Client Secret 仅显示固定掩码和末 8 位，不通过接口返回明文。
- OAuth/OIDC 失败回调固定返回 hash 路由登录页，错误提示不再闪退，并区分账户不可绑定、Provider 网络不可达与 state 过期。
- Docker 启动脚本保持 YAML 原缩进，避免后台保存配置后重启破坏 `httpConfig.port`。
- Docker 发布流程只由一个质量门事件触发，同一提交不再产生多次互相取消的镜像工作流。

### 1. 新版客户端 API 兼容增强

- 兼容 RustDesk 客户端 1.4.9 主流程 API。
- 地址簿新旧接口并存，支持备注 `note` 字段。
- 增量同步字段兼容（`username` / `hostname` / `platform` / `note`）。
- 用户/设备列表响应字段补齐（如 `display_name`、`device_group_name`）。

### 2. 新增菜单与页面

- **我的同步设备**：管理员可查看所有设备，普通用户通过用户门户登录查看自己的同步设备。
- **通讯录**：联系人和标签改为当前账户全量视图，增加地址簿名称/归属人列和表头筛选，不再通过页面级下拉框切换地址簿。
- **批量数据**：地址簿、联系人、标签和邮件模板支持 CSV 批量导入导出。
- **地址簿权限**：管理员可为指定用户创建地址簿；普通用户只能创建自己的，并且不能删除管理员代建的地址簿。
- **关于与更新**：显示运行版本、构建时间和兼容版本，允许修改在线版本检查地址。

### 3. 版本自动递增系统

- `VERSION` 文件为单一事实来源，CI 每次构建自动递增 PATCH 版本号。
- Go ldflags 注入 `appVersion` 和 `buildTime`，Vite 注入 `VITE_APP_VERSION` 和 `VITE_BUILD_TIME`。
- 首页更新日志区域显示服务端版本与构建时间，方便确认是否更新成功。

### 4. Docker 镜像自动推送

- 每次推送到 main 分支，CI 自动构建并推送镜像到 GHCR（`latest` + `main` + `sha-xxx` 标签）。
- 打 tag 时额外推送版本号标签。

### 5. 客户端附加兼容端点

- `POST /api/devices/cli`：由 no-op 升级为最小可用写入。
- `POST /api/record`：支持最小落盘协议（`new/part/tail/remove`）。
- `POST /api/sysinfo_ver`：稳定返回，减少重复 `sysinfo` 上传。
- `GET /api/login-options`：显式返回 `[]`。

### 3. 占位兼容端点（避免 404）

- `POST /api/oidc/auth`
- `GET /api/oidc/auth-query`
- `POST /lic/web/api/plugin-sign`

说明：以上端点当前为兼容实现或占位返回，用于避免客户端报错，不等同官方 Pro 完整功能。

### 4. 管理后台与文档整理

- 移除赞助/收款码相关界面与文案。
- README 默认中文展示并保留英文入口。
- 新增文档中心（使用说明 / 端口说明 / 问题排查手册）。

## 当前兼容范围（发布说明版）

当前版本可覆盖的主要流程：

- 登录 / 退出 / 当前用户信息。
- 地址簿读写（含备注）。
- 设备列表、用户列表、分组面板基础加载。
- 心跳、`sysinfo`、审计上报、审计备注更新。
- `devices/cli` 最小写入与 `record` 最小落盘。

## 已知限制（建议保留）

以下能力仍非官方 Pro 完整实现：

- OIDC 完整登录流程（当前为兼容返回）。
- 插件签名真实服务（`plugin-sign` 为兼容占位）。
- 完整分组权限模型（`accessible` 相关接口为兼容模型）。

如果发布目标是“最新客户端主流程可用”，当前版本可作为上线候选；
如果发布目标是“完整替代官方 Pro 高级能力”，建议继续补齐关键能力后再上线。

## 升级与发布前注意事项

### 1. 数据库结构同步（必须）

升级后请执行一次数据库同步：

```bash
rustdesk-api-server-pro sync
```

然后重启服务。

### 2. 发布前冒烟测试（建议）

- 登录 / 退出
- 地址簿列表与备注保存
- 设备列表与分组面板加载（不报错）
- 审计写入与备注更新
- 如启用录制：确认 `record_uploads/` 可写

### 3. 配置检查

- `server.yaml` 中 `httpConfig.port` 与对外端口一致
- `httpConfig.staticdir` 指向正确的前端 `dist` 目录
- Docker 部署时确认卷挂载正确

## 验证状态（请按实际结果更新）

发布前本地检查（示例）：

- 后端：`go build ./...` ✅
- 后端：`go test ./...` ✅
- 前端：`pnpm typecheck` ✅
- 前端：`pnpm build` ✅

## 文档入口

- 使用说明：`docs/USAGE.md`
- 端口说明：`docs/PORTS.md`
- 问题排查：`docs/TROUBLESHOOTING.md`

## 发布建议

本版本重点是“兼容可用”和“轻量部署”，并不声明完全等价官方 Pro。
建议结合自身业务场景完成一轮完整回归测试后再上线。
