import React, { useState } from 'react';
import { TableTitleComponentProps } from '../declare';
import { Button, Divider, Input } from 'antd';
import styleNames from './style/TableTitle.less';

const TableTitle: React.FC<TableTitleComponentProps> = (props) => {
  const {
    handlePlusClick,
    plusTitle,
    defaultSearchValue = '',
    handleSearch,
    handleWordChange,
  } = props;
  const [searchWord, setSearchWord] = useState(defaultSearchValue);
  const onChange = (e: React.ChangeEvent) => {
    const target = e.target;
    const value = target.value;
    setSearchWord(value);
    if (handleWordChange) {
      handleWordChange(value);
    }
  };
  const onSearch = (value: string) => {
    if (handleSearch) {
      handleSearch(value);
    }
  };
  return (
    <>
      <div className={ styleNames['op-wrap'] }>
        <Button
          className={styleNames['op-btn']}
          type="link"
          icon={plusTitle ? 'plus' : undefined}
          onClick={ plusTitle ? handlePlusClick : undefined}
        >
          { plusTitle }
        </Button>
        <Input.Search
          className={styleNames['op-input']}
          placeholder="请输入搜索关键字"
          value={ searchWord }
          onChange={ onChange }
          onSearch={ onSearch }
          allowClear
        />
      </div>
      <Divider className={styleNames['op-divider']}></Divider>
    </>
  )
};

export default TableTitle;