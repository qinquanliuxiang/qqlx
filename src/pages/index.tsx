import { GetToken } from "@/utils/localStore";
import { Outlet } from "react-router-dom";

export default function Workspace() {
  const tolen = GetToken();
  if (!tolen) {
    window.location.href = `/login?redirect=${encodeURIComponent(
      window.location.pathname
    )}`;
  }
  return <Outlet />;
}
