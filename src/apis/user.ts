import { HttpEnum } from "@/enum";
import { LoginRequestType } from "@/types/user";
import http from "@/utils/http";

export async function Login(data: LoginRequestType) {
  return await http(HttpEnum.PostMothod, "/api/v1/user/login", data);
}

// 获取当前登录用户信息
export async function GetUserInfo() {
  return await http(HttpEnum.GetMothod, "/api/v1/user/getUserInfo");
}
