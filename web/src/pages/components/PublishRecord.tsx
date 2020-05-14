import React, { useCallback } from 'react';
import { Divider, Button } from 'antd';
import styleNames from './style/PublishRecord.less';
import { formatTableTime } from '../utils/format';
import { PublishRecordItem } from '../declare';

interface Props {
  publishRecordList: PublishRecordItem[];
}

const PublishRecord: React.FC<Props> = (props) => {
  const { publishRecordList } = props;
  const getActionLabel = useCallback((action: string) => {
    switch (action) {
      case 'create':
        return '增加';
      case 'update':
        return '更新';
      default:
        return '';
    }
  }, []);
  const getActionDesc = (kind: string, precent?: number, users?: string[], groups?: string[]) => {
    if (kind === 'userPercent') {
      return `比例${ precent }`
    } else {
      return `${ users ? `用户 ${ users.join(',') }；` : ''}${ groups ? `用户 ${ groups.join(',') }；` : ''}`;
    }
  };
  return (
    <>
      {
        publishRecordList.map((item, index) => {
          return (
            <div key={ item.hid } className={ styleNames['publish-record-wrap'] }>
              <div className={ styleNames['record-dot-wrap'] }>
                <div className={ styleNames['record-dot'] }></div>
                <Divider type="vertical" style={{ height: '100%' }}></Divider>
              </div>
              <div className={ styleNames['publish-record-content-wrap'] }>
                <div>{ formatTableTime(item.createdAt) }, { item.operatorName }</div>
                <Divider style={{ margin: '5px 0' }}></Divider>
                <div className={ styleNames['publish-record-action-wrap'] }>
                  { index === 0 ? <Button type="link" block>撤回</Button> : null }
                </div>
                <ul>
                  <li>
                    <span>操作：{ `${ getActionLabel(item.action) }${ getActionDesc(item.kind, item.precent, item.users, item.groups) }` }</span>
                  </li>
                  <li>
                    <span>发布说明：{ item.desc }</span>
                  </li>
                </ul>
              </div>
            </div>
          )
        })
      }
    </>
  );
};

export default PublishRecord;