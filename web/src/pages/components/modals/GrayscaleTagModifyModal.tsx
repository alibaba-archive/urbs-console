import React from 'react';
import { Modal, Form, Button, Checkbox, Input } from 'antd';
import { FormComponentProps } from 'antd/es/form';
import styleNames from '../style/base.less';
import { DEFAULT_FORM_ITEM_LAYOUT, Label, FieldsValue, VERSION_CHANNEL, CLIENT_TYPE } from '../../declare';
import { AcUserSelect } from '../'

interface Props extends FormComponentProps {
  visible: boolean;
  isEdit: boolean,
  onCancel?: () => void;
  onOk?: (values: FieldsValue) => void;
  onOffline?: () => void;
  onDelete?: () => void;
  labelInfo?: Label;
}

const GrayscaleTagModifyModal: React.FC<Props> = (props) => {
  const {
    isEdit,
    visible,
    onCancel,
    onOk,
    labelInfo,
    form,
    onDelete,
    onOffline,
  } = props;
  const { getFieldDecorator, getFieldsValue } = form;
  const handleOnOk = () => {
    const values = getFieldsValue();
    if (onOk) {
      onOk(values);
    }
  };
  const handleOffline = () => {
    const { onOffline } = props;
    if (onOffline) {
      Modal.confirm({
        title: '操作不可逆，请再次确认',
        content: '确认下线该标签？',
        onOk: onOffline,
      });
    }
  };
  const handleDelete = () => {
    const { onDelete } = props;
    if (onDelete) {
      Modal.confirm({
        title: '操作不可逆，请再次确认',
        content: '确认删除该标签？',
        onOk: onDelete,
      });
    }
  };
  return (
    <Modal
      title={`${isEdit ? '编辑' : '添加'}环境标签`}
      visible={visible}
      onCancel={onCancel}
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
              <Button onClick={handleOffline} className={styleNames['offline-btn--color']}>下线</Button>
              <Button onClick={handleDelete} type="danger">删除</Button>
            </div>)
          }
          <div>
            <Button onClick={onCancel}>取消</Button>
            <Button type="primary" onClick={handleOnOk}>确定</Button>
          </div>
        </div>
      }
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
              initialValue: labelInfo && labelInfo.name
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
              <AcUserSelect defaultSelectedUser={labelInfo && (labelInfo.users || [])} />
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
              initialValue: labelInfo && labelInfo.desc
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
              initialValue: labelInfo && labelInfo.channels
            })(
              <Checkbox.Group options={VERSION_CHANNEL}></Checkbox.Group>
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
              initialValue: labelInfo && labelInfo.clients
            })(
              <Checkbox.Group options={CLIENT_TYPE}></Checkbox.Group>
            )
          }
        </Form.Item>
      </Form>
    </Modal>
  )
};

export default Form.create<Props>()(GrayscaleTagModifyModal);