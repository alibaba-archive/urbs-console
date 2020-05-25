import React, { useState } from 'react';
import { Tabs } from 'antd';
import { ContentTabsComponentProps, ContentTabsItem } from '../declare';
import styleNames from './style/ContentTabs.less';

const TabPane = Tabs.TabPane;

const ContentTabs: React.FC<ContentTabsComponentProps> = (props) => {
  const { tabs, activeKey, handleActiveKeyChange } = props;
  const [activeTab, setActiveTab] = useState(activeKey || tabs[0].key);
  const renderTabPane = (tabItem: ContentTabsItem) => {
    const { key, title, content } =  tabItem;
    return (
      <TabPane
        key={ key }
        tab={ title }
      >
        { content }
      </TabPane>
    )
  };
  const renderTabAction = () => {
    const activeTabItem = tabs.find(item => item.key === activeTab);
    return activeTabItem?.action;
  };
  const handleTabsActiveKeyChange = (key: string) => {
    setActiveTab(key);
    if (handleActiveKeyChange) {
      handleActiveKeyChange(key);
    }
  };
  return (
    <div className={ styleNames['tabs-wrap'] }>
      <Tabs activeKey={ activeTab } onChange={ handleTabsActiveKeyChange }>
        {
          tabs.map(tabItem => renderTabPane(tabItem))
        }
      </Tabs>
      <div className={ styleNames['tabs-action'] }>
        {
          renderTabAction()
        }
      </div>
    </div>
  )
};

export default ContentTabs;