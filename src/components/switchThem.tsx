import useGlobalStore from "@/store/global";
import { MoonOutlined, SunOutlined } from "@ant-design/icons";
import { Button } from "antd";
import { useEffect } from "react";

export default function SwitchThemComponent() {
  const { isDarkMode, setIsDarkMode } = useGlobalStore();
  useEffect(() => {
    localStorage.setItem("isDarkMode", JSON.stringify(isDarkMode));
  }, [isDarkMode]);

  return (
    <Button
      icon={isDarkMode ? <SunOutlined /> : <MoonOutlined />}
      onClick={() => setIsDarkMode(!isDarkMode)}
    >
      切换主题
    </Button>
  );
}
