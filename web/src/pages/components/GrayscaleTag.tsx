import React from 'react';
import { Table } from 'antd';
import { formatTableTime } from '../utils/format';
import { TableComponentProps, User, Label, ActionEventListeners } from '../declare';
import { Pagination } from './';

const GrayscaleTag: React.FC<TableComponentProps<any>> = (props) => {
  const { paginationProps, onAction, hideColumns } = props;
  const columns = [{
    title: '名称',
    dataIndex: 'name',
    key: 'name',
  }, {
    title: '负责人',
    dataIndex: 'users',
    key: 'users',
    render: (users: User[]) => {
      return Array.isArray(users) ? users.map(item => item.name).join() : '';
    }
  }, {
    title: '所属产品',
    dataIndex: 'product',
    key: 'product',
  }, {
    title: '描述',
    dataIndex: 'desc',
    key: 'desc',
  }, {
    title: '灰度进度',
    dataIndex: 'status',
    key: 'status',
    render: (status: number) => {
      const len = String(status).length;
      return `小于${ Math.pow(10, len) }`;
    },
  }, {
    title: '发布次数',
    dataIndex: 'release',
    key: 'release',
  }, {
    title: '更新时间',
    dataIndex: 'updatedAt',
    key: 'updatedAt',
    render: (time: string) => {
      return formatTableTime(time);
    },
  }, {
    title: '设置时间',
    dataIndex: 'createdAt',
    key: 'createdAt',
    render: (time: string) => {
      return formatTableTime(time);
    },
  }, {
    title: '',
    dataIndex: 'action',
    key: 'action',
    render: (_, record: Label) => {
      const { onDelete } = onAction ? onAction(record) : ({} as ActionEventListeners);
      return (
        <a onClick={ onDelete }>移除</a>
      );
    },
  }];
  const generateTableColumns = () => {
    if (hideColumns) {
      return columns.filter(item => !hideColumns.includes(item.key));
    }
    return columns;
  };
  return (
    <div>
      <Table rowKey="hid" { ...props } columns={ generateTableColumns() } pagination={ false }></Table>
      {
        paginationProps && (<Pagination { ...paginationProps }></Pagination>)
      }
    </div>
  );
};

export default GrayscaleTag;