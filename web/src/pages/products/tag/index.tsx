
import React, { useState, useEffect, useMemo, useCallback } from 'react';
import { connect } from 'dva';
import { TableTitle, GrayscaleTag, GrayscaleTagModifyModal, TagDetailModal } from '../../components';
import {
  TagComponentProps,
  DEFAULT_PAGE_SIZE,
  PaginationParameters,
  Label,
  FieldsValue,
} from '../../declare';

const Tag: React.FC<TagComponentProps> = (props) => {
  const {
    pageTotal,
    prePageToken,
    nextPageToken,
  } = props;
  const [curentLabel, setCurentLabel] = useState<Label>();
  const [pageSize, setPageSize] = useState(DEFAULT_PAGE_SIZE);
  const { match, dispatch, productTagsList, history } = props;
  const { params } = match;
  const productName = params.name;
  const [productTagModalVisible, setProductTagModalVisible] = useState(false);
  const [grayscaleTagModalVisible, setGrayscaleTagModalVisible] = useState(false);
  const [tagsSearchWord, setTagsSearchWord] = useState('');
  const fetchProductTags = useCallback((params: PaginationParameters, type?: string) => {
    dispatch({
      type: 'products/getProductTags',
      payload: {
        params,
        type,
        productName,
      }
    });
  }, [dispatch, productName]);
  const handleTagsSearch = (value: string) => {
    fetchProductTags({
      pageSize,
      q: value,
    }, 'del');
    setTagsSearchWord(value);
  };
  const handleTagsSearchWordChange = (value: string) => {
    setTagsSearchWord(value);
  };
  const changeProductTagModalVisible = (visible: boolean) => {
    setProductTagModalVisible(visible);
  }
  const handleOnRow = (record: Label) => {
    return {
      onDoubleClick: () => {
        setCurentLabel(record);
        changeProductTagModalVisible(true);
      }
    }
  };
  const handlePlusClick = () => {
    setCurentLabel(undefined);
    setGrayscaleTagModalVisible(true);
  };
  const handleTagModifyOk = (values: FieldsValue) => {
    if (curentLabel) {
      dispatch({
        type: 'products/updateProductTags',
        payload: {
          params: values,
          productName,
          cb: (record: Label) => {
            setCurentLabel(record);
            fetchProductTags({
              pageSize,
              q: tagsSearchWord,
            }, 'del');
            setGrayscaleTagModalVisible(false);
          },
        },
      });
    } else {
      dispatch({
        type: 'products/addProductTags',
        payload: {
          params: values,
          productName,
          cb: () => {
            fetchProductTags({
              pageSize,
              q: tagsSearchWord,
            }, 'del');
            setGrayscaleTagModalVisible(false);
          },
        },
      });
    }
  };
  const handleTagOffline = () => {
    dispatch({
      type: 'products/offlineProductTags',
      payload: {
        productName,
        label: curentLabel?.name,
        cb: () => {
          fetchProductTags({
            pageSize,
            q: tagsSearchWord,
          }, 'del');
          setGrayscaleTagModalVisible(false);
        },
      },
    });
  };
  const handleTagDelete = () => {
    dispatch({
      type: 'products/deleteProductTags',
      payload: {
        productName,
        label: curentLabel?.name,
        cb: () => {
          fetchProductTags({
            pageSize,
            q: tagsSearchWord,
          }, 'del');
          setGrayscaleTagModalVisible(false);
        },
      },
    });
  };
  useEffect(() => {
    fetchProductTags({
      pageSize,
      q: tagsSearchWord,
    });
  }, [productName, pageSize, fetchProductTags]);
  const gotoGroups = () => {
    history.push('/group')
  };
  const gotoUsers = () => {
    history.push('/user')
  };
  return (
    <div>
      <TableTitle
        plusTitle="添加灰度标签"
        handlePlusClick={ handlePlusClick }
        handleWordChange={ handleTagsSearchWordChange }
        handleSearch={ handleTagsSearch }
      />
      <GrayscaleTag
        hideColumns={ ['product', 'action'] }
        onRow={ handleOnRow }
        dataSource={ productTagsList }
        paginationProps={
          {
            pageSize,
            total: pageTotal,
            nextPageToken,
            prePageToken,
            pageSizeOptions: [10, 20, 50, 100],
            onTokenChange: (type: string, token?: string) => {
              fetchProductTags({
                pageSize,
                pageToken: token,
                q: tagsSearchWord,
              }, type);
            },
            onPageSizeChange: (size: number) => {
              setPageSize(size);
              fetchProductTags({
                pageSize: size,
                q: tagsSearchWord,
              }, 'del')
            }
          }
        }
      />
      {/* 弹窗 */}
      {
        productTagModalVisible && <TagDetailModal
          product={ productName }
          labelInfo={ curentLabel }
          visible={ productTagModalVisible }
          onCancel={ () => changeProductTagModalVisible(false) }
          onSettingEdit={ () => setGrayscaleTagModalVisible(true) }
          onGotoGroups={ gotoGroups }
          onGotoUsers={ gotoUsers }
        ></TagDetailModal>
      }
      {
        grayscaleTagModalVisible && (
          <GrayscaleTagModifyModal
            visible={ grayscaleTagModalVisible }
            isEdit={ !!curentLabel }
            labelInfo={ curentLabel }
            onOk={ handleTagModifyOk }
            onOffline={ handleTagOffline }
            onDelete={ handleTagDelete }
            onCancel={ () => setGrayscaleTagModalVisible(false) }
          />
        )
      }
    </div>
  );
}

export default connect((states) => {
  return { ...(states as any).products }
})(Tag);
