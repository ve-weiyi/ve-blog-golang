FROM golang:alpine AS builder

WORKDIR /go/src/github.com/ve-weiyi/ve-blog-golang/zero
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy

COPY service/blog/rpc/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/blog service/blog/rpc/blog.go

#FROM alpine:latest
#
#维护者信息
LABEL maintainer="791422171@qq.com"
#
WORKDIR /app
#COPY --from=0 /app/blog /app/blog
#COPY --from=0 /app/etc /app/etc
#
#EXPOSE 8888

# 启动命令，只能有一个生效
ENTRYPOINT ./blog
