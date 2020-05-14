
import React, { useEffect, useState, useMemo, useCallback } from 'react';
import { Button, Input, Modal, Icon } from 'antd';
import { connect } from 'dva';
import { ContentDetail, ContentTabs, SettingModifyModal, Setting } from '..';
import { ModuleDetailModalComponentProps, FieldsValue, PaginationParameters, DEFAULT_PAGE_SIZE, DEFAULT_MODAL_WIDTH } from '../../declare';
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
    });
  }, [fetchModuleSettingList, pageSize]);
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
        },
      },
    });
  };
  const renderModalTitle = () => {
    return (
      <div className={styleNames['modal-title']}>
        <div>{title}</div>
        <div>
          <Icon type="setting" onClick={onModuleEdit}></Icon>
        </div>
      </div>
    )
  };
  const tagTabsConfig = [{
    key: 'setting',
    title: '配置项',
    content: (
      <Setting
        dataSource={moduleSettingsList}
        paginationProps={
          {
            pageSize,
            pageSizeOptions: [10, 20, 30, 40, 50],
            total: moduleSettingPageTotal,
            prePageToken: moduleSettingPrePageToken,
            nextPageToken: moduleSettingNextPageToken,
            onTokenChange: (type, token) => {
              fetchModuleSettingList({
                pageSize,
                pageToken: token,
              }, type);
            },
            onPageSizeChange: (size) => {
              setPageSize(size);
              fetchModuleSettingList({
                pageSize: size,
              }, 'del');
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
          onClick={handleOpenModuleModifyModal}
        >
          添加配置项
        </Button>
        <Input.Search value={tabsSearchWord} placeholder="请输入搜索关键字" onChange={handleTabsSearchWordChange} onSearch={handleTabsSearch} />
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
        />
      }
    </Modal>
  );
};

export default connect((state) => {
  return { ...(state as any).products };
})(ModuleDetailModal);
