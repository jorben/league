import React from "react";
import {
  Space,
  Row,
  Col,
  Button,
  Input,
  Table,
  Form,
  Select,
  Tag,
  Popconfirm,
  Modal,
  message,
} from "antd";
import { ApiOutlined, PlusOutlined } from "@ant-design/icons";
import ApiClient from "../../../services/client";
import CONSTANTS from "../../../constants";

const SettingApi = () => {
  const [editForm] = Form.useForm();
  const [newForm] = Form.useForm();
  const [messageApi, contextHolder] = message.useMessage();
  const [editingKey, setEditingKey] = React.useState(0);
  const [isLoading, setIsLoading] = React.useState(true);
  const [isModalOpen, setIsModalOpen] = React.useState(false);
  const [dataSource, setDataSource] = React.useState([]);
  const isEditing = (record) => record.ID === editingKey;
  const [searchParam, setSearchParam] = React.useState({
    page: 1,
    size: CONSTANTS.DEFAULT_PAGESIZE,
  });

  React.useEffect(() => {
    const getApiList = async (searchParam) => {
      const query = new URLSearchParams({
        page: searchParam?.page || 1,
        size: searchParam?.size || CONSTANTS.DEFAULT_PAGESIZE,
      });
      setIsLoading(true);
      ApiClient.get(`/admin/setting/apilist?${query.toString()}`)
        .then((response) => {
          if (response.data?.code === 0) {
            setDataSource({
              Count: response.data?.data?.Count || 0,
              List: response.data?.data?.List
                ? response.data.data.List.map((item) => ({
                    ...item,
                    key: item.ID,
                  }))
                : [],
            });
          } else {
            messageApi.error(response.data?.message);
          }
        })
        .catch((error) => {
          console.log(error);
          messageApi.error("获取接口信息失败，请稍后重试！");
        })
        .finally(() => {
          setIsLoading(false);
        });
    };
    getApiList(searchParam);
  }, [searchParam, messageApi]);

  const columns = [
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
      title: "接口名称",
      dataIndex: "name",
      width: 280,
      fixed: "left",
      render: (text, record) => {
        return isEditing(record) ? (
          <Form.Item
            name="name"
            style={{
              margin: 0,
            }}
            rules={[
              {
                required: true,
                message: "请输入接口名称",
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
      title: "接口地址",
      dataIndex: "path",
      width: 320,
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
          <span>{text}</span>
        );
      },
    },
    {
      title: "请求方式",
      dataIndex: "method",
      width: 140,
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
                message: "请选择接口请求方式",
              },
            ]}
          >
            <Select
              options={[
                {
                  value: "GET",
                  label: "GET",
                },
                {
                  value: "PUT",
                  label: "PUT",
                },
                {
                  value: "POST",
                  label: "POST",
                },
                {
                  value: "DELETE",
                  label: "DELETE",
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
          messageApi.success("接口删除成功");
          setSearchParam({ ...searchParam });
        } else {
          messageApi.error(response.data?.message);
        }
      })
      .catch((error) => {
        console.log(error);
        messageApi.error("接口删除失败，请稍后重试！");
      });
  };

  const handleModalOk = async () => {
    newForm
      .validateFields()
      .then((row) => {
        ApiClient.post("/admin/setting/api", row)
          .then((response) => {
            if (response.data?.code === 0) {
              messageApi.success("新增接口成功");
              setIsModalOpen(false);
              setSearchParam({ ...searchParam });
            } else {
              messageApi.error(response.data?.message);
            }
          })
          .catch((error) => {
            console.log(error);
            messageApi.error("新增接口失败，请稍后重试！");
          });
      })
      .catch((info) => {
        console.log("Validate Failed:", info);
      });
  };

  const handleSave = async () => {
    const row = await editForm.validateFields();
    ApiClient.post("/admin/setting/api", row)
      .then((response) => {
        if (response.data?.code === 0) {
          messageApi.success("接口信息更新成功");
          setSearchParam({ ...searchParam });
          setEditingKey(0);
        } else {
          messageApi.error(response.data?.message);
        }
      })
      .catch((error) => {
        console.log(error);
        messageApi.error("更新接口信息失败，请稍后重试！");
      });
  };

  return (
    <>
      <Row>
        <Col span={12}>
          <Space>
            <ApiOutlined />
            <h3>接口</h3>
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
            添加接口
          </Button>
        </Col>
      </Row>
      <Row>
        <Col span={24}>
          <Form form={editForm} component={false}>
            <Table
              columns={columns}
              dataSource={dataSource?.List}
              scroll={{
                x: "100%",
                y: window.innerHeight - 370,
              }}
              loading={isLoading}
              pagination={{
                pageSize: searchParam.size,
                current: searchParam.page,
                // simple: false,
                showSizeChanger: true,
                hideOnSinglePage: false,
                total: dataSource.Count,
                onChange: (page, size) => {
                  setEditingKey(0);
                  setSearchParam({ page: page, size: size });
                },
              }}
            />
          </Form>
        </Col>
      </Row>
      <Modal
        title="添加接口"
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
            label="接口名称"
            name="name"
            rules={[{ required: true, message: "请输入接口名称" }]}
          >
            <Input placeholder="请输入新的接口名称" />
          </Form.Item>
          <Form.Item
            label="接口地址"
            name="path"
            rules={[{ required: true, message: "请输入接口地址" }]}
          >
            <Input placeholder="请输入新的接口地址" />
          </Form.Item>
          <Form.Item
            label="请求方法"
            name="method"
            rules={[{ required: true, message: "请选择接口请求方法" }]}
          >
            <Select
              options={[
                {
                  value: "GET",
                  label: "GET",
                },
                {
                  value: "PUT",
                  label: "PUT",
                },
                {
                  value: "POST",
                  label: "POST",
                },
                {
                  value: "DELETE",
                  label: "DELETE",
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

export default SettingApi;
