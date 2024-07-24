import React from "react";
import { useNavigate } from "react-router-dom";
import {
  Row,
  Col,
  Space,
  Button,
  message,
  Table,
  Spin,
  Divider,
  Form,
  Input,
  Tag,
  Select,
  Popconfirm,
} from "antd";
import { PlusOutlined, ApiOutlined } from "@ant-design/icons";
import ApiClient from "../../../services/client";
import CONSTANTS from "../../../constants";

const PermissionApi = () => {
  const [editForm] = Form.useForm();
  const [newForm] = Form.useForm();
  const [isLoading, setIsLoading] = React.useState(true);
  const [policyList, setPolicyList] = React.useState(null);
  const [editingKey, setEditingKey] = React.useState(0);

  const navigate = useNavigate();
  const [messageApi, contextHolder] = message.useMessage();
  const isEditing = (record) => record.ID === editingKey;
  React.useEffect(() => {
    const getPolicys = async () => {
      ApiClient.get("/admin/auth/policylist")
        .then((response) => {
          if (response.data?.code === 0) {
            setPolicyList(response.data?.data);
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
          messageApi.error("获取权限规则失败，请稍后重试！");
        })
        .finally(() => {
          setIsLoading(false);
        });
    };

    getPolicys();
  }, [isLoading, messageApi, navigate]);

  const loadingStyle = {
    padding: 50,
    background: "rgba(0, 0, 0, 0.05)",
    borderRadius: 4,
  };
  const loadingElement = <div style={loadingStyle} />;

  const colPolicy = [
    {
      title: "ID",
      dataIndex: "ID",
      width: 80,
      fixed: "left",
      render: (text, record) => {
        return isEditing(record) ? (
          <>
            <Form.Item
              name="ID"
              style={{
                margin: 0,
              }}
              hidden={true}
            >
              <Input value={text} />
            </Form.Item>
            <span>{text}</span>
          </>
        ) : (
          <span>{text}</span>
        );
      },
    },
    {
      title: "接口",
      dataIndex: "path",
      width: 280,
      fixed: "left",
      render: (text, record) => {
        return isEditing(record) ? (
          <Form.Item
            name="path"
            style={{
              margin: 0,
            }}
            rules={[
              {
                required: true,
                message: "请输入接口地址",
              },
            ]}
          >
            <Input />
          </Form.Item>
        ) : (
          <Space direction="vertical" size="2">
            <span>{record?.path_name || "未定义接口"}</span>
            <span>Path: {text}</span>
          </Space>
        );
      },
    },
    {
      title: "请求方法",
      dataIndex: "method",
      width: 320,
      render: (text, record) => {
        return isEditing(record) ? (
          <Form.Item
            name="method"
            style={{
              margin: 0,
            }}
            rules={[
              {
                required: true,
                message: "请输入请求方法",
              },
            ]}
          >
            <Input />
          </Form.Item>
        ) : (
          <Tag>{text}</Tag>
        );
      },
    },
    {
      title: "策略",
      dataIndex: "result",
      width: 140,
      render: (text, record) => {
        return isEditing(record) ? (
          <Form.Item
            name="result"
            style={{
              margin: 0,
            }}
            rules={[
              {
                required: true,
                message: "请选择权限策略",
              },
            ]}
          >
            <Select
              options={[
                {
                  value: "allow",
                  label: "allow",
                },
                {
                  value: "deny",
                  label: "deny",
                },
              ]}
            />
          </Form.Item>
        ) : (
          <Tag>{text}</Tag>
        );
      },
    },
    {
      title: "操作",
      dataIndex: "operation",
      width: 180,
      fixed: "right",
      render: (_, record) => {
        return isEditing(record) ? (
          <Space>
            <Button type="link" onClick={() => handleSave(record)}>
              保存
            </Button>
            <Button type="link" onClick={() => setEditingKey(0)}>
              取消
            </Button>
          </Space>
        ) : (
          <Space>
            <Button
              type="link"
              disabled={editingKey !== 0}
              onClick={() => handleEdit(record)}
            >
              编辑
            </Button>
            <Popconfirm
              title="确定删除该接口吗？"
              onConfirm={() => handleDelete(record)}
            >
              <Button type="link" disabled={editingKey !== 0} danger>
                删除
              </Button>
            </Popconfirm>
          </Space>
        );
      },
    },
  ];

  const handleEdit = (record) => {
    setEditingKey(record.ID);
    editForm.setFieldsValue({ ...record });
  };

  const handleDelete = async (record) => {
    const data = { ID: record.ID };
    ApiClient.post("/admin/setting/api/delete", data)
      .then((response) => {
        if (response.data?.code === 0) {
          messageApi.success("规则删除成功");
          setIsLoading(true);
        } else {
          messageApi.error(response.data?.message);
        }
      })
      .catch((error) => {
        console.log(error);
        messageApi.error("接口删除失败，请稍后重试！");
      });
  };

  const handleSave = async () => {
    const row = await editForm.validateFields();
    ApiClient.post("/admin/setting/api", row)
      .then((response) => {
        if (response.data?.code === 0) {
          messageApi.success("接口信息更新成功");
          setIsLoading(true);
          setEditingKey(0);
        } else {
          messageApi.error(response.data?.message);
        }
      })
      .catch((error) => {
        console.log(error);
        messageApi.error("更新权限规则失败，请稍后重试！");
      });
  };

  const FormatPlicy = ({ group, rules }) => (
    <>
      <Row>
        <Col span={24}>
          <Divider orientation="left">分组：{group}</Divider>
        </Col>
      </Row>
      <Row style={{ marginBottom: "20px" }}>
        <Col span={24}>
          <Table
            columns={colPolicy}
            dataSource={rules}
            loading={isLoading}
            pagination={false}
          />
        </Col>
      </Row>
    </>
  );

  const PolicyElement = policyList
    ? Object.entries(policyList).map(([group, rules]) => (
        <FormatPlicy key={group} group={group} rules={rules} />
      ))
    : "";

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
      {isLoading ? (
        <Row>
          <Spin tip="Loading">{loadingElement}</Spin>
        </Row>
      ) : (
        <Form form={editForm}>{PolicyElement}</Form>
      )}
      {contextHolder}
    </>
  );
};

export default PermissionApi;
