import { request } from '../request';

export function fetchAbPeers(params: any) {
  return request<Api.AddressBook.PeerList>({ url: '/ab/peers', params });
}

export function fetchAbList(params: any) {
  return request<Api.AddressBook.AddressBookList>({ url: '/ab/shared-profiles', params });
}

export function fetchAbPersonal() {
  return request<{ guid: string }>({ url: '/ab/personal' });
}

export function fetchAbAllList() {
  return request<Api.AddressBook.AddressBook[]>({ url: '/ab/list' });
}

export function fetchAbSharedAdd(data: Partial<Api.AddressBook.AddressBook>) {
  return request<{ guid: string }>({ url: '/ab/shared/add', method: 'post', data });
}

export function fetchAbSharedUpdate(data: Partial<Api.AddressBook.AddressBook>) {
  return request({ url: '/ab/shared/update', method: 'put', data });
}

export function fetchAbSharedDelete(guids: string[]) {
  return request({ url: '/ab/shared', method: 'delete', data: guids });
}

export function fetchAbTags(abGuid: string) {
  return request({ url: `/ab/tags/${abGuid}`, method: 'post' });
}

export function fetchAbAllTags() {
  return request<Api.AddressBook.Tag[]>({ url: '/ab/tags' });
}

export function fetchAbTagAdd(abGuid: string, data: any) {
  return request({ url: `/ab/tag/add/${abGuid}`, method: 'post', data });
}

export function fetchAbTagUpdate(abGuid: string, data: any) {
  return request({ url: `/ab/tag/update/${abGuid}`, method: 'put', data });
}

export function fetchAbTagRename(abGuid: string, data: any) {
  return request({ url: `/ab/tag/rename/${abGuid}`, method: 'put', data });
}

export function fetchAbTagDelete(abGuid: string, data: any) {
  return request({ url: `/ab/tag/${abGuid}`, method: 'delete', data });
}

export function fetchAbPeerAdd(abGuid: string, data: any) {
  return request({ url: `/ab/peer/add/${abGuid}`, method: 'post', data });
}

export function fetchAbPeerUpdate(abGuid: string, data: any) {
  return request({ url: `/ab/peer/update/${abGuid}`, method: 'put', data });
}

export function fetchAbPeerDelete(abGuid: string, data: any) {
  return request({ url: `/ab/peer/${abGuid}`, method: 'delete', data });
}
