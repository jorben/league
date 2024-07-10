import React, { useState, useEffect } from "react";
import { useNavigate, Navigate, Route, Routes } from "react-router-dom";
import { adminRoutes } from "../../../routes";
import { Layout, message, Spin, theme } from "antd";
import AdminHeader from "./AdminHeader";
import AdminFooter from "./AdminFooter";
import AdminSider from "./AdminSider";
import ApiClient from "../../../services/client";
import { LoadingOutlined } from "@ant-design/icons";
import CONSTANTS from "../../../constants";

const { Content } = Layout;

const AdminFrame = () => {
  const [collapsed, setCollapsed] = useState(false);
  const [isLoading, setLoading] = useState(true);
  const [allMenus, setAllMenus] = useState([]);
  const [messageApi, contextHolder] = message.useMessage();
  const navigate = useNavigate();
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  useEffect(() => {
    const getAdminMenus = async () => {
      ApiClient.get("/admin/menu")
        .then((response) => {
          console.log(response.data);
          if (response.data?.code === 0) {
            setLoading(false);
            setAllMenus(response.data?.data);
          } else if (
            response.data?.code === CONSTANTS.ERRCODE.ErrAuthNoLogin ||
            response.data?.code === CONSTANTS.ERRCODE.ErrAuthUnauthorized
          ) {
            messageApi.error(response.data?.message, () => {
              navigate(
                `/login?redirect_uri=${encodeURIComponent(
                  window.location.pathname
                )}`
              );
            });
          } else {
            messageApi.error(response.data?.message);
          }
        })
        .catch((error) => {
          console.log(error);
          messageApi.error("请求失败，请稍后重试！");
        });
    };
    getAdminMenus();
  }, [messageApi, navigate]);

  // console.log("In AdminFrame, menus:", allMenus);
  return (
    <Layout style={{ minHeight: "100vh" }}>
      {isLoading ? (
        <Spin
          size="large"
          indicator={<LoadingOutlined spin />}
          tip="加载中..."
          fullscreen
        />
      ) : (
        <>
          <AdminSider collapsed={collapsed} allMenus={allMenus} />
          <Layout>
            <AdminHeader collapsed={collapsed} setCollapsed={setCollapsed} />
            <Content
              style={{
                margin: "24px 16px 0 16px",
                padding: 24,
                minHeight: 280,
                // minWidth: 1640 + 48,
                background: colorBgContainer,
                borderRadius: borderRadiusLG,
              }}
            >
              <Routes>
                {adminRoutes.map((r, index) => {
                  return <Route key={index} {...r} />;
                })}
                <Route path="*" element={<Navigate to="/404" replace />} />
              </Routes>
            </Content>
            <AdminFooter />
          </Layout>
        </>
      )}
      {contextHolder}
    </Layout>
  );
};

export default AdminFrame;
