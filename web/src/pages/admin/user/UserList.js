import React from "react";
import { Col, Row, Space, Table, Badge, Tag, Drawer, Button } from "antd";
import Search from "antd/es/input/Search";
import {
  UnorderedListOutlined,
  CloseOutlined,
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
      title: "电话",
      dataIndex: "phone",
      key: "phone",
      width: 200,
      ellipsis: true,
    },
    {
      title: "绑定来源",
      dataIndex: "source",
      key: "source",
      width: 90,
      render: (text) => <Tag>Num: {text.length}</Tag>,
    },
    {
      title: "关联角色",
      dataIndex: "group",
      key: "group",
      width: 90,
      render: (text) => <Tag>Num: {text}</Tag>,
    },
    {
      title: "账号状态",
      dataIndex: "status",
      key: "status",
      width: 90,
      render: (status) => <Badge status="success" text="正常" />,
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
      source: ["google", "github"],
      group: 1,
      status: 1,
      created_at: "2024-07-03 13:08:53",
      updated_at: "2024-07-03 13:08:53",
    });
  }

  const [open, setOpen] = React.useState(false);
  const [loading, setLoading] = React.useState(true);

  const showLoading = () => {
    setOpen(true);
    setLoading(true);

    // Simple loading mock. You should add cleanup logic in real world.
    setTimeout(() => {
      setLoading(false);
    }, 400);
  };

  return (
    <>
      <Row>
        <Col span={24}>
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
            <Col>
              <Table
                columns={columns}
                dataSource={data}
                scroll={{
                  y: window.innerHeight - 370,
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
                      showLoading();
                      // console.log(e);
                    },
                    style: {
                      cursor: "pointer",
                    },
                  };
                }}
              />
            </Col>
          </Row>
        </Col>
      </Row>
      <Drawer
        closable
        destroyOnClose
        title={<span>用户详情</span>}
        placement="right"
        size="large"
        open={open}
        loading={loading}
        closeIcon={<PicLeftOutlined />}
        extra={
          <div>
            <Button
              type="text"
              icon={<CloseOutlined />}
              onClick={() => setOpen(false)}
            />
          </div>
        }
        onClose={() => setOpen(false)}
      >
        <UserDetail />
      </Drawer>
    </>
  );
};

export default UserList;
