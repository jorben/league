import React from "react";
import { Row, Col, Space, Button } from "antd";
import { PlusOutlined, MenuOutlined } from "@ant-design/icons";

const SettingMenu = () => {
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
      <Row></Row>
    </>
  );
};

export default SettingMenu;
