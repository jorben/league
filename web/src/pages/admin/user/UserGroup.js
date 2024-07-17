import React from "react";
import {
  Row,
  Col,
  Space,
  Button,
  Divider,
  Avatar,
  Spin,
  Modal,
  message,
  Input,
} from "antd";
import { UsergroupAddOutlined, PlusOutlined } from "@ant-design/icons";
import ApiClient from "../../../services/client";

const UserGroup = () => {
  const [isLoading, setIsLoading] = React.useState(true);
  const [isModalOpen, setIsModalOpen] = React.useState(false);
  const [groupList, setGroupList] = React.useState(null);

  const [messageApi, contextHolder] = message.useMessage();
  React.useEffect(() => {
    const getGroups = async () => {
      ApiClient.get("/admin/group/list")
        .then((response) => {
          if (response.data?.code === 0) {
            setGroupList(response.data?.data);
          } else {
            messageApi.error(response.data?.message);
          }
        })
        .catch((error) => {
          console.log(error);
          messageApi.error("获取用户信息失败，请稍后重试！");
        })
        .finally(() => {
          setIsLoading(false);
        });
    };
    getGroups();
  }, [messageApi]);

  const loadingStyle = {
    padding: 50,
    background: "rgba(0, 0, 0, 0.05)",
    borderRadius: 4,
  };
  const loadingElement = <div style={loadingStyle} />;

  const showModal = () => {
    setIsModalOpen(true);
  };
  const handleOk = () => {
    setIsModalOpen(false);
  };
  const handleCancel = () => {
    setIsModalOpen(false);
  };

  const FormatGroup = ({ group, users }) => (
    <>
      <Row>
        <Col span={24}>
          <Divider orientation="left">分组：{group}</Divider>
        </Col>
      </Row>
      <Row>
        <Col span={24}>
          <Space>
            {Array.isArray(users) &&
              users.length > 0 &&
              users.map((u) => (
                <Avatar
                  style={{ backgroundColor: "#fde3cf", color: "#f56a00" }}
                  key={u}
                >
                  {u}
                </Avatar>
              ))}
          </Space>
        </Col>
      </Row>
    </>
  );

  const groupElement = groupList
    ? Object.entries(groupList).map(([group, users]) => (
        <FormatGroup key={group} group={group} users={users} />
      ))
    : "";

  return (
    <>
      <Row>
        <Col span={12}>
          <Space>
            <UsergroupAddOutlined />
            <h3>用户组</h3>
          </Space>
        </Col>
        <Col span={12}>
          <Button
            type="primary"
            style={{
              float: "right",
            }}
            icon={<PlusOutlined />}
            onClick={showModal}
          >
            添加分组
          </Button>
        </Col>
      </Row>
      {isLoading ? <Spin tip="Loading">{loadingElement}</Spin> : groupElement}
      <Row>
        <Col span={24}>
          <Divider orientation="left">分组：member</Divider>
        </Col>
      </Row>
      <Row>
        <Col span={24}>
          <Space>
            <Avatar style={{ backgroundColor: "#00a2ae" }}>M</Avatar>
            <span>共计256位注册用户</span>
          </Space>
        </Col>
      </Row>
      <Modal
        title="添加分组"
        open={isModalOpen}
        onOk={handleOk}
        onCancel={handleCancel}
        okText="确认添加"
        cancelText="取消"
      >
        <Space direction="vertical">
          <Space>
            <span>组名称：</span>
            <Input placeholder="请输入分组名称，仅支持英文字母" />
          </Space>
          <Space>
            <span>用户ID：</span>
            <Input placeholder="请输入加入该分组的用户ID" />
          </Space>
        </Space>
      </Modal>
      {contextHolder}
    </>
  );
};

export default UserGroup;
