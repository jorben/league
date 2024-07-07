import axios from "axios";
import CONSTANTS from "../constants";

const ApiClient = axios.create({
  baseURL: CONSTANTS.BASEURL_API,
});

ApiClient.interceptors.request.use((config) => {
  // TODO: 计算csrf token
  const csrfToken = "1234";
  const jwt = JSON.parse(localStorage.getItem(CONSTANTS.STORAGE_KEY_JWT));
  if (csrfToken) {
    config.headers[CONSTANTS.HEADER_KEY_CSRF] = csrfToken;
  }
  if (jwt?.token) {
    config.headers[CONSTANTS.HEADER_KEY_JWT] = jwt?.token;
  }
  return config;
});

export default ApiClient;
