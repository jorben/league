
import App from "../pages/App";
import Home from "../pages/Home";
import Login from "../pages/Login";
import NotFount from "../pages/NotFount";
import AdminDashboard from "../pages/admin/dashboard/Dashboard";
import AdminUserList from "../pages/admin/user/UserList";
import { UserOutlined, HomeOutlined } from '@ant-design/icons';

export const mainRoutes = [
    {
        path: "/",
        element: <Home />,
        exact: true,
    },
    {
        path: "/login",
        element: <Login />,
    },
    {
        path: "/app",
        element: <App/>
    },
    {
        path: "/404",
        element: <NotFount />,
    }
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
        exact: true
    },
];

export const adminMenus = [
    {
      key: '/admin',
      icon: <HomeOutlined />,
      label: 'Dashboard',
    },
    {
      key: '2',
      icon: <UserOutlined />,
      label: '用户及权限',
      children: [
        {
          key: '/admin/user',
          label: '用户列表',
        },
        {
          key: '/admin/user/policy',
          label: '权限规则',
        },
      ],
    },
  ];