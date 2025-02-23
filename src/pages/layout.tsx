import { GetToken } from "@/utils/localStore";
import { Layout, Menu } from "antd";
import { Content, Header } from "antd/es/layout/layout";
import Sider from "antd/es/layout/Sider";
import { Outlet, useNavigate } from "react-router-dom";
import styles from "./index.module.css";
import TopBar from "@/components/topBar";
import { Memu } from "@/route";
import { useState } from "react";
import { MemuEnum } from "@/enum";
import { useMount } from "ahooks";

interface openKeys {
  openKey: string[];
  selectKeys: string[];
}
export default function Workspace() {
  const navigate = useNavigate();
  const tolen = GetToken();
  if (!tolen) {
    window.location.href = `/login?redirect=${encodeURIComponent(
      window.location.pathname
    )}`;
  }
  const [openKey, setOpenkey] = useState<openKeys>({
    openKey: [],
    selectKeys: [],
  });
  const menuClick = (item: { keyPath: string[] }) => {
    setOpenkey({
      openKey: [item.keyPath[1]],
      selectKeys: [item.keyPath[0]],
    });
    const paht = MemuEnum.Perfix + item.keyPath.reverse().join("/");
    navigate(paht);
  };
  const menuChange = (openKeys: string[]) => {
    setOpenkey((pre) => ({
      ...pre,
      openKey: openKeys,
    }));
  };

  useMount(() => {
    const path = location.pathname.split("/");
    setOpenkey({
      openKey: [path[2]],
      selectKeys: [path[3]],
    });
  });
  return (
    <Layout className={styles.Layout}>
      <Header className={styles.Header}>
        <TopBar />
      </Header>
      <div style={{ height: "10px" }}></div>
      <Layout>
        <Sider>
          <Menu
            style={{ height: "100%", fontWeight: "bold" }}
            items={Memu}
            mode="inline"
            onClick={menuClick}
            onOpenChange={menuChange}
            openKeys={openKey.openKey}
            selectedKeys={openKey.selectKeys}
          />
        </Sider>
        <Content>
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  );
}
