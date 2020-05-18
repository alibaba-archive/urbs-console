import React from 'react';
import { Table } from 'antd';
import { formatTableTime } from '../utils/format';
import { TableComponentProps, GroupMember, ActionEventListeners } from '../declare';
import { Pagination } from './';

const Users: React.FC<TableComponentProps<any>> = (props) => {
  const { paginationProps, onAction, hideColumns } = props;
  const columns = [{
    title: 'ID',
    dataIndex: 'user',
    key: 'user',
  }, {
    title: '同步时间',
    dataIndex: 'syncAt',
    key: 'syncAt',
    render: (time: string) => {
      return formatTableTime(time);
    },
  }, {
    title: '创建时间',
    dataIndex: 'createdAt',
    key: 'createdAt',
    render: (time: string) => {
      return formatTableTime(time);
    },
  }, {
    title: '',
    dataIndex: 'action',
    key: 'action',
    render: (_, record: GroupMember) => {
      const { onDelete } = onAction ? onAction(record) : ({} as ActionEventListeners);
      return (
        <a onClick={onDelete}>移除</a>
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
      <Table rowKey="user" {...props} columns={generateTableColumns()} pagination={false}></Table>
      {
        paginationProps && (<Pagination {...paginationProps}></Pagination>)
      }
    </div>
  );
};

export default Users;