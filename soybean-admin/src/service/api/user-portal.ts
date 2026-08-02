import { userPortalRequest } from '../request/user-portal';

export function fetchMyDevices(params: any) {
  return userPortalRequest<Api.Devices.DevicesList>({ url: '/devices/my', params });
}

export function fetchUserPortalInfo() {
  return userPortalRequest<Api.Auth.UserInfo>({ url: '/userinfo' });
}

export function fetchUserOverview() {
  return userPortalRequest<any>({ url: '/overview' });
}

export function fetchMySessions(params: any) {
  return userPortalRequest<any>({ url: '/sessions', params });
}

export function revokeMySessions(ids: number[]) {
  return userPortalRequest({ url: '/sessions', method: 'delete', data: ids });
}

export function fetchMySecurityEvents(params: any) {
  return userPortalRequest<any>({ url: '/security-events', params });
}
