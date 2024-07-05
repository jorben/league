import React from "react";
import { Col, Row, Space, Table, Tooltip, Tag } from "antd";
import Search from "antd/es/input/Search";
import {
  UnorderedListOutlined,
  DoubleRightOutlined,
  PicLeftOutlined,
} from "@ant-design/icons";
import UserDetail from "./UserDetail";

const UserList = () => {
  const columns = [
    {
      title: "用户ID",
      dataIndex: "id",
      key: "id",
      width: 80,
      fixed: "left",
    },
    {
      title: "用户昵称",
      dataIndex: "nickname",
      key: "nickname",
      width: 180,
      fixed: "left",
    },
    {
      title: "邮箱",
      dataIndex: "email",
      key: "email",
      width: 220,
      ellipsis: true,
    },
    {
      title: "关联角色",
      dataIndex: "group",
      key: "group",
      width: 90,
      render: (text) => <Tag>Num: {text}</Tag>,
    },
    {
      title: "创建时间",
      dataIndex: "created_at",
      key: "created_at",
      ellipsis: true,
    },
    {
      title: "更新时间",
      dataIndex: "updated_at",
      key: "updated_at",
      ellipsis: true,
    },
  ];

  const data = [];
  for (let i = 1; i <= 80; i++) {
    data.push({
      key: i,
      id: i,
      nickname: `Edward King ${i}`,
      email: `jorbenzhu+${i}@gmail.com`,
      group: 1,
      created_at: "2024-07-03 13:08:53",
      updated_at: "2024-07-03 13:08:53",
    });
  }

  return (
    <Row>
      <Col span={14}>
        <Row>
          <Col span={12}>
            <Space>
              <UnorderedListOutlined />
              <h3>用户列表</h3>
            </Space>
          </Col>
          <Col span={12} style={{ alignContent: "center" }}>
            <Search
              placeholder="请输入用户id/昵称/email"
              allowClear
              // onSearch={onSearch}
              style={{
                width: 280,
                float: "right",
              }}
            />
          </Col>
        </Row>

        <Row>
          <Table
            columns={columns}
            dataSource={data}
            scroll={{
              y: 690,
            }}
            pagination={{
              pageSize: 20,
              simple: true,
              showSizeChanger: false,
              hideOnSinglePage: true,
            }}
            onRow={(r) => {
              return {
                onClick: (e) => {
                  console.log(r);
                  // console.log(e);
                },
                style: {
                  cursor: "pointer",
                },
              };
            }}
          />
        </Row>
      </Col>
      <Col span={1} style={{ alignContent: "center", textAlign: "center" }}>
        <Tooltip title="点击左侧数据行查看详情">
          <DoubleRightOutlined />
        </Tooltip>
      </Col>
      <Col span={9}>
        <Row>
          <Space>
            <PicLeftOutlined />
            <h3>用户详情</h3>
          </Space>
        </Row>
        <Row>
          <UserDetail />
        </Row>
      </Col>
    </Row>
  );
};

export default UserList;
