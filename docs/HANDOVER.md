# 交接文档

> 本文档供新会话快速了解项目当前状态，避免重复探索。每次重要变更后应同步更新。

## 最近一次变更

- **日期**：2026-08-14
- **版本**：1.3.1（VERSION 文件为唯一事实来源）
- **变更内容**：修复普通用户通过收藏访问 `/#/403` 后点击"返回首页"死循环的问题

## 本次修复详情

### 问题
- **现象**：普通用户（R_USER 角色）通过浏览器收藏访问 `http://<host>:<port>/#/403` 时，点击"返回首页"按钮会跳转到 `/home`（仅 R_SUPER 可访问）→ 无权限 → 再次跳转到 403，形成死循环
- **根因**：`/403` 是 constant 路由，访问时不初始化 auth route，导致 root 路由的 redirect 未根据角色更新（普通用户应 redirect 到 `/user/profile`）
- **修复文件**：`soybean-admin/src/components/common/exception-base.vue`
- **修复方案**：在"返回首页"按钮的点击处理中，如果用户已登录但 auth route 未初始化，先调用 `authStore.initUserInfo()` 和 `routeStore.initAuthRoute()` 初始化路由，确保 root redirect 已根据角色更新后再跳转
- **验证**：`pnpm typecheck` 通过，`vite build` 通过，`go test ./...` 通过，`go vet ./...` 通过

## 项目快速概览

### 基本信息
- **项目名称**：RustDesk API Server Pro（兼容增强版）
- **技术栈**：Go + Vue3 + SQLite/MySQL
- **版本**：1.3.1
- **主分支**：main
- **工作目录**：`E:\Visual_Studio_Code\07_Rustdesk-api-server-pro`

### 架构
- 单 HTTP 端口对外提供服务
  - RustDesk 客户端 API：`/api/*`
  - 管理后台接口：`/admin/*`
  - 用户门户接口：`/user-portal/*`
  - 管理后台前端：`/`
  - plugin-sign 兼容：`/lic/web/api/plugin-sign`

### 关键目录
- `backend/`：Go 后端（Iris v12 + Xorm + SQLite/MySQL）
- `soybean-admin/`：Vue3 管理后台（Vite + Naive UI + Pinia）
- `docs/`：项目文档
- `Dockerfile`：多阶段构建（Go + Node + Alpine）
- `VERSION`：版本号（单一事实来源，CI 自动递增 PATCH）

### 鉴权中间件
- `AdminAuth`：管理员鉴权（`is_admin=true`）
- `UserAuth`：用户鉴权（允许 admin 和普通用户）
- `AdminOrUserAuth`：管理员或用户鉴权（查看权限，操作需检查 `isAdmin`）
- `ApiAuth`：客户端 API 鉴权（`is_admin=false`）

### 前端路由关键点
- 使用 hash 路由模式（`VITE_ROUTER_HISTORY_MODE=hash`）
- 静态路由模式（`VITE_AUTH_ROUTE_MODE=static`）
- 超级管理员角色：`R_SUPER`（`VITE_STATIC_SUPER_ROLE=R_SUPER`）
- 普通用户角色：`R_USER`
- 首页路由：`VITE_ROUTE_HOME=home`（超级管理员），普通用户动态更新为 `user_profile`
- constant 路由（如 403/404/500）不需要登录即可访问

## 开发与构建

### 前端
```bash
cd soybean-admin
pnpm install        # 安装依赖
pnpm typecheck      # 类型检查（vue-tsc --noEmit --skipLibCheck）
pnpm build          # 生产构建（vite build --mode prod）
# 构建产物在 soybean-admin/dist/
```

### 后端
```bash
cd backend
go test ./...       # 单元测试
go vet ./...        # 静态检查
go build -o rustdesk-api-server-pro.exe .  # Windows 编译
# 交叉编译 Linux
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o rustdesk-api-server-pro .
```

### 部署前端到后端
```bash
# 清理 backend/dist 并复制前端构建产物
Remove-Item backend\dist\* -Recurse -Force
Copy-Item soybean-admin\dist\* backend\dist -Recurse -Force
```

### 完整构建流程
1. `cd soybean-admin && pnpm typecheck && pnpm build`
2. 复制 `soybean-admin/dist/*` 到 `backend/dist/`
3. `cd backend && go build` 或交叉编译 Linux 版本
4. 部署到容器：`docker cp /tmp/rustdesk-api-server-pro rustdesk-api-server-pro:/app/rustdesk-api-server-pro`
5. `docker exec rustdesk-api-server-pro chmod 755 /app/rustdesk-api-server-pro`
6. `docker restart rustdesk-api-server-pro`
7. `docker logs rustdesk-api-server-pro --tail 100` 确认启动成功

## 部署环境

- **远程设备**：NAS / Linux Docker 主机
- **容器名**：`rustdesk-api-server-pro`
- **默认端口**：12345
- **网络模式**：host
- **外部访问**：`http://liyan-fnosnas.dynv6.net:16888/`

## 关键文档索引

| 文档 | 用途 |
|------|------|
| `README.md` | 项目总览、快速开始、配置说明 |
| `docs/PROJECT_DESCRIPTION.md` | 项目详细描述（架构、API、数据模型） |
| `docs/CURRENT_STATUS.md` | 当前能力与边界 |
| `docs/LESSONS_LEARNED.md` | 经验教训与技术笔记 |
| `docs/TROUBLESHOOTING.md` | 排障手册 |
| `docs/USAGE.md` | 使用说明（部署、升级、验证） |
| `docs/TODO.md` | 待办事项 |
| `docs/OAUTH_PROVIDERS.md` | OAuth Provider 规范 |
| `docs/FULL_LOGIC_AUDIT_REPORT.md` | 全仓功能逻辑审计报告 |
| `AGENTS.md` | AI 开发指南 |

## 已知限制（非 bug）

- plugin-sign 是兼容占位实现，不是官方真实签名服务
- `/api/devices/deploy` 返回 `NOT_ENABLED`，不写入 hbbs 部署白名单
- 企业分组、策略、accessible 权限是最小兼容模型，不是完整官方 Pro 权限模型
- 录屏上传只负责文件落盘，没有完整索引、归档、清理和权限控制
- `/api/oidc/*` 已实现客户端兼容协议，复用 `/api/oauth/*` 的服务端回调+轮询逻辑

## 常见问题快速排查

| 现象 | 原因 | 处理 |
|------|------|------|
| `/api/*` 正常但首页 404 | `httpConfig.staticdir` 指向错误 | 确认指向 `/app/dist` |
| 升级后页面报 SQL 字段不存在 | 数据库未执行 `sync` | 执行 `sync` 并重启 |
| OAuth 登录闪退 | 旧版前端缓存或 state 过期 | 强制刷新浏览器，检查回调地址 |
| 设备频繁离线 | 心跳间隔与 `deviceCheckJob.duration` 不匹配 | 调整配置 |
| 普通用户访问 403 后无法返回首页 | **已修复**（本次变更） | 升级到最新版本 |

## Git 规则

- 提交信息格式：`<type>: <description>`
  - `feat`/`fix`/`refactor`/`docs`/`test`/`chore`
- 提交前必须测试通过
- 禁止提交敏感信息

## 下一步建议

1. 部署本次修复到线上并验证
2. 完成 `docs/TODO.md` 中的高优先级待办（客户端 OAuth/WebAuth 验收、普通用户权限测试）
3. 补充单元测试覆盖率到 >70%
