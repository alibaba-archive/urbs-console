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
  },
  reducers: {
    setStateByPayload(state, { payload }: AnyAction) {
      return { ...state, ...payload };
    },
  },
  effects: {
    *getCanaryUsers({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type } = payload;
      const { prePageTokens } = yield select(state => state.users);
      const { result, nextPageToken, totalSize } = yield call(usersService.getCanaryUsers, params);      
      if (type === 'next') prePageTokens.push(nextPageToken);
      if (type === 'pre') prePageTokens.pop();
      if (type === 'del') prePageTokens.splice(0);
      const len = prePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          canaryUserList: result,
          nextPageToken,
          prePageToken: len ? prePageTokens[len - 1] : undefined,
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
      const { settingsPrePageTokens } = yield select(state => state.users);
      const { result, nextPageToken, totalSize } = yield call(usersService.getUserSettings, uid, params);
      if (type === 'next') settingsPrePageTokens.push(nextPageToken);
      if (type === 'pre') settingsPrePageTokens.pop();
      if (type === 'del') settingsPrePageTokens.splice(0);
      const len = settingsPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          settingsList: result,
          settingsNextPageToken: nextPageToken,
          settingsPrePageToken: len ? settingsPrePageTokens[len - 1] : undefined,
          settingsPrePageTokens,
          settingsPageTotal: totalSize,
        },
      });
    },
    *getUserLabels({ payload }: AnyAction, { call, select, put }: EffectsCommandMap) {
      const { params, type, uid } = payload;
      const { labelsPrePageTokens } = yield select(state => state.users);
      const { result, nextPageToken, totalSize } = yield call(usersService.getUserLabels, uid, params);
      if (type === 'next') labelsPrePageTokens.push(nextPageToken);
      if (type === 'pre') labelsPrePageTokens.pop();
      if (type === 'del') labelsPrePageTokens.splice(0);
      const len = labelsPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          labelsList: result,
          labelsNextPageToken: nextPageToken,
          labelsPrePageToken: len ? labelsPrePageTokens[len - 1] : undefined,
          labelsPrePageTokens,
          labelsPageTotal: totalSize,
        },
      });
    },
    *deleteUserLabel({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { uid, hid, cb } = payload;
      const { result } = yield call(usersService.deleteUserLabel, uid, hid);
      if (result && cb) {
        cb();
      }
    },
    *getAcUsers(_, { call, put }: EffectsCommandMap) {
      const { result } = yield call(usersService.getAcUsers);
      yield put({
        type: 'setStateByPayload',
        payload: {
          acUserList: result || []
        },
      });
    },
    *searchAcUsers({ payload }: AnyAction, { call, put }: EffectsCommandMap) {
      const { params } = payload;
      const { result } = yield call(usersService.searchAcUsers, params.key);
      yield put({
        type: 'setStateByPayload',
        payload: {
          acUserList: result || []
        },
      });
    },
  },
};

export default users;