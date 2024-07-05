import { Space, Descriptions, Button, Divider, Table, Avatar } from "antd";
import { GoogleOutlined } from "@ant-design/icons";
import React from "react";

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
      title: "头像",
      dataIndex: "avatar",
      key: "avatar",
      width: 60,
      onCell: () => {
        return {
          style: {
            textAlign: "center",
          },
        };
      },
      render: (url) => <Avatar size="large" src={url} />,
    },
    {
      title: "OpenId",
      dataIndex: "open_id",
      key: "open_id",
      width: 200,
      ellipsis: true,
    },
    {
      title: "邮箱",
      dataIndex: "email",
      key: "email",
      width: 200,
      ellipsis: true,
    },
    {
      title: "操作",
      dataIndex: "action",
      key: "action",
      width: 90,
      render: (text) => <a href="/#">解绑</a>,
    },
    // {
    //   title: "创建时间",
    //   dataIndex: "created_at",
    //   key: "created_at",
    //   ellipsis: true,
    // },
    // {
    //   title: "更新时间",
    //   dataIndex: "updated_at",
    //   key: "updated_at",
    //   ellipsis: true,
    // },
  ];

  const sourceData = [];
  for (let i = 1; i <= 3; i++) {
    sourceData.push({
      key: i,
      source: "google",
      email: `jorbenzhu+${i}@gmail.com`,
      avatar: "https://avatars.githubusercontent.com/u/2806170?v=4",
      open_id: "101146533148280428613",
    });
  }
  return (
    <Space direction="vertical">
      <Divider style={{ margin: "0 0 8px 0" }} />
      <Descriptions
        title="基本信息"
        column={2}
        // layout="vertical"
        bordered
        items={baseDetail}
        extra={<Button type="primary">Edit</Button>}
      />
      <Divider />
      <Table
        columns={sourceColumns}
        dataSource={sourceData}
        pagination={false}
      />
    </Space>
  );
};

export default UserDetail;
