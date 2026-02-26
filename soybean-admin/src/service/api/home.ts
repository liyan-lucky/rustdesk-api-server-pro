import { request } from '../request';

export function fetchStat() {
  return request<Api.Home.Stat>({ url: '/dashboard/stat' });
}

export function fetchLineCharts() {
  return request<Api.Home.LineChart>({ url: '/dashboard/line/charts' });
}

export function fetchPieCharts() {
  return request<Api.Home.PieChart[]>({ url: '/dashboard/pie/charts' });
}

export function fetchServerConfig() {
  return request<Api.Home.ServerConfig>({ url: '/dashboard/server/config' });
}

export function fetchServerConnectivity() {
  return request<Api.Home.ServerConnectivity>({ url: '/dashboard/server/connectivity' });
}
