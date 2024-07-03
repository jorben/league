
import Home from "../pages/Home";
import Login from "../pages/Login";
import NotFount from "../pages/NotFount";
import Dashboard from "../pages/admin/dashboard/Dashboard";
import List from "../pages/admin/user/List";

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
        path: "/404",
        element: <NotFount />,
    }
];

export const adminRoutes = [
    {
        path: "/",
        element: <Dashboard />,
        exact: true,
    },
    {
        path: "/user",
        element: <List />,
        exact: true
    }
]