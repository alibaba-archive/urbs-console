import React, { useState } from 'react';
import { Modal, Form, Input, Button, message } from 'antd';
import styleNames from '../style/ProductModifyModal.less';
import { Product, DEFAULT_FORM_ITEM_LAYOUT, DEFAULT_MODAL_WIDTH } from '../../declare';
import { AcUserSelect } from '../../components';

interface Props {
  visible: boolean;
  onCancel?: (e: React.MouseEvent<HTMLElement>) => void;
  onOk?: (name: string, desc: string, uids: string[]) => void;
  onOffline?: (name: string, desc: string, uids: string[]) => void;
  onDelete?: (name: string, desc: string, uids: string[]) => void;
  productInfo?: Product;
}

const ProductModifyModal: React.FC<Props> = (props) => {
  const { productInfo, visible, onCancel, onOk, onOffline, onDelete } = props;
  const isEdit = !!productInfo;
  const modalTitle = `${isEdit ? '编辑' : '添加'}产品`;
  const { name = '', desc = '', users = [] } = productInfo || {};
  const [productName, changeProductName] = useState(name);
  const [productDesc, changeProductDesc] = useState(desc);
  const [productUsers, changeProductUsers] = useState((users || []).map(user => user.uid));
  const handleOk = () => {
    if (onOk) {
      if (!(/^[0-9a-z][0-9a-z.-]{0,61}[0-9a-z]$/.test(productName))) {
        message.error('请输入符合规范的产品名称');
        return
      }
      onOk(productName, productDesc, productUsers);
    }
  };
  const handleOffline = () => {
    if (onOffline) {
      onOffline(productName, productDesc, productUsers);
    }
  };
  const handleDelete = () => {
    if (onDelete) {
      onDelete(productName, productDesc, productUsers);
    }
  };
  const handleNameChange = (e: React.ChangeEvent) => {
    const value = e.target.value;
    changeProductName(value);
  };
  const handleDescChange = (e: React.ChangeEvent) => {
    const value = e.target.value;
    changeProductDesc(value);
  };
  const handleUsersChange = (users: string[]) => {
    changeProductUsers(users);
  };
  return (
    <Modal
      visible={ visible }
      title={ modalTitle }
      onCancel={ onCancel }
      width={ DEFAULT_MODAL_WIDTH }
      destroyOnClose
      footer={
        <div
          style={{
            display: 'flex',
            justifyContent: isEdit ? 'space-between' : 'flex-end',
          }}
        >
          {
            isEdit && (<div>
              <Button onClick={ handleOffline } className={styleNames['offline-btn--color']}>下线</Button>
              {(productInfo && !~productInfo.status) && <Button onClick={ handleDelete } type="danger">删除</Button>}
            </div>)
          }
          <div>
            <Button onClick={ onCancel }>取消</Button>
            <Button type="primary" onClick={ handleOk }>确定</Button>
          </div>
        </div>
      }
    >
      <Form { ...DEFAULT_FORM_ITEM_LAYOUT }>
        <Form.Item
          label="名称"
          className={ styleNames['modify-form-item'] }
        >
          <>
            <Input onChange={ handleNameChange } value={ productName } className={ styleNames['form-name'] } disabled={ isEdit } placeholder="请输入名称"></Input>
            <span className={ styleNames['form-name__tips'] }>注：创建时可修改，编辑时不可修改</span>
          </>
        </Form.Item>
        <Form.Item
          label="负责人"
          className={ styleNames['modify-form-item'] }
        >
          <AcUserSelect defaultSelectedUser={ users } onChange={ handleUsersChange } />
        </Form.Item>
        <Form.Item
          label="描述"
          className={ styleNames['modify-form-item'] }
        >
          <Input.TextArea value={ productDesc } onChange={ handleDescChange } className={ styleNames['form-desc'] } placeholder="请输入描述"></Input.TextArea>
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default ProductModifyModal;
