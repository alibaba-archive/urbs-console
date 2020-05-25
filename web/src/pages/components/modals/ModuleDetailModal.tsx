
import React, { useEffect, useState, useMemo, useCallback } from 'react';
import { Button, Input, Modal, Icon } from 'antd';
import { connect } from 'dva';
import { ContentDetail, ContentTabs, SettingModifyModal, SettingDetailModal, Setting } from '..';
import { ModuleDetailModalComponentProps, FieldsValue, PaginationParameters, DEFAULT_PAGE_SIZE, Setting as SettingDeclare, DEFAULT_MODAL_WIDTH } from '../../declare';
import styleNames from '../style/base.less';
import { formatTableTime } from '../../utils/format';

const DEFAULT_TITLE = '功能模块';

const ModuleDetailModal: React.FC<ModuleDetailModalComponentProps> = (props) => {
  const {
    dispatch,
    product,
    moduleInfo,
    onCancel,
    onModuleEdit,
    visible,
    title = DEFAULT_TITLE,
    moduleSettingsList,
    moduleSettingNextPageToken,
    moduleSettingPrePageToken,
    moduleSettingPageTotal,
  } = props;
  const [tabsSearchWord, setTabsSearchWord] = useState('');
  const [pageSize, setPageSize] = useState(DEFAULT_PAGE_SIZE);
  const [settingModifyModalVisible, setSettingModifyModalVisible] = useState(false);
  const [moduleCanEdit, setModuleCanEdit] = useState(false);

  const [settingDetailVisible, changeSettingDetailVisible] = useState(false);
  const [curentSetting, setCurentSetting] = useState<SettingDeclare>();
  const [settingModifyVisible, changeSettingModifyVisible] = useState(false);

  const fetchModuleSettingList = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'products/getModuleSettings',
      payload: {
        type,
        module: moduleInfo?.name,
        productName: product,
        params,
      },
    });
  }, [dispatch, moduleInfo, product]);
  useEffect(() => {
    fetchModuleSettingList({
      pageSize,
      q: tabsSearchWord,
    });
  }, [fetchModuleSettingList, pageSize, tabsSearchWord]);
  useEffect(() => {
    dispatch({
      type: 'products/getPermission',
      payload: {
        cb: (canEdit: boolean) => {
          setModuleCanEdit(!!canEdit);
        },
        params: {
          product,
          module: moduleInfo?.name,
        }
      },
    })
  }, [dispatch, moduleInfo, product]);
  const handleTabsSearch = (searchWord: string) => {
    setTabsSearchWord(searchWord);
    fetchModuleSettingList({
      pageSize,
      q: searchWord,
    }, 'del');
  };
  const handleTabsSearchWordChange = (e: React.ChangeEvent) => {
    const target = e.target;
    setTabsSearchWord((target as any).value);
  };
  const handleOpenModuleModifyModal = () => {
    setSettingModifyModalVisible(true);
  };
  const handleSettingModifyModalOk = (values: FieldsValue) => {
    dispatch({
      type: 'products/addProductSettings',
      payload: {
        params: values,
        productName: product,
        cb: () => {
          setSettingModifyModalVisible(false);
          fetchModuleSettingList({
            pageSize,
            q: tabsSearchWord,
          }, 'del');
        },
      },
    });
  };
  const renderModalTitle = () => {
    return (
      <div className={styleNames['modal-title']}>
        <div>{title}</div>
        {
          moduleCanEdit && (
            <div className={styleNames['modal-icon']}>
              <Icon type="setting" onClick={onModuleEdit}></Icon>
            </div>
          )
        }
      </div>
    )
  };

  const handleOnRow = (record: SettingDeclare) => {
    return {
      onClick: () => {
        setCurentSetting(record);
        changeSettingDetailVisible(true);
      }
    };
  };

  const tagTabsConfig = [{
    key: 'setting',
    title: '配置项',
    content: (
      <Setting
        dataSource={moduleSettingsList}
        onRow={handleOnRow}
        hideColumns={['module', 'desc', 'release', 'action']}
        paginationProps={
          {
            pageSize,
            pageSizeOptions: [10, 20, 50, 100],
            total: moduleSettingPageTotal,
            prePageToken: moduleSettingPrePageToken,
            nextPageToken: moduleSettingNextPageToken,
            onTokenChange: (type, token) => {
              fetchModuleSettingList({
                pageSize,
                pageToken: token,
                q: tabsSearchWord,
              }, type);
            },
            onPageSizeChange: (size) => {
              setPageSize(size);
              fetchModuleSettingList({
                pageSize: size,
                q: tabsSearchWord,
              }, 'del');
            },
          }
        }
      />
    ),
    action: moduleCanEdit && (
      <div>
        <Button
          type="link"
          icon="plus"
          block
          onClick={handleOpenModuleModifyModal}
        >
          添加配置项
        </Button>
        <Input.Search allowClear value={tabsSearchWord} placeholder="请输入搜索关键字" onChange={handleTabsSearchWordChange} onSearch={handleTabsSearch} />
      </div>
    ),
  }];
  const moduleContentDetail = useMemo(() => {
    return (
      moduleInfo && (
        [{
          title: '名称',
          content: moduleInfo.name,
        }, {
          title: '所属产品',
          content: product,
        }, {
          title: '负责人',
          content: Array.isArray(moduleInfo.users) ? moduleInfo.users.map(item => item.name).join(',') : '',
        }, {
          title: '描述',
          content: moduleInfo.desc,
        }, {
          title: '配置项',
          content: moduleInfo.status,
        }, {
          title: '创建时间',
          content: formatTableTime(moduleInfo.createdAt),
        }, {
          title: '更新时间',
          content: formatTableTime(moduleInfo.updatedAt),
        }]
      )
    );
  }, [moduleInfo, product]);

  const fetchSettingList = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'products/getProductSettings',
      payload: {
        productName: product,
        params,
        type,
      }
    })
  }, [dispatch, product]);

  const handleSettingModifyOk = (values: FieldsValue) => {
    dispatch({
      type: 'products/updateProductSettings',
      payload: {
        params: values,
        productName: product,
        cb: (record: SettingDeclare) => {
          setCurentSetting(record);
          fetchSettingList({
            pageSize,
          }, 'del');
          changeSettingModifyVisible(false);
        },
      },
    });
  };
  const handleOfflineSetting = () => {
    dispatch({
      type: 'products/offlineProductSettings',
      payload: {
        productName: product,
        module: curentSetting?.module,
        setting: curentSetting?.name,
        cb: () => {
          fetchSettingList({
            pageSize,
          }, 'del');
          changeSettingModifyVisible(false);
          changeSettingDetailVisible(false);
        },
      },
    });
  };
  return (
    <Modal width={DEFAULT_MODAL_WIDTH} title={renderModalTitle()} visible={visible} onCancel={onCancel} footer={null}>
      <ContentDetail content={moduleContentDetail}></ContentDetail>
      <ContentTabs
        tabs={tagTabsConfig}
      ></ContentTabs>
      {/* 弹窗 */}
      {
        settingModifyModalVisible && <SettingModifyModal
          isEdit={false}
          visible={settingModifyModalVisible}
          onCancel={() => setSettingModifyModalVisible(false)}
          onOk={handleSettingModifyModalOk}
          module={moduleInfo?.name}
        />
      }
      {/* 弹窗 */}
      {
        settingModifyVisible && <SettingModifyModal
          visible={settingModifyVisible}
          isEdit={!!curentSetting}
          defaultValue={curentSetting}
          onOffline={handleOfflineSetting}
          onOk={handleSettingModifyOk}
          onCancel={() => changeSettingModifyVisible(false)}
        />
      }
      {
        settingDetailVisible && (
          <SettingDetailModal
            visible={settingDetailVisible}
            settingInfo={curentSetting}
            title="配置项"
            product={product}
            onSettingEdit={() => changeSettingModifyVisible(true)}
            onCancel={() => changeSettingDetailVisible(false)}
          />
        )
      }
    </Modal>
  );
};

export default connect((state) => {
  return { ...(state as any).products };
})(ModuleDetailModal);
