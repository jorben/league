import axios from "axios";
import CONSTAANTS from "../constants";

const ApiClient = axios.create({
  baseURL: CONSTAANTS.BASEURL_API,
});

ApiClient.interceptors.request.use((config) => {
  // TODO: 计算csrf token
  const csrfToken = "1234";
  const jwt = JSON.parse(localStorage.getItem(CONSTAANTS.STORAGE_KEY_JWT));
  if (csrfToken) {
    config.headers[CONSTAANTS.HEADER_KEY_CSRF] = csrfToken;
  }
  if (jwt?.token) {
    config.headers[CONSTAANTS.HEADER_KEY_JWT] = jwt?.token;
  }
  return config;
});

export default ApiClient;
