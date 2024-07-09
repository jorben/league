import React, { useEffect, useState } from "react";
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
  Carousel,
} from "antd";
import {
  QqOutlined,
  GoogleOutlined,
  GithubOutlined,
  LoadingOutlined,
} from "@ant-design/icons";
import Logo from "../components/Logo";
import AdminFooter from "./admin/layout/AdminFooter";
import BackgroundImg00 from "../assets/images/login_bg_00.png";
import BackgroundImg01 from "../assets/images/login_bg_01.png";
import BackgroundImg02 from "../assets/images/login_bg_02.png";
import BackgroundImg03 from "../assets/images/login_bg_03.png";
import BackgroundImg04 from "../assets/images/login_bg_04.png";
import { useLocation, useNavigate } from "react-router-dom";
import ApiClient from "../services/client";
import CONSTANTS from "../constants";

const Login = () => {
  const { token } = theme.useToken();
  const location = useLocation();
  const navigate = useNavigate();
  const [wxLoginUrl, setWxLoginUrl] = useState("");
  const searchParams = new URLSearchParams(location.search);
  const isCallback = searchParams.get("callback");

  const [messageApi, contextHolder] = message.useMessage();

  useEffect(() => {
    if (isCallback) {
      //登录回调场景
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
    } else {
      // 登录场景，获取微信地址
      const getWxLoginUrl = async () => {
        ApiClient.get("/auth/login?type=wechat&url=1")
          .then((response) => {
            if (response.data?.code === 0) {
              setWxLoginUrl(response.data?.data);
            } else {
              messageApi.error(response.data?.message);
            }
          })
          .catch((error) => {
            console.log(error);
            messageApi.error("加载微信登录二维码失败，请稍后重试！");
          });
      };
      getWxLoginUrl();
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
              <Col span={8}>
                <Carousel
                  autoplay={true}
                  dots={false}
                  fade={true}
                  speed={2000}
                  autoplaySpeed={10000}
                >
                  <img src={BackgroundImg00} alt="" width="400" />
                  <img src={BackgroundImg01} alt="" width="400" />
                  <img src={BackgroundImg02} alt="" width="400" />
                  <img src={BackgroundImg03} alt="" width="400" />
                  <img src={BackgroundImg04} alt="" width="400" />
                </Carousel>
              </Col>
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
                        {wxLoginUrl ? (
                          <iframe
                            style={{
                              margin: 0,
                              padding: 0,
                              border: "none",
                              height: 330,
                              overflow: "hidden",
                            }}
                            scrolling="no"
                            title="微信扫码登录"
                            src={wxLoginUrl}
                          />
                        ) : (
                          <>
                            <div
                              style={{
                                fontSize: "20px",
                                lineHeight: "1.6",
                                marginBottom: "15px",
                              }}
                            >
                              微信登录
                            </div>
                            <QRCode
                              value="登录二维码加载中..."
                              size="128"
                              status="loading"
                            />
                          </>
                        )}
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
                          shape="round"
                          href={CONSTANTS.BASEURL_API + "/auth/login?type=qq"}
                          icon={<QqOutlined />}
                        >
                          QQ
                        </Button>
                        <Button
                          shape="round"
                          href={
                            CONSTANTS.BASEURL_API + "/auth/login?type=google"
                          }
                          icon={<GoogleOutlined />}
                        >
                          Google
                        </Button>
                        <Button
                          shape="round"
                          href={
                            CONSTANTS.BASEURL_API + "/auth/login?type=github"
                          }
                          icon={<GithubOutlined />}
                        >
                          Github
                        </Button>
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
};

export default Login;
