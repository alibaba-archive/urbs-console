
import React, { useEffect, useState, useMemo, useCallback } from 'react';
import { Form, Modal, Table, Input, Button, message } from 'antd';
import { connect } from 'dva';
import { Label, UsersComponentProps, DEFAULT_MODAL_WIDTH, DEFAULT_FORM_ITEM_LAYOUT, DEFAULT_PAGE_SIZE, PaginationParameters, CanaryUser } from '../declare';
import { Pagination, TableTitle, ContentTabs, ContentDetail, Setting, GrayscaleTag } from '../components';
import { formatTableTime } from '../utils/format';

const Users: React.FC<UsersComponentProps> = (props) => {
  const {
    dispatch,
    canaryUserList,
    prePageToken,
    nextPageToken,
    pageTotal,
    labelsList,
    labelsPageTotal,
    labelsNextPageToken,
    labelsPrePageToken,
  } = props;
  const [currentUser, setCurrentUser] = useState<CanaryUser>();
  const [userAddModalVisible, setUserAddModalVisible] = useState(false);
  const [userDetailModalVisible, setUserDetailModalVisible] = useState(false);
  const [pageSize, setPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [labelsPageSize, setLabelsPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [batchUsers, changeBatchUsers] = useState('');
  // 获取数据
  const fetchUserList = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'users/getCanaryUsers',
      payload: {
        params,
        type,
      }
    });
  }, [dispatch]);
  const fetchUserSettings = useCallback((params: PaginationParameters, uid: string, type?: string) => {
    dispatch({
      type: 'users/getUserSettings',
      payload: {
        params,
        type,
        uid,
      }
    });
  }, [dispatch]);
  const fetchUserLabels = useCallback((params: PaginationParameters, uid: string, type?: string) => {
    dispatch({
      type: 'users/getUserLabels',
      payload: {
        params,
        type,
        uid,
      }
    });
  }, [dispatch]);
  // 操作副作用
  useEffect(() => {
    fetchUserList({
      pageSize,
    });
  }, [fetchUserList, pageSize]);

  const handleModulesSearchWordChange = (value: string) => { };
  const handleModulesSearch = (value: string) => { };
  const handleOnRow = (record: CanaryUser) => {
    return {
      onDoubleClick: () => {
        setCurrentUser(record);
        fetchUserSettings({
          pageSize,
        }, record.uid);
        fetchUserLabels({
          pageSize,
        }, record.uid);
        setUserDetailModalVisible(true);
      }
    }
  };
  const handleRefreshLabels = useCallback(() => {
    dispatch({
      type: 'users/getUserLabelsCache',
      payload: {
        params: {
          uid: currentUser?.uid,
        },
        cb: (user: CanaryUser) => setCurrentUser(user),
      },
    });
  }, [currentUser, dispatch]);
  const handleBatchUsersOk = () => {
    dispatch({
      type: 'users/addCanaryUsers',
      payload: {
        cb: () => {
          message.success('添加用户成功');
          setUserAddModalVisible(false);
          fetchUserList({
            pageSize,
          }, 'del');
        },
        params: {
          users: batchUsers.split(';'),
        },
      }
    });
  };
  const handleBatchUsersChange = (e: React.ChangeEvent) => {
    const value = e.target.value;
    changeBatchUsers(value);
  };
  const handleTokenChange = (type: string, pageToken?: string) => {
    fetchUserList({
      pageSize,
      pageToken,
    }, type);
  };
  const handlePageSizeChange = (size: number) => {
    fetchUserList({
      pageSize,
    }, 'del');
    setPageSize(size);
  };
  // 数据定义
  const contentDetailConfig = useMemo(() => {
    return currentUser ? [{
      title: 'uid',
      content: `${currentUser.uid}`,
    }, {
      title: '活跃时间',
      content: `${formatTableTime(currentUser.activeAt)}`,
    }, {
      title: '创建时间',
      content: `${formatTableTime(currentUser.createdAt)}`,
    }, {
      title: '缓存标签',
      content: (
        <div style={{ display: 'flex' }}>
          <Input.TextArea defaultValue={currentUser.labels} disabled></Input.TextArea>
          <Button icon="reload" type="link" onClick={handleRefreshLabels}>刷新缓存</Button>
        </div>
      ),
    }] : undefined;
  }, [currentUser, handleRefreshLabels]);

  const columns = [{
    title: 'ID',
    dataIndex: 'uid',
    key: 'uid',
  }, {
    title: '缓存标签',
    dataIndex: 'labels',
    key: 'labels',
  }, {
    title: '活跃时间',
    dataIndex: 'activeAt',
    key: 'activeAt',
    render: (time: string) => {
      return formatTableTime(time);
    },
  }, {
    title: '创建时间',
    dataIndex: 'createdAt',
    key: 'createdAt',
    render: (time: string) => {
      return formatTableTime(time);
    },
  }];

  const userTabsConfig = [{
    key: 'label',
    title: '灰度标签',
    content: (
      <GrayscaleTag
        dataSource={labelsList}
        hideColumns={['users', 'release', 'status', 'updatedAt']}
        onAction={
          (record: Label) => ({
            onDelete: () => {
              dispatch({
                type: 'users/deleteUserLabel',
                payload: {
                  uid: currentUser?.uid,
                  hid: record.hid,
                  cb: () => {
                    fetchUserLabels({
                      pageSize: labelsPageSize,
                    }, (currentUser?.uid) as string, 'del');
                  }
                }
              })
            },
          })
        }
        paginationProps={{
          total: labelsPageTotal,
          nextPageToken: labelsNextPageToken,
          prePageToken: labelsPrePageToken,
          pageSizeOptions: [10, 20, 30, 40, 50],
          onTokenChange: (type: string, token?: string) => {
            fetchUserLabels({
              pageSize: labelsPageSize,
              pageToken: token,
            }, (currentUser?.uid) as string, type)
          },
          onPageSizeChange: (size: number) => {
            fetchUserLabels({
              pageSize: size,
            }, (currentUser?.uid) as string, 'del');
            setLabelsPageSize(size);
          },
        }}
      />
    ),
    action: (
      <Input
        placeholder="请输入搜索关键字"
      />
    ),
  }, {
    key: 'setting',
    title: '配置项',
    content: (<Setting></Setting>),
    action: (
      <Input
        placeholder="请输入搜索关键字"
      />
    ),
  }];

  return (
    <div>
      <TableTitle
        plusTitle="添加用户"
        handlePlusClick={() => setUserAddModalVisible(true)}
        handleSearch={handleModulesSearch}
        handleWordChange={handleModulesSearchWordChange}
      />
      <Table
        rowKey="uid"
        onRow={handleOnRow}
        pagination={false}
        columns={columns}
        dataSource={canaryUserList}
      />
      <Pagination
        pageSize={pageSize}
        pageSizeOptions={[10, 20, 30, 40]}
        onTokenChange={handleTokenChange}
        prePageToken={prePageToken}
        nextPageToken={nextPageToken}
        onPageSizeChange={handlePageSizeChange}
        total={pageTotal}
      />
      {/* 弹窗 */}
      {
        userAddModalVisible && (
          <Modal
            title="添加用户"
            visible={userAddModalVisible}
            onCancel={() => setUserAddModalVisible(false)}
            onOk={handleBatchUsersOk}
            destroyOnClose
          >
            <Form {...DEFAULT_FORM_ITEM_LAYOUT}>
              <Form.Item label="批量用户" help="添加多个用户使用英文;间隔">
                <Input onChange={handleBatchUsersChange} value={batchUsers} placeholder="请输入用户uid"></Input>
              </Form.Item>
            </Form>
          </Modal>
        )
      }
      <Modal
        footer={null}
        title="用户"
        visible={userDetailModalVisible}
        onCancel={() => setUserDetailModalVisible(false)}
        width={DEFAULT_MODAL_WIDTH}
        destroyOnClose
      >
        <ContentDetail content={contentDetailConfig}></ContentDetail>
        <ContentTabs
          tabs={userTabsConfig}
        ></ContentTabs>
      </Modal>
    </div>
  );
};

export default connect((state) => {
  const {
    canaryUserList,
    prePageToken,
    nextPageToken,
    pageTotal,
    labelsList,
    labelsPageTotal,
    labelsNextPageToken,
    labelsPrePageToken,
  } = (state as any).users;
  return {
    canaryUserList,
    prePageToken,
    nextPageToken,
    pageTotal,
    labelsList,
    labelsPageTotal,
    labelsNextPageToken,
    labelsPrePageToken,
  };
})(Users);
