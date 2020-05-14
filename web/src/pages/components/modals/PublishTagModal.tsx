import React from 'react';
import { Modal, Form, Button, Radio, Row, Col, Input, InputNumber, Select } from 'antd';
import { FormComponentProps } from 'antd/es/form';
import { DEFAULT_FORM_ITEM_LAYOUT, FieldsValue, UserPercentRule } from '../../declare';

interface FormDefaultValue {
  name: string;
  product: string;
  group: string[],
  user: string[],
}

interface Props extends FormComponentProps {
  visible: boolean;
  onCancel?: () => void;
  onOk?: (values: FieldsValue) => void;
  label?: string;
  product?: string;
  module?: string;
  onGotoGroups?: () => void;
  onGotoUsers?: () => void;
  defauleRule?: UserPercentRule;
  title: string | React.ReactNode;
  grayscale?: string[];
}

const PublishTagModal: React.FC<Props> = (props) => {
  const {
    visible,
    onCancel,
    onOk,
    form,
    label,
    product,
    onGotoGroups,
    onGotoUsers,
    defauleRule,
    title,
    module,
    grayscale,
  } = props;
  const { getFieldDecorator, getFieldValue, getFieldsValue } = form;
  const handleOnOk = () => {
    const values = getFieldsValue();
    const kind = getFieldValue('kind');
    if (kind === 'batch') {
      values.groups = values.groups.split(',');
      values.users = values.users.split(',');
    }
    if (kind === 'userPercent') {
      values.rule = {
        value: values.percent ? +values.percent : undefined,
      }
    }
    if (onOk) {
      onOk(values);
    }
  };
  return (
    <Modal
      title={ title }
      visible={ visible }
      onCancel={ onCancel }
      destroyOnClose
      onOk={ handleOnOk }
    >
      <Form { ...DEFAULT_FORM_ITEM_LAYOUT }>
        <Form.Item
          label="名称"
          style={{
            margin: '0',
          }}
        >
          <span>{ label }</span>
        </Form.Item>
        <Form.Item
          label="所属产品"
          style={{
            margin: '0',
          }}
        >
          <span>{ product }</span>
        </Form.Item>
        {
          module ? (<Form.Item
            label="所属模块"
            style={{
              margin: '0',
            }}
          >
            <span>{ product }</span>
          </Form.Item>) : null
        }
        {
          (Array.isArray(grayscale) && grayscale.length) ? <Form.Item
            label="灰度值"
            style={{
              margin: '0',
            }}
          >
            {
              getFieldDecorator('value', {
                initialValue: []
              })(
                <Select>
                  {
                    grayscale.map((item, index) => (
                      <Select.Option key={ `${item}_${ index }` } value={ item }>{ item }</Select.Option>
                    ))
                  }
                </Select>
              )
            }
          </Form.Item> : null
        }
        <Form.Item
          label="发布类型"
          style={{
            margin: '0',
          }}
        >
          {
            getFieldDecorator('kind', {
              initialValue: 'batch'
            })(
              <Radio.Group>
                <Radio value="batch">批量</Radio>
                <Radio value="userPercent">比例</Radio>
              </Radio.Group>
            )
          }
        </Form.Item>
        {
          getFieldValue('kind') === 'batch' ? (
            <>
              <Form.Item
                label="批量群组"
                style={{
                  margin: '0',
                }}
              >
                {
                  getFieldDecorator('groups', {
                    initialValue: ''
                  })(
                    <Row>
                      <Col span={ 18 }>
                        <Input placeholder="输入多个使用英文 , 分隔"></Input>
                      </Col>
                      <Col span={ 6 }>
                        <Button onClick={ onGotoGroups } type="link" block>去创建群组</Button>
                      </Col>
                    </Row>
                  )
                }
              </Form.Item>
              <Form.Item
                label="批量用户"
                style={{
                  margin: '0',
                }}
              >
                {
                  getFieldDecorator('users', {
                    initialValue: ''
                  })(
                    <Row>
                      <Col span={ 18 }>
                        <Input placeholder="输入多个使用英文 , 分隔"></Input>
                      </Col>
                      <Col span={ 6 }>
                        <Button onClick={ onGotoUsers } type="link" block>去添加用户</Button>
                      </Col>
                    </Row>
                  )
              }
              </Form.Item>
            </>
          ) : (
            <Form.Item
              label="比例添加"
              style={{
                margin: '0',
              }}
            >
              {
                getFieldDecorator('percent')(
                  <Row>
                    <Col span={ 8 }>
                      <InputNumber defaultValue={ defauleRule ? defauleRule.rule.value : 0 } style={{width: '100%'}} min={ 0 } max={ 100 } placeholder="请输入1～100"></InputNumber>
                    </Col>
                    <Col span={ 9 } push={ 1 }>
                      %，本次新增比例用户
                    </Col>
                  </Row>
                )
              }
            </Form.Item>
          )
        }
        <Form.Item
          label="发布说明"
          style={{
            margin: '0',
          }}
        >
          {
            getFieldDecorator('desc')(
              <Input.TextArea placeholder="请输入发布说明"></Input.TextArea>
            )
          }
        </Form.Item>
      </Form>
    </Modal>
  )
};

export default Form.create<Props>({
  onFieldsChange: (props, fields) => {
    const { kind } = fields;
    const { form } = props;
    if (kind === 'batch') {
      form.resetFields(['percent']);
    }
    if (kind === 'userPercent') {
      form.resetFields(['groups', 'users']);
    }
  },
})(PublishTagModal);