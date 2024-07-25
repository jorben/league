import React from "react";
import { useNavigate } from "react-router-dom";
import {
  Row,
  Col,
  Space,
  Button,
  message,
  Table,
  Divider,
  Form,
  Input,
  Tag,
  Select,
  Popconfirm,
  Modal,
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
  const [isModalOpen, setIsModalOpen] = React.useState(false);
  const [groupList, setGroupList] = React.useState([]);

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

    const getGroups = async () => {
      ApiClient.get("/admin/group/list")
        .then((response) => {
          if (response.data?.code === 0) {
            const list = Object.keys(response.data?.data);
            setGroupList([...list, "member", "anyone"]);
          } else {
            messageApi.error(response.data?.message);
          }
        })
        .catch((error) => {
          console.log(error);
          messageApi.error("获取用户组失败，请稍后重试！");
        });
    };
    getGroups();
    getPolicys();
  }, [isLoading, messageApi, navigate]);

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
            <Form.Item
              name="subject"
              style={{
                margin: 0,
              }}
              hidden={true}
            >
              <Input value={record?.subject} />
            </Form.Item>
            <span>{text}</span>
          </>
        ) : (
          <span>{text}</span>
        );
      },
    },
    {
      title: "请求方法",
      dataIndex: "method",
      width: 180,
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
      title: "接口",
      dataIndex: "path",
      width: 400,
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
            <span>{record?.path_name || "未命名接口"}</span>
            <span>Path: {text}</span>
          </Space>
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
                  label: "允许",
                },
                {
                  value: "deny",
                  label: "拒绝",
                },
              ]}
            />
          </Form.Item>
        ) : text === "allow" ? (
          <Tag color="green">允许</Tag>
        ) : (
          <Tag color="volcano">拒绝</Tag>
        );
      },
    },
    {
      title: "备注",
      dataIndex: "comment",
      render: (text, record) => {
        return isEditing(record) ? (
          <Form.Item
            name="comment"
            style={{
              margin: 0,
            }}
          >
            <Input />
          </Form.Item>
        ) : (
          <span>{text}</span>
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
              title="确定删除该规则吗？"
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
    ApiClient.post("/admin/auth/policy/delete", data)
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
    editForm
      .validateFields()
      .then((row) => {
        ApiClient.post("/admin/auth/policy", row)
          .then((response) => {
            if (response.data?.code === 0) {
              messageApi.success("规则更新成功");
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
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const handleModalOk = async () => {
    newForm
      .validateFields()
      .then((row) => {
        ApiClient.post("/admin/auth/policy", row)
          .then((response) => {
            if (response.data?.code === 0) {
              messageApi.success("新增规则成功");
              setIsModalOpen(false);
              setIsLoading(true);
            } else {
              messageApi.error(response.data?.message);
            }
          })
          .catch((error) => {
            console.log(error);
            messageApi.error("新增规则失败，请稍后重试！");
          });
      })
      .catch((info) => {
        console.log("Validate Failed:", info);
      });
  };

  const FormatPlicy = ({ group, rules }) => {
    const rulesWithKey = rules.map((r) => ({ ...r, key: r.ID }));
    return (
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
              dataSource={rulesWithKey}
              loading={isLoading}
              pagination={false}
              //   sticky={true}
            />
          </Col>
        </Row>
      </>
    );
  };

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
            onClick={() => setIsModalOpen(true)}
          >
            添加规则
          </Button>
        </Col>
      </Row>
      <Form form={editForm}>{PolicyElement}</Form>
      <Modal
        title="添加规则"
        open={isModalOpen}
        onOk={handleModalOk}
        onCancel={() => setIsModalOpen(false)}
        okText="确认添加"
        cancelText="取消"
      >
        <Form
          form={newForm}
          labelCol={{
            span: 6,
          }}
          wrapperCol={{
            span: 16,
          }}
          style={{ margin: "40px 0" }}
        >
          <Form.Item
            label="所属分组"
            name="subject"
            rules={[{ required: true, message: "请选择规则所属分组" }]}
          >
            <Select
              options={(groupList || []).map((g) => ({
                value: g,
                label: g,
              }))}
            />
          </Form.Item>
          <Form.Item
            label="请求方法"
            name="method"
            rules={[{ required: true, message: "请输入接口请求方法" }]}
          >
            <Input placeholder="GET | POST" />
          </Form.Item>
          <Form.Item
            label="接口地址"
            name="path"
            rules={[{ required: true, message: "请输入接口地址" }]}
          >
            <Input placeholder="请输入接口地址" />
          </Form.Item>
          <Form.Item
            label="策略"
            name="result"
            rules={[{ required: true, message: "请选择规则策略" }]}
          >
            <Select
              options={[
                {
                  value: "allow",
                  label: "允许",
                },
                {
                  value: "deny",
                  label: "拒绝",
                },
              ]}
            />
          </Form.Item>
          <Form.Item label="备注" name="comment">
            <Input placeholder="请输入备注" />
          </Form.Item>
        </Form>
      </Modal>
      {contextHolder}
    </>
  );
};

export default PermissionApi;
