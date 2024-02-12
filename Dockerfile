## Build
FROM golang:1.20-alpine AS build-server
WORKDIR /build
COPY ./server/go.mod ./server/go.sum ./
RUN go mod download && go mod verify
COPY ./server .
RUN go build -o ./server

FROM node:lts-alpine AS builder-web
WORKDIR /build
COPY ./web .
RUN echo "node" `node -v` && echo "yarn" `yarn -v` && yarn && yarn build

## Deploy
FROM ubuntu 
WORKDIR /app
COPY --from=build-server /build/server .
COPY --from=builder-web /build/dist ./static
COPY ./server/setting.yaml docker-entrypoint.sh ./
VOLUME ["/app/data"]
EXPOSE 8080
CMD [ "sh","./docker-entrypoint.sh" ]