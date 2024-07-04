import React from 'react'
import { Layout } from 'antd';
const {Footer} = Layout
const AdminFooter = (noBackground=false) => {
    return (
        <Footer style={{ 
            textAlign: 'center',
            background: noBackground ? 'none' : '#f5f5f5',
        }}>Copyright &copy; 2020-{new Date().getFullYear()} League All Rights Reserved.</Footer>
    )
}

export default AdminFooter;