import { api } from "boot/axios";

/**
 * @description -封装User类型的接口方法
 */
export class UserService {
  // 模块一
  /**
   * @description 用户登录
   * @param {string} username - 用户名
   * @return {HttpResponse} result
   */
  static async login(params) {
    // 接口一
    return api.post("/login", params);
  }
  static async list(params) {
    return api.post("/api/v1/user/list", params);
  }

  // createOrUpdate 存在ID更新,否则新增
  static createOrUpdate(params) {
    return api.post("/api/v1/user/add_update", params);
  }

  static get(id) {
    return api.get("/api/v1/user?id=" + id, {});
  }

  static delete(id) {
    return api.delete("/api/v1/user?id=" + id, {});
  }

  static async options_manger() {
    return api.get("/api/v1/user/manger", {});
  }

  static async options_staff() {
    return api.get("/api/v1/user/staff", {});
  }
}
