import App from "../pages/App";
import Home from "../pages/Home";
import Login from "../pages/Login";
import NotFount from "../pages/NotFount";
import AdminDashboard from "../pages/admin/dashboard/Dashboard";
import AdminUserList from "../pages/admin/user/UserList";
import AdminUserGroup from "../pages/admin/user/UserGroup";

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
];
