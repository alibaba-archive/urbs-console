import React, { useState } from 'react';
import { Tooltip, Icon, message } from 'antd';
import { connect } from 'dva';
import { ProductModifyModal } from '../../pages/components';

const Products = ({dispatch}) => {
  const [productAddVisible, changeProductAddVisible] = useState(false);
  const handleModifyOk = (name: string, desc: string, uids: string[]) => {
    dispatch({
      type: 'products/createProduct',
      payload: {
        cb: () => {
          message.success('添加企业成');
          changeProductAddVisible(false);
          dispatch({
            type: 'products/getProducts',
            payload: {},
          });
        },
        params: {
          name,
          desc,
          uids,
        }
      }
    })
  };
  return (
    <>
      <span>
        Urbs 灰度平台
      </span>
      <Tooltip title="添加产品" placement="right">
        <Icon type="plus"onClick={ () => changeProductAddVisible(true) } ></Icon>
      </Tooltip>
      {
        productAddVisible && <ProductModifyModal
          visible={ productAddVisible }
          onCancel={ () => changeProductAddVisible(false) }
          onOk={ handleModifyOk }
        />
      }
    </>
  )
};

export default connect()(Products)