import React from "react";
import { useNavigate } from "react-router-dom";
import { Row, Col, Space, Button, Tabs, message } from "antd";
import { PlusOutlined, MenuOutlined, NumberOutlined } from "@ant-design/icons";
import ApiClient from "../../../services/client";
import CONSTANTS from "../../../constants";
import SettingMenuTable from "./SettingMenuTable";

const SettingMenu = () => {
  const [isLoading, setIsLoading] = React.useState(true);

  const [tabItems, setTabItems] = React.useState([]);
  const [messageApi, contextHolder] = message.useMessage();
  const navigate = useNavigate();

  React.useEffect(() => {
    const getAllMenus = async () => {
      ApiClient.get("/admin/setting/menulist")
        .then((response) => {
          if (response.data?.code === 0) {
            const data = Object.entries(response.data?.data).map(
              ([t, menus]) => ({
                key: t,
                label:
                  t === "admin" ? "后台菜单" : t === "index" ? "前台菜单" : t,
                icon: <NumberOutlined />,
                children: (
                  <SettingMenuTable
                    menus={menus}
                    isLoading={isLoading}
                    setIsLoading={setIsLoading}
                  />
                ),
              })
            );
            setTabItems(data);
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
        })
        .finally(() => {
          setIsLoading(false);
        });
    };
    getAllMenus();
  }, [isLoading, messageApi, navigate]);

  return (
    <>
      <Row>
        <Col span={12}>
          <Space>
            <MenuOutlined />
            <h3>菜单</h3>
          </Space>
        </Col>
        <Col span={12}>
          <Button
            type="primary"
            style={{
              float: "right",
            }}
            icon={<PlusOutlined />}
            // onClick={() => setIsModalOpen(true)}
          >
            添加菜单
          </Button>
        </Col>
      </Row>
      <Row>
        <Col span={24}>
          <Tabs items={tabItems} tabPosition="left" type="line" />
        </Col>
      </Row>
      {contextHolder}
    </>
  );
};

export default SettingMenu;
