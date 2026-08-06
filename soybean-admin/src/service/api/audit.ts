import { request } from '../request';

export function fetchAuditLogList(params: any) {
  return request<Api.Audit.AuditLogList>({ url: '/audit/list', params });
}

export function fetchAuditFileTransferLogList(params: any) {
  return request<Api.Audit.AuditFileTransferList>({ url: '/audit/file-transfer-list', params });
}

export function fetchAuditLogClear() {
  return request<any>({ url: '/audit/clear', method: 'delete' });
}

export function fetchAuditFileTransferLogClear() {
  return request<any>({ url: '/audit/file-transfer-clear', method: 'delete' });
}

export function fetchContainerLogList(params: any) {
  return request<any>({ url: '/container-logs/list', params });
}

export function fetchContainerLogClear() {
  return request<any>({ url: '/container-logs/clear', method: 'delete' });
}
