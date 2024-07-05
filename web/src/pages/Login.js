import React from 'react'
import { Col, Layout, Row, Divider, theme, Space, Button, QRCode, Flex } from 'antd';
import {QqOutlined,GoogleOutlined, GithubOutlined} from '@ant-design/icons'
import Logo from '../components/Logo';
import AdminFooter from './admin/layout/AdminFooter';
import BackgroundImg from '../assets/images/loginbackground2@2x.png'

function Login() {
    const { token } = theme.useToken();
    return (
        <Layout style={{width: '100%', minHeight:"100vh", background:"rgba(80,80,80,.2)", alignItems: 'center', justifyContent: "center"}}>
            <div style={{
                width: "1200px",
                height: "680px",
                boxShadow: "0 0 12px 12px rgba(10,10,10,0.1)",
                background: token.colorBgContainer,
            }}>
                <Row>
                    <Col span={8} style={{height:680, backgroundImage:`url(${BackgroundImg})`, backgroundSize: 'cover'}}></Col>
                    <Col span={16}>
                        <Row style={{padding:"12px", justifyContent:"center"}}><Logo collapsed={false} theme="light"/></Row>
                        <Row><Col span={20} offset={2}><Divider style={{margin:"0 0 20px 0"}}/></Col></Row>
                        <Row>
                            <Col span={16} offset={4}>

                            <Row style={{justifyContent:"center", textAlign:"center"}}>
                                <Flex vertical>
                                    <div><h2>微信登录</h2></div>
                                    <QRCode value="https://league.yation.com/login?type=wechat" size="112" status="loading" />
                                    <div style={{color:"#666", marginTop:"12px"}}><p>使用微信扫一扫登录</p><p>"League"</p></div>
                                </Flex>
                            </Row>
                            <Row>
                                <Divider style={{margin:"32px 0 24px 0"}} />
                                <Space>
                                    <span>其他登录方式：</span>
                                    <Button shape="circle" href='/auth/login?type=qq' icon={<QqOutlined />} />
                                    <Button shape="circle" href='/auth/login?type=google' icon={<GoogleOutlined />} />
                                    <Button shape='circle' href='/auth/login?type=github' icon={<GithubOutlined />} />
                                </Space>
                                
                            </Row>
                            </Col>
                        </Row>
                    </Col>
                </Row>
            </div>
            <AdminFooter noBackground={true}/>
        </Layout>
    )
}

export default Login
