
import React, { useEffect, useState, useMemo, useCallback } from 'react';
import { Table } from 'antd';
import { connect } from 'dva';
import { ModulesComponentProps, DEFAULT_PAGE_SIZE, User, PaginationParameters, Module, FieldsValue } from '../../declare';
import { ModuleModifyModal, ModuleDetailModal, TableTitle, Pagination } from '../../components';
import { formatTableTime } from '../../utils/format';

const Modules: React.FC<ModulesComponentProps> = (props) => {
  const {
    dispatch,
    match,
    productModulesList,
    modulePageTotal,
    moduleNextPageToken,
    modulePrePageToken,
  } = props;
  const { params } = match;
  const productName = params.name;
  const [curentModule, setCurentModule] = useState<Module>();
  const [moduleModalVisible, setModuleModalVisible] = useState(false);
  const [moduleDetailVisible, setModuleDetailVisible] = useState(false);
  const [pageSize, setPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [modulesSearchWord, setModulesSearchWord] = useState('');
  const fetchModuleList = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'products/getProductModules',
      payload: {
        productName,
        params,
        type,
      }
    })
  }, [dispatch, productName]);
  useEffect(() => {
    fetchModuleList({
      pageSize,
      q: modulesSearchWord,
    });
  }, [fetchModuleList, modulesSearchWord, pageSize]);
  const handleModulesSearchWordChange = (value: string) => {
    setModulesSearchWord(value);
  };
  const handleModulesSearch = (value: string) => {
    fetchModuleList({
      pageSize,
      q: value,
    }, 'del');
  };
  const handleOnRow = (record: Module) => {
    return {
      onClick: () => {
        setCurentModule(record);
        setModuleDetailVisible(true);
      }
    }
  };
  const columns = useMemo(() => ([{
    title: '名称',
    dataIndex: 'name',
    key: 'name',
  }, {
    title: '负责人',
    dataIndex: 'users',
    key: 'users',
    render: (users: User[]) => {
      return Array.isArray(users) ? users.map(item => item.name).join() : '';
    }
  }, {
    title: '描述',
    dataIndex: 'desc',
    key: 'desc',
  }, {
    title: '配置项数',
    dataIndex: 'status',
    key: 'status',
  }, {
    title: '更新时间',
    dataIndex: 'updatedAt',
    key: 'updatedAt',
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
  }]), []);
  const handlePlusClick = () => {
    setCurentModule(undefined);
    setModuleModalVisible(true);
  };
  const handleModuleModifyOk = (values: FieldsValue) => {
    if (curentModule) {
      dispatch({
        type: 'products/updateProductModules',
        payload: {
          params: values,
          productName,
          cb: (record: Module) => {
            setCurentModule(record);
            fetchModuleList({
              pageSize,
              q: modulesSearchWord,
            }, 'del');
            setModuleModalVisible(false);
          },
        },
      });
    } else {
      dispatch({
        type: 'products/addProductModules',
        payload: {
          params: values,
          productName,
          cb: () => {
            fetchModuleList({
              pageSize,
              q: modulesSearchWord,
            }, 'del');
            setModuleModalVisible(false);
          },
        },
      });
    }
  };
  const handleModuleOffline = () => {
    dispatch({
      type: 'products/offlineProductModules',
      payload: {
        productName,
        module: curentModule?.name,
        cb: () => {
          fetchModuleList({
            pageSize,
            q: modulesSearchWord,
          }, 'del');
          setModuleModalVisible(false);
          setModuleDetailVisible(false);
        },
      },
    });
  };
  return (
    <div>
      <TableTitle
        plusTitle="添加功能模块"
        handlePlusClick={handlePlusClick}
        handleWordChange={handleModulesSearchWordChange}
        handleSearch={handleModulesSearch}
      />
      <Table
        rowKey="name"
        onRow={handleOnRow}
        pagination={false}
        columns={columns}
        dataSource={productModulesList}
      ></Table>
      <Pagination
        pageSize={pageSize}
        total={modulePageTotal}
        pageSizeOptions={[10, 20, 50, 100]}
        nextPageToken={moduleNextPageToken}
        prePageToken={modulePrePageToken}
        onTokenChange={
          (type: string, token?: string) => {
            fetchModuleList({
              pageSize,
              pageToken: token,
              q: modulesSearchWord,
            }, type);
          }
        }
        onPageSizeChange={
          (size: number) => {
            setPageSize(size);
            fetchModuleList({
              pageSize: size,
              q: modulesSearchWord,
            }, 'del');
          }
        }
      />
      {/* 弹窗 */}
      {
        moduleModalVisible && <ModuleModifyModal
          visible={moduleModalVisible}
          isEdit={!!curentModule}
          moduleInfo={curentModule}
          onOk={handleModuleModifyOk}
          onCancel={() => setModuleModalVisible(false)}
          onOffline={handleModuleOffline}
        />
      }
      {
        moduleDetailVisible && <ModuleDetailModal
          moduleInfo={curentModule}
          product={productName}
          visible={moduleDetailVisible}
          onCancel={() => setModuleDetailVisible(false)}
          onModuleEdit={() => setModuleModalVisible(true)}
        ></ModuleDetailModal>
      }
    </div>
  );
};

export default connect((state) => {
  return { ...(state as any).products };
})(Modules);
