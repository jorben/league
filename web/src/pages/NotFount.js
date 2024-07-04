import React from 'react'
import { Result, Button } from 'antd';
import {HomeOutlined} from '@ant-design/icons'



const NotFount = () => (
  <Result
    status="404"
    title="404"
    subTitle="Sorry, the page you visited does not exist."
    extra={<a href="/" ><Button type="primary" icon={<HomeOutlined />}>回到首页</Button></a>}
  />
);

export default NotFount
