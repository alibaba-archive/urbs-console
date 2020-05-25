import request from '../utils/request';
import { PaginationParameters, FieldsValue, Group } from '../declare';
import { serviceApiPrefix, generateQuery } from './utils';

export function getGroups(params: PaginationParameters) {
  return request(`${serviceApiPrefix}/groups${ generateQuery(params) }`);
};

export function addGroups(params: FieldsValue) {
  return request.post(`${serviceApiPrefix}/groups:batch`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({groups: [params]}),
  });
};

export function updateGroups(params: Group) {
  return request.put(`${serviceApiPrefix}/groups/${params.uid}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params),
  });
};

export function deleteGroups(uid: string) {
  return request.delete(`${serviceApiPrefix}/groups/${uid}`);
};

export function getGroupLabels(uid: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/groups/${uid}/labels${ generateQuery(params) }`);
};

export function deleteGroupLabel (product: string, label: string, uid: string) {
  return request.delete(`${serviceApiPrefix}/products/${product}/labels/${label}/groups/${uid}`);
}

export function getGroupSettings(uid: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/groups/${uid}/settings${ generateQuery(params) }`);
};

export function getGroupMembers(uid: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/groups/${uid}/members${ generateQuery(params) }`);
};

export function addGroupMembers(uid: string, users: string[]) {
  return request.post(`${serviceApiPrefix}/groups/${uid}/members:batch`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({users})
  });
};

export function deleteGroupMembers(uid: string, user: string) {
  return request.delete(`${serviceApiPrefix}/groups/${uid}/members?user=${user}`);
};

export function deleteGroupSetting (product: string, module: string, setting: string, uid: string) {
  return request.delete(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/groups/${uid}`);
}

export function rollbackGroupSetting (product: string, module: string, setting: string, uid: string) {
  return request.put(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/groups/${uid}:rollback`);
}

export function getPermission () {
  return request.post(`${serviceApiPrefix}/ac/permission:check`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({}),
  });
}
