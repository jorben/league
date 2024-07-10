import React, { useState, useEffect } from "react";
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  LogoutOutlined,
} from "@ant-design/icons";
import {
  Button,
  Space,
  Avatar,
  Flex,
  Layout,
  theme,
  Popover,
  message,
  Skeleton,
} from "antd";
import { useNavigate } from "react-router-dom";
import ApiClient from "../../../services/client";
import CONSTANTS from "../../../constants";

const { Header } = Layout;

const AdminHeader = ({ collapsed, setCollapsed }) => {
  const {
    token: { colorBgContainer },
  } = theme.useToken();

  const navigate = useNavigate();

  const menus = (
    <Space direction="vertical">
      <Button
        type="text"
        icon={<LogoutOutlined />}
        onClick={() => {
          localStorage.removeItem(CONSTANTS.STORAGE_KEY_JWT);
          navigate(
            `/login?redirect_uri=${encodeURIComponent(
              window.location.pathname
            )}`
          );
        }}
      >
        退出登录
      </Button>
    </Space>
  );

  const [loading, setLoading] = useState(true);
  const [userinfo, setUserinfo] = useState(null);
  const [messageApi, contextHolder] = message.useMessage();

  useEffect(() => {
    const getUserInfo = async () => {
      ApiClient.get("/user/current")
        .then((response) => {
          // console.log("/user/current", response.data);
          if (response.data?.code === 0) {
            setLoading(false);
            setUserinfo(response.data?.data);
          } else {
            messageApi.error(response.data?.message);
          }
        })
        .catch((error) => {
          console.log(error);
          messageApi.error("获取用户信息失败，请稍后重试！");
        });
    };
    getUserInfo();
  }, [messageApi]);

  return (
    <>
      <Header style={{ padding: "0 24px 0 0", background: colorBgContainer }}>
        <Button
          type="text"
          icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
          onClick={() => setCollapsed(!collapsed)}
          style={{
            fontSize: "16px",
            width: 64,
            height: 64,
          }}
        />
        <Popover placement="bottom" content={menus}>
          <Button type="text" style={{ height: 64, float: "right" }}>
            <Space wrap size={16}>
              {loading ? (
                <Skeleton.Avatar active size="large" loading={loading} />
              ) : (
                <Avatar size="large" src={userinfo?.avatar} />
              )}
              <Flex
                vertical
                gap={2}
                style={{ lineHeight: "20px", textAlign: "left" }}
              >
                {loading ? (
                  <Skeleton.Button active size="small" />
                ) : (
                  <strong>{userinfo?.nickname}</strong>
                )}

                {loading ? (
                  <Skeleton.Input active size="small" />
                ) : (
                  <span>{userinfo?.email}</span>
                )}
              </Flex>
            </Space>
          </Button>
        </Popover>
      </Header>
      {contextHolder}
    </>
  );
};

export default AdminHeader;
