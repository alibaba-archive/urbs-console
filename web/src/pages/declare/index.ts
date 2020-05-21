import { RouteComponentProps } from 'react-router-dom';
import { TableProps } from 'antd/es/table';
import { ModalProps } from 'antd/es/modal';
import { Dispatch, AnyAction } from 'redux';

export const DEFAULT_MODAL_WIDTH = '800px';
export const DEFAULT_PAGE_SIZE = 20;
export const DEFAULT_PAGE_INDEX = 1;
export const DEFAULT_FORM_ITEM_LAYOUT = {
  labelCol: {
    xs: { span: 24 },
    sm: { span: 4 },
  },
  wrapperCol: {
    xs: { span: 24 },
    sm: { span: 20 },
  },
};
export const VERSION_CHANNEL = [{
  label: 'stable',
  value: 'stable',
}, {
  label: 'beta',
  value: 'beta',
}, {
  label: 'dev',
  value: 'dev',
}];

export const CLIENT_TYPE = [{
  label: 'web',
  value: 'web',
}, {
  label: 'ios',
  value: 'ios',
}, {
  label: 'android',
  value: 'android',
}];
export interface User {
  uid: string;
  name: string;
}

export interface Product {
  createdAt: string;
  updatedAt: string;
  name: string;
  desc: string;
  users: User[];
  status: number;
}

export interface ProductStatistics {
  labels: number;
  modules: number;
  settings: number;
  release: number;
}
interface RouteParams {
  name: string;
}

export interface BaseContainerComponentProps extends RouteComponentProps<RouteParams> {
  dispatch: Dispatch<AnyAction>;
}

export interface SettingComponentProps extends BaseContainerComponentProps {
  productSettingsList: Setting[];
  settingPrePageToken?: string;
  settingNextPageToken?: string;
  settingPageTotal?: number;
}

export interface ModulesComponentProps extends BaseContainerComponentProps {
  productModulesList: Module[];
  modulePrePageToken?: string;
  moduleNextPageToken?: string;
  modulePageTotal?: number;
}

export interface PublishRecordItem {
  operatorName: string;
  action: string;
  desc: string;
  kind: string;
  percent: number;
  groups: string[];
  users: string[];
  createdAt: string;
  hid: string;
}
interface ContentDetailItem {
  title: React.ReactNode;
  content: React.ReactNode;
}
export interface ContentDetailComponentProps {
  content?: ContentDetailItem[];
}

export interface ContentTabsItem {
  key: string;
  title: React.ReactNode;
  content: React.ReactNode;
  action?: React.ReactNode;
}
export interface ContentTabsComponentProps {
  tabs: ContentTabsItem[];
  handleActiveKeyChange?: (key: string) => void;
  activeKey?: string;
}

export interface UserPercentRule {
  hid: string;
  kind: string;
  rule: {
    value: number
  };
}

export interface ModuleDetailModalComponentProps {
  dispatch: Dispatch<AnyAction>;
  visible: boolean;
  product: string;
  moduleSettingsList: Setting[];
  moduleSettingNextPageToken?: string;
  moduleSettingPrePageToken?: string;
  moduleSettingPageTotal?: number,
  moduleInfo?: Module;
  onCancel?: (e: React.MouseEvent) => void;
  title?: React.ReactNode;
  onModuleEdit: () => void;
}

export interface PaginationParams {
  pageSize: number;
  skip: number;
}

export interface TableTitleComponentProps {
  plusTitle: React.ReactNode;
  defaultSearchValue?: string;
  handlePlusClick?: React.MouseEventHandler<HTMLElement>;
  handleWordChange?: (searchValue: string) => void;
  handleSearch?: (searchValue: string) => void;
}

export interface TagComponentProps extends BaseContainerComponentProps {
  productTagsList: Label[];
  prePageToken?: string;
  nextPageToken?: string;
  pageTotal?: number;
}

export enum TagTabsKey {
  Publish = 'publish',
  Group = 'group',
  User = 'user',
}

export interface SettingDetailComponentProps extends ModalProps {
  product: string;
  onSettingEdit: () => void;
  dispatch: Dispatch<AnyAction>;
  settingInfo?: Setting;
  onGotoGroups?: () => void;
  onGotoUsers?: () => void;
}

export interface TagDetailComponentProps extends ModalProps {
  product: string;
  onSettingEdit: () => void;
  dispatch: Dispatch<AnyAction>;
  labelInfo?: Label;
  onGotoGroups?: () => void;
  onGotoUsers?: () => void;
}

export interface UserSelectComponentProps {
  acUserList: User[];
  dispatch: Dispatch<AnyAction>;
  placeholder?: string;
  defaultSelectedUser?: User[];
  onChange?: (users: User[]) => void;
}

export interface ProductsComponentProps extends BaseContainerComponentProps {
  productList: Product[];
  productStatistics: ProductStatistics;
}

export interface CanaryUser {
  uid: string;
  createdAt: string;
  activeAt: string;
  labels: string;
}

export interface Module {
  name: string;
  desc: string;
  createdAt: string;
  updatedAt: string;
  offlineAt: string;
  users: User[];
  status: number;
}

export interface Setting {
  hid: string;
  name: string;
  desc: string;
  module: string;
  product: string;
  createdAt: string;
  updatedAt: string;
  offlineAt: string;
  users: User[];
  status: number;
  release: number;
  channels?: string[];
  clients?: string[];
  values?: string[];
}

export interface Label {
  hid: string;
  name: string;
  desc: string;
  product: string;
  createdAt: string;
  activeAt: string;
  updatedAt: string;
  users: User[];
  status: number;
  release: number;
  channels?: string[];
  clients?: string[];
}

export interface UsersComponentProps extends BaseContainerComponentProps {
  dispatch: Dispatch<AnyAction>;
  canaryUserList: CanaryUser[];
  prePageToken?: string;
  nextPageToken?: string;
  pageTotal?: number;
  labelsList: Label[],
  labelsPageTotal?: number;
  labelsNextPageToken?: string;
  labelsPrePageToken?: string;
}

export interface ACComponentProps extends BaseContainerComponentProps {
  dispatch: Dispatch<AnyAction>;
  acUserList: User[];
  onChange?: (users: string[]) => void;
  acPrePageToken?: string;
  acNextPageToken?: string;
  pageTotal?: number;
}

export interface AcUserSelectComponentProps {
  defaultSelectedUser?: User[];
  dispatch: Dispatch<AnyAction>;
  onChange?: (users: string[]) => void;
}

export interface PaginationParameters {
  pageSize: number;
  pageToken?: string;
  q?: string;
}

export interface PaginationComponentProps {
  pageSize?: number;
  pageSizeOptions?: number[];
  total?: number;
  prePageToken?: string;
  nextPageToken?: string;
  onTokenChange?: (type: string, token?: string) => void;
  onPageSizeChange?: (size: number) => void;
}

export interface ActionEventListeners {
  onDelete: (e: React.MouseEvent) => void;
  onRollback?: (e: React.MouseEvent) => void;
}

export interface TableComponentProps<T> extends TableProps<T> {
  hideColumns?: string[];
  paginationProps?: PaginationComponentProps;
  onAction?: (record: T) => ActionEventListeners;
}

export interface Group {
  uid: string;
  group: string;
  kind: string;
  desc: string;
  status: number;
  syncAt: string;
  updatedAt: string;
  createdAt: string;
}

export interface GroupMember {
  user: string;
  createdAt: string;
  syncAt: string;
}

export interface GroupsComponentProps {
  dispatch: Dispatch<AnyAction>;
  groupList: Group[];
  prePageToken?: string;
  nextPageToken?: string;
  pageTotal?: number;
  labelsList: Label[];
  labelsPrePageToken?: string;
  labelsNextPageToken?: string;
  labelsPageTotal?: number;
  membersList: GroupMember[];
  membersPrePageToken?: string;
  membersNextPageToken?: string;
  membersPageTotal?: number;
}

export interface FieldsValue {
  [field: string]: any;
}
