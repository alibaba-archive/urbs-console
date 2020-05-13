import React, { useEffect, useCallback } from 'react';
import { Menu, Icon } from 'antd';
import { Link } from 'umi';
import { connect } from 'dva';

const SubMenu = Menu.SubMenu;

const Products = ({dispatch, productList, isRootRoute, history, onActiveKeyChange, activeKeys, defaultOpenSub}) => {
  useEffect(() => {
    dispatch({
      type: 'products/getProducts',
      payload: {
        needRedirect: isRootRoute,
        redirectTo: (name: string) =>  history.push(`/products/${ name }`),
      },
    });
  }, [dispatch, history, isRootRoute]);
  const renderSubTitle = useCallback((title: string) => {
    return (
      <span to={ `/products/${ title }` }>
        <Icon type="ant-design" />
        { title }
      </span>
    );
  }, []);
  const handleSubClick = (key: string) => {
    history.push(`/products/${ key }`);
    onActiveKeyChange(`/products/${ key }`);
  };
  const handleItemClick = (key: string) => {
    onActiveKeyChange(key);
  };

  if (productList && productList.length > 0) {
    const menus = productList.map(elem => {
        return <SubMenu
          key= { `/products/${ elem.name }` }
          title={ renderSubTitle(elem.name) }
          onTitleClick={ () => handleSubClick(elem.name) }
        >
          <Menu.Item key={ `/products/${ elem.name }/tag` }>
            <Link to={ `/products/${ elem.name }/tag` }>
              灰度标签
            </Link>
          </Menu.Item>
          <Menu.Item key={ `/products/${ elem.name }/module` }>
            <Link to={ `/products/${ elem.name }/module` }>
              功能模块
            </Link>
          </Menu.Item>
          <Menu.Item key={ `/products/${ elem.name }/setting` }>
            <Link to={ `/products/${ elem.name }/setting` }>
              配置项
            </Link>
          </Menu.Item>
        </SubMenu>
    });
    return (
      <Menu
        mode="inline"
        defaultOpenKeys={ defaultOpenSub }
        selectedKeys={activeKeys}
        onSelect={ ({key}) => handleItemClick(key) }
      >
        { menus }
      </Menu>
    )
  } else {
    return <></>
  }
  
}

function mapStateToProps(state) {
  const { productList } = state.products;  
  return {
    productList,
  }
}

export default connect(mapStateToProps)(Products)