package errcode

import (
	"errors"
	"fmt"
	"sync"
)

type Entry struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Module      string `json:"module"`
	Description string `json:"description"`
	Solution    string `json:"solution"`
}

var (
	mu      sync.RWMutex
	entries = make(map[string]Entry)
)

func register(code, message, module, description, solution string) Entry {
	e := Entry{Code: code, Message: message, Module: module, Description: description, Solution: solution}
	mu.Lock()
	entries[code] = e
	mu.Unlock()
	return e
}

func New(code, message string) error {
	return errors.New(code + ": " + message)
}

func Errorf(code, format string, args ...interface{}) error {
	return fmt.Errorf(code+": "+format, args...)
}

func List() []Entry {
	mu.RLock()
	defer mu.RUnlock()
	result := make([]Entry, 0, len(entries))
	for _, e := range entries {
		result = append(result, e)
	}
	return result
}

// ── Authentication (ERR-1xxx) ────────────────────────────────────────

var (
	ERR1001 = register("ERR-1001", "CaptchaError", "auth", "管理员登录验证码错误", "检查验证码输入是否正确，或刷新验证码后重试")
	ERR1002 = register("ERR-1002", "UserNotExists", "auth", "管理员登录用户不存在", "确认用户名是否正确，或联系管理员创建账号")
	ERR1003 = register("ERR-1003", "UsernameOrPasswordError", "auth", "管理员登录用户名或密码错误", "检查用户名和密码是否正确，注意大小写")
	ERR1004 = register("ERR-1004", "Username Or Password Error", "auth", "客户端登录用户名或密码错误", "检查 RustDesk 客户端输入的用户名和密码")
	ERR1005 = register("ERR-1005", "No Email Address", "auth", "客户端登录缺少邮箱地址", "该用户未设置邮箱，无法使用邮箱验证码登录，请联系管理员补充邮箱")
	ERR1006 = register("ERR-1006", "Verification Code Error", "auth", "验证码错误", "检查验证码是否正确、是否过期，重新获取验证码后重试")
	ERR1007 = register("ERR-1007", "Admin required!", "auth", "需要管理员权限", "当前用户不是管理员，无法执行此操作")
	ERR1008 = register("ERR-1008", "token required", "auth", "缺少认证 Token", "请求必须携带有效的 Authorization Token")
	ERR1009 = register("ERR-1009", "id required", "auth", "缺少设备 ID", "CLI 设备操作必须提供设备 ID")
	ERR1010 = register("ERR-1010", "unauthorized", "auth", "未授权访问", "请先登录，或检查 Token 是否有效")
)

var ErrUnauthorized = New(ERR1010.Code, ERR1010.Message)

// ── OAuth Admin (ERR-2xxx) ───────────────────────────────────────────

var (
	ERR2001 = register("ERR-2001", "provider not found", "oauth", "OAuth Provider 未找到", "确认 Provider 名称拼写正确，检查 server.yaml 或管理后台中的 OAuth 配置")
	ERR2002 = register("ERR-2002", "provider disabled", "oauth", "OAuth Provider 已禁用", "在管理后台 → 系统管理 → 第三方登录中启用该 Provider")
	ERR2003 = register("ERR-2003", "missing code or state", "oauth", "OAuth 回调缺少授权码或状态参数", "这通常由 Provider 端配置错误导致，检查回调地址是否正确配置")
	ERR2004 = register("ERR-2004", "state invalid or expired", "oauth", "OAuth 状态无效或已过期", "授权流程超时（默认 180 秒），请重新发起登录；若频繁出现请检查服务器时钟")
	ERR2005 = register("ERR-2005", "failed to generate state", "oauth", "生成 OAuth 状态令牌失败", "系统随机数生成异常，检查系统熵源或重启服务")
	ERR2006 = register("ERR-2006", "failed to generate ticket", "oauth", "生成 OAuth 一次性票据失败", "系统随机数生成异常，检查系统熵源或重启服务")
	ERR2007 = register("ERR-2007", "ticket required", "oauth", "缺少 OAuth 票据参数", "客户端必须先完成授权流程获取票据，再调用兑换接口")
	ERR2008 = register("ERR-2008", "ticket invalid or expired", "oauth", "OAuth 票据无效或已过期", "票据使用后即失效（一次性），请重新发起授权流程")
	ERR2009 = register("ERR-2009", "oauth ticket user not available", "oauth", "OAuth 票据对应的用户不可用", "用户账号可能已被禁用或删除，请联系管理员")
	ERR2010 = register("ERR-2010", "oauth issuer required", "oauth", "OIDC Provider 缺少 issuer 配置", "在 Provider 配置中填写正确的 issuer URL（如 https://sso.example.com）")
	ERR2011 = register("ERR-2011", "oauth metadata missing required endpoints", "oauth", "OIDC 发现文档缺少必要端点", "检查 issuer URL 是否正确，确认 Provider 的 .well-known/openid-configuration 可访问")
	ERR2012 = register("ERR-2012", "oauth token response missing token", "oauth", "Token 交换响应中缺少 access_token", "检查 Client ID 和 Client Secret 是否正确，确认 Provider 的 Token 端点正常")
	ERR2013 = register("ERR-2013", "oauth userinfo subject mismatch", "oauth", "UserInfo 返回的 sub 与 ID Token 不一致", "Provider 实现有误，联系 Provider 方或改用非 OIDC 模式")
	ERR2014 = register("ERR-2014", "oauth subject claim missing", "oauth", "OAuth 用户信息缺少 subject 标识", "检查 Provider 返回的 claims 中是否包含 sub 字段，或配置 subjectClaim 映射")
	ERR2015 = register("ERR-2015", "verified github email required", "oauth", "GitHub 账号需要已验证的邮箱", "在 GitHub Settings → Emails 中设置一个已验证的公开邮箱")
	ERR2016 = register("ERR-2016", "verified oauth email required", "oauth", "OAuth 账号需要已验证的邮箱", "在 Provider 端设置已验证的邮箱地址")
	ERR2017 = register("ERR-2017", "email domain not allowed", "oauth", "邮箱域名不在允许列表中", "在管理后台 → 第三方登录 → 编辑 Provider → 允许的邮箱域名中添加你的邮箱域名，或清空该字段以允许所有域名")
	ERR2018 = register("ERR-2018", "qq openid response invalid", "oauth", "QQ OpenID 响应无效", "检查 QQ 互联应用的 APP ID 和 APP Secret 是否正确")
	ERR2019 = register("ERR-2019", "oauth issuer required for id token", "oauth", "验证 ID Token 需要 issuer 配置", "在 Provider 配置中填写 issuer URL")
	ERR2020 = register("ERR-2020", "oauth jwks_uri missing for id token", "oauth", "验证 ID Token 需要 jwks_uri 配置", "配置 issuer URL 后系统会自动发现 jwks_uri，或手动填写 jwksUri")
	ERR2021 = register("ERR-2021", "invalid id token", "oauth", "ID Token 验证失败", "检查 issuer、clientId 配置是否与 Provider 匹配，确认 JWKS 端点可访问")
	ERR2022 = register("ERR-2022", "bound admin user not available", "oauth", "绑定的管理员账号不可用", "该管理员账号可能已被禁用或删除，请联系其他管理员处理")
	ERR2023 = register("ERR-2023", "no bindable oauth account", "oauth", "没有可绑定的 OAuth 账号", "当前 OAuth 用户未与任何本地账号关联，且未开启自动创建用户，请联系管理员")
	ERR2024 = register("ERR-2024", "failed to allocate unique username", "oauth", "无法分配唯一用户名", "自动创建用户时用户名冲突，请手动创建用户后绑定 OAuth")
	ERR2025 = register("ERR-2025", "oauth redirect url missing and request base unavailable", "oauth", "OAuth 回调地址缺失且无法自动推断", "在 Provider 配置中填写 redirectUrl，或确保请求包含正确的 Origin 头")
	ERR2026 = register("ERR-2026", "invalid state format", "oauth", "OAuth state 格式无效", "授权流程异常，请重新发起登录")
	ERR2027 = register("ERR-2027", "state signature mismatch", "oauth", "OAuth state 签名不匹配", "服务端 signKey 可能已变更，导致旧 state 失效，请重新发起登录")
	ERR2028 = register("ERR-2028", "state payload incomplete", "oauth", "OAuth state 数据不完整", "授权流程异常，请重新发起登录")
	ERR2029 = register("ERR-2029", "state expired", "oauth", "OAuth state 已过期", "授权流程超时，请重新发起登录")
	ERR2030 = register("ERR-2030", "oauth discovery failed with status %d", "oauth", "OIDC 发现端点返回非 200 状态码", "检查 issuer URL 是否正确，确认 Provider 的 .well-known/openid-configuration 可访问")
	ERR2031 = register("ERR-2031", "token exchange failed with status %d", "oauth", "Token 交换返回非 200 状态码", "检查 Client ID、Client Secret 和 redirectUrl 是否正确")
	ERR2032 = register("ERR-2032", "qq userinfo failed with ret %d", "oauth", "QQ 用户信息接口返回错误", "检查 QQ 互联应用配置和用户授权范围")
	ERR2033 = register("ERR-2033", "qq api failed with status %d", "oauth", "QQ API 请求失败", "检查 QQ 互联应用配置和网络连通性")
	ERR2034 = register("ERR-2034", "userinfo failed with status %d", "oauth", "UserInfo 端点返回非 200 状态码", "检查 Provider 的 userinfo 端点是否正常，确认 access_token 有效")
)

// ── OAuth Admin Controller (ERR-21xx) ────────────────────────────────

var (
	ERR2101 = register("ERR-2101", "InvalidAccountId", "oauth-admin", "无效的 OAuth 账号 ID", "使用正确的账号 ID 进行操作")
	ERR2102 = register("ERR-2102", "UnsupportedOAuthProvider", "oauth-admin", "不支持的 OAuth Provider 类型", "目前支持 github、google、oidc、qq 类型")
	ERR2103 = register("ERR-2103", "InvalidProviderName", "oauth-admin", "无效的 Provider 标识名", "标识名只能包含字母、数字、连字符和下划线")
	ERR2104 = register("ERR-2104", "ClientIdRequired", "oauth-admin", "Client ID 不能为空", "填写 OAuth 应用注册时获得的 Client ID")
	ERR2105 = register("ERR-2105", "ProviderNameExists", "oauth-admin", "Provider 标识名已存在", "使用不同的标识名，或编辑已有 Provider")
	ERR2106 = register("ERR-2106", "ClientSecretRequired", "oauth-admin", "Client Secret 不能为空", "填写 OAuth 应用注册时获得的 Client Secret")
	ERR2107 = register("ERR-2107", "ProviderNotFound", "oauth-admin", "Provider 未找到", "确认 Provider 名称正确")
	ERR2108 = register("ERR-2108", "ProviderNotEnabled", "oauth-admin", "Provider 未启用", "先启用 Provider 再执行测试")
	ERR2109 = register("ERR-2109", "ProviderNotEnabledOrIncomplete", "oauth-admin", "Provider 未启用或配置不完整", "检查 Provider 的启用状态和必要配置项")
	ERR2110 = register("ERR-2110", "InvalidRedirectUrl", "oauth-admin", "回调地址 URL 格式无效", "填写完整的 HTTPS 回调地址，如 https://your-domain/admin/auth/oauth/github/callback")
)

// ── OAuth Client API (ERR-22xx) ──────────────────────────────────────

var (
	ERR2201 = register("ERR-2201", "provider required", "oauth-client", "客户端 OAuth 请求缺少 Provider 名称", "指定要使用的 Provider 标识名")
	ERR2202 = register("ERR-2202", "id and uuid required", "oauth-client", "客户端 OAuth 请求缺少设备 ID 和 UUID", "RustDesk 客户端必须发送 rustdesk_id 和 uuid 参数")
	ERR2203 = register("ERR-2203", "provider not available for client login", "oauth-client", "该 Provider 不允许客户端登录", "将 Provider 的账户角色设为 user 以允许客户端登录")
	ERR2204 = register("ERR-2204", "poll token missing from state", "oauth-client", "OAuth state 中缺少 poll token", "授权流程异常，请重新发起登录")
	ERR2205 = register("ERR-2205", "poll token required", "oauth-client", "轮询请求缺少 poll token", "客户端必须提供 start 接口返回的 poll_token")
	ERR2206 = register("ERR-2206", "client ticket expected", "oauth-client", "票据类型不匹配，期望客户端票据", "使用客户端 OAuth 流程获取的票据来兑换，不要使用管理后台的票据")
	ERR2207 = register("ERR-2207", "failed to generate poll token", "oauth-client", "生成轮询令牌失败", "系统随机数生成异常，检查系统熵源或重启服务")
	ERR2208 = register("ERR-2208", "oauth_account_not_bound", "oauth-client", "OAuth 账号未绑定本地用户", "在管理后台绑定用户，或开启 autoCreateUser 自动创建")
	ERR2209 = register("ERR-2209", "oauth_provider_unreachable", "oauth-client", "OAuth Provider 不可达", "检查服务器能否访问 Provider 的 Token/UserInfo 端点")
	ERR2210 = register("ERR-2210", "oauth_state_expired", "oauth-client", "OAuth 授权状态已过期", "授权超时，请重新发起登录")
	ERR2211 = register("ERR-2211", "oauth_provider_not_for_client", "oauth-client", "该 Provider 不支持客户端登录", "将 Provider 的账户角色设为 user")
	ERR2212 = register("ERR-2212", "oauth_auth_failed", "oauth-client", "OAuth 认证失败", "检查 Provider 配置和网络连通性，查看服务端日志获取详细错误")
)

// ── OIDC (ERR-3xxx) ─────────────────────────────────────────────────

var (
	ERR3001 = register("ERR-3001", "oidc disabled", "oidc", "OIDC 登录已禁用", "在 server.yaml 中设置 oidc.enabled: true")
	ERR3002 = register("ERR-3002", "failed to generate state", "oidc", "生成 OIDC 状态令牌失败", "系统随机数生成异常，检查系统熵源或重启服务")
	ERR3003 = register("ERR-3003", "missing code or state", "oidc", "OIDC 回调缺少授权码或状态参数", "检查 OIDC 回调地址配置是否正确")
	ERR3004 = register("ERR-3004", "state invalid or expired", "oidc", "OIDC 状态无效或已过期", "授权流程超时，请重新发起登录")
	ERR3005 = register("ERR-3005", "failed to generate ticket", "oidc", "生成 OIDC 票据失败", "系统随机数生成异常，检查系统熵源或重启服务")
	ERR3006 = register("ERR-3006", "ticket required", "oidc", "缺少 OIDC 票据参数", "完成授权流程后再调用兑换接口")
	ERR3007 = register("ERR-3007", "ticket invalid or expired", "oidc", "OIDC 票据无效或已过期", "票据一次性使用，请重新发起授权")
	ERR3008 = register("ERR-3008", "oidc issuer required", "oidc", "OIDC issuer 配置缺失", "在 server.yaml 中填写 oidc.issuer")
	ERR3009 = register("ERR-3009", "oidc metadata missing required endpoints", "oidc", "OIDC 发现文档缺少必要端点", "检查 issuer URL，确认 .well-known/openid-configuration 可访问")
	ERR3010 = register("ERR-3010", "oidc token response missing token", "oidc", "OIDC Token 响应缺少 access_token", "检查 clientId、clientSecret 和 redirectUrl 配置")
	ERR3011 = register("ERR-3011", "oidc subject claim missing", "oidc", "OIDC 用户信息缺少 subject 标识", "检查 Provider 返回的 claims 中是否包含 sub 字段")
	ERR3012 = register("ERR-3012", "email domain not allowed", "oidc", "OIDC 邮箱域名不在允许列表中", "在 server.yaml 的 oidc.allowedEmailDomains 中添加你的邮箱域名，或设为空列表允许所有域名")
	ERR3013 = register("ERR-3013", "invalid id token", "oidc", "OIDC ID Token 验证失败", "检查 issuer、clientId 配置，确认 JWKS 端点可访问")
	ERR3014 = register("ERR-3014", "id token issuer invalid", "oidc", "ID Token 的 issuer 与配置不匹配", "确认 server.yaml 中的 oidc.issuer 与 Provider 实际 issuer 一致")
	ERR3015 = register("ERR-3015", "id token audience invalid", "oidc", "ID Token 的 audience 与 clientId 不匹配", "确认 oidc.clientId 与 Provider 注册的应用 ID 一致")
	ERR3016 = register("ERR-3016", "id token expired", "oidc", "ID Token 已过期", "检查服务器时钟是否准确，与 Provider 时钟偏差不应超过 5 分钟")
	ERR3017 = register("ERR-3017", "id token issued-at invalid", "oidc", "ID Token 的 iat 声明无效", "检查服务器时钟是否准确")
	ERR3018 = register("ERR-3018", "bound admin user not available", "oidc", "OIDC 绑定的管理员账号不可用", "该管理员可能已被禁用或删除")
	ERR3019 = register("ERR-3019", "no bindable admin account", "oidc", "没有可绑定的 OIDC 管理员账号", "当前 OIDC 用户未与管理员关联，且未开启 autoCreateAdmin，请联系管理员")
	ERR3020 = register("ERR-3020", "failed to allocate unique username", "oidc", "OIDC 自动创建用户名失败", "用户名冲突，请手动创建用户后绑定")
	ERR3021 = register("ERR-3021", "oidc redirect url missing and request base unavailable", "oidc", "OIDC 回调地址缺失且无法推断", "在 server.yaml 中填写 oidc.redirectUrl")
	ERR3022 = register("ERR-3022", "oidc discovery failed with status %d", "oidc", "OIDC 发现端点返回非 200", "检查 issuer URL 和网络连通性")
	ERR3023 = register("ERR-3023", "token exchange failed with status %d", "oidc", "OIDC Token 交换返回非 200", "检查 clientId、clientSecret 和 redirectUrl")
	ERR3024 = register("ERR-3024", "userinfo failed with status %d", "oidc", "OIDC UserInfo 端点返回非 200", "检查 access_token 是否有效，userinfo 端点是否正常")
	ERR3025 = register("ERR-3025", "id token kid missing", "oidc", "ID Token 缺少 kid 头部", "Provider 未在 ID Token 中包含 kid，无法匹配 JWKS 密钥")
	ERR3026 = register("ERR-3026", "unsupported id token alg %s", "oidc", "ID Token 使用了不支持的签名算法", "目前仅支持 RS256 算法，联系 Provider 方或改用非 OIDC 模式")
	ERR3027 = register("ERR-3027", "oidc jwks_uri missing", "oidc", "OIDC 配置缺少 jwks_uri", "配置 issuer 后自动发现，或手动填写 jwksUri")
	ERR3028 = register("ERR-3028", "oidc jwks fetch failed with status %d", "oidc", "JWKS 端点请求失败", "检查 jwks_uri 是否可访问")
	ERR3029 = register("ERR-3029", "matching oidc jwk not found", "oidc", "未找到匹配 kid 的 JWKS 密钥", "检查 Provider 的 JWKS 端点是否包含对应密钥")
	ERR3030 = register("ERR-3030", "invalid rsa jwk", "oidc", "RSA JWK 格式无效", "Provider 返回的 JWKS 格式不标准")
	ERR3031 = register("ERR-3031", "invalid rsa exponent", "oidc", "RSA JWK 指数无效", "Provider 返回的 JWKS 格式不标准")
	ERR3032 = register("ERR-3032", "OIDC_NOT_SUPPORTED", "oidc", "OIDC 兼容接口不支持", "本项目使用独立的 OAuth Provider 体系，不兼容 RustDesk 官方 OIDC 接口")
	ERR3033 = register("ERR-3033", "OIDC_QUERY_NOT_SUPPORTED", "oidc", "OIDC 查询接口不支持", "本项目使用独立的 OAuth Provider 体系，不支持 OIDC 查询")
)

// ── User Management (ERR-4xxx) ───────────────────────────────────────

var (
	ERR4001 = register("ERR-4001", "UsernameEmpty", "user", "用户名不能为空", "填写有效的用户名")
	ERR4002 = register("ERR-4002", "UserExists", "user", "用户名已存在", "使用不同的用户名")
	ERR4003 = register("ERR-4003", "PasswordEmpty", "user", "密码不能为空", "填写有效的密码")
	ERR4004 = register("ERR-4004", "TFA_Validate_Err", "user", "两步验证码错误", "检查 2FA 验证码是否正确，注意时间同步")
	ERR4005 = register("ERR-4005", "DataError", "user", "数据错误", "检查提交的数据格式是否正确")
	ERR4006 = register("ERR-4006", "UserNotExists", "user", "用户不存在", "确认用户名是否正确")
	ERR4007 = register("ERR-4007", "NoUserIds", "user", "未选择要操作的用户", "至少选择一个用户")
)

// ── Address Book (ERR-5xxx) ──────────────────────────────────────────

var (
	ERR5001 = register("ERR-5001", "address book not found", "addressbook", "地址簿未找到", "确认地址簿 ID 是否正确")
	ERR5002 = register("ERR-5002", "name required", "addressbook", "名称不能为空", "填写有效的名称")
	ERR5003 = register("ERR-5003", "tag already exists", "addressbook", "标签已存在", "使用不同的标签名")
	ERR5004 = register("ERR-5004", "tag not found", "addressbook", "标签未找到", "确认标签名称是否正确")
	ERR5005 = register("ERR-5005", "old and new required", "addressbook", "重命名标签需要提供旧名称和新名称", "同时填写原标签名和新标签名")
	ERR5006 = register("ERR-5006", "device id required", "addressbook", "设备 ID 不能为空", "提供有效的设备 ID")
	ERR5007 = register("ERR-5007", "peer not found", "addressbook", "设备未找到", "确认设备 ID 是否正确")
	ERR5008 = register("ERR-5008", "user not found", "addressbook", "用户未找到", "确认用户名是否正确")
	ERR5009 = register("ERR-5009", "personal address book is read-only", "addressbook", "个人地址簿为只读", "个人地址簿由系统自动管理，不可手动修改")
	ERR5010 = register("ERR-5010", "personal address book cannot be deleted", "addressbook", "个人地址簿不可删除", "个人地址簿由系统自动管理")
	ERR5011 = register("ERR-5011", "administrator-created address book cannot be deleted by user", "addressbook", "管理员创建的地址簿用户无权删除", "请联系管理员删除")
	ERR5012 = register("ERR-5012", "invalid rule", "addressbook", "无效的访问规则", "规则值必须为 1（允许）或 2（拒绝）")
	ERR5013 = register("ERR-5013", "target type and target guid required", "addressbook", "规则缺少目标类型和目标 GUID", "填写完整的目标类型和目标标识")
	ERR5014 = register("ERR-5014", "guid and target required", "addressbook", "规则缺少 GUID 和目标", "填写完整的规则标识和目标信息")
	ERR5015 = register("ERR-5015", "rule not found", "addressbook", "访问规则未找到", "确认规则 ID 是否正确")
	ERR5016 = register("ERR-5016", "NoPeerIds", "addressbook", "未选择要操作的设备", "至少选择一个设备")
	ERR5017 = register("ERR-5017", "NoTagNames", "addressbook", "未选择要操作的标签", "至少选择一个标签")
	ERR5018 = register("ERR-5018", "Number of equipment in excess of licenses", "addressbook", "设备数量超过授权许可数", "增加授权设备数或删除多余设备")
	ERR5019 = register("ERR-5019", "rule must be 1, 2 or 3", "addressbook", "规则值必须为 1、2 或 3", "1=允许，2=拒绝，3=密码访问")
	ERR5020 = register("ERR-5020", "exceed_max_devices", "addressbook", "设备数超过最大限制", "增加授权设备数或删除多余设备")
)

// ── Enterprise/Compat (ERR-6xxx) ─────────────────────────────────────

var (
	ERR6001 = register("ERR-6001", "device group not found", "enterprise", "设备组未找到", "确认设备组名称是否正确")
	ERR6002 = register("ERR-6002", "user group not found", "enterprise", "用户组未找到", "确认用户组名称是否正确")
	ERR6003 = register("ERR-6003", "strategy not found", "enterprise", "策略未找到", "确认策略名称是否正确")
	ERR6004 = register("ERR-6004", "name required", "enterprise", "名称不能为空", "填写有效的名称")
	ERR6005 = register("ERR-6005", "strategy_guid required", "enterprise", "策略 GUID 不能为空", "提供有效的策略标识")
	ERR6006 = register("ERR-6006", "device not found", "enterprise", "设备未找到", "确认设备 ID 是否正确")
	ERR6007 = register("ERR-6007", "user not found", "enterprise", "用户未找到", "确认用户标识是否正确")
	ERR6008 = register("ERR-6008", "user required", "enterprise", "用户参数不能为空", "提供有效的用户标识")
	ERR6009 = register("ERR-6009", "INVALID_INPUT", "enterprise", "无效的输入参数", "检查请求参数格式")
	ERR6010 = register("ERR-6010", "NOT_ENABLED", "enterprise", "功能未启用", "在 server.yaml 中启用相应功能")
)

// ── Record (ERR-7xxx) ───────────────────────────────────────────────

var (
	ERR7001 = register("ERR-7001", "type required", "record", "记录类型不能为空", "指定有效的记录类型")
	ERR7002 = register("ERR-7002", "file required", "record", "文件名不能为空", "提供有效的文件名")
	ERR7003 = register("ERR-7003", "invalid offset", "record", "无效的文件偏移量", "偏移量必须为非负整数")
	ERR7004 = register("ERR-7004", "unsupported record op", "record", "不支持的记录操作", "仅支持 R（读取）和 W（写入）操作")
	ERR7005 = register("ERR-7005", "invalid record write size", "record", "无效的写入大小", "写入数据大小与声明不匹配")
	ERR7006 = register("ERR-7006", "record file exceeds max size %d bytes", "record", "记录文件超过最大大小限制", "清理旧记录文件或增大容量限制")
)

// ── Mail Template (ERR-8xxx) ─────────────────────────────────────────

var (
	ERR8001 = register("ERR-8001", "MailTemplateNameEmpty", "mail", "邮件模板名称不能为空", "填写有效的模板名称")
	ERR8002 = register("ERR-8002", "MailTemplateSubjectEmpty", "mail", "邮件模板主题不能为空", "填写有效的邮件主题")
	ERR8003 = register("ERR-8003", "MailTemplateContentsEmpty", "mail", "邮件模板内容不能为空", "填写有效的邮件内容")
	ERR8004 = register("ERR-8004", "DataError", "mail", "数据错误", "检查提交的数据格式")
	ERR8005 = register("ERR-8005", "MailTemplateNotFound", "mail", "邮件模板未找到", "确认模板 ID 是否正确")
)

// ── Session/Token (ERR-9xxx) ─────────────────────────────────────────

var (
	ERR9001 = register("ERR-9001", "NoSessionIds", "session", "未选择要操作的会话", "至少选择一个会话")
	ERR9002 = register("ERR-9002", "NoTokenIds", "token", "未选择要操作的 Token", "至少选择一个 Token")
	ERR9003 = register("ERR-9003", "UUIDEmpty", "mail-log", "邮件日志 UUID 不能为空", "提供有效的 UUID")
)

// ── Dashboard/Connectivity (ERR-A0xx) ────────────────────────────────

var (
	ERRA001 = register("ERR-A001", "invalid connectivity target", "dashboard", "无效的连通性检测目标", "目标必须是 hbbs、hbbr 或自定义地址")
	ERRA002 = register("ERR-A002", "empty target", "dashboard", "检测目标为空", "填写有效的目标地址")
	ERRA003 = register("ERR-A003", "empty host", "dashboard", "主机地址为空", "填写有效的主机地址")
	ERRA004 = register("ERR-A004", "invalid url", "dashboard", "URL 格式无效", "填写有效的 HTTP/HTTPS URL")
	ERRA005 = register("ERR-A005", "key is empty", "dashboard", "API Key 为空", "在 server.yaml 中配置 signKey")
)

// ── Infrastructure/Startup (ERR-B0xx) ────────────────────────────────

var (
	ERRB001 = register("ERR-B001", "db engine is nil", "infra", "数据库引擎未初始化", "检查数据库配置（driver、dsn）是否正确，确认数据库服务正在运行")
	ERRB002 = register("ERR-B002", "unsafe signKey", "infra", "signKey 不安全", "在 server.yaml 中设置至少 32 字符的随机密钥，不要使用默认值")
	ERRB003 = register("ERR-B003", "job db engine is nil", "infra", "定时任务数据库引擎未初始化", "检查数据库配置")
	ERRB004 = register("ERR-B004", "invalid device check job duration", "infra", "设备检查定时任务间隔无效", "在 server.yaml 中设置有效的 jobsConfig.deviceCheckJob.duration（分钟）")
	ERRB005 = register("ERR-B005", "create scheduler failed", "infra", "创建定时调度器失败", "检查 gocron 调度器配置，查看服务端日志获取详细错误")
	ERRB006 = register("ERR-B006", "create device check job failed", "infra", "创建设备检查定时任务失败", "检查定时任务配置，查看服务端日志获取详细错误")
	ERRB010 = register("ERR-B010", "database error", "infra", "数据库操作错误", "检查数据库连接和配置，查看服务端日志获取详细错误")
)

// ── HTTP/File Utils (ERR-C0xx) ───────────────────────────────────────

var (
	ERRC001 = register("ERR-C001", "only support http, https, socks5 proxy protocols", "util", "仅支持 http/https/socks5 代理协议", "修改代理配置为支持的协议")
	ERRC002 = register("ERR-C002", "http response too large", "util", "HTTP 响应过大", "检查请求的 URL 是否返回了异常大的响应")
	ERRC003 = register("ERR-C003", "http get failed with status %d", "util", "HTTP GET 请求失败", "检查目标 URL 是否可访问")
	ERRC004 = register("ERR-C004", "download failed with status %d", "util", "文件下载失败", "检查下载 URL 是否可访问")
	ERRC005 = register("ERR-C005", "download too large", "util", "下载文件过大", "检查文件大小是否超过限制")
	ERRC006 = register("ERR-C006", "empty zip entry name", "util", "ZIP 条目名称为空", "ZIP 文件格式异常")
	ERRC007 = register("ERR-C007", "zip entry is a symlink", "util", "ZIP 条目是符号链接", "出于安全考虑，不允许 ZIP 中包含符号链接")
	ERRC008 = register("ERR-C008", "zip entry uses absolute path", "util", "ZIP 条目使用绝对路径", "出于安全考虑，不允许 ZIP 条目使用绝对路径")
	ERRC009 = register("ERR-C009", "zip entry escapes destination", "util", "ZIP 条目路径逃逸", "出于安全考虑，不允许 ZIP 条目路径超出目标目录")
)

// ── RustDesk Server Helper (ERR-D0xx) ────────────────────────────────

var (
	ERRD001 = register("ERR-D001", "unsupported operating system", "rustdesk", "不支持的操作系统", "目前仅支持 Linux、Windows、macOS")
	ERRD002 = register("ERR-D002", "hbbr start error", "rustdesk", "hbbr 启动失败", "检查 hbbr 二进制文件是否存在且可执行，确认端口未被占用")
	ERRD003 = register("ERR-D003", "hbbs start error", "rustdesk", "hbbs 启动失败", "检查 hbbs 二进制文件是否存在且可执行，确认端口未被占用")
	ERRD004 = register("ERR-D004", "write hbbr pid file error", "rustdesk", "写入 hbbr PID 文件失败", "检查 rustdesk-server 目录的写入权限")
	ERRD005 = register("ERR-D005", "write hbbs pid file error", "rustdesk", "写入 hbbs PID 文件失败", "检查 rustdesk-server 目录的写入权限")
)

// ── GitHub API (ERR-E0xx) ────────────────────────────────────────────

var (
	ERRE001 = register("ERR-E001", "get releases request failed", "github", "获取 GitHub Releases 请求失败", "检查网络连通性和 GitHub API 访问权限")
	ERRE002 = register("ERR-E002", "decode releases response failed", "github", "解析 GitHub Releases 响应失败", "GitHub API 返回了非预期格式，检查 API 版本兼容性")
	ERRE003 = register("ERR-E003", "get latest release request failed", "github", "获取最新 Release 请求失败", "检查网络连通性和 GitHub API 访问权限")
	ERRE004 = register("ERR-E004", "decode latest release response failed", "github", "解析最新 Release 响应失败", "GitHub API 返回了非预期格式，检查 API 版本兼容性")
	ERRE005 = register("ERR-E005", "get release by tag request failed", "github", "按标签获取 Release 请求失败", "检查网络连通性和 GitHub API 访问权限")
	ERRE006 = register("ERR-E006", "decode release by tag response failed", "github", "按标签解析 Release 响应失败", "GitHub API 返回了非预期格式，检查 API 版本兼容性")
)

// ── Audit (ERR-F0xx) ────────────────────────────────────────────────

var (
	ERRF001 = register("ERR-F001", "guid required", "audit", "审计记录 GUID 不能为空", "提供有效的审计记录 GUID")
)
