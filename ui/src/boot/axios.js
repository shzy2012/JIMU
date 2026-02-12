import { boot } from "quasar/wrappers";
import { Cookies } from "quasar";
import axios from "axios";
import { useRouter } from "vue-router";
import { Endpoint } from "./config";

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
const api = axios.create({
  baseURL: Endpoint,
  timeout: 0, // 设置超时时间为0，表示请求不会超时
});

const router = useRouter();

// Add a request interceptor
api.interceptors.request.use(
  function (config) {
    if (["login", "open"].some((v) => config.url.includes(v))) {
      return config;
    }

    const u = Cookies.get("u");
    if (!u || u.token == "") {
      //拦截
      // 检查当前页面路径
      const currentPath = window.location.pathname;
      // 如果已经在 /open 页面，不进行跳转，让请求失败（避免循环刷新）
      if (currentPath.includes("/open")) {
        // 不设置 token，让请求继续但会失败，由页面处理错误
        return config;
      }
      // 其他页面跳转到登录页
      window.location.href = "/login";
      window.location.reload();
    }

    config.headers = {
      ...config.headers,
      // Authorization: `Bearer ${u.token}`,
      token: u.token,
    };
    return config;
  },
  function (error) {
    // Do something with request error
    return Promise.reject(error);
  }
);

// Add a response interceptor
api.interceptors.response.use(
  function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data
    return response;
  },
  function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    if (error.response && error.response.status === 401) {
      // 场景:
      // 1: 服务端判断用户非法,走401状态,跳转到login
      // 2: 进入到login页面后,因为本地存在用户信息,再次跳转到跟页面/
      // 3: 提现循环跳转
      // 解决发方法: 将非法的本地用户清除掉,避免循环形成
      Cookies.remove("u", { path: "/" }); //勿删、勿动
      // 检查当前页面路径
      const currentPath = window.location.pathname;
      // 如果已经在 /open 页面，不进行跳转（避免循环刷新），让页面处理错误
      if (currentPath.includes("/open")) {
        return Promise.reject(error);
      }
      // 其他页面跳转到登录页
      window.location.href = "/login";
    }
    return Promise.reject(error);
  }
);

export default boot(({ app }) => {
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios;
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api;
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
});

export { api };
