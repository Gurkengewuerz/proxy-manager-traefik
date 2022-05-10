import request from "../request";

const getRouters = () => {
  return request.get("/router")
}

const deleteRouter = (id) => {
  return request.delete("/router", {data: {id}})
}

const putRouter = (id) => {
  return undefined
}

const postRouter = (id) => {
  return request.post("/router", {data: {id}})
}

export {
  getRouters,
  deleteRouter,
  putRouter,
  postRouter,
}