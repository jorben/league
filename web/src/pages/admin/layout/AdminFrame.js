import React, {useState} from 'react'
import {Navigate, Route, Routes} from 'react-router-dom'
import {adminRoutes} from '../../../routes'
import { Layout, theme } from 'antd';
import AdminHeader from './AdminHeader';
import AdminFooter from './AdminFooter';
import AdminSider from './AdminSider';

const { Content } = Layout;


const AdminFrame = () => {
  const [collapsed, setCollapsed] = useState(false);
  const { token: { colorBgContainer, borderRadiusLG }} = theme.useToken();
  return (
    <Layout style={{minHeight:"100vh"}}>
      <AdminSider collapsed={collapsed}/>
      <Layout>
          <AdminHeader collapsed={collapsed} setCollapsed={setCollapsed}/>
          <Content
          style={{
            margin: '24px 16px 0 16px',
            padding: 24,
            minHeight: 280,
            background: colorBgContainer,
            borderRadius: borderRadiusLG,
          }}
        >
          <Routes>
            {adminRoutes.map((r, index) => {
                return <Route key={index} {...r} />
            })}
            <Route path="*" element={<Navigate to="/404" replace />} />
          </Routes>
        </Content>
        <AdminFooter />
      </Layout>
    </Layout>
  )
}

export default AdminFrame
