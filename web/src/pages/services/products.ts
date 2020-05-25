import request from '../utils/request';
import { PaginationParameters, FieldsValue } from '../declare';
import { serviceApiPrefix, generateQuery } from './utils';

export function getProductsTag(product: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/products/${product}/labels${generateQuery(params)}`);
};

export function addProductsTag(product: string, params: FieldsValue) {
  return request.post(`${serviceApiPrefix}/products/${product}/labels`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params),
  });
};

export function offlineProductsTag(product: string, label: string) {
  return request.put(`${serviceApiPrefix}/products/${product}/labels/${label}:offline`);
};

export function offlineProductsSetting(product: string, module: string, setting: string) {
  return request.put(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}:offline`);
};

export function publishProductsTag(product: string, label: string, params: FieldsValue) {
  const { kind } = params;
  const partUrl = kind === 'batch' ? ':assign' : '/rules';
  return request.post(`${serviceApiPrefix}/products/${product}/labels/${label}${partUrl}`, {
    headers: {
      'Content-type': 'application/json'
    },
    body: JSON.stringify(params)
  });
};

export function getLabelLogs(product: string, label: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/products/${product}/labels/${label}/logs`);
};

export function getLabelGroups(product: string, label: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/products/${product}/labels/${label}/groups${generateQuery(params)}`);
};

export function getLabelUsers(product: string, label: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/products/${product}/labels/${label}/users${generateQuery(params)}`);
};

export function getPublishRules(product: string, label: string) {
  return request(`${serviceApiPrefix}/products/${product}/labels/${label}/rules`);
};

export function getPublishSettingRules(product: string, module: string, setting: string) {
  return request(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/rules`);
};

export function updateProductTagRule(
  product: string,
  label: string,
  rule: string,
  params: FieldsValue,
) {
  return request.put(`${serviceApiPrefix}/products/${product}/labels/${label}/rules/${rule}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params)
  });
};

export function publishProductsSetting(product: string, module: string, setting: string, params: FieldsValue) {
  const { kind } = params;
  const partUrl = kind === 'batch' ? ':assign' : '/rules';
  return request.post(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}${partUrl}`, {
    headers: {
      'Content-type': 'application/json'
    },
    body: JSON.stringify(params)
  });
};

export function updateProductSettingRule(
  product: string,
  module: string,
  setting: string,
  rule: string,
  params: FieldsValue,
) {
  return request.put(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/rules/${rule}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params)
  });
};

export function deleteProductsTag(product: string, label: string) {
  return request.delete(`${serviceApiPrefix}/products/${product}/labels/${label}`);
};

export function updateProductsTag(product: string, params: FieldsValue) {
  return request.put(`${serviceApiPrefix}/products/${product}/labels/${params.name}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params),
  });
};

export function getProductStatistics(product: string) {
  return request(`${serviceApiPrefix}/products/${product}/statistics`);
};

export function getProductsModule(product: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/products/${product}/modules${generateQuery(params)}`);
};

export function addProductsModule(product: string, params: FieldsValue) {
  return request.post(`${serviceApiPrefix}/products/${product}/modules`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params),
  });
};

export function updateProductsModule(product: string, params: FieldsValue) {
  return request.put(`${serviceApiPrefix}/products/${product}/modules/${params.name}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params),
  });
};

export function getProductsSetting(product: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/products/${product}/settings${generateQuery(params)}`);
};

export function getModuleSetting(product: string, module: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/products/${product}/modules/${module}/settings${generateQuery(params)}`);
};

export function addProductsSetting(product: string, params: FieldsValue) {
  return request.post(`${serviceApiPrefix}/products/${product}/modules/${params.module}/settings`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params),
  });
};

export function updateProductsSetting(product: string, params: FieldsValue) {
  return request.put(`${serviceApiPrefix}/products/${product}/modules/${params.module}/settings/${params.name}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params),
  });
};

export function getSettingGroups(product: string, module: string, setting: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/groups${generateQuery(params)}`);
};

export function getSettingUsers(product: string, module: string, setting: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/users${generateQuery(params)}`);
};

export function getSettingLogs(product: string, module: string, setting: string, params: PaginationParameters) {
  return request(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/logs`);
};

export function offlineProductsModule(product: string, module: string) {
  return request.put(`${serviceApiPrefix}/products/${product}/modules/${module}:offline`);
};

export function getProductList() {
  return request(`${serviceApiPrefix}/products`);
};

export function updateProduct(name: string, desc: string, uids: string[]) {
  return request.put(`${serviceApiPrefix}/products/${name}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      desc,
      uids,
    })
  });
};

export function createProduct(name: string, desc: string, uids: string[]) {
  return request.post(`${serviceApiPrefix}/products`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      name,
      desc,
      uids,
    })
  });
};

export function offlineProduct(name: string) {
  return request.put(`${serviceApiPrefix}/products/${name}:offline`);
};

export function deleteProduct(name: string) {
  return request.delete(`${serviceApiPrefix}/products/${name}`);
};

export function getPermission(params: object) {
  return request.post(`${serviceApiPrefix}/ac/permission:check`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(params),
  });
}

export function recallLabelLogs(product: string, label: string, hid: string) {
  return request.post(`${serviceApiPrefix}/products/${product}/labels/${label}:recall`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ hid }),
  });
};

export function deleteLabelGroup(product: string, label: string, gid: string) {
  return request.delete(`${serviceApiPrefix}/products/${product}/labels/${label}/groups/${gid}`);
};

export function deleteLabeUser(product: string, label: string, uid: string) {
  return request.delete(`${serviceApiPrefix}/products/${product}/labels/${label}/users/${uid}`);
};

export function recallSettingLogs(product: string, module: string, setting: string, hid: string) {
  return request.post(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}:recall`, {
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ hid }),
  });
};

export function deleteSettingGroup(product: string, module: string, setting: string, gid: string) {
  return request.delete(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/groups/${gid}`);
};

export function deleteSettingUser(product: string, module: string, setting: string, uid: string) {
  return request.delete(`${serviceApiPrefix}/products/${product}/modules/${module}/settings/${setting}/users/${uid}`);
};