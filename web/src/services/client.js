import axios from "axios";

const ApiClient = axios.create({
  baseURL: "http://127.0.0.1:8080/api",
});

ApiClient.interceptors.request.use((config) => {
  // TODO: 计算csrf token
  const csrfToken = "1234";
  const jwt =
    "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJsZWFndWUiLCJleHAiOjE3MjAyOTIxOTMsIm5iZiI6MTcyMDI4NDY5MywiaWF0IjoxNzIwMjg0OTkzLCJqdGkiOiIxIn0.4Q2L3FfuxrBQ7d05NiY7_dNqZi_ckCM36lv2FSR3YLuUgkTeNrY8Wp5GGxt-GVh6";
  if (csrfToken) {
    config.headers["X-Csrf-Token"] = csrfToken;
  }
  if (jwt) {
    config.headers["X-Token"] = jwt;
  }
  return config;
});

export default ApiClient;
