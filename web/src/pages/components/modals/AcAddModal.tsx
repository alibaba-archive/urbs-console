import React from 'react';
import { Modal, Form, Input, Button, message } from 'antd';
import { FormComponentProps } from 'antd/es/form';
import { Group, User, DEFAULT_FORM_ITEM_LAYOUT, FieldsValue } from '../../declare';

interface Props extends FormComponentProps {
    visible: boolean;
    onCancel?: (e: React.MouseEvent<HTMLElement>) => void;
    onOk?: (fieldsValue: FieldsValue, isEdit?: boolean) => void;
    onDel?: (e: React.MouseEvent<HTMLElement>) => void;
    User?: User;
}

const AcAddModal: React.FC<Props> = (props) => {
    const { User, visible, onCancel, onOk, onDel, form } = props;
    const { getFieldDecorator } = form;
    const isEdit = !!User;
    const modalTitle = `${isEdit ? '编辑' : '添加'}用户`;
    const handleOnOk = () => {
        const values = form.getFieldsValue();
        const { uid, name } = values;
        if (!uid) {
            message.info('请输入名称');
            return
        }
        if (!name) {
            message.info('请输入用户名');
            return
        }
        if (onOk) {
            onOk(values, isEdit);
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
                            <Button type="danger" onClick={onDel}>删除</Button>
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
                <Form.Item label="ID">
                    {
                        getFieldDecorator('uid', {
                            initialValue: User ? User.uid : undefined,
                        })(
                            <Input disabled={isEdit} placeholder="请输入uid"></Input>
                        )
                    }
                </Form.Item>
                <Form.Item label="名称">
                    {
                        getFieldDecorator('name', {
                            initialValue: User ? User.name : undefined,
                        })(
                            <Input placeholder="请输入用户名"></Input>
                        )
                    }
                </Form.Item>
            </Form>
        </Modal>
    );
};

export default Form.create<Props>({
    onFieldsChange(props, field, allFields) {
        console.log(props, field, allFields);
    },
    onValuesChange(props, changeValues, allValues) {
        console.log(props, changeValues, allValues);
    },
})(AcAddModal);
