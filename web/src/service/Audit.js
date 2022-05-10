import request from "../request";

const getAudit = () => {
  return request.get("/audit")
}

export {
  getAudit
}