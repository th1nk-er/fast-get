## Build
FROM golang:1.20-alpine AS build-server
WORKDIR /build
COPY ./server/go.mod ./server/go.sum ./
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go mod download && go mod verify
COPY ./server .
RUN go build -o ./server

FROM node:lts-alpine AS builder-web
WORKDIR /build
COPY ./web .
RUN echo "node" `node -v` && echo "yarn" `yarn -v` && yarn && yarn build

## Deploy
FROM nginx
WORKDIR /app
COPY --from=build-server /build/server .
COPY --from=builder-web /build/dist .
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY ./server/setting.yaml docker-entrypoint.sh ./
RUN mkdir data
VOLUME ["/app/data"]
CMD [ "sh","./docker-entrypoint.sh" ]