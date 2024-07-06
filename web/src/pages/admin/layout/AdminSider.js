import React from "react";
import { Layout, Menu } from "antd";
import Logo from "../../../components/Logo";
import { useNavigate, useLocation } from "react-router-dom";
import * as Icons from "@ant-design/icons";

const { Sider } = Layout;

const AdminSider = ({ collapsed, allMenus }) => {
  // console.log("In AdminSider, menus:", allMenus);

  const navigate = useNavigate();
  const { pathname } = useLocation();

  // const iconMenus = (allMenus) => {
  //   return allMenus.map(node => ({node.icon ? Icons[node.icon] : '', ...node}))
  // }
  const transformMenuIcons = (menus) => {
    return menus.map((item) => ({
      ...item,
      icon: item.icon ? React.createElement(Icons[item.icon] || null) : null,
      ...(item.children && { children: transformMenuIcons(item.children) }),
    }));
  };

  const iconMenus = transformMenuIcons(allMenus);

  const findMenuKeyByPath = (currentPath, menus) => {
    for (const menu of menus) {
      // 1. 检查当前层级是否匹配
      if (menu.key === currentPath) {
        return menu.key;
      }

      // 2. 递归查找子菜单
      if (menu.children) {
        const foundKey = findMenuKeyByPath(currentPath, menu.children);
        if (foundKey) {
          return foundKey;
        }
      }
    }

    // 2. 没有完全匹配，则递归查找父级路径
    if (currentPath.includes("/")) {
      const parentPath = currentPath.substring(0, currentPath.lastIndexOf("/"));
      return findMenuKeyByPath(parentPath, menus);
    }
    return "";
  };

  const findParentMenuKeys = (targetKey, menus) => {
    const parentKeys = [];

    const findRecursive = (key, menus) => {
      for (const menu of menus) {
        if (menu.children) {
          const foundIndex = menu.children.findIndex(
            (child) => child.key === key
          );
          if (foundIndex !== -1) {
            parentKeys.push(menu.key);
            findRecursive(menu.key, iconMenus); // 继续查找当前父节点的父节点
            break; // 找到目标节点后，不再遍历其兄弟节点
          } else {
            findRecursive(key, menu.children); // 递归查找子菜单
          }
        }
      }
    };

    findRecursive(targetKey, menus);
    return parentKeys.reverse(); // 返回顺序是从顶级父节点到目标节点的父节点
  };

  const selectedKey = findMenuKeyByPath(pathname, iconMenus);
  const openKeys = findParentMenuKeys(selectedKey, iconMenus);

  return (
    <Sider trigger={null} collapsible collapsed={collapsed}>
      <Logo collapsed={collapsed} theme="black" />
      <Menu
        theme="dark"
        mode="inline"
        defaultSelectedKeys={["/admin"]}
        defaultOpenKeys={openKeys}
        selectedKeys={[selectedKey]}
        onClick={(e) => {
          navigate(e.key);
        }}
        items={iconMenus}
      />
    </Sider>
  );
};

export default AdminSider;
