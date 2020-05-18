import React, { useState, useMemo } from 'react';
import { Link } from 'umi';
import { Menu, Icon } from 'antd';
import styles from './index.less';
import Products from './components/Products';
import Title from './components/Title';

const BasicLayout: React.FC = props => {
  const { history } = props;
  const [activeKey, changeActiveKey] = useState(window.location.pathname);
  const handleActiveKeyChange = (key: string) => {
    changeActiveKey(key);
  };
  const defaultOpenSubs = useMemo(() => {
    const pathname = window.location.pathname;
    const pathReg = /^\/products(?:\/([^\/#\?]+?))(?:\/([^\/#\?]+?))[\/#\?]?$/i;
    const routePath = pathReg.exec(pathname);
    return routePath ? [`/products/${routePath[1]}`] : [];
  }, []);
  return (
    <div className={styles.parent}>
      <div className={styles.sidebar}>
        <div className={styles.title}>
          <Title></Title>
        </div>
        <div className={styles.sidemenu}>
          <Products
            isRootRoute={window.location.pathname === '/'}
            history={history}
            onActiveKeyChange={handleActiveKeyChange}
            activeKeys={[activeKey]}
            defaultOpenSub={defaultOpenSubs}
          />
        </div>
        <div className={styles.quicklink}>
          <Menu
            selectedKeys={[activeKey]}
            onSelect={({ key }) => handleActiveKeyChange(key)}
          >
            <Menu.Item
              key="/group"
            >
              <Link to="/group">
                <Icon type="container" />
                群组
              </Link>
            </Menu.Item>
            <Menu.Item
              key="/user"
            >
              <Link to="/user">
                <Icon type="user"></Icon>
                用户
              </Link>
            </Menu.Item>
            <Menu.Item
              key="/ac"
            >
              <Link to="/ac">
                <Icon type="container" />
                管理
              </Link>
            </Menu.Item>
            <Menu.Item
              key="/help"
            >
              <Link to="/help">
                <Icon type="question-circle" />
                使用帮助
              </Link>
            </Menu.Item>
          </Menu>
        </div>
      </div>
      <div className={styles.mainframe}>
        {props.children}
      </div>
    </div>
  );
};

export default BasicLayout;
