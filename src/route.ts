import { createBrowserRouter } from "react-router";
import User from "./pages/user";
import Role from "./pages/role";
import Root from "./root";
import Workspace from "./pages";
import Login from "./pages/login";
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

export default router;
