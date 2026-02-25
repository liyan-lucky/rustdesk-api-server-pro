# 使用说明（部署、初始化、升级、验证）

本文档面向当前仓库状态（RustDesk API Server Pro 兼容增强版），用于指导部署、初始化、升级与上线前验证。

## 1. 项目组成

- `backend/`：Go 后端 API 服务（客户端接口 + 管理后台接口）
- `soybean-admin/`：管理后台前端（构建后静态文件由后端同端口提供）
- `docker/`：容器相关配置

## 2. 运行方式概览

支持两种常见方式：

- 二进制部署（本地或服务器直接运行）
- Docker / Docker Compose 部署

无论哪种方式，核心流程都是：

1. 准备配置（`backend/server.yaml`）
2. 同步数据库结构（`sync`）
3. 启动服务（`start`）
4. 用浏览器和客户端做冒烟验证

## 3. 二进制部署（推荐先本地验证）

### 3.1 后端准备

在 `backend/` 目录执行：

```powershell
go build -o rustdesk-api-server-pro.exe .
```

首次运行前准备配置文件：

- 如果 `backend/server.yaml` 不存在，程序会按默认配置生成
- 也可以手动修改现有 `backend/server.yaml`

关键配置项：

- `db.driver`：`sqlite`（默认）或 `mysql`
- `db.dsn`：数据库连接字符串
- `httpConfig.port`：HTTP 监听端口（如 `:12345`）
- `httpConfig.staticdir`：前端静态文件目录（如 `dist` 或 `/app/dist`）
- `smtpConfig.*`：邮箱功能相关配置（可选）

### 3.2 数据库结构同步（必须）

每次升级前后建议执行一次：

```powershell
.\rustdesk-api-server-pro.exe sync
```

说明：

- 本仓库最近补充了地址簿兼容字段（如 `peer.note`），未执行 `sync` 可能导致运行时报错
- 执行 `sync` 后请重启服务

### 3.3 启动服务

```powershell
.\rustdesk-api-server-pro.exe start
```

启动成功后：

- 客户端接口与管理后台都在同一个 HTTP 端口上提供
- 管理后台前端入口通常为：`http://<服务器IP>:<端口>/`

## 4. 前端开发/构建（如需修改管理后台）

目录：`soybean-admin/`

安装依赖：

```powershell
pnpm install
```

开发检查：

```powershell
pnpm typecheck
pnpm build
```

说明：

- 当前仓库状态下，`pnpm typecheck` 和 `pnpm build` 应可通过
- 构建产物位于 `soybean-admin/dist`
- 部署时需保证后端 `httpConfig.staticdir` 指向正确的 `dist` 目录

## 5. Docker / Docker Compose（常见部署方式）

仓库内已包含：

- `Dockerfile`
- `docker-compose.yaml`

建议流程：

1. 根据你的环境修改 `docker-compose.yaml`（端口映射、卷挂载、配置路径）
2. 确认 `server.yaml` 中 `httpConfig.staticdir` 与容器内路径一致（示例通常为 `/app/dist`）
3. 启动容器
4. 进入容器执行一次 `sync`（若镜像启动流程未自动执行）

如果你用容器部署，务必确认以下目录有持久化：

- 数据库文件目录（SQLite 时）
- `server.yaml`
- `record_uploads/`（如果客户端启用录制上传）

## 6. 上线前最小验证（强烈建议）

### 6.1 服务端可用性

- 后端日志无持续 panic / SQL 错误
- 浏览器打开首页能看到管理后台页面
- `server.yaml` 中端口与静态目录配置正确

### 6.2 RustDesk 客户端主流程

- 登录成功
- 退出成功
- 地址簿列表正常加载
- 地址簿备注（`note`）可保存并刷新回显
- 设备列表正常显示
- 分组面板可打开（即使暂无真实分组数据，也不应报 404）
- 审计记录能写入，审计备注可更新

### 6.3 兼容端点验证（按需）

- `devices/cli`：客户端更新设备信息/备注后无报错
- `record`：启用录制时服务端 `record_uploads/` 可写入文件
- `oidc`：如未使用，可接受返回“不支持”的兼容响应
- `plugin-sign`：如未使用插件签名，可接受兼容占位返回

## 7. 升级建议（当前仓库适用）

升级本项目版本时建议执行：

1. 备份数据库与 `server.yaml`
2. 更新代码
3. 重新构建后端与前端（如你有二次开发）
4. 执行 `sync`
5. 重启服务
6. 用最新版 RustDesk 客户端做一轮冒烟测试

## 8. 关于“官方 Pro 完整能力”的说明

当前版本重点是兼容新版客户端主流程，以下能力仍不是官方 Pro 等价实现：

- 完整 OIDC 登录流程
- 官方插件签名服务
- 完整分组权限模型（`accessible` 相关接口）

如果你的发布目标是“主流程可用 + 自建可维护”，当前版本可用；
如果目标是“完全替代官方 Pro”，建议继续按需补齐以上能力。
