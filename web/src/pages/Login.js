import React, { useEffect } from "react";
import {
  Col,
  Layout,
  Row,
  Divider,
  theme,
  Space,
  Button,
  QRCode,
  Flex,
  Spin,
  message,
} from "antd";
import {
  QqOutlined,
  GoogleOutlined,
  GithubOutlined,
  LoadingOutlined,
} from "@ant-design/icons";
import Logo from "../components/Logo";
import AdminFooter from "./admin/layout/AdminFooter";
import BackgroundImg from "../assets/images/loginbackground2@2x.png";
import { useLocation, useNavigate } from "react-router-dom";
import ApiClient from "../services/client";
import CONSTANTS from "../constants";

function Login() {
  const { token } = theme.useToken();
  const location = useLocation();
  const navigate = useNavigate();
  const searchParams = new URLSearchParams(location.search);
  const isCallback = searchParams.get("callback");

  const [messageApi, contextHolder] = message.useMessage();

  useEffect(() => {
    if (isCallback) {
      const loginCallback = async () => {
        ApiClient.get("/auth/callback" + location.search)
          .then((response) => {
            // console.log("/auth/callback", response.data);
            if (response.data?.code === 0) {
              // 存储jwt
              localStorage.setItem(
                CONSTANTS.STORAGE_KEY_JWT,
                JSON.stringify(response.data?.data)
              );
              // 跳转页面
              // TODO: 支持跳转回登录来源页面，并做同源校验
              navigate("/admin");
            } else {
              messageApi.error(response.data?.message);
            }
          })
          .catch((error) => {
            console.log(error);
            messageApi.error("获取用户信息失败，请稍后重试！");
          });
      };
      loginCallback();
    }
  }, [messageApi, isCallback, navigate, location]);

  return (
    <Layout
      style={{
        width: "100%",
        minHeight: "100vh",
        background: "rgba(80,80,80,.2)",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      {isCallback ? (
        <Spin
          size="large"
          indicator={<LoadingOutlined spin />}
          tip="加载中..."
          fullscreen
        />
      ) : (
        <>
          <div
            style={{
              width: "1200px",
              height: "680px",
              boxShadow: "0 0 12px 12px rgba(10,10,10,0.1)",
              background: token.colorBgContainer,
            }}
          >
            <Row>
              <Col
                span={8}
                style={{
                  height: 680,
                  backgroundImage: `url(${BackgroundImg})`,
                  backgroundSize: "cover",
                }}
              ></Col>
              <Col span={16}>
                <Row style={{ padding: "12px", justifyContent: "center" }}>
                  <Logo collapsed={false} theme="light" />
                </Row>
                <Row>
                  <Col span={20} offset={2}>
                    <Divider style={{ margin: "0 0 20px 0" }} />
                  </Col>
                </Row>
                <Row>
                  <Col span={16} offset={4}>
                    <Row
                      style={{ justifyContent: "center", textAlign: "center" }}
                    >
                      <Flex vertical>
                        <div>
                          <h2>微信登录</h2>
                        </div>
                        <QRCode
                          value="https://league.yation.com/login?type=wechat"
                          size="112"
                          status="loading"
                        />
                        <div style={{ color: "#666", marginTop: "12px" }}>
                          <p>使用微信扫一扫登录</p>
                          <p>"League"</p>
                        </div>
                      </Flex>
                    </Row>
                    <Row>
                      <Divider style={{ margin: "32px 0 24px 0" }} />
                      <Space>
                        <span>其他登录方式：</span>
                        <Button
                          shape="circle"
                          href={CONSTANTS.BASEURL_API + "/auth/login?type=qq"}
                          icon={<QqOutlined />}
                        />
                        <Button
                          shape="circle"
                          href={
                            CONSTANTS.BASEURL_API + "/auth/login?type=google"
                          }
                          icon={<GoogleOutlined />}
                        />
                        <Button
                          shape="circle"
                          href={
                            CONSTANTS.BASEURL_API + "/auth/login?type=github"
                          }
                          icon={<GithubOutlined />}
                        />
                      </Space>
                    </Row>
                  </Col>
                </Row>
              </Col>
            </Row>
          </div>
          <AdminFooter noBackground={true} />
        </>
      )}
      {contextHolder}
    </Layout>
  );
}

export default Login;
