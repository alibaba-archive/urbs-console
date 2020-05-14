import React, { useState, useEffect, useCallback, useMemo } from 'react';
import { Icon, Modal, Button, Input, Form, message } from 'antd';
import { connect } from 'dva';
import { GrayscaleTag, Setting, Users, UserGroup, TableTitle, GroupModifyModal, ContentTabs, ContentDetail } from '../components';
import { DEFAULT_FORM_ITEM_LAYOUT, DEFAULT_PAGE_SIZE, PaginationParameters, Group, GroupsComponentProps, FieldsValue, Label, GroupMember } from '../declare';
import { formatTableTime } from '../utils/format';
import styleNames from '../components/style/base.less';

const Groups: React.FC<GroupsComponentProps> = (props) => {
  const {
    dispatch,
    groupList,
    prePageToken,
    nextPageToken,
    pageTotal,
    labelsList,
    labelsPrePageToken,
    labelsNextPageToken,
    labelsPageTotal,
    membersList,
    membersPrePageToken,
    membersNextPageToken,
    membersPageTotal,
  } = props;
  const [pageSize, setPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [currentGroup, setCurrentGroup] = useState<Group>();
  const [labelsPageSize, setLabelsPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [batchUsers, setBatchUsers] = useState('');
  const [settingsPageSize, setSettingsPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [membersPageSize, setMembersPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [groupModifyVisible, changeGroupModifyVisible] = useState(false);
  const [groupDetailVisible, changeGroupDetailVisible] = useState(false);
  const [groupUserAddVisible, changeGroupUserAddVisible] = useState(false);

  const fetchGroupList = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'groups/getGroups',
      payload: {
        params,
        type,
      }
    });
  }, [dispatch]);

  const fetchLabelList = useCallback((params: PaginationParameters, uid: string, type?: string) => {
    dispatch({
      type: 'groups/getGroupLabels',
      payload: {
        params,
        uid,
        type,
      }
    });
  }, [dispatch]);

  const fetchSettingList = useCallback((params: PaginationParameters, uid: string, type?: string) => {
    dispatch({
      type: 'groups/getGroupSettings',
      payload: {
        params,
        uid,
        type,
      }
    });
  }, [dispatch]);

  const fetchMemberList = useCallback((params: PaginationParameters, uid: string, type?: string) => {
    dispatch({
      type: 'groups/getGroupMembers',
      payload: {
        params,
        uid,
        type,
      }
    });
  }, [dispatch]);

  const handleGroupSearch = (key: string) => {
    fetchGroupList({
      pageSize,
      q: key
    }, 'del');
  };

  const handlePageSizeChange = (size: number) => {
    fetchGroupList({
      pageSize,
    }, 'del');
    setPageSize(size);
  };

  const handleTokenChange = (type: string, pageToken?: string) => {
    fetchGroupList({
      pageSize,
      pageToken,
    }, type);
  };

  const handleAddGroup = (fieldsValue: FieldsValue, isEdit?: boolean) => {
    if (isEdit) {
      dispatch({
        type: 'groups/updateGroups',
        payload: {
          params: {...currentGroup, ...fieldsValue},
          cb: () => {
            fetchGroupList({
              pageSize,
            }, 'del');
            message.success('修改群组成功');
            changeGroupModifyVisible(false);
          }
        }
      });
    } else {
      dispatch({
        type: 'groups/addGroups',
        payload: {
          params: fieldsValue,
          cb: () => {
            fetchGroupList({
              pageSize,
            }, 'del');
            message.success('添加群组成功');
            changeGroupModifyVisible(false);
          }
        }
      });
    }
  };

  const handleDelGroup = () => {
    dispatch({
      type: 'groups/deleteGroups',
      payload: {
        params: currentGroup,
        cb: () => {
          fetchGroupList({
            pageSize,
          }, 'del');
          message.success('删除群组成功');
          changeGroupModifyVisible(false);
          changeGroupDetailVisible(false);
        }
      }
    });
  };

  const handleBatchUsersChange = (e: React.ChangeEvent) => {
    const value = e.target.value;
    setBatchUsers(value);
  };

  const handleBatchUsersOk = () => {
    if (!batchUsers) {
      message.info('请输入用户uid');
      return;
    }
    dispatch({
      type: 'groups/addGroupMembers',
      payload: {
        uid: currentGroup?.uid,
        cb: () => {
          message.success('添加用户成功');
          changeGroupUserAddVisible(false);
          fetchMemberList({
            pageSize,
          }, currentGroup?.uid as string, 'del');
        },
        params: {
          users: batchUsers.split(';'),
        },
      }
    });
  };

  useEffect(() => {
    fetchGroupList({
      pageSize,
    });
  }, [fetchGroupList, pageSize]);

  const handleOnRow = (record: Group) => {
    return {
      onDoubleClick: () => {
        fetchLabelList({
          pageSize: labelsPageSize,
        }, record.uid);
        fetchMemberList({
          pageSize: membersPageSize,
        }, record.uid);
        fetchSettingList({
          pageSize: settingsPageSize,
        }, record.uid);
        setCurrentGroup(record);
        changeGroupDetailVisible(true);
      }
    };
  };
  const tabsConfig = [{
    key: 'label',
    title: '灰度标签',
    content: (
      <GrayscaleTag
        hideColumns={ ['users', 'status', 'release', 'updatedAt'] }
        dataSource={ labelsList }
        onAction={
          (record: Label) => ({
            onDelete: () => {
              dispatch({
                type: 'groups/deleteGroupLabel',
                payload: {
                  uid: currentGroup?.uid,
                  hid: record.hid,
                  cb: () => {
                    fetchLabelList({
                      pageSize: labelsPageSize,
                    }, (currentGroup?.uid) as string, 'del');
                  }
                }
              });
            },
          })
        }
        paginationProps={
          {
            total: labelsPageTotal,
            prePageToken: labelsPrePageToken,
            nextPageToken: labelsNextPageToken,
            onPageSizeChange: (size: number) => {
              setLabelsPageSize(size);
              fetchLabelList({
                pageSize: size
              }, currentGroup?.uid as string, 'del');
            },
            onTokenChange: (type: string, token?: string) => {
              fetchLabelList({
                pageSize: labelsPageSize,
                pageToken: token,
              }, currentGroup?.uid as string, type);
            },
          }
        }
      />
    ),
    action: (
      <Input.Search
        key="group"
        placeholder="请输入搜索关键字"
      />
    ),
  }, {
    key: 'setting',
    title: '配置项',
    content: (
      <Setting
        hideColumns={ [] }
      />
    ),
    action: (
      <Input.Search
        key="group"
        placeholder="请输入搜索关键字"
      />
    ),
  }, {
    key: 'user',
    title: '用户',
    content: (
      <Users
        dataSource={ membersList }
        onAction={
          (record: GroupMember) => ({
            onDelete: () => {
              dispatch({
                type: 'groups/deleteGroupMembers',
                payload: {
                  uid: currentGroup?.uid,
                  params: {
                    user: record.user
                  },
                  cb: () => {
                    fetchMemberList({
                      pageSize: membersPageSize,
                    }, (currentGroup?.uid) as string, 'del');
                    message.success('移除用户成功');
                  }
                }
              });
            },
          })
        }
        paginationProps={
          {
            total: membersPageTotal,
            prePageToken: membersPrePageToken,
            nextPageToken: membersNextPageToken,
            onPageSizeChange: (size: number) => {
              setMembersPageSize(size);
              fetchMemberList({
                pageSize: size
              }, currentGroup?.uid as string, 'del');
            },
            onTokenChange: (type: string, token?: string) => {
              fetchMemberList({
                pageSize: membersPageSize,
                pageToken: token,
              }, currentGroup?.uid as string, type);
            },
          }
        }
      />
    ),
    action: (
      <div>
        <Button
          type="link"
          icon="plus"
          block
          onClick={ () => changeGroupUserAddVisible(true) }
        >
          添加成员
        </Button>
        <Input.Search
          key="user"
          placeholder="请输入搜索关键字"
        />
      </div>
    ),
  }];
  const contentDetail = useMemo(() => {
    return currentGroup ? [{
      title: 'uid',
      content: currentGroup.uid,
    }, {
      title: '类型',
      content: currentGroup.kind,
    }, {
      title: '描述',
      content: currentGroup.desc,
    }, {
      title: '成员数量',
      content: currentGroup.status,
    }, {
      title: '同步时间',
      content: (
        <div style={{ display: 'flex', position: 'relative', top: '-5px' }}>
          <Button style={{ padding: '0' }} disabled type="link">{ formatTableTime(currentGroup.syncAt || '') }</Button>
          <Button icon="reload" type="link">重新同步</Button>
        </div>
      ),
    }, {
      title: '创建日期',
      content: formatTableTime(currentGroup.createdAt || ''),
    }] : undefined;
}, [currentGroup]);

  const handlePlusClick = () => {
    setCurrentGroup(undefined);
    changeGroupModifyVisible(true);
  };
  return (
    <div>
      <TableTitle
        plusTitle="添加群组"
        handlePlusClick={ handlePlusClick }
        handleSearch={ handleGroupSearch }
      />
      <UserGroup
        onRow={ handleOnRow }
        hideColumns={ ['action'] }
        dataSource={ groupList }
        paginationProps={{
          pageSize,
          pageSizeOptions: [10, 20, 30, 40],
          nextPageToken,
          prePageToken,
          total: pageTotal,
          onTokenChange: handleTokenChange,
          onPageSizeChange: handlePageSizeChange
        }}></UserGroup>
      {/* 弹窗 */}
      {
        groupModifyVisible && <GroupModifyModal
          visible={ groupModifyVisible }
          onCancel={ () => changeGroupModifyVisible(false) }
          groupInfo={ currentGroup }
          onOk={ handleAddGroup }
          onDel={ handleDelGroup }
        />
      }
      {/* 弹窗 */}
      <Modal
        visible={ groupDetailVisible }
        title={
          <div className={styleNames['modal-title']}>
            <div>群组</div>
            <div>
              <Icon type="setting" onClick={ () => changeGroupModifyVisible(true) }></Icon>
            </div>
          </div>
        }
        width="50%"
        footer={ null }
        onCancel={ () => changeGroupDetailVisible(false) }
      >
        <ContentDetail
          content={ contentDetail }
        />
        <ContentTabs
          tabs={ tabsConfig }
        />
      </Modal>
      {
        groupUserAddVisible && (<Modal
          title="添加成员"
          visible={ groupUserAddVisible }
          onCancel={ () => changeGroupUserAddVisible(false) }
          onOk={ handleBatchUsersOk }
        >
          <Form { ...DEFAULT_FORM_ITEM_LAYOUT }>
            <Form.Item
              label="群组"
              style={{marginBottom: '0'}}
            >
              <span>{currentGroup?.uid}</span>
            </Form.Item>
            <Form.Item
              label="类型"
              style={{marginBottom: '0'}}
            >
              <span>{currentGroup?.kind}</span>
            </Form.Item>
            <Form.Item
              label="批量用户"
              style={{marginBottom: '0'}}
              help="添加多个用户使用英文;间隔"
            >
              <Input onChange={ handleBatchUsersChange } value={ batchUsers } placeholder="请输入用户uid"></Input>
            </Form.Item>
          </Form>
        </Modal>)
      }
    </div>
  );
}

export default connect((state) => {
  const {
    groupList,
    prePageToken,
    nextPageToken,
    pageTotal,
    labelsList,
    labelsPrePageToken,
    labelsNextPageToken,
    labelsPageTotal,
    membersList,
    membersPrePageToken,
    membersNextPageToken,
    membersPageTotal,
  } = (state as any).groups;
  return {
    groupList,
    prePageToken,
    nextPageToken,
    pageTotal,
    labelsList,
    labelsPrePageToken,
    labelsNextPageToken,
    labelsPageTotal,
    membersList,
    membersPrePageToken,
    membersNextPageToken,
    membersPageTotal,
  };
})(Groups);