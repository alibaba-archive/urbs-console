import request from '../utils/request';
import { PaginationParameters, FieldsValue, Group } from '../declare';
import { serviceApiPrefix, generateQuery } from './utils';

export function getGroups(params: PaginationParameters) {
  return request(`${serviceApiPrefix}/groups${generateQuery(params)}`);
};

export function addGroups(params: FieldsValue) {
  return request.post(`${serviceApiPrefix}/groups:batch`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ groups: [params] }),
  });
};

export function updateGroups(params: Group) {
  return request.put(`${serviceApiPrefix}/groups/${params.uid}?kind=${params.kind}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params),
  });
};

export function deleteGroups(kind: string, uid: string) {
  return request.delete(`${serviceApiPrefix}/groups/${uid}?kind=${kind}`);
};

export function getGroupLabels(uid: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/groups/${uid}/labels${generateQuery(params)}`);
};

export function deleteGroupLabel(product: string, label: string, kind: string, uid: string) {
  return request.delete(`${serviceApiPrefix}/products/${product}/labels/${label}/groups/${uid}?kind=${kind}`);
}

export function getGroupSettings(uid: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/groups/${uid}/settings${generateQuery(params)}`);
};

export function getGroupMembers(uid: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/groups/${uid}/members${generateQuery(params)}`);
};

export function addGroupMembers(kind: string, uid: string, users: string[]) {
  return request.post(`${serviceApiPrefix}/groups/${uid}/members:batch?kind=${kind}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ users })
  });
};

export function deleteGroupMembers(kind: string, uid: string, user: string) {
  return request.delete(`${serviceApiPrefix}/groups/${uid}/members?user=${user}&kind=${kind}`);
};

export function deleteGroupSetting(product: string, module: string, setting: string, kind: string, uid: string) {
  return request.delete(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/groups/${uid}?kind=${kind}`);
}

export function rollbackGroupSetting(product: string, module: string, setting: string, kind: string, uid: string) {
  return request.put(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/groups/${uid}:rollback?kind=${kind}`);
}

export function getPermission() {
  return request.post(`${serviceApiPrefix}/ac/permission:check`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({}),
  });
}
