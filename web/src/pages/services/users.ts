import request from '../utils/request';
import { PaginationParameters, FieldsValue } from '../declare';
import { serviceApiPrefix, generateQuery } from './utils';

export function addAcUsers(users: FieldsValue) {
  return request.post(`${serviceApiPrefix}/ac/users`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ users }),
  });
};

export function updateAcUser(uid: string, name: FieldsValue) {
  return request.put(`${serviceApiPrefix}/ac/users/${uid}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ name }),
  });
};

export function deleteAcUser(uid: string) {
  return request.delete(`${serviceApiPrefix}/ac/users/${uid}`);
}



export function getAcUsers() {
  return request(`${serviceApiPrefix}/ac/users?pageSize=1000`);
};

export function getAcUsersList(params?: PaginationParameters) {
  return request(`${serviceApiPrefix}/ac/users${generateQuery(params ? params : {})}`);
};


export function searchAcUsers(searchKey: string) {
  return request(`${serviceApiPrefix}/ac/users:search?key=${searchKey}`);
};

export function getCanaryUsers(params: PaginationParameters) {
  return request(`${serviceApiPrefix}/users${generateQuery(params)}`);
};

export function addCanaryUsers(users: string[]) {
  return request.post(`${serviceApiPrefix}/users:batch`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ users }),
  });
}

export function getUserLabelsCache(uid: string) {
  return request.put(`${serviceApiPrefix}/users/${uid}/labels:cache`);
}

export function getUserLabels(uid: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/users/${uid}/labels${generateQuery(params)}`);
}

export function getUserSettings(uid: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/users/${uid}/settings${generateQuery(params)}`);
}

export function deleteUserLabel(product: string, label: string, uid: string) {
  return request.delete(`${serviceApiPrefix}/products/${product}/labels/${label}/users/${uid}`);
}

export function deleteUserSetting(product: string, module: string, setting: string, uid: string) {
  return request.delete(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/users/${uid}`);
}

export function rollbackUserSetting(product: string, module: string, setting: string, uid: string) {
  return request.put(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/users/${uid}:rollback`);
}

export function getPermission() {
  return request.post(`${serviceApiPrefix}/ac/permission:check`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({}),
  });
}
