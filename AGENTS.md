# AGENTS.md - 项目开发指南

## 项目信息

- **项目名称**: RustDesk API Server Pro（兼容增强版）
- **技术栈**: Go + Vue3 + SQLite/MySQL
- **版本**: v1.2.29
- **主分支**: main
- **工作目录**: `E:\Visual_Studio_Code\07 Rustdesk-api-server-pro`

## 开发规则

### 通用规则
- 所有错误消息必须带 `ERR-xxxx` 编码前缀
- 鉴权中间件返回 401/406 时必须使用 `StopWithJSON`
- 用户使用简体中文交互
- 所有代码注释和文档使用简体中文
- 禁止提交敏感信息（密钥、密码、token等）

### Git 规则
- 主分支: `main`
- 提交前必须测试通过
- 提交信息格式: `<type>: <description>`
  - `feat`: 新功能
  - `fix`: 修复bug
  - `refactor`: 重构
  - `docs`: 文档更新
  - `test`: 测试相关
  - `chore`: 构建/工具相关

### 代码规范
- Go: 遵循 Go 官方规范，使用 `gofmt` 格式化
- Vue: 遵循 Vue3 Composition API 规范
- 禁止使用 `any` 类型（除非必要）
- 函数必须有文档注释

### 测试规范
- 新功能必须编写单元测试
- 修复bug必须添加回归测试
- 测试覆盖率目标: >70%

## 部署环境

### 远程设备
- **地址**: `ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net`
- **类型**: 群晖NAS
- **部署方式**: Docker
- **容器名**: `rustdesk-api-server-pro`
- **端口**: 16888
- **网络模式**: host

### 部署路径
- **数据目录**: `/vol1/docker/rustdesk-api-server-pro/`
- **更新脚本**: `/vol1/docker/rustdesk-api-server-pro/update-rustdesk-api.sh`
- **配置文件**: `/vol1/docker/rustdesk-api-server-pro/server.yaml`
- **数据库**: `/vol1/docker/rustdesk-api-server-pro/server.db`

### 部署命令
```bash
# 编译
cd backend && go build -o rustdesk-api-server.exe

# 上传
scp -P 2828 backend/rustdesk-api-server.exe LiYan@liyan-fnosnas.dynv6.net:/vol1/docker/rustdesk-api-server-pro/

# 重启容器
ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net 'docker restart rustdesk-api-server-pro'

# 查看日志
ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net 'docker logs rustdesk-api-server-pro --tail 100'
```

### Git 代理配置
```bash
git config --global http.proxy http://127.0.0.1:4088
```

## 项目结构

```
rustdesk-api-server-pro/
├── backend/                 # Go 后端
│   ├── app/
│   │   ├── controller/      # 控制器
│   │   │   ├── admin/       # 管理后台接口
│   │   │   └── api/         # RustDesk 客户端接口
│   │   ├── middleware/      # 中间件
│   │   ├── model/           # 数据模型
│   │   └── route.go         # 路由配置
│   ├── internal/
│   │   ├── service/         # 业务逻辑
│   │   ├── core/            # 核心功能
│   │   └── errcode/         # 错误码定义
│   ├── config/              # 配置
│   └── util/                # 工具函数
├── soybean-admin/           # Vue3 管理后台前端
│   ├── src/
│   │   ├── views/           # 页面组件
│   │   ├── router/          # 路由配置
│   │   └── service/         # API 服务
│   └── package.json
├── docs/                    # 文档
├── docker/                  # Docker 相关
└── VERSION                  # 版本号（单一事实来源）
```

## 关键模块

### 后端中间件
- `AdminAuth`: 管理员鉴权（要求 `is_admin=true`）
- `UserAuth`: 用户鉴权（允许 admin 和普通用户）
- `AdminOrUserAuth`: 管理员或用户鉴权（允许查看，操作需检查 `isAdmin`）

### OAuth 统一回调
- 回调URL: `/admin/auth/oauth/{provider}/callback`
- 客户端和admin共用回调端点
- 通过 `pollToken` 区分客户端登录和admin登录
- 客户端回调返回HTML页面，包含 `rustdesk://oauth/callback?poll_token=xxx` 链接

### 错误码规范
- ERR1xxx: 通用错误
- ERR2xxx: OAuth/OIDC 相关错误
- ERR22xx: 客户端 OAuth 相关错误
- 所有错误码定义在 `backend/internal/errcode/`

## 常见任务

### 添加新API
1. 在 `backend/app/controller/` 创建控制器
2. 在 `backend/app/route.go` 注册路由
3. 在 `backend/internal/service/` 实现业务逻辑
4. 添加单元测试
5. 更新文档

### 修改鉴权规则
1. 修改 `backend/app/middleware/auth.go`
2. 更新 `backend/app/route.go` 中的中间件配置
3. 测试所有受影响的API

### 更新前端
1. 修改 `soybean-admin/src/` 下的代码
2. 构建前端: `cd soybean-admin && npm run build`
3. 复制 `dist/` 到 `backend/dist/`
4. 重新编译后端

### 部署更新
1. 编译后端
2. 上传到设备
3. 重启容器
4. 查看日志确认启动成功

## 故障排查

### 查看日志
```bash
ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net 'docker logs rustdesk-api-server-pro --tail 200 -f'
```

### 检查容器状态
```bash
ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net 'docker ps -a | grep rustdesk-api-server-pro'
```

### 进入容器
```bash
ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net 'docker exec -it rustdesk-api-server-pro sh'
```

### 检查数据库
```bash
ssh -p 2828 LiYan@liyan-fnosnas.dynv6.net 'docker exec -it rustdesk-api-server-pro sqlite3 /app/data/server.db'
```

## 参考文档

- 项目详细描述: `docs/PROJECT_DESCRIPTION.md`
- 经验教训: `docs/LESSONS_LEARNED.md`
- 使用说明: `docs/USAGE.md`
- Docker 说明: `docs/DOCKER.md`
- 排障手册: `docs/TROUBLESHOOTING.md`
- OAuth Provider 规范: `docs/OAUTH_PROVIDERS.md`
