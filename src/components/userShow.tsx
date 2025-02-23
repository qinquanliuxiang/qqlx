import { GetUserInfo } from "@/apis/user";
import useGlobalStore from "@/store/global";
import { User } from "@/types/user";
import { useRequest } from "ahooks";
import { App, Avatar, Dropdown, MenuProps, Modal } from "antd";
import { useState } from "react";

export default function UserShow() {
  const { message } = App.useApp();
  const { user, setUser } = useGlobalStore();
  const [modalOpen, setModalOpen] = useState<boolean>(false);

  useRequest(GetUserInfo, {
    onSuccess: (res) => {
      if (res.code === 0) {
        setUser(res.data as User);
      } else {
        message.error(res.err);
      }
    },
    onError: (err) => {
      message.error(err.message);
    },
  });
  const items: MenuProps["items"] = [
    {
      key: "1",
      label: (
        <a
          onClick={() => {
            setModalOpen(true);
            // getUserByIdRun({ id: user?.user.id });
          }}
        >
          个人信息
        </a>
      ),
    },
    {
      key: "2",
      label: (
        <a
          onClick={() => {
            setModalOpen(true);
          }}
        >
          修改密码
        </a>
      ),
    },
    {
      key: "3",
      label: <a>退出登录</a>,
    },
  ];

  return (
    <>
      <Dropdown
        placement="bottomRight"
        menu={{ items, selectable: true }}
        trigger={["click"]}
      >
        <a onClick={(e) => e.preventDefault()}>
          <Avatar src={user.avatar} size={40} />
        </a>
      </Dropdown>
      <Modal></Modal>
    </>
  );
}
