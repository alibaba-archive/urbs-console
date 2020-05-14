
import React, { useEffect, useState, useMemo, useCallback } from 'react';
import { Button, Input, Modal, Icon } from 'antd';
import { connect } from 'dva';
import { PublishRecord, ContentDetail, ContentTabs, UserGroup, Users, GrayscaleTagModifyModal, PublishTagModal } from '../';
import { TagDetailComponentProps, TagTabsKey, PaginationParameters, FieldsValue, UserPercentRule, DEFAULT_PAGE_SIZE } from '../../declare';
import styles from '../style/TagDetailModal.less';
import { formatTableTime } from '../../utils/format';

const DEFAULT_TITLE = '灰度标签';

const TagDetailModal: React.FC<TagDetailComponentProps> = (props) => {
  const {
    product,
    onSettingEdit,
    onCancel,
    visible,
    title = DEFAULT_TITLE,
    labelInfo,
    dispatch,
    onGotoGroups,
    onGotoUsers,
    labelLogsList,
    labelGroupsList,
    labelGroupsNextPageToken,
    labelGroupsPrePageToken,
    labelGroupsPageTotal,
    labelUsersList,
    labelUsersNextPageToken,
    labelUsersPrePageToken,
    labelUsersPageTotal,
  } = props as any;
  const [labelGroupPageSize, changeLabelGroupPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [labelUserPageSize, changeLabelUserPageSize] = useState(DEFAULT_PAGE_SIZE);

  const [userPercentRule, changeUserPercentRule] = useState<UserPercentRule>();
  const [tabsActiveKey, setTabsActiveKey] = useState(String(TagTabsKey.Publish));
  const [tabsSearchWord, setTabsSearchWord] = useState('');
  const [publishTagModalVisible, setPublishTagModalVisible] = useState(false);
  const [grayscaleTagModalVisible, setGrayscaleTagModalVisible] = useState(false);
  const fetchLabelLogs = useCallback(() => {
    dispatch({
      type: 'products/getLabelLogs',
      payload: {
        params: {
          pageSize: 100
        },
        product,
        label: labelInfo?.name,
      },
    });
  }, [dispatch, labelInfo, product]);
  const fetchLabelGroups = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'products/getLabelGroups',
      payload: {
        type,
        params,
        product,
        label: labelInfo?.name,
      },
    });
  }, [dispatch, labelInfo, product]);
  const fetchLabelUsers = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'products/getLabelUsers',
      payload: {
        type,
        params,
        product,
        label: labelInfo?.name,
      },
    });
  }, [dispatch, labelInfo, product]);
  useEffect(() => {
    fetchLabelLogs();
    fetchLabelGroups({
      pageSize: labelGroupPageSize,
    });
    fetchLabelUsers({
      pageSize: labelUserPageSize,
    });
  }, [fetchLabelLogs, fetchLabelGroups, fetchLabelUsers, labelGroupPageSize, labelUserPageSize]);
  const handleTabsActiveKeyChange = (activeKey: string) => {
    setTabsActiveKey(activeKey);
    setTabsSearchWord('');
    switch (activeKey) {
      case TagTabsKey.Publish:
        fetchLabelLogs();
        break;
      case TagTabsKey.Group:
        fetchLabelGroups({
          pageSize: labelGroupPageSize,
        }, 'del');
        break;
      case TagTabsKey.User:
        fetchLabelUsers({
          pageSize: labelUserPageSize,
        }, 'del');
        break;
      default:
        break;
    }
  };
  const handleTabsSearch = (searchWord: string) => {
    switch (tabsActiveKey) {
      case TagTabsKey.Group:
        fetchLabelGroups({
          pageSize: labelGroupPageSize,
          q: searchWord,
        });
        break;
      case TagTabsKey.User:
        fetchLabelUsers({
          pageSize: labelUserPageSize,
          q: searchWord,
        });
        break;
      default:
        break;
    }
  };
  const handleTabsSearchWordChange = (e: React.ChangeEvent) => {
    const nativeEvent = e.nativeEvent;
    const target = nativeEvent.target || nativeEvent.srcElement;
    setTabsSearchWord((target as any).value);
  };
  const changePublishTagModalVisible = (visible: boolean) => {
    setPublishTagModalVisible(visible);
  };
  const handleOpenPublishTagModal = () => {
    dispatch({
      type: 'products/getPublishRules',
      payload: {
        product,
        label: labelInfo?.name,
        cb: (rule?: UserPercentRule) => {
          changeUserPercentRule(rule);
          changePublishTagModalVisible(true);
        }
      }
    });
  };
  const handleOpenPublishTagModalOk = (values: FieldsValue) => {
    if (userPercentRule && values.kind === 'userPercent') {
      dispatch({
        type: 'products/updateProductTagRule',
        payload: {
          product,
          label: labelInfo?.name,
          rule: userPercentRule.hid,
          params: values,
          cb: () => {
            fetchLabelLogs();
            changePublishTagModalVisible(false);
          }
        },
      });
    } else {
      dispatch({
        type: 'products/publishProductTags',
        payload: {
          product,
          label: labelInfo?.name,
          params: values,
          cb: () => {
            changePublishTagModalVisible(false);
          }
        },
      });
    }
  };
  const handleOpenPublishTagModalCancel = () => {
    changePublishTagModalVisible(false);
  };
  const renderModalTitle = () => {
    return (
      <div className={styles['tag-modal-title']}>
        <div>{title}</div>
        <div>
          <Icon type="setting" onClick={onSettingEdit}></Icon>
        </div>
      </div>
    )
  };
  const tagTabsConfig = [{
    key: TagTabsKey.Publish,
    title: '发布记录',
    content: (
      <PublishRecord
        publishRecordList={labelLogsList}
      />
    ),
    action: (
      <Button
        type="link"
        icon="plus"
        block
        onClick={handleOpenPublishTagModal}
      >
        添加灰度发布
      </Button>
    ),
  }, {
    key: TagTabsKey.Group,
    title: '群组',
    content: (
      <UserGroup
        dataSource={labelGroupsList}
        hideColumns={['syncAt']}
        paginationProps={
          {
            total: labelGroupsPageTotal,
            nextPageToken: labelGroupsNextPageToken,
            prePageToken: labelGroupsPrePageToken,
            pageSize: labelGroupPageSize,
            pageSizeOptions: [10, 20, 30, 40],
            onPageSizeChange: (size) => {
              changeLabelGroupPageSize(size);
              fetchLabelGroups({
                pageSize: size,
              }, 'del');
            },
            onTokenChange: (type, token) => {
              fetchLabelGroups({
                pageSize: labelGroupPageSize,
                pageToken: token,
              }, type);
            }
          }
        }
      />
    ),
    action: (
      <Input.Search
        value={tabsSearchWord}
        placeholder="请输入搜索关键字"
        onChange={handleTabsSearchWordChange}
        onSearch={handleTabsSearch}
      />
    ),
  }, {
    key: TagTabsKey.User,
    title: '用户',
    content: (
      <Users
        dataSource={labelUsersList}
        hideColumns={['syncAt']}
        paginationProps={
          {
            total: labelUsersPageTotal,
            nextPageToken: labelUsersNextPageToken,
            prePageToken: labelUsersPrePageToken,
            pageSize: labelUserPageSize,
            pageSizeOptions: [10, 20, 30, 40],
            onPageSizeChange: (size) => {
              changeLabelUserPageSize(size);
              fetchLabelUsers({
                pageSize: size,
              }, 'del');
            },
            onTokenChange: (type, token) => {
              fetchLabelUsers({
                pageSize: labelGroupPageSize,
                pageToken: token,
              }, type);
            }
          }
        }
      />
    ),
    action: (
      <Input.Search
        value={tabsSearchWord}
        placeholder="请输入搜索关键字"
        onChange={handleTabsSearchWordChange}
        onSearch={handleTabsSearch}
      />
    ),
  }];
  const labelContentDetail = useMemo(() => {
    return labelInfo && ([
      {
        title: '名称',
        content: labelInfo.name,
      },
      {
        title: '所属产品',
        content: product,
      },
      {
        title: '负责人',
        content: Array.isArray(labelInfo.users) ? labelInfo.users.map(item => item.name).join(',') : '',
      },
      {
        title: '版本通道',
        content: Array.isArray(labelInfo.channels) ? labelInfo.channels.join(',') : '空',
      },
      {
        title: '端类型',
        content: Array.isArray(labelInfo.clients) ? labelInfo.clients.join(',') : '空',
      },
      {
        title: '发布次数',
        content: labelInfo.release,
      },
      {
        title: '灰度进度',
        content: labelInfo.status,
      },
      {
        title: '创建时间',
        content: formatTableTime(labelInfo.createdAt),
      },
      {
        title: '更新时间',
        content: formatTableTime(labelInfo.updatedAt),
      },
    ]);
  }, [labelInfo, product]);
  return (
    <Modal title={renderModalTitle()} visible={visible} onCancel={onCancel} footer={null}>
      <ContentDetail content={labelContentDetail}></ContentDetail>
      <ContentTabs
        activeKey={tabsActiveKey}
        handleActiveKeyChange={handleTabsActiveKeyChange}
        tabs={tagTabsConfig}
      />
      {/* 弹窗 */}
      {
        publishTagModalVisible && <PublishTagModal
          title="发布灰度标签"
          visible={publishTagModalVisible}
          onCancel={handleOpenPublishTagModalCancel}
          onOk={handleOpenPublishTagModalOk}
          label={labelInfo?.name}
          product={product}
          onGotoGroups={onGotoGroups}
          onGotoUsers={onGotoUsers}
          defauleRule={userPercentRule}
        ></PublishTagModal>
      }
      <GrayscaleTagModifyModal
        visible={grayscaleTagModalVisible}
        isEdit={true}
        onCancel={() => setGrayscaleTagModalVisible(false)}
        onOk={() => setGrayscaleTagModalVisible(false)}
      ></GrayscaleTagModifyModal>
    </Modal>
  );
};

export default connect((state) => {
  return {
    ...(state as any).products
  };
})(TagDetailModal);
