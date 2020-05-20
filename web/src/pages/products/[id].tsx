import React, { useEffect, useState, useMemo } from 'react';
import { Icon, Tooltip, message } from 'antd';
import { connect } from 'dva';
import { ProductModifyModal } from '../components';
import { ProductsComponentProps } from '../declare'
import styles from './[id].less';

const DEFAULT_TITLE = '编辑产品';

enum OperationType {
  update = 'update',
  delete = 'delete',
  offline = 'offline',
}

const operationTitle = {
  [OperationType.update]: '更新',
  [OperationType.delete]: '删除',
  [OperationType.offline]: '下线',
}

const Products: React.FC<ProductsComponentProps> = (props) => {
  const { match, dispatch, productStatistics, productList, history } = props;
  const { params } = match;
  const { name: productName } = params;
  const [productEditVisible, setProductEditVisible] = useState(false);
  const [productCanEdit, setProductCanEdit] = useState(false);
  const currentProductStatistics = useMemo(() => {
    const {
      labels = 0,
      modules = 0,
      settings = 0,
      release = 0,
    } = productStatistics;
    return [{
      key: 'label',
      label: '灰度标签数',
      count: labels,
    }, {
      key: 'module',
      label: '功能模块数',
      count: modules,
    }, {
      key: 'setting',
      label: '配置项数',
      count: settings,
    }, {
      key: 'publish',
      label: '发布次数',
      count: release,
    }]
  }, [productStatistics]);
  const currentProductDetail = useMemo(() => {
    return productList.find(product => product.name === productName);
  }, [productName, productList]);
  useEffect(() => {
    dispatch({
      type: 'products/getProductStatistics',
      payload: {
        params: {
          productName: productName,
        }
      }
    });
  }, [dispatch, productName]);
  useEffect(() => {
    dispatch({
      type: 'products/getPermission',
      payload: {
        cb: (canEdit: boolean) => {
          setProductCanEdit(!!canEdit);
        },
        params: {
          product: productName,
        }
      },
    });
  }, [dispatch, productName]);
  const operatingProduct = (name: string, desc: string, uids: string[], type: OperationType) => {
    dispatch({
      type: `products/${type}Product`,
      payload: {
        cb: () => {
          message.success(`产品${ operationTitle[type] }成功`);
          setProductEditVisible(false);
          dispatch({
            type: 'products/getProducts',
            payload: {},
          });
          if (OperationType.offline === type) {
            const products = productList.find(product => product.name !== productName);
            history.push(`/products/${ products ? `${products.name}` : '/help' }`)
          }
          if (OperationType.update === type) {
            dispatch({
              type: 'products/getPermission',
              payload: {
                cb: (canEdit: boolean) => {
                  setProductCanEdit(!!canEdit);
                },
                params: {
                  product: productName,
                }
              },
            });
          }
        },
        params:{
          name,
          desc,
          uids,
        },
      },
    });
  };
  return (
    <div className={ styles.normal }>
      <ul className={ styles['product-detail-wrap'] }>
        <li>
          <div>
            {currentProductDetail?.name}
          </div>
          {
            productCanEdit && <div className={ styles['product-detail-desc'] }>
              <Tooltip placement="right" title={ DEFAULT_TITLE }>
                <Icon type="setting" onClick={ () => setProductEditVisible(true) }></Icon>
              </Tooltip>
            </div>
          }
        </li>
        <li>
          <div>
            负责人 :
          </div>
          <div className={ styles['product-detail-desc'] }>
            {
              currentProductDetail?.users ? currentProductDetail.users.map(user => user.name).join('，') : ''
            }
          </div>
        </li>
        <li>
          <div>
            描述 :
          </div>
          <div className={ styles['product-detail-desc'] }>
            {
              currentProductDetail?.desc
            }
          </div>
        </li>
      </ul>
      <ul className={ styles['product-counts-wrap'] }>
        {
          currentProductStatistics.map(item => (
            <li key={ item.key }>
              <div className={ styles.counts }>
                { item.count }
              </div>
              <div className={ styles['counts-desc'] }>
                { item.label }
              </div>
            </li>
          ))
        }
      </ul>
      {/* 弹窗 */}
      {
        productEditVisible && (<ProductModifyModal
          visible={ productEditVisible }
          productInfo={ currentProductDetail }
          onCancel={ () => setProductEditVisible(false) }
          onOk={ (name, desc, uids) => operatingProduct(name, desc, uids, OperationType.update) }
          onOffline={ (name, desc, uids) => operatingProduct(name, desc, uids, OperationType.offline) }
          onDelete={ (name, desc, uids) => operatingProduct(name, desc, uids, OperationType.delete) }
        />)
      }
    </div>
  );
};

export default connect((state) => {
  const { productStatistics, productList } = (state as any).products;
  return { productStatistics, productList };
})(Products);
