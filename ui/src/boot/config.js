import { ref } from "vue";
import { Cookies } from "quasar";

if (process.env.PROD) {
  console.log("[Quasar] Running PROD.");
}
if (process.env.DEV) {
  console.log("[Quasar] Running DEV.");
}

let Endpoint = "http://127.0.0.1:8000";
if (process.env.PROD) {
  Endpoint = "";
}

export { Endpoint };
export const RightDrawerOpen = ref(false);
export const ChatChanged = ref("");

export class LocalStorageManager {
  constructor(key) {
    this.key = key;
  }
  set(value) {
    localStorage.setItem(this.key, JSON.stringify(value));
  }
  get() {
    const storedValue = localStorage.getItem(this.key);
    return storedValue ? JSON.parse(storedValue) : null;
  }
  del() {
    localStorage.removeItem(this.key);
  }
}

export const HasAuth = () => {
  const u = GetU();
  if (u) {
    if (u.role == "operator" || u.role == "admin") {
      return true;
    }
  }
  return false;
  return true;
};

// 获取
export const GetU = function () {
  return Cookies.get("u", { path: "/" });
};

// 移除
export const DelU = function () {
  Cookies.remove("u", { path: "/" });
};

// 设置
export const SetU = function (uData) {
  Cookies.set("u", uData, { path: "/" });
};

export const uuidv4 = function () {
  return "10000000-1000-4000-8000-100000000000".replace(/[018]/g, (c) =>
    (
      +c ^
      (crypto.getRandomValues(new Uint8Array(1))[0] & (15 >> (+c / 4)))
    ).toString(16)
  );
};
