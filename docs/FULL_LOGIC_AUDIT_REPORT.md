# Rustdesk-api-server-pro 全功能逻辑审计报告

**审计日期**: 2026-08-07  
**审计范围**: 后端全部控制器/服务/中间件 + 前端路由/状态/API/页面  
**审计版本**: v1.2.16  

---

## 一、审计总览

| 模块 | Critical | High | Medium | Low | 小计 |
|------|----------|------|--------|-----|------|
| M1-认证授权 | 2 | 4 | 3 | 10 | 19 |
| M2-用户与设备 | 5 | 7 | 7 | 4 | 23 |
| M3-日志与监控 | 2 | 5 | 9 | 8 | 24 |
| M4-配置与系统 | 6 | 8 | 16 | 10 | 40 |
| M5-前端逻辑 | 7 | 16 | 27 | 12 | 62 |
| **合计** | **22** | **40** | **62** | **44** | **168** |

> 注: M5中去除了3个与M1-M4重复的问题，实际独立问题约165个。

---

## 二、Critical 问题清单（22个）

### 后端 - 认证授权（M1）

| ID | 文件:行 | 描述 | 影响 |
|----|---------|------|------|
| M1-C1 | `middleware/auth.go` | AdminAuth isAdmin回退：非管理员token通过AdminAuth时isAdmin默认设为true | 权限提升：普通用户获得管理员权限 |
| M1-C2 | `middleware/auth.go` | AdminAuth与ApiAuth的token提取方式不一致（Bearer前缀） | 部分请求认证失败 |

### 后端 - 用户与设备（M2）

| ID | 文件:行 | 描述 | 影响 |
|----|---------|------|------|
| M2-C1 | `middleware/auth.go` | 同M1-C1，AdminAuth回退权限提升 | 权限提升 |
| M2-C2 | `controller/admin/users.go` | 用户CRUD缺权限检查 | 普通用户可增删改用户 |
| M2-C3 | `controller/admin/users.go` | 用户删除不级联且可删自己 | 数据残留/自我删除导致系统不可用 |
| M2-C4 | `repository/xorm_system_repository.go` | 设备心跳不更新updated_at | 离线检测完全失效 |
| M2-C5 | `controller/admin/devices.go` | 设备禁用不断开已建立连接 | 禁用设备仍可连接 |

### 后端 - 日志与监控（M3）

| ID | 文件:行 | 描述 | 影响 |
|----|---------|------|------|
| M3-C1 | `controller/admin/mail_template.go` | 邮件模板不存在时nil指针panic | 服务崩溃 |
| M3-C2 | `controller/admin/audit.go` | 文件传输审计查询用了错误表前缀 | SQL错误，功能不可用 |

### 后端 - 配置与系统（M4）

| ID | 文件:行 | 描述 | 影响 |
|----|---------|------|------|
| M4-C1 | `controller/admin/token.go` | HandleClear未保护当前管理员token | 管理员清空自己的token导致登出 |
| M4-C2 | `controller/admin/sessions.go` | HandleKill可终止自己且无审计 | 自我终止导致管理中断 |
| M4-C3 | `controller/admin/sessions.go` | HandleKill用RemoveElement(ids,1)硬编码逻辑错误 | 保护当前session逻辑失效 |
| M4-C4 | `controller/admin/sessions.go` | HandleClear与HandleList范围不一致 | 清除范围超出预期 |
| M4-C5 | `config/config.go` | 配置写入非真原子(Windows) | 配置文件可能被清空 |
| M4-C6 | `config/config.go` | GetServerConfig自动写入默认配置 | 可能覆盖已有配置 |

### 前端 - 路由与权限（M5）

| ID | 文件:行 | 描述 | 影响 |
|----|---------|------|------|
| M5-C1 | `hooks/common/router.ts:97` | redirectFromLogin开放重定向漏洞 | 钓鱼攻击/OAuth令牌窃取 |
| M5-C2 | `login/modules/pwd-login.vue:84` | OAuth登录URL传递redirect未校验 | OAuth开放重定向 |
| M5-C3 | `store/modules/auth/index.ts:43` | $reset恢复旧Token导致登出后状态不一致 | 登出后仍显示已登录 |
| M5-C4 | `router/guard/route.ts:44` | 已登录用户OAuth票据丢失 | OAuth绑定流程中断 |
| M5-C5 | `store/modules/auth/index.ts:65` | Token未在Bootstrap失败时清理 | 半登录状态 |
| M5-C6 | `router/elegant/routes.ts:294` | system_server路由缺roles限制 | 普通用户可修改服务器配置 |
| M5-C7 | `router/elegant/routes.ts:122` | audit_loginlogs路由缺roles限制 | 普通用户可查看所有登录日志 |

---

## 三、High 问题清单（40个）

### 后端 - 认证授权（M1: 4个）
- M1-H1: 并发map无锁读取panic (`service/oauth_provider_service.go`)
- M1-H2: OIDC内存存储多实例不可用 (`service/oidc_auth_service.go`)
- M1-H3: OAuth用户绑定TOCTOU竞态 (`service/oauth_provider_service.go`)
- M1-H4: OAuthAccount缺(provider,subject)唯一约束

### 后端 - 用户与设备（M2: 7个）
- M2-H1: 设备心跳并发竞态产生重复记录
- M2-H2: 2FA验证码检查逻辑错误
- M2-H3: 邮箱验证码重放攻击
- M2-H4: 2FA码永不过期
- M2-H5: 分页参数无边界校验除零panic
- M2-H6: 地址簿权限检查不一致
- M2-H7: 地址簿规则更新信息泄露

### 后端 - 日志与监控（M3: 5个）
- M3-H1: 分页TotalPage计算错误(整除多1页)
- M3-H2: 分页参数无边界校验
- M3-H3: MailService单例模式失效
- M3-H4: 容器日志中间件每请求创建goroutine无缓冲
- M3-H5: 日志表无自动轮转磁盘无限增长

### 后端 - 配置与系统（M4: 8个）
- M4-H1: 配置热更新不一致(c.Cfg vs GetServerConfig)
- M4-H2: OAuth Provider重名校验失效
- M4-H3: secret保留逻辑可被绕过
- M4-H4: HandleDeleteProvider不校验name格式
- M4-H5: HandleDeleteAccount不校验id范围
- M4-H6: SSRF风险(探测不限制内网IP)
- M4-H7: extractErrCode字符串解析脆弱
- M4-H8: Error启动goroutine无panic保护

### 前端 - 路由与权限（M5: 10个）
- M5-H1: resetStore在导航守卫内调用router.push引发竞态
- M5-H2: initUserInfo未await resetStore导致竞态
- M5-H3: initAuthRoute无并发保护
- M5-H4: 请求实例handleLogout未await resetStore
- M5-H5: clearAuthStorage未清理isAdmin/userType/globalTabs
- M5-H6: audit_loginlogs对非管理员可访问（同M5-C7，路由层面）
- M5-H7: system_server对非管理员可访问（同M5-C6，路由层面）
- M5-H8: 无Token刷新机制（假死状态）
- M5-H9: OAuth票据可重放/未单次消费
- M5-H10: OIDC票据成功后未清理URL

### 前端 - API与页面（M5: 6个）
- M5-H11: API函数大量使用any类型（address-book/audit/devices等）
- M5-H12: 用户编辑表单密码输入框未设type="password"
- M5-H13: user-login.vue绕过Auth Store直接window.location.href跳转
- M5-H14: pwd-login.vue OAuth登录双击竞态条件
- M5-H15: oauth/index.vue保存配置无表单验证
- M5-H16: workspace系列页面load缺try/finally，loading可能卡死
- M5-H17: card-data.vue卡片详情加载竞态条件
- M5-H18: config-operation.vue Clipboard实例未销毁内存泄漏

---

## 四、修复优先级排序

### P0 - 立即修复（安全漏洞，影响生产）

1. **M1-C1/M2-C1**: AdminAuth isAdmin回退权限提升 → 后端
2. **M5-C6/M5-H7**: system_server路由缺roles → 前端
3. **M5-C7/M5-H6**: audit_loginlogs路由缺roles → 前端
4. **M5-C1**: redirectFromLogin开放重定向 → 前端
5. **M5-C2**: OAuth登录URL开放重定向 → 前端
6. **M5-C3**: $reset恢复旧Token → 前端
7. **M5-H5**: clearAuthStorage未清理isAdmin → 前端
8. **M2-C2**: 用户CRUD缺权限检查 → 后端
9. **M4-C1**: HandleClear未保护当前管理员 → 后端
10. **M4-C3**: HandleKill RemoveElement逻辑错误 → 后端

### P1 - 尽快修复（功能缺陷/数据安全）

11. **M5-C5**: Token未在Bootstrap失败时清理 → 前端
12. **M5-C4**: 已登录用户OAuth票据丢失 → 前端
13. **M2-C4**: 设备心跳不更新updated_at → 后端
14. **M2-C3**: 用户删除不级联且可删自己 → 后端
15. **M3-C1**: 邮件模板nil指针panic → 后端
16. **M3-C2**: 文件传输审计错误表前缀 → 后端
17. **M4-C5**: 配置写入非真原子 → 后端
18. **M5-H12**: 密码输入框明文显示 → 前端
19. **M5-H16**: workspace页面loading卡死 → 前端
20. **M5-H13**: user-login.vue绕过Auth Store → 前端

### P2 - 计划修复（健壮性/体验）

21-40: 其余High问题
41-168: Medium和Low问题

---

## 五、前端修复方案（可立即执行）

### 修复1: 路由权限缺失（M5-C6, M5-C7）

**文件**: `soybean-admin/src/router/elegant/routes.ts`

```diff
// audit_loginlogs 添加 roles
{
  name: 'audit_loginlogs',
  meta: {
+   roles: ['R_SUPER'],
    ...
  }
}

// system_server 添加 roles
{
  name: 'system_server',
  meta: {
+   roles: ['R_SUPER'],
    ...
  }
}
```

### 修复2: 开放重定向（M5-C1）

**文件**: `soybean-admin/src/hooks/common/router.ts`

```diff
async function redirectFromLogin() {
  const redirect = route.value.query?.redirect as string;
-  if (redirect) {
+  if (redirect && redirect.startsWith('/') && !redirect.startsWith('//')) {
    return routerPush(redirect);
  }
  return toHome();
}
```

### 修复3: $reset恢复旧Token（M5-C3）

**文件**: `soybean-admin/src/store/modules/auth/index.ts`

```diff
async function resetStore() {
  clearAuthStorage();
-  authStore.$reset();
+  token.value = '';
+  Object.assign(userInfo, { userId: '', userName: '', roles: [], buttons: [] });
+  isLogin.value = false;
  if (!route.meta.constant) {
    await toLogin();
  }
  tabStore.cacheTabs();
  routeStore.resetStore();
}
```

### 修复4: clearAuthStorage未清理全部（M5-H5）

**文件**: `soybean-admin/src/store/modules/auth/shared.ts`

```diff
export function clearAuthStorage() {
  localStg.remove('token');
  localStg.remove('refreshToken');
+  localStg.remove('isAdmin');
+  localStg.remove('userType');
+  localStg.remove('globalTabs');
}
```

### 修复5: Token未在Bootstrap失败时清理（M5-C5）

**文件**: `soybean-admin/src/store/modules/auth/index.ts`

```diff
async function applyTokenAndBootstrap(loginToken) {
  localStg.set('token', loginToken.token);
  const pass = await getUserInfo();
  if (pass) {
    token.value = loginToken.token;
    return true;
  }
+  clearAuthStorage();
  return false;
}

async function login(model, redirect = true) {
  ...
  if (!error) {
    const pass = await applyTokenAndBootstrap(loginToken);
    if (pass) {
      ...
    }
+   else {
+     resetStore();
+   }
  } else {
    resetStore();
  }
  ...
}
```

### 修复6: 密码输入框明文（M5-H12）

**文件**: `soybean-admin/src/views/user/list/components/edit.vue`

```diff
<NInput
  v-model:value="model.password"
+ type="password"
+ show-password-on="click"
  :placeholder="$t('page.user.list.inputPassword')"
/>
```

### 修复7: workspace页面loading卡死（M5-H16）

**文件**: `soybean-admin/src/views/workspace/sessions/index.vue`（及devices/security/overview）

```diff
async function load() {
+  try {
    loading.value = true;
    const { data: r } = await fetchMySessions({ ... });
    if (r) {
      data.value = r.records;
      total.value = r.total;
    }
+  } finally {
    loading.value = false;
+  }
}
```

### 修复8: OAuth登录双击竞态（M5-H14）

**文件**: `soybean-admin/src/views/_builtin/login/modules/pwd-login.vue`

```diff
async function handleOAuthLogin(provider: Api.Auth.OAuthProvider) {
+  if (activeProvider.value) return;
  activeProvider.value = provider.name;
  ...
}
```

---

## 六、后端修复方案（需CI/CD部署）

### 修复1: AdminAuth权限提升（M1-C1）

**文件**: `backend/app/middleware/auth.go`

核心问题：非管理员token通过AdminAuth时，isAdmin应设为false而非true。需检查具体代码确认修复方案。

### 修复2: 用户CRUD权限检查（M2-C2）

**文件**: `backend/app/controller/admin/users.go`

所有用户管理接口应检查当前用户是否为管理员。

### 修复3: HandleKill逻辑错误（M4-C3）

**文件**: `backend/app/controller/admin/sessions.go`

RemoveElement(ids, 1)应改为RemoveElement(ids, currentSessionID)。

### 修复4: 设备心跳更新updated_at（M2-C4）

**文件**: `backend/internal/repository/xorm_system_repository.go`

心跳更新时应同时更新updated_at字段。

---

## 七、建议的修复批次

### 批次1: 前端安全修复（可立即构建部署）
- M5-C6: system_server路由roles
- M5-C7: audit_loginlogs路由roles
- M5-C1: redirectFromLogin开放重定向
- M5-C2: OAuth登录URL校验
- M5-C3: $reset恢复旧Token
- M5-H5: clearAuthStorage清理
- M5-C5: Bootstrap失败清理
- M5-H12: 密码输入框type
- M5-H16: workspace loading try/finally
- M5-H14: OAuth双击竞态

### 批次2: 后端安全修复（需CI/CD）
- M1-C1: AdminAuth权限提升
- M2-C2: 用户CRUD权限
- M4-C1: HandleClear保护
- M4-C3: HandleKill逻辑
- M2-C4: 设备心跳updated_at
- M3-C1: 邮件模板nil指针
- M3-C2: 文件传输审计表前缀

### 批次3: 前端功能修复
- M5-C4: OAuth票据丢失
- M5-H13: user-login.vue统一登录
- M5-H15: OAuth配置表单验证
- M5-H8: Token刷新机制
- M5-H9/H10: OAuth票据清理

### 批次4: 后端功能修复
- M2-C3: 用户删除级联
- M2-C5: 设备禁用断开
- M4-C5: 配置原子写入
- 其余High问题

---

**报告生成时间**: 2026-08-07  
**下一步**: 按批次1执行前端安全修复，构建并部署验证
