FROM golang:1.18-alpine  AS go-builder

WORKDIR /build/
COPY ../server ./server/
COPY ../go.mod ./
COPY ../go.sum ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server traefikmanager/server

FROM node:16-alpine as npm-build
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat

WORKDIR /build/
COPY ../web ./

ENV NODE_ENV production
RUN npm install --prefer-offline && npm run build


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app/
COPY --from=go-builder /build/server ./
RUN mkdir -p /app/web/
COPY --from=npm-build /build/build/ ./web/build/

CMD ["./server"]