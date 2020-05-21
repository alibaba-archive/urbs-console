
import React, { useEffect, useState, useMemo, useCallback } from 'react';
import { Form, Modal, Table, Input, Button, message } from 'antd';
import { connect } from 'dva';
import { Pagination, TableTitle, AcAddModal } from '../components';
import { User, ACComponentProps, FieldsValue, DEFAULT_MODAL_WIDTH, DEFAULT_FORM_ITEM_LAYOUT, DEFAULT_PAGE_SIZE, PaginationParameters, CanaryUser } from '../declare';

const AC: React.FC<ACComponentProps> = (props) => {
    const {
        dispatch,
        acUserList,
        acPrePageToken,
        acNextPageToken,
        pageTotal,
    } = props;
    const [currentUser, setCurrentUser] = useState<User>();
    const [userAddModalVisible, setUserAddModalVisible] = useState(false);
    const [pageSize, setPageSize] = useState(DEFAULT_PAGE_SIZE);
    const columns = [{
        title: 'ID',
        dataIndex: 'uid',
        key: 'uid',
    }, {
        title: '名称',
        dataIndex: 'name',
        key: 'name',
    }, {
        title: '创建时间',
        dataIndex: 'createdAt',
        key: 'createdAt',
    }
    ];
    useEffect(() => {
        dispatch({
            type: 'users/getAcUsersList',
            payload: {
                params: { pageSize: pageSize }
            }
        });
    }, [dispatch, pageSize]);
    const fetchUserList = useCallback((params: PaginationParameters, type?: string) => {
        dispatch({
            type: 'users/getAcUsersList',
            payload: {
                params,
                type,
            }
        });
    }, [dispatch]);
    const handleAddUserOk = (values: FieldsValue, isEdit?: boolean) => {
        if (isEdit) {
            dispatch({
                type: 'users/updateAcUser',
                payload: {
                    cb: () => {
                        console.log("updateAcUser")
                        message.success('更新用户成功');
                        setUserAddModalVisible(false);
                    },
                    name: values.name,
                    uid: values.uid,
                }
            });
        }
        else {
            dispatch({
                type: 'users/addAcUsers',
                payload: {
                    cb: () => {
                        message.success('添加用户成功');
                        setUserAddModalVisible(false);
                    },
                    users: [values],
                }
            });
        }
    };
    const handleTokenChange = (type: string, pageToken?: string) => {
        fetchUserList({
            pageSize,
            pageToken,
        }, type);
    };
    const handleDelUser = (values: FieldsValue) => {
        dispatch({
            type: 'users/deleteAcUser',
            payload: {
                cb: () => {
                    message.success('添加用户成功');
                    setUserAddModalVisible(false);
                },
                uid: currentUser?.uid,
            }
        });
    };
    const handleOnRow = (record: User) => {
        return {
            onDoubleClick: () => {
                setCurrentUser(record);
                setUserAddModalVisible(true);
            }
        };
    };
    const handlePlusClick = () => {
        setCurrentUser(undefined);
        setUserAddModalVisible(true);
    };
    const handleSearch = (value: string) => {
        if (value != "") {
            dispatch({
                type: 'users/searchAcUsers',
                payload: {
                    params: {
                        key: value,
                    },
                },
            });
        } else {
            dispatch({
                type: 'users/getAcUsersList',
                payload: {
                    params: { pageSize: pageSize }
                }
            });
        }
    };
    const handleSearchWordChange = (value: string) => { };
    const handlePageSizeChange = (size: number) => {
        fetchUserList({
            pageSize: size,
        }, 'del');
        setPageSize(size);
    };
    return (
        <div>
            <TableTitle
                plusTitle="添加用户"
                handlePlusClick={() => handlePlusClick()}
                handleSearch={handleSearch}
                handleWordChange={handleSearchWordChange}
            />
            <Table
                rowKey="uid"
                onRow={handleOnRow}
                pagination={false}
                columns={columns}
                dataSource={acUserList}
            />
            <Pagination
                pageSize={pageSize}
                pageSizeOptions={[10, 20, 50, 100]}
                onTokenChange={handleTokenChange}
                prePageToken={acPrePageToken}
                nextPageToken={acNextPageToken}
                onPageSizeChange={handlePageSizeChange}
                total={pageTotal}
            />
            {
                userAddModalVisible && <AcAddModal
                    visible={userAddModalVisible}
                    onCancel={() => setUserAddModalVisible(false)}
                    User={currentUser}
                    onOk={handleAddUserOk}
                    onDel={handleDelUser}
                />
            }
        </div >
    );
};

export default connect((state) => {
    const {
        acUserList,
        acPrePageToken,
        acNextPageToken,
    } = (state as any).users;
    return {
        acPrePageToken,
        acNextPageToken,
        acUserList
    };
})(AC);