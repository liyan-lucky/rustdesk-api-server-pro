import { request } from '../request';

/**
 * Login
 *
 * @param userName User name
 * @param password Password
 */
export function fetchLogin(model: Api.Form.LoginForm) {
  return request<Api.Auth.LoginToken>({
    url: '/auth/login',
    method: 'post',
    data: {
      username: model.username,
      password: model.password,
      code: model.code,
      captchaId: model.captchaId
    }
  });
}

export function fetchCaptcha() {
  return request<Api.Auth.Captcha>({ url: '/auth/captcha' });
}

export function fetchOidcLoginUrl(redirect?: string) {
  return request<Api.Auth.OidcLoginUrl>({
    url: '/auth/oidc/url',
    params: redirect ? { redirect } : undefined
  });
}

export function fetchOidcTicketToken(ticket: string) {
  return request<Api.Auth.LoginToken>({
    url: '/auth/oidc/token',
    params: { ticket }
  });
}

export function fetchOAuthProviders() {
  return request<Api.Auth.OAuthProvider[]>({
    url: '/auth/oauth/providers'
  });
}

export function fetchOAuthLoginUrl(provider: string, redirect?: string) {
  return request<Api.Auth.OidcLoginUrl>({
    url: '/auth/oauth/url',
    params: {
      provider,
      ...(redirect ? { redirect } : {})
    }
  });
}

export function fetchOAuthTicketToken(ticket: string) {
  return request<Api.Auth.LoginToken>({
    url: '/auth/oauth/token',
    params: { ticket }
  });
}

/** Get user info */
export function fetchGetUserInfo() {
  return request<Api.Auth.UserInfo>({ url: '/userinfo' });
}
