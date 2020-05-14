import { AnyAction } from 'redux';
import * as productsService from '../services/products';
import { Model, EffectsCommandMap } from 'dva';

const products: Model = {
  namespace: 'products',
  state: {
    productModulesList: [],
    modulePrePageTokens: [],
    productList: [],
    productStatistics: {},
    productTagsList: [],
    prePageTokens: [],
    productSettingsList: [],
    settingPrePageTokens: [],
    moduleSettingsList: [],
    moduleSettingPrePageTokens: [],
    labelLogsList: [],
    settingLogsList: [],
    labelGroupsList: [],
    labelGroupsPrePageTokens: [],
    labelUsersList: [],
    labelUsersPrePageTokens: [],
    settingGroupsList: [],
    settingGroupsPrePageTokens: [],
    settingUsersList: [],
    settingUsersPrePageTokens: [],
  },
  reducers: {
    setStateByPayload(state, { payload }: AnyAction) {
      return { ...state, ...payload };
    },
    setProductStatistics (state, { payload }: AnyAction) {
      const { productStatistics } = payload;
      return {...state, productStatistics};
    },
    setProductsList (state, { payload }: AnyAction) {
      return { ...state, productList: payload.productList };
    },
  },
  effects: {
    *getLabelGroups({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type, product, label } = payload;
      const { pageToken } = params;
      const { labelGroupsPrePageTokens } = yield select(state => state.products);
      const preLen = labelGroupsPrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(productsService.getLabelGroups, product, label, params);      
      if (type === 'next') labelGroupsPrePageTokens.push(pageToken);
      if (type === 'pre') labelGroupsPrePageTokens.pop();
      if (type === 'del') labelGroupsPrePageTokens.splice(0);
      const curLen = labelGroupsPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          labelGroupsList: result,
          labelGroupsNextPageToken: nextPageToken,
          labelGroupsPrePageToken: curLen ? labelGroupsPrePageTokens[curLen - 1] : (preLen ? '' : undefined),
          labelGroupsPrePageTokens,
          labelGroupsPageTotal: totalSize,
        },
      });
    },
    *getLabelUsers({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type, product, label } = payload;
      const { pageToken } = params;
      const { labelUsersPrePageTokens } = yield select(state => state.products);
      const preLen = labelUsersPrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(productsService.getLabelUsers, product, label, params);      
      if (type === 'next') labelUsersPrePageTokens.push(pageToken);
      if (type === 'pre') labelUsersPrePageTokens.pop();
      if (type === 'del') labelUsersPrePageTokens.splice(0);
      const curLen = labelUsersPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          labelUsersList: result,
          labelUsersNextPageToken: nextPageToken,
          labelUsersPrePageToken: curLen ? labelUsersPrePageTokens[curLen - 1] : (preLen ? '' : undefined),
          labelUsersPrePageTokens,
          labelUsersPageTotal: totalSize,
        },
      });
    },
    *getProductTags({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type, productName } = payload;
      const { pageToken } = params;
      const { prePageTokens } = yield select(state => state.products);
      const preLen = prePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(productsService.getProductsTag, productName, params);      
      if (type === 'next') prePageTokens.push(pageToken);
      if (type === 'pre') prePageTokens.pop();
      if (type === 'del') prePageTokens.splice(0);
      const curLen = prePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          productTagsList: result,
          nextPageToken,
          prePageToken: curLen ? prePageTokens[curLen - 1] : (preLen ? '' : undefined),
          prePageTokens,
          pageTotal: totalSize,
        },
      });
    },
    *updateProductTags({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, productName, cb } = payload;
      const { result } = yield call(productsService.updateProductsTag, productName, params);      
      if (result) {
        cb(result);
      }
    },
    *addProductTags({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, productName, cb } = payload;
      const { result } = yield call(productsService.addProductsTag, productName, params);      
      if (result) {
        cb();
      }
    },
    *publishProductTags({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, product, label, cb } = payload;
      const { result } = yield call(productsService.publishProductsTag, product, label, params);
      if (result) {
        cb();
      }
    },
    *publishProductSettings({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, product, module, setting, cb } = payload;
      const { result } = yield call(productsService.publishProductsSetting, product, module, setting, params);
      if (result) {
        cb();
      }
    },
    *getPublishRules({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { product, label, cb } = payload;
      const { result } = yield call(productsService.getPublishRules, product, label);
      if (result) {
        const userPercentRule = result.find(item => item.kind === 'userPercent');
        cb(userPercentRule);
      }
    },
    *getPublishSettingRules({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { product, module, setting, cb } = payload;
      const { result } = yield call(productsService.getPublishSettingRules, product, module, setting);
      if (result) {
        const userPercentRule = result.find(item => item.kind === 'userPercent');
        cb(userPercentRule);
      }
    },
    *updateProductTagRule({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { product, label, rule, params, cb } = payload;
      const { result } = yield call(productsService.updateProductTagRule, product, label, rule, params);
      if (result) {
        cb();
      }
    },
    *updateProductSettingRule({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { product, module, setting, rule, params, cb } = payload;
      const { result } = yield call(productsService.updateProductSettingRule, product, module, setting, rule, params);
      if (result) {
        cb();
      }
    },
    *getLabelLogs({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { params, product, label } = payload;
      const { result } = yield call(productsService.getLabelLogs, product, label, params);
      console.log('getLabelLogs', result);
      yield put({
        type: 'setStateByPayload',
        payload: {
          labelLogsList: result,
        },
      });
    },
    *getSettingLogs({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { params, product, module, setting } = payload;
      const { result } = yield call(productsService.getSettingLogs, product, module, setting, params);
      console.log('getSettingLogs', result);
      yield put({
        type: 'setStateByPayload',
        payload: {
          settingLogsList: result,
        },
      });
    },
    *offlineProductTags({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { label, productName, cb } = payload;
      const { result } = yield call(productsService.offlineProductsTag, productName, label);      
      if (result) {
        cb();
      }
    },
    *deleteProductTags({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { label, productName, cb } = payload;
      const { result } = yield call(productsService.deleteProductsTag, productName, label);      
      if (result) {
        cb();
      }
    },
    *getProductStatistics ({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { result } = yield call(productsService.getProductStatistics, payload.params.productName);
      yield put({
        type: 'setProductStatistics',
        payload: {
          productStatistics: result || {},
        }
      });
    },
    *getProductModules ({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type, productName } = payload;
      const { pageToken } = params;
      const { modulePrePageTokens } = yield select(state => state.products);
      const preLen = modulePrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(productsService.getProductsModule, productName, params);      
      if (type === 'next') modulePrePageTokens.push(pageToken);
      if (type === 'pre') modulePrePageTokens.pop();
      if (type === 'del') modulePrePageTokens.splice(0);
      const curLen = modulePrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          productModulesList: result,
          moduleNextPageToken: nextPageToken,
          modulePrePageToken: curLen ? modulePrePageTokens[curLen - 1] : (preLen ? '' : undefined),
          modulePrePageTokens,
          modulePageTotal: totalSize,
        },
      });
    },
    *updateProductModules({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, productName, cb } = payload;
      const { result } = yield call(productsService.updateProductsModule, productName, params);      
      if (result) {
        cb(result);
      }
    },
    *addProductModules({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, productName, cb } = payload;
      const { result } = yield call(productsService.addProductsModule, productName, params);      
      if (result) {
        cb();
      }
    },
    *offlineProductModules({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { module, productName, cb } = payload;
      const { result } = yield call(productsService.offlineProductsModule, productName, module);      
      if (result) {
        cb();
      }
    },
    *getProductSettings ({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type, productName } = payload;
      const { pageToken } = params;
      const { settingPrePageTokens } = yield select(state => state.products);
      const preLen = settingPrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(productsService.getProductsSetting, productName, params);      
      if (type === 'next') settingPrePageTokens.push(pageToken);
      if (type === 'pre') settingPrePageTokens.pop();
      if (type === 'del') settingPrePageTokens.splice(0);
      const curLen = settingPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          productSettingsList: result,
          settingNextPageToken: nextPageToken,
          settingPrePageToken: curLen ? settingPrePageTokens[curLen - 1] : (preLen ? '' : undefined),
          settingPrePageTokens,
          settingPageTotal: totalSize,
        },
      });
    },
    *getModuleSettings ({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type, productName, module } = payload;
      const { pageToken } = params;
      const { moduleSettingPrePageTokens } = yield select(state => state.products);
      const preLen = moduleSettingPrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(productsService.getModuleSetting, productName, module, params);      
      if (type === 'next') moduleSettingPrePageTokens.push(pageToken);
      if (type === 'pre') moduleSettingPrePageTokens.pop();
      if (type === 'del') moduleSettingPrePageTokens.splice(0);
      const curLen = moduleSettingPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          moduleSettingsList: result,
          moduleSettingNextPageToken: nextPageToken,
          moduleSettingPrePageToken: curLen ? moduleSettingPrePageTokens[curLen - 1] : (preLen ? '' : undefined),
          moduleSettingPrePageTokens,
          moduleSettingPageTotal: totalSize,
        },
      });
    },
    *updateProductSettings({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, productName, cb } = payload;
      const { result } = yield call(productsService.updateProductsSetting, productName, params);      
      if (result) {
        cb(result);
      }
    },
    *addProductSettings({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, productName, cb } = payload;
      const { result } = yield call(productsService.addProductsSetting, productName, params);      
      if (result) {
        cb();
      }
    },
    *offlineProductSettings({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { setting, module, productName, cb } = payload;
      const { result } = yield call(productsService.offlineProductsSetting, productName, module, setting);      
      if (result) {
        cb();
      }
    },
    *getSettingGroups({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type, product, module, setting } = payload;
      const { pageToken } = params;
      const { settingGroupsPrePageTokens } = yield select(state => state.products);
      const preLen = settingGroupsPrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(productsService.getSettingGroups, product, module, setting, params);      
      if (type === 'next') settingGroupsPrePageTokens.push(pageToken);
      if (type === 'pre') settingGroupsPrePageTokens.pop();
      if (type === 'del') settingGroupsPrePageTokens.splice(0);
      const curLen = settingGroupsPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          settingGroupsList: result,
          settingGroupsNextPageToken: nextPageToken,
          settingGroupsPrePageToken: curLen ? settingGroupsPrePageTokens[curLen - 1] : (preLen ? '' : undefined),
          settingGroupsPrePageTokens,
          settingGroupsPageTotal: totalSize,
        },
      });
    },
    *getSettingUsers({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type, product, module, setting } = payload;
      const { pageToken } = params;
      const { settingUsersPrePageTokens } = yield select(state => state.products);
      const preLen = settingUsersPrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(productsService.getSettingUsers, product, module, setting, params);      
      if (type === 'next') settingUsersPrePageTokens.push(pageToken);
      if (type === 'pre') settingUsersPrePageTokens.pop();
      if (type === 'del') settingUsersPrePageTokens.splice(0);
      const curLen = settingUsersPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          settingUsersList: result,
          settingUsersNextPageToken: nextPageToken,
          settingUsersPrePageToken: curLen ? settingUsersPrePageTokens[curLen - 1] : (preLen ? '' : undefined),
          settingUsersPrePageTokens,
          settingUsersPageTotal: totalSize,
        },
      });
    },
    *getProducts ({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { needRedirect, redirectTo } = payload;
      const { result } = yield call(productsService.getProductList);
      yield put({
        type: 'setProductsList',
        payload: {
          productList: result,
        },
      });
      if (needRedirect) {
        redirectTo(result[0].name)
      }
    },
    *updateProduct ({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { params, cb } = payload;
      const { name, desc, uids } = params;
      const { result } = yield call(productsService.updateProduct, name, desc, uids);
      if (result) {
        cb();
      }
    },
    *createProduct ({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { params, cb } = payload;
      const { name, desc, uids } = params;
      const { result } = yield call(productsService.createProduct, name, desc, uids);
      if (result) {
        cb();
      }
    },
    *offlineProduct ({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { params, cb } = payload;
      const { name } = params;
      const { result } = yield call(productsService.offlineProduct, name);
      if (result) {
        cb();
      }
    },
    *deleteProduct ({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { params, cb } = payload;
      const { name } = params;
      const { result } = yield call(productsService.deleteProduct, name);
      if (result) {
        cb();
      }
    },
  },
};

export default products;