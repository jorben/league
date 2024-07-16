import React, { useEffect, useRef, useState } from "react";
import { PlusOutlined } from "@ant-design/icons";
import { Input, Tag, theme, message } from "antd";
import ApiClient from "../../../services/client";
// import { TweenOneGroup } from "rc-tween-one";
const UserDetailGroup = ({ userId, group }) => {
  const { token } = theme.useToken();
  const [tags, setTags] = useState([]);
  const [inputVisible, setInputVisible] = useState(false);
  const [inputValue, setInputValue] = useState("");
  const inputRef = useRef(null);
  const [messageApi, contextHolder] = message.useMessage();

  useEffect(() => {
    if (inputVisible) {
      inputRef.current?.focus();
    }
  }, [inputVisible]);

  useEffect(() => {
    setTags(group ?? []);
  }, [group]);
  const handleClose = (removedTag) => {
    exitGroup(removedTag);
  };
  const showInput = () => {
    setInputVisible(true);
  };
  const handleInputChange = (e) => {
    setInputValue(e.target.value);
  };
  const handleInputConfirm = () => {
    if (inputValue && tags.indexOf(inputValue) === -1) {
      // 加入用户组
      joinGroup(inputValue);
    }
    setInputVisible(false);
    setInputValue("");
  };
  const forMap = (tag) => (
    <span
      key={tag}
      style={{
        display: "inline-block",
      }}
    >
      <Tag
        color="blue"
        closable
        onClose={(e) => {
          e.preventDefault();
          handleClose(tag);
        }}
      >
        {tag}
      </Tag>
    </span>
  );
  const tagChild = tags ? tags.map(forMap) : "";
  const tagPlusStyle = {
    background: token.colorBgContainer,
    borderStyle: "dashed",
  };
  const joinGroup = async (groupName) => {
    const data = { id: userId, group: groupName };
    ApiClient.post("/admin/user/join_group", data)
      .then((response) => {
        if (response.data?.code === 0) {
          setTags([...tags, groupName]);
          messageApi.success("已加入角色");
        } else {
          messageApi.error(response.data?.message);
        }
      })
      .catch((error) => {
        console.log(error);
        messageApi.error("请求失败，请稍后重试！");
      });
  };
  const exitGroup = async (groupName) => {
    const data = { id: userId, group: groupName };
    ApiClient.post("/admin/user/exit_group", data)
      .then((response) => {
        if (response.data?.code === 0) {
          const newTags = tags.filter((tag) => tag !== groupName);
          setTags(newTags);
          messageApi.success("已退出角色");
        } else {
          messageApi.error(response.data?.message);
        }
      })
      .catch((error) => {
        console.log(error);
        messageApi.error("请求失败，请稍后重试！");
      });
  };
  return (
    <>
      <div
        style={{
          marginBottom: 16,
        }}
      >
        <Tag>Member</Tag>
        {tagChild}
      </div>
      {inputVisible ? (
        <Input
          ref={inputRef}
          type="text"
          size="small"
          style={{
            width: 78,
          }}
          value={inputValue}
          onChange={handleInputChange}
          onBlur={handleInputConfirm}
          onPressEnter={handleInputConfirm}
        />
      ) : (
        <Tag color="geekblue" onClick={showInput} style={tagPlusStyle}>
          <PlusOutlined /> 加入角色
        </Tag>
      )}
      {contextHolder}
    </>
  );
};

export default UserDetailGroup;
