import axios from "axios";
import { HttpEnum } from "@/enum";
import { GetToken, RemoveToken } from "./localStore";

const viteEnv = import.meta.env.VITE_ENV || "dev";
let host = "";

switch (viteEnv) {
  case "dev":
    host = "http://www.qqlx.net:8080";
    break;
  case "pro": {
    const protocol = window.location.protocol;
    const hostname = window.location.hostname;
    const prot = window.location.port;
    host = `${protocol}//${hostname}:${prot}/`;
    break;
  }
}

axios.defaults.timeout = 100000;
axios.defaults.baseURL = host;

// 请求拦截器
axios.interceptors.request.use(
  function (config) {
    if (!config.headers[HttpEnum.ContentType]) {
      config.headers[HttpEnum.ContentType] = HttpEnum.ContentTypeJson;
    }
    const token = GetToken();
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

// 添加响应拦截器
axios.interceptors.response.use(
  function (response) {
    // 检查 response 和 response.data 是否存在
    if (response && response.status === 200 && response.data) {
      return response.data; // 确保 status 为 200 并且返回 data
    }
  },
  function (error) {
    if (error?.response?.status === 401) {
      RemoveToken();
      // 保存当前路径
      const originalPath = window.location.pathname;
      // 将路径存入 sessionStorage 或 localStorage
      sessionStorage.setItem("originalPath", originalPath);
      // 重定向到登录页
      window.location.href = "/login";

      // // 登录成功后，获取存储的原始路径
      // const originalPath = sessionStorage.getItem('originalPath');

      // // 重定向用户到原始路径或主页
      // if (originalPath) {
      //   navigate(originalPath);
      //   sessionStorage.removeItem('originalPath');  // 清除存储的路径
      // } else {
      //   navigate('/');
      // }
    }
    return Promise.reject(error);
  }
);

// get方法
export function get(url: string, params: unknown) {
  return new Promise((resolve, reject) => {
    axios
      .get(url, {
        params: params,
      })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        reject(error);
      });
  });
}

// post方法
export function post(url: string, data: unknown) {
  return new Promise((resolve, reject) => {
    axios.post(url, data).then(
      (response) => {
        resolve(response);
      },
      (err) => {
        reject(err);
      }
    );
  });
}

export interface responseType {
  code: number;
  msg: string;
  err: string;
  data: unknown;
}

//统一接口处理，返回数据
export default function (method: string, url: string, param?: unknown) {
  return new Promise<responseType>((resolve, reject) => {
    switch (method) {
      case HttpEnum.GetMothod:
        get(url, param)
          .then((response) => {
            resolve(response as responseType);
          })
          .catch((error) => {
            reject(error);
          });
        break;
      case HttpEnum.PostMothod:
        post(url, param)
          .then((response) => {
            resolve(response as responseType);
          })
          .catch((error) => {
            reject(error);
          });
        break;
    }
  });
}
