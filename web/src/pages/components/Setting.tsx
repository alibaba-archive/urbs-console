import React from 'react';
import { Table } from 'antd';
import { formatTableTime } from '../utils/format';
import { TableComponentProps, User, Setting } from '../declare';
import { Pagination } from './';

const Settings: React.FC<TableComponentProps<Setting>> = (props) => {
  const { hideColumns, paginationProps, onAction } = props;
  const columns = [{
    title: '名称',
    dataIndex: 'name',
    key: 'name',
  }, {
    title: '负责人',
    dataIndex: 'users',
    key: 'users',
    render: (users: User[]) => {
      return Array.isArray(users) ? users.map(user => user.name).join(',') : '';
    },
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
      return `${ formatTableTime(time) }`;
    },
  }, {
    title: '创建时间',
    dataIndex: 'createdAt',
    key: 'createdAt',
    render: (time: string) => {
      return `${ formatTableTime(time) }`;
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

export default Settings;