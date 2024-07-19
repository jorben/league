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
  Form,
} from "antd";
import { UsergroupAddOutlined, PlusOutlined } from "@ant-design/icons";
import { useNavigate } from "react-router-dom";
import ApiClient from "../../../services/client";
import CONSTANTS from "../../../constants";

const UserGroup = () => {
  const [isLoading, setIsLoading] = React.useState(true);
  const [isModalOpen, setIsModalOpen] = React.useState(false);
  const [groupList, setGroupList] = React.useState(null);
  const [memberCount, setMemberCount] = React.useState(0);
  const [form] = Form.useForm();

  const navigate = useNavigate();
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

    const getMembers = async () => {
      ApiClient.get("/admin/user/list")
        .then((response) => {
          if (response.data?.code === 0) {
            setMemberCount(response.data?.data?.Count || 0);
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

    getGroups();
    getMembers();
  }, [isLoading, messageApi, navigate]);

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
    form
      .validateFields()
      .then((formValue) => {
        const data = {
          id: formValue?.userId ? Number(formValue.userId) : 0,
          group: formValue?.groupName,
          new: 1,
        };
        ApiClient.post("/admin/user/join_group", data)
          .then((response) => {
            if (response.data?.code === 0) {
              messageApi.success("添加分组成功");
              setIsLoading(true);
            } else {
              messageApi.error(response.data?.message);
            }
          })
          .catch((error) => {
            console.log(error);
            messageApi.error("请求失败，请稍后重试！");
          });
        setIsModalOpen(false);
        form.resetFields();
      })
      .catch((info) => {
        console.log("Validate Failed:", info);
      });
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
      <Row style={{ marginBottom: "20px" }}>
        <Col span={24}>
          <Space split={<Divider type="vertical" />}>
            {Array.isArray(users) &&
              users.length > 0 &&
              users.map((u) => (
                <Space key={u.ID}>
                  <Avatar size="large" src={u?.avatar} />
                  <Space direction="vertical" size="4">
                    <span>用户ID: {u?.ID}</span>
                    <span>昵称: {u?.nickname}</span>
                  </Space>
                </Space>
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
            <span>共计 {memberCount} 位注册用户</span>
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
        <Form
          form={form}
          labelCol={{
            span: 6,
          }}
          wrapperCol={{
            span: 16,
          }}
          style={{ margin: "40px 0" }}
        >
          <Form.Item
            label="分组名称"
            name="groupName"
            rules={[
              { required: true, message: "请输入分组名称" },
              { pattern: /^[A-Za-z]+$/, message: "分组名称只能包含英文字母" },
            ]}
          >
            <Input placeholder="请输入新的分组名称" />
          </Form.Item>
          <Form.Item
            label="用户ID"
            name="userId"
            rules={[
              { required: true, message: "请输入一个添加到该分组的用户ID" },
              { pattern: /^[1-9][0-9]*$/, message: "用户ID只能是数字" },
            ]}
          >
            <Input placeholder="请输入一个需要添加到该分组的用户ID" />
          </Form.Item>
        </Form>
      </Modal>
      {contextHolder}
    </>
  );
};

export default UserGroup;
