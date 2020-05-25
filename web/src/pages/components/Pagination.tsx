import React from 'react';
import { Select, Icon } from 'antd';
import classNames from 'classnames';
import { PaginationComponentProps } from '../declare';
import styleNames from './style/Pagination.less';

enum PageTurn {
  next = 'next',
  pre = "pre",
}

const Pagination: React.FC<PaginationComponentProps> = (props) => {
  const { pageSize = 20, total, prePageToken, nextPageToken, pageSizeOptions, onTokenChange, onPageSizeChange } = props;
  const handleClick = (type: PageTurn, token?: string) => {
    if (onTokenChange && typeof token === 'string') {
      onTokenChange(type, token);
    }
  };
  const handleSelect = (value: number | string) => {
    if (onPageSizeChange) {
      onPageSizeChange(+value);
    }
  };
  return (
    <ul className={ styleNames.pagination }>
      {
        total && (
          <li className={ styleNames['pagination-total-text'] }>
            共 { total } 条
          </li>
        )
      }
      <li
        onClick={ () => handleClick(PageTurn.pre, prePageToken) }
        className={ classNames(styleNames['pagination-prev'], { [styleNames['pagination-disabled']]: typeof prePageToken !== 'string' }) }
      >
        <a>
          <Icon type="left" /> 上一页
        </a>
      </li>
      <li className={ styleNames['pagination-item'] }>
        <a> / </a>
      </li>
      <li
        onClick={ () => handleClick(PageTurn.next, nextPageToken) }
        className={ classNames(styleNames['pagination-next'], { [styleNames['pagination-disabled']]: typeof nextPageToken !== 'string' }) }
      >
        <a>
          下一页 <Icon type="right" />
        </a>
      </li>
      {
        (pageSizeOptions && Array.isArray(pageSizeOptions) && pageSizeOptions.length) && (
          <li className={ styleNames['pagination-options'] }>
            <Select
              defaultValue={ pageSize }
              className={ styleNames['pagination-options-size-changer'] }
              onSelect={ handleSelect }
            >
              {
                pageSizeOptions.map((item, index) => {
                  return (
                    <Select.Option
                      key={ `${item}_${index}` }
                      value={ item }
                    >
                      { item }条 / 页
                    </Select.Option>
                  )
                })
              }
            </Select>
          </li>
        )
      }
    </ul>
  );
};

export default Pagination;