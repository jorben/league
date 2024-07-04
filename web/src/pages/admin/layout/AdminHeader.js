import React from 'react'
import {MenuFoldOutlined, MenuUnfoldOutlined, UserOutlined, LogoutOutlined} from '@ant-design/icons';
import { Button, Space, Avatar, Flex,Layout, theme, Dropdown } from 'antd';

const { Header } = Layout;

const AdminHeader = ({ collapsed, setCollapsed }) => {
    const {
      token: { colorBgContainer },
    } = theme.useToken();
    const items = [
        {
          label: <a href="/auth/logout">退出登录</a>,
          key: '0',
          icon: <LogoutOutlined />
        },
      ];

    return (
        <Header style={{ padding: '0 24px 0 0', background: colorBgContainer }}>
            <Button
              type="text"
              icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
              onClick={() => setCollapsed(!collapsed)}
              style={{
                fontSize: '16px',
                width: 64,
                height: 64,
              }}
            />
            <Dropdown menu={{items}} placement='bottom' arrow >
                <Button type="text" style={{height: 64, float:'right'}}>
                    <Space wrap size={16} >
                        <Avatar size="large" icon={<UserOutlined/>} />
                        <Flex vertical gap={2} style={{lineHeight:'20px', textAlign: 'left'}}>
                            <strong>Jorben Zhu</strong>
                            <span>jorbenzhu@gmail.com</span>
                        </Flex>
                    </Space>
                </Button>
            </Dropdown>
          </Header>
    )}

    export default AdminHeader;