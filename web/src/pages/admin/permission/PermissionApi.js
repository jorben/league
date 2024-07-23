import React from "react";
import { Row, Col, Space, Button } from "antd";
import { PlusOutlined, ApiOutlined } from "@ant-design/icons";

const PermissionApi = () => {
  return (
    <>
      <Row>
        <Col span={12}>
          <Space>
            <ApiOutlined />
            <h3>接口权限</h3>
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
            添加规则
          </Button>
        </Col>
      </Row>
      <Row></Row>
    </>
  );
};

export default PermissionApi;
