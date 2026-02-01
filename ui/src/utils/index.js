// 获取url指定参数
export const getQueryParam = (key) => {
  const newKey = String(key).toLowerCase();
  const queryString = window.location.href.split("?")[1];
  console.log(window.location.href);
  const urlParams = new URLSearchParams(queryString);
  return urlParams.get(newKey);
};

// 透传参数数组
const pass = ["channel"];

// 生成获取url指定参数的函数
const urlparams = (key) => {
  const newKey = String(key).toLowerCase();
  const queryString = window.location.href.split("?")[1];
  const urlParams = new URLSearchParams(queryString);
  return urlParams.get(newKey);
};

// url要剪切拼入的参数
const getSliceParam = () => {
  const arr = [];
  pass.forEach((v) => {
    const val = getQueryParam(v);
    const str = v + "=" + val;
    if (val || val === 0) {
      arr.push(str);
    }
  });
  return arr.join("&");
};

const getIntoParam = (obj) => {
  const arr = [];
  for (const key in obj) {
    if (Object.hasOwnProperty.call(obj, key)) {
      const val = obj[key];
      arr.push(`${key}=${val}`);
    }
  }
  return arr.join("&");
};
const goto = (url, obj) => {
  let newUrl = "";
  if (url.indexOf("http://") > -1 || url.indexOf("https://") > -1) {
    newUrl = url;
  } else {
    newUrl = `${window.location.origin}${url === "/" ? "" : url}`;
  }
  const iStr = getIntoParam(obj);
  const fStr = getSliceParam();
  if (iStr && fStr) {
    newUrl += `?${fStr}${iStr ? "&" + iStr : ""}`;
  }
  window.location.href = newUrl;
};

// 解决window.open函数在ios失效的情况
function navigation(url) {
  var u = navigator.userAgent;
  var isAndroid = u.indexOf("Android") > -1 || u.indexOf("Adr") > -1; //android终端
  var isiOS = !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/); //ios终端

  if (isiOS) {
    //ios终端
    window.location.href = url;
  } else if (isAndroid) {
    //android终端
    window.open(url, "_blank", "noreferrer");
  } else {
    window.open(url, "_blank", "noreferrer");
  }
}

// array 去重
const unique = (arr) => [...new Set(arr)];

const utc2local = (utc) => {
  const utcDate = new Date(utc);
  return utcDate.toLocaleString();
};

const clipboard = (text) => {
  if (window.isSecureContext && navigator.clipboard) {
    console.log("navigator clipboard");
    navigator.clipboard
      .writeText(text)
      .then(function () {})
      .catch(function (err) {
        console.log("navigator clipboard=>", err);
      });
  } else {
    const textArea = document.createElement("textarea");
    textArea.value = text;
    // Move textarea out of the viewport so it's not visible
    textArea.style.position = "absolute";
    textArea.style.left = "-999999px";
    document.body.prepend(textArea);
    textArea.select();
    try {
      document.execCommand("copy");
    } catch (error) {
      console.error(error);
    } finally {
      textArea.remove();
    }
  }
};

export const utils = {
  goto,
  navigation,
  unique,
  utc2local,
  urlparams,
  clipboard,
};
