import { useEffect } from "react";
import { App, ConfigProvider, theme } from "antd";
import { RouterProvider } from "react-router-dom";
import router from "./route";
import useGlobalStore from "./store/global";

function MyApp() {
  const { isDarkMode, setIsDarkMode } = useGlobalStore();
  const model = localStorage.getItem("isDarkMode");
  useEffect(() => {
    if (model === null) {
      localStorage.setItem("isDarkMode", "false");
    } else {
      setIsDarkMode(JSON.parse(model));
    }
  }, []);
  return (
    <ConfigProvider
      theme={{
        algorithm: isDarkMode
          ? [theme.darkAlgorithm, theme.compactAlgorithm]
          : [theme.defaultAlgorithm, theme.compactAlgorithm],
      }}
    >
      <App className="app">
        <RouterProvider router={router} />
      </App>
    </ConfigProvider>
  );
}

export default MyApp;
