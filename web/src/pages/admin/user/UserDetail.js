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
  message,
} from "antd";
import {
  QuestionCircleOutlined,
  DeleteOutlined,
  DisconnectOutlined,
} from "@ant-design/icons";
import React from "react";
import UserDetailGroup from "./UserDetailGroup";
import CONSTANTS from "../../../constants";
import ApiClient from "../../../services/client";
import BrandIcon from "../../../components/BrandIcon";
import moment from "moment";
import { useNavigate } from "react-router-dom";

const UserDetail = ({ user }) => {
  const [messageApi, contextHolder] = message.useMessage();
  const [reload, setReload] = React.useState(false);
  const [userDetail, setUserDetail] = React.useState({
    base: [],
    group: [],
    source: [],
  });
  const navigate = useNavigate();
  const formatBaseInfo = (userDetail) => [
    {
      key: "id",
      label: "用户ID",
      children: userDetail?.ID,
    },
    {
      key: "nickname",
      label: "头像昵称",
      children: (
        <Space>
          <Avatar src={userDetail?.avatar} size="large" />
          <span>{userDetail?.nickname}</span>
        </Space>
      ),
    },
    {
      key: "email",
      label: "邮箱",
      children: userDetail?.email,
    },
    {
      key: "phone",
      label: "手机",
      children: userDetail?.phone,
    },
    {
      key: "bio",
      label: "简介",
      children: userDetail?.bio,
      span: 2,
    },
    {
      key: "created_at",
      label: "创建时间",
      children: moment(userDetail?.CreatedAt).format("YYYY-MM-DD HH:mm:ss"),
    },
    {
      key: "updated_at",
      label: "更新时间",
      children: moment(userDetail?.UpdatedAt).format("YYYY-MM-DD HH:mm:ss"),
    },
  ];

  const unbindSource = (source) => {
    const data = { id: userDetail?.ID, source: source };
    ApiClient.post("/admin/user/source", data)
      .then((response) => {
        if (response.data?.code === 0) {
          setReload(true);
          messageApi.success("解绑成功");
        } else {
          messageApi.error(response.data?.message);
        }
      })
      .catch((error) => {
        console.log(error);
        messageApi.error("请求失败，请稍后重试！");
      });
  };

  const updateStatus = (status) => {
    const data = { id: userDetail?.ID, status: status };
    ApiClient.post("/admin/user/status", data)
      .then((response) => {
        if (response.data?.code === 0) {
          setUserDetail({ ...userDetail, status: status });
          messageApi.success("状态更新成功");
        } else {
          messageApi.error(response.data?.message);
        }
      })
      .catch((error) => {
        console.log(error);
        messageApi.error("请求失败，请稍后重试！");
      });
  };

  const formatUserDetail = React.useCallback((user) => {
    return {
      ...user,
      base: formatBaseInfo(user),
      source: user?.source
        ? user?.source.map((s, i) => ({ ...s, key: i }))
        : [],
    };
  }, []);

  React.useEffect(() => {
    const getUserDetail = async (id) => {
      if (!id) {
        messageApi.error("缺少用户ID，请刷新后重试");
      } else {
        ApiClient.get(`admin/user/detail?id=${id}`)
          .then((response) => {
            // console.log(response.data);
            if (response.data?.code === 0) {
              setUserDetail(formatUserDetail(response.data?.data));
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
      }
      setReload(false);
    };
    if (!reload) {
      // 首次加载，从上游组件获取信息
      setUserDetail(formatUserDetail(user));
    } else {
      getUserDetail(user?.ID);
    }
  }, [user, reload, messageApi, navigate, formatUserDetail]);

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
      render: (source) => <BrandIcon name={source} size={30} />,
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
          <Avatar size="large" src={record?.avatar} />
          <span>{record?.nickname}</span>
        </Space>
      ),
    },
    {
      title: "OpenId",
      dataIndex: "open_id",
      key: "open_id",
      ellipsis: true,
    },
    {
      title: "邮箱",
      dataIndex: "email",
      key: "email",
      ellipsis: true,
    },
    {
      title: "创建时间",
      dataIndex: "CreatedAt",
      key: "created_at",
      ellipsis: true,
      render: (t) => (t ? moment(t).format("YYYY-MM-DD HH:mm:ss") : "-"),
    },
    {
      title: "更新时间",
      dataIndex: "UpdatedAt",
      key: "updated_at",
      ellipsis: true,
      render: (t) => (t ? moment(t).format("YYYY-MM-DD HH:mm:ss") : "-"),
    },
    {
      title: "操作",
      dataIndex: "action",
      key: "action",
      width: 110,
      fixed: "right",
      render: (_, record) => (
        <Button
          type="text"
          onClick={() => {
            unbindSource(record?.source);
          }}
          icon={<DisconnectOutlined />}
          danger
        >
          解绑
        </Button>
      ),
    },
  ];

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
        items={userDetail?.base}
      />
      <Divider orientation="left">登录来源</Divider>
      <Table
        columns={sourceColumns}
        dataSource={userDetail?.source}
        pagination={false}
        bordered
        scroll={{
          x: "150%",
        }}
      />
      <Divider orientation="left">关联角色</Divider>
      <UserDetailGroup group={userDetail?.group} />
      <Divider orientation="left">账户状态</Divider>
      <Row>
        <Col span={12}>
          <Space>
            <span>禁用用户：</span>
            <Switch
              checked={userDetail?.status === 1}
              checkedChildren="开启"
              unCheckedChildren="关闭"
              onChange={(checked) => {
                updateStatus(checked ? 1 : 0);
              }}
            />
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
      {contextHolder}
    </Space>
  );
};

export default UserDetail;
