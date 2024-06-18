FROM golang:alpine as builder

WORKDIR /go/src/github.com/Dbinggo/HireSphere/server

COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

#LABEL MAINTAINER=""

WORKDIR /server

COPY --from=builder /go/src/github.com/Dbinggo/HireSphere/server ./
COPY --from=builder /go/src/github.com/Dbinggo/HireSphere/server/resource ./resource/
COPY --from=builder /go/src/github.com/Dbinggo/HireSphere/server/config.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.yaml