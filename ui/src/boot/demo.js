import { api } from "boot/axios";

// 测试Guest
export class DemoService {
  static select() {
    return api.get("/api/v1/guest/demo/select", {});
  }

  static list(params) {
    return api.post("/api/v1/guest/demo/list", params);
  }

  static addOrUpdate(params) {
    return api.post("/api/v1/guest/demo", params);
  }

  static delete(id) {
    return api.delete("/api/v1/guest/demo?id=" + id, {});
  }
}
