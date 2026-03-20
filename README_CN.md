# 中文说明入口

中文主文档请优先查看 [README.md](./README.md)。

当前仓库的中文说明重点包括：

- 管理后台多语言优化
- 服务器配置入口从首页调整到左侧菜单
- 第三方登录支持骨架
  支持 `oidc`、`google`、`github`
- 旧 `oidc` 配置兼容，新项目推荐使用 `oauth.providers`

常用文档入口：

- 中文主文档：[README.md](./README.md)
- English: [README_EN.md](./README_EN.md)
- 使用说明：`docs/USAGE.md`
- Docker 说明：`docs/DOCKER.md`
- 排障文档：`docs/TROUBLESHOOTING.md`

第三方登录回调地址规则：

- 多 provider：`/admin/auth/oauth/<provider>/callback`
- 旧版 OIDC：`/admin/auth/oidc/callback`

如果你的前端是通过反向代理单独托管 `dist`，在启用第三方登录前请重新构建前端并替换静态资源目录。
