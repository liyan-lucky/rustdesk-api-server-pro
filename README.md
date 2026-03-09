# RustDesk API Server Pro（兼容增强版）

[简体中文（默认）](./README.md) | [English](./README_EN.md)

基于 [RustDesk](https://github.com/rustdesk/rustdesk) 客户端调用行为的第三方 API 服务端实现，并提供 Web 管理后台（`soybean-admin`）。

> <span style="color:#ff4d4f;font-weight:700;">警告：本分支部分兼容性更新由 ChatGPT 生成/辅助完成。请在使用前自行审查代码并充分测试，生产环境务必谨慎。</span>

![Dashboard](./img/1.jpeg "Dashboard")

## 文档中心（建议优先阅读）

- [使用说明（部署、初始化、升级、验证）](./docs/USAGE.md)
- [Docker 安装与配置参考](./docs/DOCKER.md)
- [端口与访问路径说明（HTTP/API/Admin/SMTP）](./docs/PORTS.md)
- [问题排查手册（常见问题与日志定位）](./docs/TROUBLESHOOTING.md)
- [发布说明模板（GitHub Release 可直接复用）](./RELEASE_NOTES.md)

## Docker 部署示例

```bash
mkdir -p /opt/rustdesk-api-server-pro/data
cd /opt/rustdesk-api-server-pro

curl -L -o server.yaml https://raw.githubusercontent.com/liyan-lucky/rustdesk-api-server-pro/master/backend/server.yaml

docker run -d \
  --name rustdesk-api-server-pro \
  --restart unless-stopped \
  --network host \
  -e ADMIN_USER=admin \
  -e ADMIN_PASS='ChangeMe123!' \
  -v $(pwd)/server.yaml:/app/server.yaml \
  -v $(pwd)/data:/app/data \
  ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest

docker logs -f rustdesk-api-server-pro
```

## 当前状态（兼容增强版）

- 当前分支定位为“兼容增强版”，目标是贴近最新 RustDesk 客户端常用 API 调用流程。
- 已补齐大量新版本客户端兼容点（地址簿、设备列表、分组面板基础请求、审计、`sysinfo`、`devices/cli`、`record` 等）。
- 前端管理后台已移除赞助/收款码相关内容。
- README 默认中文展示，保留英文入口。
- 项目仍保留后续重构计划（见 issue #30），当前版本可用于“可上线兼容版”。

## 项目定位

本项目不是官方 Pro 服务端，也不追求在短期内完整复刻官方全部高级能力。当前设计目标：

- 兼容优先：优先保证新客户端主流程不报错、不 404，关键字段可正常读写。
- 轻量优先：支持单机部署，默认 SQLite，可选 MySQL。
- 迭代优先：通过兼容层持续适配客户端 API 变化，降低升级成本。

适用场景：

- 自建 RustDesk API 服务端 + 管理后台。
- 研究 RustDesk 客户端 API 调用逻辑并进行二次开发。

不适用场景：

- 强依赖官方 Pro 完整 OIDC 登录流程。
- 强依赖官方完整分组权限模型。
- 强依赖官方插件签名服务的生产校验链路。

## 与新版客户端兼容范围（当前代码状态）

已覆盖或兼容处理的主要接口与流程：

- 账号相关：`/api/login`、`/api/logout`、`/api/currentUser`、`/api/login-options`
- 地址簿：`/api/ab`、`/api/ab/settings`、`/api/ab/personal`、`/api/ab/shared/profiles`、`/api/ab/peers`、`/api/ab/tag/*`、`/api/ab/peer/*`
- 分组/设备面板基础请求：`/api/device-group/accessible`、`/api/users?accessible=...`、`/api/peers?accessible=...`
- 心跳与系统信息：`/api/heartbeat`、`/api/sysinfo`、`/api/sysinfo_ver`
- 审计：`/api/audit/conn`、`/api/audit/file`、`/api/audit/alarm`、`PUT /api/audit`
- 兼容端点：`POST /api/devices/cli`、`POST /api/record`、`POST /api/oidc/auth`、`GET /api/oidc/auth-query`、`POST /lic/web/api/plugin-sign`

## 已补齐的关键兼容点

- 地址簿 `note` 字段支持（读/写/同步）。
- 新版增量更新字段兼容（`username`、`hostname`、`platform`、`note`）。
- `display_name`、`device_group_name` 字段补齐。
- `devices/cli` 从 no-op 升级为最小可用写入。
- `record` 从 no-op 升级为最小落盘协议实现（`new/part/tail/remove`）。
- `sysinfo_ver` 改为稳定返回，减少重复 `sysinfo` 上传。
- `PUT /api/audit` 备注更新兼容。

## 非完整官方 Pro 实现（发布前请知悉）

以下能力当前属于“兼容实现”或“占位返回”，用于避免客户端报错，不等同官方 Pro：

- OIDC（`/api/oidc/*`）
- 插件签名（`/lic/web/api/plugin-sign`）
- 完整分组权限模型（`accessible` 相关接口为兼容模型）

发布判断建议：

- 若目标是“最新客户端主流程可用”：当前版本可发布。
- 若目标是“完整替代官方 Pro 高级能力”：建议继续补齐关键能力后再发布。

## 技术栈

- 后端：Go（Iris）
- 前端：Vue 3 + Vite + Naive UI（`soybean-admin`）
- 数据库：SQLite（默认）/ MySQL（可选）

## 目录结构

- `backend/`：Go 后端 API 服务
- `soybean-admin/`：管理后台前端
- `docker/`：容器相关配置
- `docs/`：部署与排障文档
- `img/`：README 图片资源

## 发布前最小检查（建议）

1. 执行数据库结构同步：`rustdesk-api-server-pro sync`
2. 重启服务并确认 `server.yaml` 端口与静态目录配置正确
3. 用新版客户端做冒烟测试（登录、地址簿、设备列表、分组、审计）
4. 检查日志是否存在持续报错（OIDC/record/权限/字段缺失等）

## 备注

- GitHub 页面按钮/标签（如 `README`、`Commits`）显示语言由 GitHub 与浏览器语言决定，仓库代码无法直接修改。
- 仓库首页只展示概要信息，详细部署与排障请查看“文档中心”。