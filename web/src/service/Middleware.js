import request from "../request";

const getMiddlewares = () => {
  return request.get("/middleware")
}

const deleteMiddleware = (id) => {
  return request.delete("/middleware", {data: {id}})
}

const putMiddleware = (id) => {
  return undefined
}

const postMiddleware = (id) => {
  return request.post("/middleware", {data: {id}})
}

export {
  getMiddlewares,
  deleteMiddleware,
  putMiddleware,
  postMiddleware,
}