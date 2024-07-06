import axios from "axios";

const ApiClient = axios.create({
  baseURL: "http://127.0.0.1:8080/api",
});

ApiClient.interceptors.request.use((config) => {
  // TODO: 计算csrf token
  const csrfToken = "1234";
  if (csrfToken) {
    config.headers["X-Csrf-Token"] = csrfToken;
  }
  return config;
});

export default ApiClient;
