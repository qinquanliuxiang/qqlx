import { useEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";

export default function Root() {
  const usenavigate = useNavigate();
  const path = window.location.pathname;
  useEffect(() => {
    switch (path) {
      case "/":
        usenavigate("/workspace/ram/user");
        break;
      case "/workspace":
        usenavigate("/workspace/ram/user");
        break;
      case "/workspace/":
        usenavigate("/workspace/ram/user");
        break;
    }
  }, [path, usenavigate]);
  return <Outlet />;
}
