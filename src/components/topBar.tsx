import { Flex, Space } from "antd";
import SwitchThemComponent from "./switchThem";
import UserShow from "./userShow";

export default function TopBar() {
  return (
    <Flex flex={1} justify="space-between" align="center" gap={100}>
      <div></div>

      <Space size="large">
        <UserShow />
        <SwitchThemComponent />
      </Space>
    </Flex>
  );
}
