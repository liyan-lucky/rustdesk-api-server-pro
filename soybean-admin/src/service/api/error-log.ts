import { request } from '../request';

export function fetchErrorLogList(params: any) {
  return request<Api.Audit.ErrorLogList>({ url: '/error-logs/list', params });
}

export function fetchErrorLogClear() {
  return request<any>({ url: '/error-logs/clear', method: 'delete' });
}
