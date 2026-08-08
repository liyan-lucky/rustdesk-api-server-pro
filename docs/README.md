# RustDesk API Server Pro 文档索引

## 文档体系

文档结构规范详见 [DOCUMENTATION_GUIDE.md](./DOCUMENTATION_GUIDE.md)。

## 核心文档

1. `CURRENT_STATUS.md`：当前仓库状态、架构、能力边界、待办事项、已知问题和部署历史。
2. `PROJECT_DESCRIPTION.md`：项目定位、系统组成、接口边界、数据模型和维护重点。
3. `PROJECT_WORKFLOW.md`：项目结构、主分支开发、验证、提交和发布规范。
4. `LESSONS_LEARNED.md`：经验教训与技术笔记（含 OAuth 回调、中间件、错误码、部署流程、性能/安全建议）。
5. `DOCUMENTATION_GUIDE.md`：文档体系规范，定义文档职责划分和维护规范。

## 部署与运维

6. `DOCKER.md`：Docker / Docker Compose 安装、配置、升级和排查。
7. `OPENWRT_ONE_CONTAINER.md`：OpenWrt x86 / 软路由一体化部署与对齐更新。
8. `PORTS.md`：端口与访问路径说明。
9. `TROUBLESHOOTING.md`：常见问题与排查。
10. `USAGE.md`：使用说明。

## 开发与设计

11. `OAUTH_PROVIDERS.md`：OAuth Provider 规范。
12. `API_MATRIX.md`：API 兼容矩阵。
13. `OFFICIAL_API_COMPATIBILITY_PLAN.md`：官方 API 兼容计划。
14. `AUTO_COMPAT_UPDATE_DESIGN.md`：自动兼容更新设计。
15. `AUDIT_LOG_DESIGN.md`：审计日志设计。
16. `FULL_LOGIC_AUDIT_REPORT.md`：逻辑审计报告。
17. `RELEASE_PROCESS.md`：发布流程。
18. `SECRET_HANDLING.md`：密钥处理规范。
19. `V3_QUALITY.md`：前端质量门禁。

## 合规与清理

20. `LEGAL_COMPLIANCE.md`：合规指南。
21. `CLEANUP_HISTORY.md`：清理历史记录。

## 辅助

22. `NEXT_CONVERSATION_PROMPT.md`：将当前任务转交到新对话时可直接粘贴的续接提示词。
23. `compat/rustdesk-current.json`：RustDesk 客户端兼容清单。

## 当前主事实

- 当前主工作分支：`main`。
- 本地开发、提交、推送和正式发布统一在 `main` 进行；远端仅额外保留只读快照用途的 `backup`，不保留其他长期或临时分支。
- 当前备份分支：`backup`，由 `.github/workflows/force-backup-main.yml` 手动输入 `YES` 后强制覆盖。
- 当前推荐部署方式：单容器一体化服务。
- 当前默认端口：`12345/tcp`，以 `server.yaml` 的 `httpConfig.port` 为准。
- 当前镜像内置管理后台前端，旧 `rustdesk-web` / nginx 前端容器不再是必需组件。
- 当前版本：`v1.2.29`（由 `VERSION` 文件控制）。
- 当前第三方登录统一使用 `oauth.providers`；GitHub 主流程、安全回调、Secret 掩码提示已实现。测试设备仍存在 `github.com:443` 出站超时；微信/QQ 尚属计划。详见 [OAUTH_PROVIDERS.md](./OAUTH_PROVIDERS.md)。
- 当前通讯录交互统一为账户级全量列表、地址簿名称列和表头筛选；地址簿/联系人/标签及邮件模板新增型页面提供 CSV 导入导出。
- 地址簿归属由后端强制校验：管理员可代用户创建，普通用户只能自建且不能删除管理员代建项。
- "关于与更新"页面的在线版本检查地址允许按发布站点修改，配置仅保存在浏览器。

## 文档维护要求

- 部署方式、端口、Docker/OpenWrt 脚本、OAuth 行为、数据库或默认配置变化时，必须同步更新 `CURRENT_STATUS.md`、根 README 和相关专项文档。
- 不要把生产数据库、密钥、token、真实账号密码、OAuth secret 或生产配置提交到仓库。
- OpenWrt / Docker 命令应保持 host 网络、`/mnt/docker` 数据目录、中文 label 和端口 label 风格。
- 版本号变更时必须同步更新 `VERSION`、`CURRENT_STATUS.md`、`compat/rustdesk-current.json`。
- 新增文档时必须在本索引和 `DOCUMENTATION_GUIDE.md` 中注册。
- 删除文档时必须同步更新所有引用。
