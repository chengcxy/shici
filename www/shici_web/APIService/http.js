/**
 * 封装Axios
 * 处理请求、响应错误信息
 */
import { Message } from "element-ui"; //引用饿了么UI消息组件
import axios from "axios"; //引用axios
import Cookies from "js-cookie";

axios.defaults.headers.post["Content-Type"] =
  "application/x-www-form-urlencoded";

// create an axios instance
const service = axios.create({
  baseURL: "/api", // 所有异步请求都加上/api,nginx转发到后端Springboot
  withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000 // request timeout
});

// request interceptor
service.interceptors.request.use(
  config => {
    config.headers["Token"] = Cookies.get("token") || "";
    return config;
  },
  err => {
    return Promise.reject(err);
  }
);
// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
   */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data; //res is my own data

    return res;
  },
  error => {
    if (error.response.status && error.response.status === 413) {
      Message({
        message: "图片太大",
        type: "error",
        duration: 5 * 1000
      });
    } else {
      Message({
        message: "连接超时",
        type: "error",
        duration: 5 * 1000
      });
    }
    return Promise.reject(error);
  }
);

export default service; //导出封装后的axios
