import React from 'react';
import { Row, Col } from 'antd';
import { ContentDetailComponentProps } from '../declare';
import styleNames from './style/ContentDetail.less';

const ContentDetail: React.FC<ContentDetailComponentProps> = (props) => {
  const { content } = props;
  return (
      <div className={ styleNames['detail-wrap'] }>
        {
          content && content.map((item, index) => (
            <Row
              key={ `${ item.title }_${ index }` }
              className={ styleNames['detail-row'] }
            >
              <Col span={ 3 } className={ styleNames['detail-content__title'] }>
                { item.title }
              </Col>
              <Col span={ 21 } className={ styleNames['detail-content__desc'] }>
                { item.content }
              </Col>
            </Row>
          ))
        }
      </div>
  )
};

export default ContentDetail;