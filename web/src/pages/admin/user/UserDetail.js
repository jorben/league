import {
  Space,
  Descriptions,
  Divider,
  Table,
  Avatar,
  Switch,
  Tooltip,
  Button,
  Row,
  Col,
} from "antd";
import {
  GoogleOutlined,
  QuestionCircleOutlined,
  DeleteOutlined,
} from "@ant-design/icons";
import React from "react";
import UserDetailGroup from "./UserDetailGroup";

const UserDetail = () => {
  const baseDetail = [
    {
      key: "id",
      label: "用户ID",
      children: "1234",
    },
    {
      key: "nickname",
      label: "用户昵称",
      children: "脚本大哥",
    },
    {
      key: "email",
      label: "邮箱",
      children: "jorbenzhu@gmail.com",
    },
    {
      key: "phone",
      label: "手机",
      children: "",
    },
    {
      key: "bio",
      label: "简介",
      children:
        "这里是个人简介，很长的简介，看看超长的效果如何。这里是个人简介，很长的简介，看看超长的效果如何。这里是个人简介，很长的简介，看看超长的效果如何。",
      span: 2,
    },
    {
      key: "created_at",
      label: "创建时间",
      children: "2024-07-03 13:08:53",
    },
    {
      key: "updated_at",
      label: "更新时间",
      children: "2024-07-03 13:08:53",
    },
  ];

  const sourceColumns = [
    {
      title: "渠道",
      dataIndex: "source",
      key: "source",
      width: 60,
      fixed: "left",
      onCell: () => {
        return {
          style: {
            textAlign: "center",
          },
        };
      },
      render: (source) => <GoogleOutlined />,
    },
    {
      title: "头像昵称",
      dataIndex: "avatar",
      key: "avatar",
      width: 90,
      fixed: "left",
      onCell: () => {
        return {
          style: {
            textAlign: "center",
          },
        };
      },
      render: (_, record) => (
        <Space direction="vertical">
          <Avatar size="large" src={record.avatar} />
          <span>{record.nickname}</span>
        </Space>
      ),
    },
    {
      title: "OpenId",
      dataIndex: "open_id",
      key: "open_id",
      // width: 200,
      ellipsis: true,
    },
    {
      title: "邮箱",
      dataIndex: "email",
      key: "email",
      // width: 200,
      ellipsis: true,
    },
    {
      title: "创建时间",
      dataIndex: "created_at",
      key: "created_at",
      // width: 200,
      ellipsis: true,
    },
    {
      title: "更新时间",
      dataIndex: "updated_at",
      key: "updated_at",
      // width: 200,
      ellipsis: true,
    },
    {
      title: "操作",
      dataIndex: "action",
      key: "action",
      width: 90,
      fixed: "right",
      render: (text) => <a href="/#">解绑</a>,
    },
  ];

  const sourceData = [];
  for (let i = 1; i <= 3; i++) {
    sourceData.push({
      key: i,
      source: "google",
      email: `jorbenzhu+${i}@gmail.com`,
      avatar: "https://avatars.githubusercontent.com/u/2806170?v=4",
      open_id: "101146533148280428613",
      nickname: "渠道昵称",
      created_at: "2024-07-03 13:08:53",
      updated_at: "2024-07-03 13:08:53",
    });
  }
  return (
    <Space direction="vertical">
      <Descriptions
        title={
          <Divider style={{ marginTop: 0 }} orientation="left">
            基本信息
          </Divider>
        }
        column={2}
        bordered
        items={baseDetail}
      />
      <Divider orientation="left">登录来源</Divider>
      <Table
        columns={sourceColumns}
        dataSource={sourceData}
        pagination={false}
        bordered
        scroll={{
          x: "150%",
        }}
      />
      <Divider orientation="left">关联角色</Divider>
      <UserDetailGroup />
      <Divider orientation="left">账户状态</Divider>
      <Row>
        <Col span={12}>
          <Space>
            <span>禁用用户：</span>
            <Switch checkedChildren="开启" unCheckedChildren="关闭" />
            <Tooltip title="开启后用户将无法登录">
              <QuestionCircleOutlined />
            </Tooltip>
          </Space>
        </Col>
        <Col span={12} style={{ textAlign: "right" }}>
          <Button type="primary" danger icon={<DeleteOutlined />}>
            删除用户
          </Button>
        </Col>
      </Row>
    </Space>
  );
};

export default UserDetail;
