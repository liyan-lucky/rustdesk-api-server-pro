# 变更日志

## v1.2.29 (2026-08-07)

### 新功能

#### 1. 普通用户查看权限增强
- **目标**: 普通用户可以查看日志审计和系统设置信息（只读），但不能执行操作
- **实现**:
  - 新建 `AdminOrUserAuth` 中间件（`backend/app/middleware/auth.go:145`）
  - 修改路由：audit/system 控制器移到 AdminOrUserAuth（`backend/app/route.go`）
  - 15个操作类API添加 isAdmin 检查，非admin返回401
  - 前端路由 roles 修改：audit/system 子路由改为 `['R_SUPER', 'R_USER']`（oauth/tokens除外）
  - 前端操作按钮添加 isAdmin 控制：所有清除/编辑/添加按钮

#### 2. 客户端OAuth统一回调
- **目标**: 统一回调URL，让客户端和admin共用回调端点
- **实现**:
  - `resolveClientCallbackURL` 使用 admin 回调URL
  - `ConsumeUnifiedCallback` 统一回调处理函数
  - `HandleOauthCallback` 根据 pollToken 返回HTML页面或302重定向
  - OAuth回调页面显示按钮链接到 `rustdesk://oauth/callback?poll_token=xxx`
  - 添加HTML注释调试信息

### 修改的文件

#### 后端
- `backend/app/middleware/auth.go` - 新增 AdminOrUserAuth 中间件
- `backend/app/route.go` - 修改路由中间件配置
- `backend/app/controller/admin/auth.go` - 修改 HandleOauthCallback 和 renderOAuthCallbackPage
- `backend/internal/service/oauth_provider_service.go` - 新增 ConsumeUnifiedCallback 和 resolveClientCallbackURL
- `backend/app/controller/admin/audit.go` - 添加 isAdmin 检查
- `backend/app/controller/admin/dashboard.go` - 添加 isAdmin 检查
- `backend/app/controller/admin/mail_template.go` - 添加 isAdmin 检查
- `backend/app/controller/admin/mail_logs.go` - 添加 isAdmin 检查
- `backend/app/controller/admin/token.go` - 添加 isAdmin 检查
- `backend/app/controller/admin/oauth.go` - 添加 isAdmin 检查
- `backend/app/controller/admin/security_audit.go` - 添加 isAdmin 检查
- `backend/app/controller/admin/error_log.go` - 添加 isAdmin 检查
- `backend/app/controller/admin/container_log.go` - 添加 isAdmin 检查

#### 前端
- `soybean-admin/src/router/elegant/routes.ts` - 修改路由 roles 配置
- `soybean-admin/src/views/audit/*/components/table-header.vue` - 添加清除按钮 isAdmin 控制
- `soybean-admin/src/views/audit/system-logs/index.vue` - 添加清除按钮 isAdmin 控制
- `soybean-admin/src/views/system/mail_template/index.vue` - 添加编辑列 isAdmin 控制
- `soybean-admin/src/views/system/mail_template/components/table-header.vue` - 添加添加按钮 isAdmin 控制
- `soybean-admin/src/views/home/modules/server-connection-config.vue` - 添加编辑按钮 isAdmin 控制

#### 文档
- `VERSION` - 更新到 1.2.29
- `README.md` - 更新最近更新
- `AGENTS.md` - 新建项目开发指南
- `CHANGELOG.md` - 新建变更日志
- `docs/CURRENT_STATUS.md` - 更新为统一状态文档（合并原 PROJECT_STATUS.md）
- `docs/LESSONS_LEARNED.md` - 更新为统一经验文档（合并原 TECHNICAL_NOTES.md）

### 测试状态
- ✅ 已部署到设备
- ✅ 普通用户查看权限功能正常
- ⏳ 客户端OAuth回调等待用户测试确认

### 已知问题
- 客户端OAuth回调：浏览器安全提示显示当前页面URL，需要用户确认这是正常行为

## v1.2.28 (2026-08-06)

### 新功能
- OAuth统一回调基础实现
- 客户端和admin共用回调端点

### 修改的文件
- `backend/internal/service/oauth_provider_service.go` - 新增统一回调处理
- `backend/app/controller/admin/auth.go` - 修改回调处理逻辑

## v1.2.27 (2026-08-05)

### 新功能
- AdminOrUserAuth中间件基础实现

### 修改的文件
- `backend/app/middleware/auth.go` - 新增中间件
- `backend/app/route.go` - 应用中间件

## 历史版本

更早的版本变更请查看 Git 提交历史。
