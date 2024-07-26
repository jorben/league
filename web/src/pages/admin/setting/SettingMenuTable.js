import React from "react";
import { Space, Button, Form, Input, Table, Popconfirm, message } from "antd";
import ApiClient from "../../../services/client";

const SettingMenuTable = ({ type, menus, isLoading, setIsLoading }) => {
  const [editForm] = Form.useForm();
  const [editingKey, setEditingKey] = React.useState(0);
  const isEditing = (record) => record.id === editingKey;
  const [messageApi, contextHolder] = message.useMessage();
  const columns = [
    {
      title: "菜单名称",
      dataIndex: "label",
      width: 400,
      fixed: "left",
      render: (text, record) => {
        return isEditing(record) ? (
          <>
            <Form.Item
              name="id"
              style={{
                margin: 0,
              }}
              hidden={true}
            >
              <Input value={text} />
            </Form.Item>
            <Form.Item
              name="parent"
              style={{
                margin: 0,
              }}
              hidden={true}
            >
              <Input value={record?.parent} />
            </Form.Item>
            <Form.Item
              name="label"
              style={{
                margin: 0,
              }}
              rules={[
                {
                  required: true,
                  message: "请输入菜单名称",
                },
              ]}
            >
              <Input />
            </Form.Item>
          </>
        ) : (
          <span>{text}</span>
        );
      },
    },
    {
      title: "菜单Icon",
      dataIndex: "icon",
      width: 180,
      render: (text, record) => {
        return isEditing(record) ? (
          <Form.Item
            name="icon"
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
      title: "菜单Path",
      dataIndex: "key",
      render: (text, record) => {
        return isEditing(record) ? (
          <Form.Item
            name="key"
            style={{
              margin: 0,
            }}
            rules={[
              {
                required: true,
                message: "请输入菜单Path",
              },
            ]}
          >
            <Input />
          </Form.Item>
        ) : (
          <span>{text}</span>
        );
      },
    },
    {
      title: "排序",
      dataIndex: "order",
      width: 100,
      render: (text, record) => {
        return isEditing(record) ? (
          <Form.Item
            name="order"
            style={{
              margin: 0,
            }}
            rules={[
              {
                required: true,
                message: "请输入菜单排序",
              },
              { pattern: /^[0-9]+$/, message: "排序只能是数字" },
            ]}
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
              title="确定删除该菜单吗？"
              onConfirm={() => handleDelete(record)}
            >
              <Button
                type="link"
                disabled={
                  editingKey !== 0 ||
                  (Array.isArray(record?.children) &&
                    record?.children.length > 0)
                }
                danger
              >
                删除
              </Button>
            </Popconfirm>
          </Space>
        );
      },
    },
  ];

  const handleEdit = (record) => {
    setEditingKey(record.id);
    editForm.setFieldsValue({ ...record });
  };

  const handleDelete = async (record) => {
    const data = { ID: record.ID };
    ApiClient.post("/admin/menu/delete", data)
      .then((response) => {
        if (response.data?.code === 0) {
          messageApi.success("菜单删除成功");
          setIsLoading(true);
        } else {
          messageApi.error(response.data?.message);
        }
      })
      .catch((error) => {
        console.log(error);
        messageApi.error("菜单删除失败，请稍后重试！");
      });
  };

  const handleSave = async () => {
    editForm
      .validateFields()
      .then((row) => {
        ApiClient.post("/admin/menu", row)
          .then((response) => {
            if (response.data?.code === 0) {
              messageApi.success("菜单更新成功");
              setIsLoading(true);
              setEditingKey(0);
            } else {
              messageApi.error(response.data?.message);
            }
          })
          .catch((error) => {
            console.log(error);
            messageApi.error("更新菜单失败，请稍后重试！");
          });
      })
      .catch((error) => {
        console.log(error);
      });
  };
  return (
    <>
      <Form form={editForm} component={false}>
        <Table
          columns={columns}
          dataSource={menus}
          loading={isLoading}
          pagination={false}
        />
      </Form>
      {contextHolder}
    </>
  );
};

export default SettingMenuTable;
