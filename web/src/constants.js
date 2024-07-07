const CONSTAANTS = {
  ERRCODE: {
    ErrAuthNoLogin: -10005, //未登录
    ErrAuthUnauthorized: -10006, //未授权
  },
  BASEURL_API: "http://127.0.0.1:8080/api", // 后端接口baseurl
  STORAGE_KEY_JWT: "jwt", // jwt存储key名称
  HEADER_KEY_JWT: "X-Token", // jwt header key
  HEADER_KEY_CSRF: "X-Csrf-Token",
  DEFAULT_PAGESIZE: 20, // 默认的pagesize
};

export default CONSTAANTS;
