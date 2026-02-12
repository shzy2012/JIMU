import { api } from "boot/axios";

// 销售出库单
export class SaloutstockService {
  // 获取订单列表
  static list(params = {}) {
    return api.post("/api/v1/saloutstock/list", params);
  }

  // 获取订单详情
  static get(billNo) {
    return api.get("/api/v1/saloutstock", {
      params: { id: billNo },
    });
  }

  // 完成订单（审核订单）
  static completeOrder(params) {
    return api.post("/api/v1/saloutstock", params);
  }
}
