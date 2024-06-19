FROM golang:alpine as builder

WORKDIR /server

COPY  . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && cd cmd \
    && go build -o  server . \
    && mv ./server ../

FROM alpine:latest

#LABEL MAINTAINER=""

WORKDIR /server

COPY --from=builder /server/server ./
COPY --from=builder /server/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml