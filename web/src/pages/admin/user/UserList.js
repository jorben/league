import React from "react";
import {
  Col,
  Row,
  Space,
  Table,
  Badge,
  Tag,
  Drawer,
  Button,
  Avatar,
  message,
} from "antd";
import Search from "antd/es/input/Search";
import {
  UnorderedListOutlined,
  CloseOutlined,
  PicLeftOutlined,
} from "@ant-design/icons";
import UserDetail from "./UserDetail";
import CONSTANTS from "../../../constants";
import ApiClient from "../../../services/client";
import { useNavigate } from "react-router-dom";
import moment from "moment";
import BrandIcon from "../../../components/BrandIcon";

const UserList = () => {
  const columns = [
    {
      title: "用户ID",
      dataIndex: "ID",
      key: "id",
      width: 80,
      fixed: "left",
    },
    {
      title: "用户头像",
      dataIndex: "avatar",
      key: "avatar",
      width: 100,
      fixed: "left",
      render: (src) => <Avatar size="large" src={src} />,
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
      render: (source) => (
        <Space>
          {source &&
            source.map((s, i) => (
              <BrandIcon key={i} name={s?.source} size={24} />
            ))}
        </Space>
      ),
    },
    {
      title: "关联角色",
      dataIndex: "group",
      key: "group",
      width: 90,
      render: (text) => <Tag>Num: {text?.length ? text.length + 1 : 1}</Tag>,
    },
    {
      title: "账号状态",
      dataIndex: "status",
      key: "status",
      width: 90,
      render: (status) =>
        status ? (
          <Badge status="warning" text="禁用" />
        ) : (
          <Badge status="success" text="正常" />
        ),
    },
    {
      title: "创建时间",
      dataIndex: "CreatedAt",
      key: "created_at",
      ellipsis: true,
      render: (t) => moment(t).format("YYYY-MM-DD HH:mm:ss"),
    },
    {
      title: "更新时间",
      dataIndex: "UpdatedAt",
      key: "pdated_at",
      ellipsis: true,
      render: (t) => moment(t).format("YYYY-MM-DD HH:mm:ss"),
    },
  ];

  const navigate = useNavigate();
  const [messageApi, contextHolder] = message.useMessage();
  const [openDrawer, setOpenDrawer] = React.useState(false);
  const [loading, setLoading] = React.useState(true);
  const [showUser, setShowUser] = React.useState(null);
  const [userList, setUserList] = React.useState({ Count: 0, List: [] });
  const [searchParam, setSearchParam] = React.useState({
    key: "",
    page: 1,
    size: CONSTANTS.DEFAULT_PAGESIZE,
  });

  const showUserDetail = (user) => {
    setShowUser(user);
    setOpenDrawer(true);
  };

  React.useEffect(() => {
    const getUserList = async (searchParam) => {
      const query = new URLSearchParams({
        search: searchParam?.key || "",
        page: searchParam?.page || 1,
        size: searchParam?.size || CONSTANTS.DEFAULT_PAGESIZE,
      });
      ApiClient.get(`/admin/user/list?${query.toString()}`)
        .then((response) => {
          // console.log(response.data);
          setLoading(false);
          if (response.data?.code === 0) {
            setUserList({
              Count: response.data?.data?.Count || 0,
              List: response.data?.data?.List
                ? response.data.data.List.map((item) => ({
                    ...item,
                    key: item.ID,
                  }))
                : [],
            });
          } else if (
            response.data?.code === CONSTANTS.ERRCODE.ErrAuthNoLogin ||
            response.data?.code === CONSTANTS.ERRCODE.ErrAuthUnauthorized
          ) {
            messageApi.error(response.data?.message, () => {
              navigate("/login");
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
    setLoading(true);
    getUserList(searchParam);
  }, [messageApi, navigate, searchParam]);

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
                onSearch={(value) => {
                  setSearchParam({ ...searchParam, key: value });
                }}
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
                dataSource={userList.List}
                loading={loading}
                scroll={{
                  x: "100%",
                  y: window.innerHeight - 370,
                }}
                pagination={{
                  pageSize: searchParam.size,
                  current: searchParam.page,
                  simple: true,
                  showSizeChanger: false,
                  hideOnSinglePage: true,
                  total: userList.Count,
                  onChange: (page, size) =>
                    setSearchParam({ ...searchParam, page: page, size: size }),
                }}
                onRow={(r) => {
                  return {
                    onClick: (e) => {
                      showUserDetail(r);
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
        open={openDrawer}
        loading={loading}
        closeIcon={<PicLeftOutlined />}
        extra={
          <div>
            <Button
              type="text"
              icon={<CloseOutlined />}
              onClick={() => {
                setOpenDrawer(false);
                setSearchParam({ ...searchParam });
              }}
            />
          </div>
        }
        onClose={() => {
          setOpenDrawer(false);
          setSearchParam({ ...searchParam });
        }}
      >
        <UserDetail user={showUser} />
      </Drawer>
      {contextHolder}
    </>
  );
};

export default UserList;
