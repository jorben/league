import App from "../pages/App";
import Home from "../pages/Home";
import Login from "../pages/Login";
import NotFount from "../pages/NotFount";
import AdminDashboard from "../pages/admin/dashboard/Dashboard";
import AdminUserList from "../pages/admin/user/UserList";
import AdminUserGroup from "../pages/admin/user/UserGroup";
import AdminSettingApi from "../pages/admin/setting/SettingApi";
import AdminSettingMenu from "../pages/admin/setting/SettingMenu";
import AdminPermissionApi from "../pages/admin/permission/PermissionApi";
import AdminPermissionMenu from "../pages/admin/permission/PermissionMenu";
import AdminPermissionData from "../pages/admin/permission/PermissionData";

export const mainRoutes = [
  {
    path: "/",
    element: <Home />,
    exact: true,
  },
  {
    path: "/index.html",
    element: <Home />,
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/app",
    element: <App />,
  },
  {
    path: "/404",
    element: <NotFount />,
  },
];

// children with /admin
export const adminRoutes = [
  {
    path: "/",
    element: <AdminDashboard />,
    exact: true,
  },
  {
    path: "/user",
    element: <AdminUserList />,
    exact: true,
  },
  {
    path: "/user/group",
    element: <AdminUserGroup />,
    exact: true,
  },
  {
    path: "/permission/menu",
    element: <AdminPermissionMenu />,
    exact: true,
  },
  {
    path: "/permission/api",
    element: <AdminPermissionApi />,
    exact: true,
  },
  {
    path: "/permission/data",
    element: <AdminPermissionData />,
    exact: true,
  },
  {
    path: "/setting/api",
    element: <AdminSettingApi />,
    exact: true,
  },
  {
    path: "/setting/menu",
    element: <AdminSettingMenu />,
    exact: true,
  },
];
