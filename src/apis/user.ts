import { Http } from "@/enum";
import { LoginRequestType } from "@/types/login";
import http from "@/utils/http";

export async function Login(data: LoginRequestType) {
  return await http(Http.PostMothod, "/api/v1/user/login", data);
}
