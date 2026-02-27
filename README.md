# RustDesk API Server Pro（兼容增强版）

[简体中文（默认）](./README.md) | [English](./README_EN.md)

基于开源 [RustDesk](https://github.com/rustdesk/rustdesk) 客户端调用行为的第三方 API 服务端实现，并提供 Web 管理后台（`soybean-admin`）。

> 警告：本分支部分兼容性更新由 ChatGPT 生成/辅助完成。请在使用前自行审查代码并充分测试，生产环境务必谨慎。

![Dashboard](./img/1.jpeg "Dashboard")

## 文档中心（建议优先查阅）

- [使用说明（部署、初始化、升级、验证）](./docs/USAGE.md)
- [Docker 安装与配置参考（安装命令、默认参数、卷、端口）](./docs/DOCKER.md)
- [端口与访问路径说明（HTTP/API/Admin/SMTP）](./docs/PORTS.md)
- [问题排查手册（常见问题与日志定位）](./docs/TROUBLESHOOTING.md)
- [发版说明模板（GitHub Release 可直接复用）](./RELEASE_NOTES.md)

## Docker 部署示例（命令）

```bash
mkdir -p /opt/rustdesk-api-server-pro/data
cd /opt/rustdesk-api-server-pro

curl -L -o server.yaml https://raw.githubusercontent.com/liyan-lucky/rustdesk-api-server-pro/master/backend/server.yaml

docker run -d \\
  --name rustdesk-api-server-pro \\
  --restart unless-stopped \\
  --network host \\
  -e ADMIN_USER=admin \\
  -e ADMIN_PASS='ChangeMe123!' \\
  -v $(pwd)/server.yaml:/app/server.yaml \\
  -v $(pwd)/data:/app/data \\
  ghcr.io/liyan-lucky/rustdesk-api-server-pro:latest

docker logs -f rustdesk-api-server-pro
```

- 详细 Docker 安装命令、Compose 示例、默认参数、端口与卷说明：`docs/DOCKER.md`
- 默认示例使用 `host` 网络，实际监听端口以 `server.yaml` 中 `httpConfig.port` 为准（常见 `:12345`）

## 当前状态（兼容增强版）

- 当前分支定位为“兼容增强版”，目标是尽量贴近最新版 RustDesk 客户端的常用 API 调用流程
- 已完成新版客户端常见接口兼容补齐（地址簿、设备列表、分组面板基础请求、审计、sysinfo、devices/cli、record 等）
- 前端管理后台已移除赞助/收款码相关界面与文案
- 仓库主页默认语言为中文（可通过顶部链接切换英文说明）
- 项目仍保留后续重构计划（见 issue #30），但当前版本可以作为“可上线兼容版”使用

## 项目定位（详细说明）

本项目不是官方 Pro 服务端，也不追求在短期内完全复刻官方全部高级能力。当前版本的设计目标是：

- 兼容优先：优先保证新版客户端主流程不报错、不 404、关键字段能正确读写
- 轻量优先：支持单机部署，默认 SQLite，适合私有化自建或中小规模场景
- 迭代优先：通过兼容层持续补齐官方客户端 API 变化，降低升级成本

适用场景：

- 自建 RustDesk API 服务端 + 管理后台
- 内网/私有环境的设备管理与基础审计
- 研究 RustDesk 客户端 API 调用逻辑并做二次开发

不适用或需谨慎评估的场景：

- 强依赖官方 Pro 的完整 OIDC 登录流程
- 强依赖官方插件签名服务（`plugin-sign`）的生产校验链路
- 强依赖官方完整分组权限模型（`device-group/accessible`、`users/peers?accessible=` 的细粒度权限）

## 与新版客户端兼容范围（当前代码状态）

已覆盖或兼容处理的主要接口与流程：

- 账号相关：`/api/login`、`/api/logout`、`/api/currentUser`、`/api/login-options`
- 地址簿（新旧接口并存）：
  - `/api/ab`
  - `/api/ab/settings`
  - `/api/ab/personal`
  - `/api/ab/shared/profiles`
  - `/api/ab/peers`
  - `/api/ab/tags/{guid}`
  - `/api/ab/peer/*`
  - `/api/ab/tag/*`
- 分组/设备面板基础请求：
  - `/api/device-group/accessible`
  - `/api/users?accessible=...`
  - `/api/peers?accessible=...`
- 同步与状态：
  - `/api/heartbeat`
  - `/api/sysinfo`
  - `/api/sysinfo_ver`
- 审计：
  - `/api/audit/conn`
  - `/api/audit/file`
  - `/api/audit/alarm`
  - `PUT /api/audit`（备注更新兼容）
- 客户端附加兼容端点：
  - `POST /api/devices/cli`（最小可用写入）
  - `POST /api/record`（最小落盘协议：`new/part/tail/remove`）
  - `POST /api/oidc/auth`（兼容返回）
  - `GET /api/oidc/auth-query`（兼容返回）
  - `POST /lic/web/api/plugin-sign`（兼容占位返回）

## 已补齐的关键兼容点（相对旧版第三方实现）

- 地址簿 `note` 字段支持（读/写/同步）
- 新版地址簿增量更新字段兼容（如 `username`、`hostname`、`platform`、`note`）
- `display_name` 字段补齐（用户/登录相关响应）
- `device_group_name` 字段补齐（设备列表兼容字段）
- `devices/cli` 从 no-op 升级为最小可用写入（设备与地址簿部分字段回写）
- `record` 从 no-op 升级为最小落盘协议实现
- `sysinfo_ver` 改为稳定返回值，减少客户端重复上传 `sysinfo`
- `PUT /api/audit` 备注更新兼容

## 非完整官方 Pro 实现（发布前请知悉）

以下能力目前是“兼容实现”或“占位返回”，用于避免客户端报错，但不等同官方 Pro：

- OIDC（`/api/oidc/*`）：返回兼容结构，未实现完整登录流程
- 插件签名（`/lic/web/api/plugin-sign`）：兼容占位，不是官方签名服务
- 分组权限模型：`device-group/accessible`、`users/peers?accessible=` 为兼容模型，非官方完整权限逻辑

发布判断建议：

- 如果目标是“最新版客户端主流程可用”：当前版本可发布
- 如果目标是“完全替代官方 Pro 高级能力”：建议继续补齐 OIDC、插件签名、完整权限模型后再发布

## 技术栈

- 后端：Go（Iris）
- 前端：Vue 3 + Vite + Naive UI（`soybean-admin`）
- 数据库：SQLite（默认）/ MySQL（可选）

## 项目结构

- `backend/`：Go 后端 API 服务
- `soybean-admin/`：管理后台前端
- `docker/`：容器相关配置
- `docs/`：使用说明、端口说明、问题排查手册
- `img/`：README 图片资源

## 发布前最小检查（建议）

1. 执行数据库结构同步：`rustdesk-api-server-pro.exe sync`
2. 重启服务，确认 `server.yaml` 中端口与静态目录配置正确
3. 用最新版客户端做一轮冒烟测试（登录、地址簿、设备列表、分组面板、审计）
4. 检查日志中是否存在持续报错（OIDC/record/权限/数据库字段缺失等）

## 说明

- GitHub 页面按钮/标签（如 `README`、`Commits`）的显示语言由 GitHub 与浏览器语言决定，仓库代码无法直接修改
- 仓库首页仅展示总览信息；详细部署与排障请查看上方“文档中心”
