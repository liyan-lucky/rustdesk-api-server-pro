import { userPortalRequest } from '../request/user-portal';

export function fetchMyDevices(params: any) {
  return userPortalRequest<Api.Devices.DevicesList>({ url: '/devices/my', params });
}

export function fetchUserPortalInfo() {
  return userPortalRequest<Api.Auth.UserInfo>({ url: '/userinfo' });
}
