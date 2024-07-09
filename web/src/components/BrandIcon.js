import React from "react";
import {
  GoogleOutlined,
  GithubOutlined,
  QqOutlined,
  WechatOutlined,
  QuestionCircleOutlined,
} from "@ant-design/icons";
import CONSTANTS from "../constants";

const BrandIcon = ({ name, size }) => {
  if (name === CONSTANTS.USER_PROVIDER.GOOGLE) {
    return <GoogleOutlined style={{ fontSize: size + "px" }} />;
  } else if (name === CONSTANTS.USER_PROVIDER.GITHUB) {
    return <GithubOutlined style={{ fontSize: size + "px" }} />;
  } else if (name === CONSTANTS.USER_PROVIDER.QQ) {
    return <QqOutlined style={{ fontSize: size + "px" }} />;
  } else if (name === CONSTANTS.USER_PROVIDER.WECHAT) {
    return <WechatOutlined style={{ fontSize: size + "px" }} />;
  } else {
    return <QuestionCircleOutlined style={{ fontSize: size + "px" }} />;
  }
};

export default BrandIcon;
