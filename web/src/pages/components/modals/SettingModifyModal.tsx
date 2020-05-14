import React from 'react';
import { Modal, Form, Button, Checkbox, Input } from 'antd';
import { FormComponentProps } from 'antd/es/form';
import styleNames from '../style/base.less';
import { DEFAULT_FORM_ITEM_LAYOUT, Setting, FieldsValue, CLIENT_TYPE, VERSION_CHANNEL } from '../../declare';
import { AcUserSelect } from '../'

interface Props extends FormComponentProps {
  visible: boolean;
  isEdit: boolean,
  onCancel?: () => void;
  onOk?: (values: FieldsValue) => void;
  onOffline?: () => void;
  defaultValue?: Setting;
}

const SettingModifyModal: React.FC<Props> = (props) => {
  const { isEdit, visible, onCancel, onOk, defaultValue, form, onOffline } = props;
  const { getFieldDecorator, getFieldsValue } = form;
  const handleOnOk = () => {
    const fieldsValues = getFieldsValue();
    if (onOk) {
      fieldsValues.values = fieldsValues.values.split(',');
      onOk(fieldsValues);
    }
  };
  return (
    <Modal
      title={`${ isEdit ? '编辑' : '添加'}配置项`}
      visible={ visible }
      onCancel={ onCancel }
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
              <Button onClick={ onOffline } className={styleNames['offline-btn--color']}>下线</Button>
            </div>)
          }
          <div>
            <Button onClick={ onCancel }>取消</Button>
            <Button type="primary" onClick={ handleOnOk }>确定</Button>
          </div>
        </div>
      }
    >
      <Form { ...DEFAULT_FORM_ITEM_LAYOUT }>
        <Form.Item
          label="名称"
          style={{
            margin: '0',
          }}
        >
          {
            getFieldDecorator('name', {
              initialValue: defaultValue && defaultValue.name
            })(
              <Input placeholder="请输入名称" disabled={ isEdit }></Input>
            )
          }
        </Form.Item>
        <Form.Item
          label="所属模块"
          style={{
            margin: '0',
          }}
        >
          {
            getFieldDecorator('module', {
              initialValue: defaultValue && defaultValue.module
            })(
              <Input placeholder="请输入所属模块" disabled={ isEdit }></Input>
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
              <AcUserSelect defaultSelectedUser={ defaultValue && defaultValue.users } ></AcUserSelect>
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
              initialValue: defaultValue && defaultValue.desc
            })(
              <Input.TextArea placeholder="请输入描述"></Input.TextArea>
            )
          }
        </Form.Item>
        <Form.Item
          label="版本通道"
          style={{
            margin: '0',
          }}
        >
          {
            getFieldDecorator('channels', {
              initialValue: defaultValue ? [defaultValue.channels] : []
            })(
              <Checkbox.Group options={ VERSION_CHANNEL }></Checkbox.Group>
            )
          }
        </Form.Item>
        <Form.Item
          label="端类型"
          style={{
            margin: '0',
          }}
        >
          {
            getFieldDecorator('clients', {
              initialValue: defaultValue ? [defaultValue.clients] : []
            })(
              <Checkbox.Group options={ CLIENT_TYPE }></Checkbox.Group>
            )
          }
        </Form.Item>
        <Form.Item
          label="可选值"
          style={{
            margin: '0',
          }}
        >
          {
            getFieldDecorator('values', {
              initialValue: defaultValue ? defaultValue.values?.join(',') : []
            })(
              <Input placeholder="多个值以英文 , 分隔"></Input>
            )
          }
        </Form.Item>
      </Form>
    </Modal>
  )
};

export default Form.create<Props>()(SettingModifyModal);