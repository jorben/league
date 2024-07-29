import React from "react";
import { useNavigate } from "react-router-dom";
import {
  Row,
  Col,
  Space,
  Button,
  Tabs,
  Modal,
  Form,
  Input,
  InputNumber,
  Select,
  message,
} from "antd";
import { PlusOutlined, MenuOutlined, NumberOutlined } from "@ant-design/icons";
import ApiClient from "../../../services/client";
import CONSTANTS from "../../../constants";
import SettingMenuTable from "./SettingMenuTable";

const SettingMenu = () => {
  const [isLoading, setIsLoading] = React.useState(true);
  const [isModalOpen, setIsModalOpen] = React.useState(false);
  const [tabItems, setTabItems] = React.useState([]);
  const [menuTypes, setMenuTypes] = React.useState([]);
  const [currentTab, setCurrentTab] = React.useState("admin");
  const [messageApi, contextHolder] = message.useMessage();
  const [newForm] = Form.useForm();
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
            setMenuTypes(Object.keys(response.data?.data));
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

  const handleModalOk = async () => {
    newForm
      .validateFields()
      .then((row) => {
        ApiClient.post("/admin/setting/menu", row)
          .then((response) => {
            if (response.data?.code === 0) {
              messageApi.success("新增菜单项成功");
              setIsModalOpen(false);
              setIsLoading(true);
            } else {
              messageApi.error(response.data?.message);
            }
          })
          .catch((error) => {
            console.log(error);
            messageApi.error("新增接口失败，请稍后重试！");
          });
      })
      .catch((info) => {
        console.log("Validate Failed:", info);
      });
  };

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
            onClick={() => setIsModalOpen(true)}
          >
            添加菜单
          </Button>
        </Col>
      </Row>
      <Row>
        <Col span={24}>
          <Tabs
            defaultActiveKey={currentTab}
            items={tabItems}
            tabPosition="left"
            type="line"
            onChange={(key) => {
              setCurrentTab(key);
              // console.log("change to:", currentTab);
            }}
          />
        </Col>
      </Row>
      <Modal
        title="添加菜单项"
        open={isModalOpen}
        onOk={handleModalOk}
        onCancel={() => setIsModalOpen(false)}
        okText="确认添加"
        cancelText="取消"
      >
        <Form
          form={newForm}
          labelCol={{
            span: 6,
          }}
          wrapperCol={{
            span: 16,
          }}
          style={{ margin: "40px 0" }}
          initialValues={{ type: currentTab, order: 1 }}
        >
          <Form.Item
            label="所属菜单"
            name="type"
            rules={[{ required: true, message: "请选择所属菜单" }]}
          >
            <Select
              options={menuTypes.map((mt) => ({
                value: mt,
                label:
                  mt === "admin"
                    ? "后台菜单"
                    : mt === "index"
                    ? "前台菜单"
                    : mt,
              }))}
              defaultValue={currentTab}
            />
          </Form.Item>
          <Form.Item
            label="菜单名称"
            name="label"
            rules={[{ required: true, message: "请输入菜单名称" }]}
          >
            <Input placeholder="请输入菜单名称" />
          </Form.Item>
          <Form.Item
            label="菜单Path"
            name="key"
            rules={[{ required: true, message: "请输入菜单Path" }]}
          >
            <Input placeholder="请输入菜单Path" />
          </Form.Item>
          <Form.Item
            label="菜单排序"
            name="order"
            rules={[
              {
                required: true,
                message: "请输入菜单排序",
              },
              { pattern: /^[0-9]+$/, message: "排序只能是数字" },
            ]}
          >
            <InputNumber defaultValue={1} min={0} />
          </Form.Item>
          <Form.Item label="父节点Path" name="parent">
            <Input placeholder="一级菜单请留空" />
          </Form.Item>
          <Form.Item label="菜单ICON" name="icon">
            <Input placeholder="请输入ICON名称" />
          </Form.Item>
        </Form>
      </Modal>
      {contextHolder}
    </>
  );
};

export default SettingMenu;
