import { createBrowserRouter } from "react-router";
import User from "./pages/user";
import Role from "./pages/role";
import Root from "./root";
import Workspace from "./pages/layout";
import Login from "./pages/login";
import { MenuProps } from "antd";
type MenuItem = Required<MenuProps>["items"][number];

const router = createBrowserRouter([
  {
    path: "/",
    Component: Root,
    children: [
      {
        path: "workspace",
        Component: Workspace,
        children: [
          {
            path: "ram",
            children: [
              {
                path: "user",
                Component: User,
              },
              {
                path: "role",
                Component: Role,
              },
            ],
          },
        ],
      },
    ],
  },
  {
    path: "/login",
    Component: Login,
  },
]);

export const Memu: MenuItem[] = [
  {
    key: "ram",
    icon: null,
    label: "用户",
    children: [
      {
        key: "user",
        icon: null,
        label: "用户列表",
      },
      {
        key: "role",
        icon: null,
        label: "角色列表",
      },
    ],
  },
];

export default router;
