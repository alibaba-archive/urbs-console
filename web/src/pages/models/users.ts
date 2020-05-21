import { AnyAction } from 'redux';
import * as usersService from '../services/users';
import { Model, EffectsCommandMap } from 'dva';

const users: Model = {
  namespace: 'users',
  state: {
    canaryUserList: [],
    prePageTokens: [],
    acUserList: [],
    labelsList: [],
    labelsPrePageTokens: [],
    settingsList: [],
    settingsPrePageTokens: [],
    acPrePageTokens: [],
    acNextPageTokens: [],
  },
  reducers: {
    setStateByPayload(state, { payload }: AnyAction) {
      return { ...state, ...payload };
    },
  },
  effects: {
    *getCanaryUsers({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type } = payload;
      const { pageToken } = params;
      const { prePageTokens } = yield select(state => state.users);
      const preLen = prePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(usersService.getCanaryUsers, params);
      if (type === 'next') prePageTokens.push(pageToken);
      if (type === 'pre') prePageTokens.pop();
      if (type === 'del') prePageTokens.splice(0);
      const curLen = prePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          canaryUserList: result,
          nextPageToken,
          prePageToken: curLen ? prePageTokens[curLen - 1] : (preLen ? '' : undefined),
          prePageTokens,
          pageTotal: totalSize,
        },
      });
    },
    *addCanaryUsers({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, cb } = payload;
      const { result } = yield call(usersService.addCanaryUsers, params.users);
      if (result) {
        cb();
      }
    },
    *getUserLabelsCache({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, cb } = payload;
      const { result } = yield call(usersService.getUserLabelsCache, params.uid);
      if (result) {
        cb(result);
      }
    },
    *getUserSettings({ payload }: AnyAction, { call, select, put }: EffectsCommandMap) {
      const { params, type, uid } = payload;
      const { pageToken } = params;
      const { settingsPrePageTokens } = yield select(state => state.users);
      const preLen = settingsPrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(usersService.getUserSettings, uid, params);
      if (type === 'next') settingsPrePageTokens.push(pageToken);
      if (type === 'pre') settingsPrePageTokens.pop();
      if (type === 'del') settingsPrePageTokens.splice(0);
      const curLen = settingsPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          settingsList: result,
          settingsNextPageToken: nextPageToken,
          settingsPrePageToken: curLen ? settingsPrePageTokens[curLen - 1] : (preLen ? '' : undefined),
          settingsPrePageTokens,
          settingsPageTotal: totalSize,
        },
      });
    },
    *getUserLabels({ payload }: AnyAction, { call, select, put }: EffectsCommandMap) {
      const { params, type, uid } = payload;
      const { pageToken } = params;
      const { labelsPrePageTokens } = yield select(state => state.users);
      const preLen = labelsPrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(usersService.getUserLabels, uid, params);
      if (type === 'next') labelsPrePageTokens.push(pageToken);
      if (type === 'pre') labelsPrePageTokens.pop();
      if (type === 'del') labelsPrePageTokens.splice(0);
      const curLen = labelsPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          labelsList: result,
          labelsNextPageToken: nextPageToken,
          labelsPrePageToken: curLen ? labelsPrePageTokens[curLen - 1] : (preLen ? '' : undefined),
          labelsPrePageTokens,
          labelsPageTotal: totalSize,
        },
      });
    },
    *deleteUserLabel({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { uid, hid, cb } = payload;
      yield call(usersService.deleteUserLabel, uid, hid);
      if (cb) {
        cb();
      }
    },
    *addAcUsers({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { users, cb } = payload;
      yield call(usersService.addAcUsers, users);
      yield put({
        type: 'setStateByPayload',
        payload: {},
      });
      if (cb) {
        cb();
      }
    },
    *deleteAcUser({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { uid, cb } = payload;
      const result = yield call(usersService.deleteAcUser, uid);
      if (result && cb) {
        cb();
      }
    },
    *updateAcUser({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { uid, name, cb } = payload;
      const result = yield call(usersService.updateAcUser, uid, name);
      if (result && cb) {
        cb();
      }
    },
    *getAcUsers({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { cb } = payload;
      const { result } = yield call(usersService.getAcUsers);
      yield put({
        type: 'setStateByPayload',
        payload: {
          acUserList: result || []
        },
      });
      if (cb) {
        cb(result || []);
      }
    },
    *getAcUsersList({ payload }, { call, select, put }: EffectsCommandMap) {
      const { params, type } = payload
      const { pageToken } = params;
      const { acPrePageTokens } = yield select(state => state.users);
      const preLen = acPrePageTokens.length;
      const { result, nextPageToken } = yield call(usersService.getAcUsersList, params);
      if (type === 'next') acPrePageTokens.push(pageToken);
      if (type === 'pre') acPrePageTokens.pop();
      if (type === 'del') acPrePageTokens.splice(0);
      const curLen = acPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          acNextPageToken: nextPageToken,
          acPrePageToken: curLen ? acPrePageTokens[curLen - 1] : (preLen ? '' : undefined),
          acPrePageTokens,
          acUserList: result || []
        },
      });
    },
    *searchAcUsers({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { params, cb } = payload;
      const { result } = yield call(usersService.searchAcUsers, params.key);
      yield put({
        type: 'setStateByPayload',
        payload: {
          acUserList: result || []
        },
      });
      if (cb) {
        cb(result || []);
      }
    },
    *deleteUserSetting({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { uid, product, module, setting, cb } = payload;
      const { result } = yield call(usersService.deleteUserSetting, product, module, setting, uid);
      if (result && cb) {
        cb();
      }
    },
    *rollbackUserSetting({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { uid, product, module, setting, cb } = payload;
      const { result } = yield call(usersService.rollbackUserSetting, product, module, setting, uid);
      if (result && cb) {
        cb();
      }
    },
    *getPermission ({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { cb } = payload;
      const { result } = yield call(usersService.getPermission);
      if (cb) {
        cb(result);
      }
    },
  },
};

export default users;