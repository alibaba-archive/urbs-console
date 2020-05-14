import React from 'react';
import { Modal, Form, Input, Select, Button, message } from 'antd';
import { FormComponentProps } from 'antd/es/form';
import { Group, DEFAULT_FORM_ITEM_LAYOUT, FieldsValue } from '../../declare';

interface Props extends FormComponentProps {
  visible: boolean;
  onCancel?: (e: React.MouseEvent<HTMLElement>) => void;
  onOk?: (fieldsValue: FieldsValue, isEdit?: boolean) => void;
  onDel?: (e: React.MouseEvent<HTMLElement>) => void;
  groupInfo?: Group;
}

const ModuleModifyModal: React.FC<Props> = (props) => {
  const { groupInfo, visible, onCancel, onOk, onDel, form } = props;
  const { getFieldDecorator } = form;
  const isEdit = !!groupInfo;
  const modalTitle = `${isEdit ? '编辑' : '添加'}群组`;
  const handleOnOk = () => {
    const values = form.getFieldsValue();
    const {uid, kind} = values;
    if (!uid) {
      message.info('请输入名称');
      return
    }
    if (!kind) {
      message.info('请选择类型');
      return
    }
    if (onOk) {
      onOk(values, isEdit);
    }
  };
  return (
    <Modal
      visible={ visible }
      title={ modalTitle }
      onCancel={ onCancel }
      footer={
        <div
          style={{
            display: 'flex',
            justifyContent: isEdit ? 'space-between' : 'flex-end',
          }}
        >
          {
            isEdit && (<div>
              <Button type="danger" onClick={ onDel }>删除</Button>
            </div>)
          }
          <div>
            <Button onClick={ onCancel }>取消</Button>
            <Button type="primary" onClick={ handleOnOk }>确定</Button>
          </div>
        </div>
      }
      destroyOnClose
    >
      <Form { ...DEFAULT_FORM_ITEM_LAYOUT }>
        <Form.Item label="名称">
          {
            getFieldDecorator('uid', {
              initialValue: groupInfo ? groupInfo.uid : undefined,
            })(
              <Input disabled={ isEdit } placeholder="请输入uid"></Input>
            )
          }
        </Form.Item>
        <Form.Item label="类型">
          {
            getFieldDecorator('kind', {
              initialValue: groupInfo ? groupInfo.kind : undefined
            })(
              <Select placeholder="请选择类型" disabled={ isEdit }>
                <Select.Option value="organization">organization</Select.Option>
              </Select>
            )
          }
        </Form.Item>
        <Form.Item label="描述">
          {
            getFieldDecorator('desc', {
              initialValue: groupInfo ? groupInfo.desc : undefined
            })(
              <Input.TextArea placeholder="请输入描述"></Input.TextArea>
            )
          }
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default Form.create<Props>({
  onFieldsChange (props, field, allFields) {
    console.log(props, field, allFields);
  },
  onValuesChange (props, changeValues, allValues) {
    console.log(props, changeValues, allValues);
  },
})(ModuleModifyModal);
