import React, { useCallback } from 'react';
import { Divider, Button, Modal } from 'antd';
import styleNames from './style/PublishRecord.less';
import { formatTableTime } from '../utils/format';
import { PublishRecordItem } from '../declare';

interface Props {
  publishRecordList: PublishRecordItem[];
  onReback?: (hid: string) => void;
  canRecall: boolean;
}

const PublishRecord: React.FC<Props> = (props) => {
  const { publishRecordList, canRecall } = props;
  const getActionLabel = useCallback((action: string) => {
    switch (action) {
      case 'create':
        return '增加';
      case 'update':
        return '更新';
      case 'cleanup':
        return '清空';
      default:
        return '';
    }
  }, []);
  const getActionDesc = (kind: string, percent?: number, users?: string[], groups?: string[]) => {
    if (kind === 'userPercent') {
      return `用户比例到 ${percent}%`
    } else if (kind === 'newUserPercent') {
      return `新用户比例到 ${percent}%`
    } else if (kind === 'childLabelUserPercent') {
      return `灰中灰用户比例到 ${percent}%`
    } else {
      return `${users ? `用户 ${users.join(',')}；` : ''}${groups ? `群组 ${groups.join(',')}；` : ''}`;
    }
  };
  const handleReback = (item: PublishRecordItem) => {
    return () => {
      const { onReback, publishRecordList } = props;
      if (onReback && publishRecordList.length) {
        Modal.confirm({
          title: '操作不可逆，请再次确认',
          content: '确认撤回？',
          onOk: () => onReback(item.hid),
        });
      }
    }
  };
  return (
    <>
      {
        publishRecordList.map((item, index) => {
          return (
            <div key={item.hid} className={styleNames['publish-record-wrap']}>
              <div className={styleNames['record-dot-wrap']}>
                <div className={styleNames['record-dot']}></div>
                <Divider type="vertical" style={{ height: '100%' }}></Divider>
              </div>
              <div className={styleNames['publish-record-content-wrap']}>
                <div>{formatTableTime(item.createdAt)}, {item.operatorName}</div>
                <Divider style={{ margin: '5px 0' }}></Divider>
                <div className={styleNames['publish-record-action-wrap']}>
                  {canRecall && !('percent' in item) ? <Button type="link" block onClick={handleReback(item)}>撤回</Button> : null}
                </div>
                <ul>
                  <li>
                    <span>操作：{`${getActionLabel(item.action)}${getActionDesc(item.kind, item.percent, item.users, item.groups)}`}</span>
                  </li>
                  <li>
                    <span>发布说明：{item.desc}</span>
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