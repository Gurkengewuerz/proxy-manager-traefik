import request from "../request";

const getAuthInfo = () => {
  return request.get("/auth/info")
}

export {
  getAuthInfo
}