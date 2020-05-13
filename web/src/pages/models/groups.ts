import { AnyAction } from 'redux';
import { Model, EffectsCommandMap } from 'dva';
import * as groupsService from '../services/groups';

const groups: Model = {
  namespace: 'groups',
  state: {
    groupList: [],
    prePageTokens: [],
    labelsList: [],
    labelsPrePageTokens: [],
    membersList: [],
    membersPrePageTokens: [],
    settingsList: [],
    settingsPrePageTokens: [],
  },
  reducers: {
    setStateByPayload(state, { payload }: AnyAction) {
      return { ...state, ...payload };
    },
  },
  effects: {
    *getGroups({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type } = payload;
      const { pageToken } = params;
      const { prePageTokens } = yield select(state => state.groups);
      const preLen = prePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(groupsService.getGroups, params);
      if (type === 'next') prePageTokens.push(pageToken);
      if (type === 'pre') prePageTokens.pop();
      if (type === 'del') prePageTokens.splice(0);
      const curLen = prePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          groupList: result,
          nextPageToken,
          prePageToken: curLen ? prePageTokens[curLen - 1] : (preLen ? '' : undefined),
          prePageTokens,
          pageTotal: totalSize,
        },
      });
    },
    *addGroups({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, cb } = payload;
      const { result } = yield call(groupsService.addGroups, params);
      if (result) {
        cb();
      }
    },
    *updateGroups({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, cb } = payload;
      const { result } = yield call(groupsService.updateGroups, params);
      if (result) {
        cb();
      }
    },
    *deleteGroups({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { params, cb } = payload;
      const { result } = yield call(groupsService.deleteGroups, params.uid);
      if (result) {
        cb();
      }
    },
    *getGroupLabels({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, type, uid } = payload;
      const { pageToken } = params;
      const { labelsPrePageTokens } = yield select(state => state.groups);
      const preLen = labelsPrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(groupsService.getGroupLabels, uid, params);
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
    *deleteGroupLabel({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { uid, hid, cb } = payload;
      const { result } = yield call(groupsService.deleteGroupLabel, uid, hid);
      if (result && cb) {
        cb();
      }
    },
    *getGroupSettings({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, uid, type } = payload;
      const { pageToken } = params;
      const { settingsPrePageTokens } = yield select(state => state.groups);
      const preLen = settingsPrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(groupsService.getGroupSettings, uid, params);
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
    *getGroupMembers({ payload }: AnyAction, { call, put, select }: EffectsCommandMap) {
      const { params, uid, type } = payload;
      const { pageToken } = params;
      const { membersPrePageTokens } = yield select(state => state.groups);
      const preLen = membersPrePageTokens.length;
      const { result, nextPageToken, totalSize } = yield call(groupsService.getGroupMembers, uid, params);
      if (type === 'next') membersPrePageTokens.push(pageToken);
      if (type === 'pre') membersPrePageTokens.pop();
      if (type === 'del') membersPrePageTokens.splice(0);
      const curLen = membersPrePageTokens.length;
      yield put({
        type: 'setStateByPayload',
        payload: {
          membersList: result,
          membersNextPageToken: nextPageToken,
          membersPrePageToken: curLen ? membersPrePageTokens[curLen - 1] : (preLen ? '' : undefined),
          membersPrePageTokens,
          membersPageTotal: totalSize,
        },
      });
    },
    *addGroupMembers({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { uid, params, cb } = payload;
      const { result } = yield call(groupsService.addGroupMembers, uid, params.users);
      if (result && cb) {
        cb();
      }
    },
    *deleteGroupMembers({ payload }: AnyAction, { call }: EffectsCommandMap) {
      const { uid, params, cb } = payload;
      const { result } = yield call(groupsService.deleteGroupMembers, uid, params.user);
      if (result && cb) {
        cb();
      }
    },
  },
};

export default groups;