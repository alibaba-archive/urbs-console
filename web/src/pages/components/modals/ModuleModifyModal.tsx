import React from 'react';
import { Modal, Form, Input, Button } from 'antd';
import { FormComponentProps } from 'antd/es/form';
import { DEFAULT_FORM_ITEM_LAYOUT, Module, FieldsValue } from '../../declare';
import { AcUserSelect } from '../';
import styleNames from '../style/base.less';

interface Props extends FormComponentProps {
  visible: boolean;
  isEdit: boolean,
  onCancel?: (e: React.MouseEvent<HTMLElement>) => void;
  onOk?: (values: FieldsValue) => void;
  onOffline?: () => void;
  moduleInfo?: Module;
}

const ModuleModifyModal: React.FC<Props> = (props) => {
  const { isEdit, onOffline, moduleInfo, visible, onCancel, onOk, form } = props;
  const { getFieldDecorator, getFieldsValue } = form;
  const modalTitle = `${isEdit ? '编辑' : '添加'}功能模块`;
  const handleOnOk = () => {
    const values = getFieldsValue();
    if (onOk) {
      onOk(values);
    }
  };
  return (
    <Modal
      visible={visible}
      title={modalTitle}
      onCancel={onCancel}
      footer={
        <div
          style={{
            display: 'flex',
            justifyContent: isEdit ? 'space-between' : 'flex-end',
          }}
        >
          {
            isEdit && (<div>
              <Button onClick={onOffline} className={styleNames['offline-btn--color']}>下线</Button>
            </div>)
          }
          <div>
            <Button onClick={onCancel}>取消</Button>
            <Button type="primary" onClick={handleOnOk}>确定</Button>
          </div>
        </div>
      }
      destroyOnClose
    >
      <Form {...DEFAULT_FORM_ITEM_LAYOUT}>
        <Form.Item
          label="名称"
          style={{
            margin: '0',
          }}
        >
          {
            getFieldDecorator('name', {
              initialValue: moduleInfo ? moduleInfo.name : '',
              rules: [{
                required: true,
                message: '请输入名称',
              }],
            })(
              <Input disabled={isEdit} placeholder="请输入名称"></Input>
            )
          }
        </Form.Item>
        <Form.Item
          label="负责人"
          style={{
            margin: '0',
          }}
        >
          {
            getFieldDecorator('uids')(
              <AcUserSelect defaultSelectedUser={moduleInfo ? moduleInfo.users : []} />
            )
          }
        </Form.Item>
        <Form.Item
          label="描述"
          style={{
            margin: '0',
          }}
        >
          {
            getFieldDecorator('desc', {
              initialValue: moduleInfo ? moduleInfo.desc : ''
            })(
              <Input.TextArea placeholder="请输入描述"></Input.TextArea>
            )
          }
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default Form.create<Props>()(ModuleModifyModal);
