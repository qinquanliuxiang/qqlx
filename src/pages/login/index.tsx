import { Login } from "@/apis/user";
import { UrlParams } from "@/enum";
import { LoginRequestType, LoginResponseType } from "@/types/login";
import { GetToken } from "@/utils/localStore";
import { useRequest } from "ahooks";
import { App, Button, Card, Form, Input, Layout } from "antd";
import { useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import styles from "./index.module.css";
import SwitchThemComponent from "@/components/switchThem";
export default function LoginPage() {
  const token = GetToken();
  const location = useLocation();
  const navigate = useNavigate();
  const { message } = App.useApp();
  const { run: onFinish, loading: loginLoading } = useRequest(Login, {
    manual: true,
    onSuccess: (res) => {
      if (res.code === 0) {
        if (res.data) {
          const r = res.data as LoginResponseType;
          localStorage.setItem("token", r.token);
          localStorage.setItem("user", JSON.stringify(r.user));
        }
        navigate("/workspace/ram/user");
      } else {
        message.open({
          type: "error",
          content: res.err,
        });
      }
    },
    onError: (err) => {
      message.open({
        type: "error",
        content: err.message,
      });
    },
  });

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
    <Layout>
      <div className={styles.login}>
        <Card
          extra={<SwitchThemComponent />}
          title={<p className={styles.p}>欢迎登录</p>}
          styles={{
            body: {
              width: "400px",
            },
          }}
        >
          <Form
            name="login"
            style={{ maxWidth: 600 }}
            onFinish={onFinish}
            autoComplete="off"
          >
            <Form.Item<LoginRequestType>
              name="email"
              rules={[
                { required: true, message: "邮箱为空" },
                {
                  pattern:
                    /^[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)*@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/,
                  message: "邮箱格式不正确",
                },
              ]}
            >
              <Input placeholder="请输入邮箱" />
            </Form.Item>

            <Form.Item<LoginRequestType>
              name="password"
              rules={[
                { required: true, message: "密码为空" },
                { min: 8, message: "密码至少8位" },
              ]}
            >
              <Input.Password placeholder="请输入密码" />
            </Form.Item>
            <Form.Item>
              <Button
                block
                loading={loginLoading}
                type="primary"
                htmlType="submit"
              >
                登录
              </Button>
            </Form.Item>
          </Form>
        </Card>
      </div>
    </Layout>
  );
}
