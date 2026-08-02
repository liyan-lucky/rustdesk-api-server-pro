# Backend

Go/Iris/Xorm 服务端，同时提供 RustDesk 客户端 API（`/api/*`）、管理端 API（`/admin/*`）、普通用户门户 API（`/user-portal/*`）和前端静态资源。

数据库模型变更后先运行 `go run . sync`。通讯录管理接口位于 `/admin/ab/*`：`GET /admin/ab/peers` 与 `GET /admin/ab/tags` 提供当前账户全量视图；`POST /admin/ab/shared/add` 允许管理员通过 `user_id` 为指定用户创建地址簿。`address_book.created_by_admin` 是删除权限的后端强制边界。

最低验证：

```bash
go test ./...
go vet ./...
```
