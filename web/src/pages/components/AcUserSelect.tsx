import React, { useState, useEffect } from 'react';
import { connect } from 'dva';
import { debounce } from 'lodash';
import { Select, Spin } from 'antd';
import { AcUserSelectComponentProps, User } from '../declare';

const AcUserSelect: React.FC<AcUserSelectComponentProps> = (props) => {
  const { dispatch, onChange, defaultSelectedUser = [] } = props;
  const [selectedUser, setSelectedUser] = useState(defaultSelectedUser.map(user => `${ user.uid }/${ user.name }`));
  const [fetching, setFetching] = useState(false);
  const [userList, setUserList] = useState<User[]>([]);
  useEffect(() => {
    dispatch({
      type: 'users/getAcUsers',
      payload: {
        cb: (users: User[]) => {
          setUserList(users);
        }
      }
    });
  }, [dispatch]);
  const handleSearch = debounce((value: string) => {
    setFetching(true);
    if (value) {
      dispatch({
        type: 'users/searchAcUsers',
        payload: {
          params: {
            key: value,
          },
          cb: (users: User[]) => {
            setUserList(users);
            setTimeout(() => setFetching(false), 1000);
          }
        },
      });
    } else {
      dispatch({
        type: 'users/getAcUsers',
        payload: {
          cb: (users: User[]) => {
            setUserList(users);
            setTimeout(() => setFetching(false), 1000);
          }
        }
      });
    }
  }, 800);
  const handleChange = (value: string[]) => {
    if (onChange) {
      const users = value.map(user => user.split('/')[0])
      onChange(users);
    }
    setSelectedUser(value);
  };
  return (
    <Select
      mode="multiple"
      placeholder="请输入搜索关键字"
      onSearch={ handleSearch }
      value={ selectedUser }
      onChange={ handleChange }
      notFoundContent={ fetching ? <Spin size="small" /> : 'Not Found' }
      getPopupContainer={ () => window.document.body }
    >
      {
        (!!userList.length) && userList.map(user => {
          return (
            <Select.Option
              key={ user.uid }
              value={ `${ user.uid }/${ user.name }` }
            >
              { `${ user.name }` }
            </Select.Option>
          )
        })
      }
    </Select>
  );
};

export default connect()(AcUserSelect);