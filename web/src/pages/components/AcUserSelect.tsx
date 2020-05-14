import React, { useState, useEffect } from 'react';
import { connect } from 'dva';
import { debounce } from 'lodash';
import { Select } from 'antd';
import { AcUserSelectComponentProps } from '../declare';

const AcUserSelect: React.FC<AcUserSelectComponentProps> = (props) => {
  const { dispatch, onChange, acUserList, defaultSelectedUser = [] } = props;
  const [selectedUser, changeSelectedUser] = useState(defaultSelectedUser.map(user => user.uid));
  useEffect(() => {
    dispatch({
      type: 'users/getAcUsers'
    });
  }, [dispatch]);
  const handleDropdownVisibleChange = (open: boolean) => {
    if (open) {
      dispatch({
        type: 'users/getAcUsers'
      });
    }
  };
  const handleSearch = debounce((value: string) => {
    if (value) {
      dispatch({
        type: 'users/searchAcUsers',
        payload: {
          params: {
            key: value,
          },
        },
      });
    } else {
      dispatch({
        type: 'users/getAcUsers'
      });
    }
  }, 500);
  const handleSelect = (value: string) => {
    const selected = [...selectedUser, value];
    changeSelectedUser(selected);
    if (onChange) {
      onChange(selected);
    }
  };
  const handleDeSelect = (value: string) => {
    const selected = selectedUser.filter(item => item !== value);
    changeSelectedUser(selected);
    if (onChange) {
      onChange(selected);
    }
  };
  return (
    <Select
      mode="multiple"
      placeholder="请输入搜索关键字"
      onDropdownVisibleChange={ handleDropdownVisibleChange }
      onSearch={ handleSearch }
      value={ selectedUser }
      onSelect={ handleSelect }
      onDeselect={ handleDeSelect }
    >
      {
        acUserList && acUserList.map(user => {
          return (
            <Select.Option
              key={ user.uid }
              value={ user.uid }
            >
              { user.name }
            </Select.Option>
          )
        })
      }
    </Select>
  );
};

export default connect(state => {
  const { acUserList } = state.users;
  return { acUserList }
})(AcUserSelect);