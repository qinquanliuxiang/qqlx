import { UrlParams } from "@/enum";
import { GetToken } from "@/utils/localStore";
import { Button, Card, Form, Input, message } from "antd";
import { useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";

export default function Login() {
  const token = GetToken();
  const location = useLocation();
  const navigate = useNavigate();
  const [messageApi, contextHolder] = message.useMessage();
  useEffect(() => {
    if (token) {
      const searchParams = new URLSearchParams(location.search);
      if (searchParams.has(UrlParams.Redirect)) {
        const redirect = searchParams.get(UrlParams.Redirect);
        if (redirect) {
          navigate(redirect);
        }
      } else {
        navigate("/workspace/ram/user");
      }
    }
  }, [location.search, navigate, token]);
  return (
    <>
      {contextHolder}
      <Card
        title={getTitle()}
        styles={{
          header: {
            padding: "0px",
          },
          body: {
            width: "350px",
          },
        }}
      >
        <Form
          name="login"
          style={{ maxWidth: 600 }}
          onFinish={onFinish}
          autoComplete="off"
        >
          <Form.Item<FieldType>
            name="name"
            rules={[{ required: true, message: "请输入用户名!" }]}
          >
            <Input placeholder="用户名" />
          </Form.Item>

          <Form.Item<FieldType>
            name="password"
            rules={[{ required: true, message: "请输入密码!" }]}
          >
            <Input.Password placeholder="密码" />
          </Form.Item>
          <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
            <Button loading={loading} type="primary" htmlType="submit">
              登录
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </>
  );
}
