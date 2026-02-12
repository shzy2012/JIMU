import { route } from "quasar/wrappers";
import { useQuasar } from "quasar";
import {
  createRouter,
  createMemoryHistory,
  createWebHistory,
  createWebHashHistory,
} from "vue-router";
import routes from "./routes";
import { GetU } from "../boot/config";
import { utils } from "../utils/index";
window.utils = utils; //挂载全局公共方法
/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Router instance.
 */

export default route(function (/* { store, ssrContext } */) {
  const createHistory = process.env.SERVER
    ? createMemoryHistory
    : process.env.VUE_ROUTER_MODE === "history"
    ? createWebHistory
    : createWebHashHistory;

  const Router = createRouter({
    scrollBehavior: () => ({ left: 0, top: 0 }),
    routes,

    // Leave this as is and make changes in quasar.conf.js instead!
    // quasar.conf.js -> build -> vueRouterMode
    // quasar.conf.js -> build -> publicPath
    history: createHistory(
      process.env.MODE === "ssr" ? void 0 : process.env.VUE_ROUTER_BASE
    ),
  });

  Router.beforeEach((to, from, next) => {
    console.log("资源加载中. 请稍等...");
    useQuasar().loading.show({
      message: "资源加载中. 请稍等...",
    });

    // 忽略无需登录的页面
    if (["login", "open"].some((v) => to.path.toLowerCase().includes(v))) {
      return next();
    }

    // 检查是否登录,未登录,则调转到/login页面
    const u = GetU();
    if (!u || u.token == "") {
      if (to.path.toLowerCase().includes("open")) {
        return next({ path: "/open" });
      } else {
        return next({ path: "/login" });
      }
    }
    next();
  });

  Router.afterEach(() => {
    useQuasar().loading.hide();
  });

  return Router;
});
