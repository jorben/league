import React from 'react'
import { Layout, Menu } from 'antd';
import {UploadOutlined, UserOutlined, VideoCameraOutlined} from '@ant-design/icons';
import Logo from '../../../components/Logo';
const {Sider} = Layout

const AdminSider = ({collapsed}) => {
    return (
        <Sider trigger={null} collapsible collapsed={collapsed}>
        <Logo collapsed={collapsed} theme="black"/>
        <Menu
          theme="dark"
          mode="inline"
          defaultSelectedKeys={['1']}
          items={[
            {
              key: '1',
              icon: <UserOutlined />,
              label: 'nav 1',
            },
            {
              key: '2',
              icon: <VideoCameraOutlined />,
              label: 'nav 2',
            },
            {
              key: '3',
              icon: <UploadOutlined />,
              label: 'nav 3',
            },
          ]}
        />
      </Sider>
    )
}

export default AdminSider;