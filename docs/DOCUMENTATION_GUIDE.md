# 文档体系规范

本文档定义项目的文档结构、职责划分和维护规范，确保文档不重复、不遗漏、与项目现状同步。

## 文档分类

### 一、根目录文档

| 文档 | 职责 | 维护时机 |
|------|------|----------|
| `README.md` | 项目主文档（中文），面向用户和开发者 | 每次发布或重大变更 |
| `README_EN.md` | 项目主文档（英文） | 与 README.md 同步 |
| `AGENTS.md` | 项目开发指南，AI 代理和开发者使用 | 开发规则变更时 |
| `CHANGELOG.md` | 变更日志，按版本记录功能变更 | 每次版本发布 |
| `TODO.md` | 待办事项 | 持续更新 |
| `RELEASE_NOTES.md` | 发布说明 | 每次版本发布 |
| `VERSION` | 版本号文件（单一事实来源） | CI 自动递增 |

### 二、合规文档（根目录）

| 文档 | 职责 |
|------|------|
| `COMPLIANCE.md` | 合规总览 |
| `CONTRIBUTING.md` | 贡献指南 |
| `CODE_OF_CONDUCT.md` | 行为准则 |
| `DISCLAIMER.md` | 免责声明 |
| `PRIVACY.md` | 隐私政策 |
| `SECURITY.md` | 安全政策 |
| `SUPPORT.md` | 支持说明 |
| `THIRD_PARTY_NOTICES.md` | 第三方声明 |

### 三、docs/ 目录文档

| 文档 | 职责 | 维护时机 |
|------|------|----------|
| `README.md` | 文档索引 | 文档结构变更时 |
| `CURRENT_STATUS.md` | 当前仓库状态（架构、能力、待办、已知问题） | 每次开发后 |
| `LESSONS_LEARNED.md` | 经验教训与技术笔记 | 每次踩坑或技术决策后 |
| `PROJECT_DESCRIPTION.md` | 项目详细描述 | 架构变更时 |
| `PROJECT_WORKFLOW.md` | 项目协作规范 | 流程变更时 |
| `USAGE.md` | 使用说明 | 功能变更时 |
| `DOCKER.md` | Docker 部署 | 部署变更时 |
| `OPENWRT_ONE_CONTAINER.md` | OpenWrt 部署 | 部署变更时 |
| `PORTS.md` | 端口说明 | 端口变更时 |
| `TROUBLESHOOTING.md` | 排障手册 | 新问题发现时 |
| `OAUTH_PROVIDERS.md` | OAuth 规范 | Provider 变更时 |
| `API_MATRIX.md` | API 兼容矩阵 | API 变更时 |
| `OFFICIAL_API_COMPATIBILITY_PLAN.md` | API 兼容计划 | 兼容策略变更时 |
| `AUTO_COMPAT_UPDATE_DESIGN.md` | 自动兼容设计 | 设计变更时 |
| `AUDIT_LOG_DESIGN.md` | 审计日志设计 | 设计变更时 |
| `FULL_LOGIC_AUDIT_REPORT.md` | 逻辑审计报告 | 审计完成后 |
| `RELEASE_PROCESS.md` | 发布流程 | 流程变更时 |
| `SECRET_HANDLING.md` | 密钥处理 | 安全策略变更时 |
| `LEGAL_COMPLIANCE.md` | 合规指南 | 合规要求变更时 |
| `CLEANUP_HISTORY.md` | 清理历史 | 每次清理后 |
| `V3_QUALITY.md` | 前端质量门禁 | 质量标准变更时 |
| `NEXT_CONVERSATION_PROMPT.md` | 续接提示词 | 每次开发后 |
| `compat/rustdesk-current.json` | 兼容清单 | 版本变更时 |

### 四、子项目文档

| 文档 | 职责 |
|------|------|
| `backend/README.md` | 后端说明 |
| `soybean-admin/README.md` | 前端说明 |
| `soybean-admin/docs/i18n-coverage-report.md` | i18n 覆盖率报告 |
| `soybean-admin/docs/i18n-fallback-todo.md` | i18n 回退待办 |
| `soybean-admin/docs/I18N_V3.md` | i18n V3 规范 |

## 维护规范

### 1. 单一职责原则
- 每个文档只负责一个明确的职责
- 状态信息统一在 `CURRENT_STATUS.md`，不分散到多个文档
- 经验教训统一在 `LESSONS_LEARNED.md`，不分散到多个文档
- 开发指南统一在 `AGENTS.md`，不另建快速开始文档

### 2. 版本号同步
- `VERSION` 文件是版本号的单一事实来源
- 所有文档中引用的版本号必须与 `VERSION` 一致
- `docs/compat/rustdesk-current.json` 中的 `compat_server_version` 必须同步
- `CURRENT_STATUS.md` 中的版本号必须同步

### 3. 文档更新时机
- 新增功能：更新 `CHANGELOG.md`、`CURRENT_STATUS.md`、相关专项文档
- 修复 bug：更新 `CHANGELOG.md`、`CURRENT_STATUS.md`（如影响已知问题）
- 踩坑经验：更新 `LESSONS_LEARNED.md`
- 架构变更：更新 `PROJECT_DESCRIPTION.md`、`CURRENT_STATUS.md`
- 部署变更：更新 `DOCKER.md`、`CURRENT_STATUS.md`、`AGENTS.md`

### 4. 文档引用规范
- 文档间引用使用相对路径
- 删除文档时必须同步更新所有引用
- 新增文档时必须在 `docs/README.md` 索引中注册

### 5. 语言规范
- 所有文档使用简体中文（除 `README_EN.md` 和合规文档外）
- 代码注释使用简体中文
- 错误消息使用简体中文并带 `ERR-xxxx` 编码前缀

## 已废弃文档

以下文档已合并/删除，不再维护：

| 已删除文档 | 合并到 | 原因 |
|------------|--------|------|
| `PROJECT_STATUS.md` | `docs/CURRENT_STATUS.md` | 状态信息重复 |
| `TECHNICAL_NOTES.md` | `docs/LESSONS_LEARNED.md` | 经验笔记分散 |
| `QUICK_START.md` | `AGENTS.md` | 内容已被覆盖 |
| `README_CN.md` | `README.md` | 仅入口页，冗余 |
