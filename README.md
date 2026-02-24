# Rustdesk Api Server Pro

[简体中文](./README.md) | [English](./README_EN.md)

基于开源 [RustDesk](https://github.com/rustdesk/rustdesk) 客户端的 API 服务端实现，并提供配套 Web 管理后台（`soybean-admin`）。

> 警告：本分支部分兼容性更新由 ChatGPT 生成/辅助完成。请在使用前自行审查代码并充分测试，生产环境务必谨慎。

![Dashboard](./img/1.jpeg "Dashboard")

## 当前状态

- 当前分支定位为“兼容增强版”，重点提升对新版 RustDesk 客户端 API 的兼容性
- 已移除赞助/收款码相关界面与文案
- 仓库默认 README 已改为中文显示（顶部可切换英文）
- 项目后续仍计划重写，进度见：<https://github.com/lantongxue/rustdesk-api-server-pro/issues/30>

## 项目定位（详细说明）

本项目是面向 RustDesk 客户端生态的第三方 API 服务端实现，目标是在尽量保持轻量化部署体验的前提下，为客户端提供可用的 API 接口与基础管理后台能力。

与官方服务端/官方 Pro 服务端相比，本项目当前更偏向：

- 兼容优先：优先保证客户端请求不报错、主流程可用
- 轻量优先：默认 sqlite、单机部署成本低
- 可维护优先：通过兼容层逐步补齐新版客户端调用点

这意味着它适合用于：

- 私有化自建环境的 API 对接与管理后台
- 学习/研究 RustDesk 客户端 API 调用方式
- 在官方 Pro 能力之外，做定制化接口扩展

同时也意味着在某些高级能力上（例如完整 OIDC、官方插件签名、完整分组权限模型）仍与官方 Pro 存在差异。

## 兼容性说明（基于当前代码状态）

- 已兼容新版客户端常用 API 流程：
  - 登录 / 登出 / 当前用户 / 登录选项
  - 地址簿（新旧接口）与地址簿备注字段
  - 设备列表 / 用户列表 / 分组面板基础请求
  - 心跳 / sysinfo / 审计上报与审计备注更新
  - `devices/cli` 最小可用更新能力
  - `record` 最小上传落盘协议（`new/part/tail/remove`）
- 新增兼容端点（避免 404）：
  - `/api/oidc/auth`
  - `/api/oidc/auth-query`
  - `/lic/web/api/plugin-sign`

### 已补齐的关键兼容点（相对老版本第三方实现）

- 地址簿 `note` 字段读写与同步
- 新版地址簿增量更新字段（如 `username`、`hostname`、`platform`）
- 分组面板相关请求兼容（`device-group/accessible`、`users/peers accessible`）
- `devices/cli` 最小写入能力（设备/地址簿部分字段回写）
- `record` 上传协议最小落盘实现（用于避免新版客户端录制上传时报错）
- `sysinfo_ver` 稳定返回，减少客户端重复上传 `sysinfo`
- 审计备注更新接口兼容（`PUT /api/audit`）

## 非完整实现（发布前请知悉）

- `OIDC` 仍为兼容返回（未实现完整登录流程）
- `plugin-sign` 为兼容占位实现（非官方签名服务）
- `device-group/accessible`、`users?accessible=`、`peers?accessible=` 为兼容模型，不等同官方 Pro 权限逻辑

### 对发布决策的影响（建议）

- 如果你的目标是“新版客户端主流程可用”：当前版本可以发布
- 如果你的目标是“完整替代官方 Pro 的高级能力”：建议继续补齐 OIDC、插件签名、完整权限模型后再发布

## 技术栈

- 后端：Go（Iris）
- 前端：Vue 3 + Vite + Naive UI（`soybean-admin`）
- 数据库：SQLite（默认）/ MySQL（可选）

## 项目结构

- `backend/`：Go 后端 API 服务
- `soybean-admin/`：管理后台前端
- `docker/`：容器相关配置
- `img/`：README 图片资源

## 发布前建议（最小流程）

1. 执行数据库结构同步：`rustdesk-api-server-pro.exe sync`
2. 重启服务
3. 用最新版客户端做冒烟测试（登录、地址簿、分组面板、设备列表、审计）

### 推荐冒烟测试项（更详细）

- 账号登录/退出、`currentUser` 返回正常
- 地址簿新增/编辑/删除、备注字段保存与回显
- 分组面板加载不报错（即使无真实分组数据）
- 设备列表可加载，字段显示正常
- 审计日志可记录，审计备注可修改
- 如启用录制：确认 `record_uploads/` 目录可写

## 说明

- GitHub 页面按钮/标签语言（如 `README`、`Commits`）由 GitHub/浏览器语言决定，仓库无法直接修改。
- 本 README 仅用于说明当前分支状态与兼容范围，详细开发/部署细节可按需补充到独立文档。
