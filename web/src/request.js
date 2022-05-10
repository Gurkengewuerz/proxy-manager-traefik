import axios from 'axios'
import setupAxiosInterceptors from "./axios-interceptor";

const BACKEND_URL = "http://localhost:4664"

const request = axios.create({
  baseURL: BACKEND_URL,
})

setupAxiosInterceptors(request, () => {
  if (localStorage.getItem("authenticationToken")) {
    localStorage.removeItem("authenticationToken");
  }

  window.location = BACKEND_URL + "/login";
});

export default request