# 使用golang:alpine作为基础镜像
FROM golang:alpine AS builder

# 设置当前阶段标签，当前阶段是用于构建 Golang 项目的
LABEL stage=gobuilder

# 用阿里云的镜像地址替换默认的Alpine Linux软件包下载地址
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 设置Golang环境变量
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

# 安装必要的依赖
RUN apk update --no-cache
RUN apk add --no-cache git
RUN apk add --no-cache tzdata

# 设置新构建容器的工作目录
WORKDIR /

# 下载项目
RUN git init
RUN git clone https://github.com/ve-weiyi/ve-blog-golang.git

# 切换到项目目录
WORKDIR /ve-blog-golang/server

# 切换到目标分支并拉取最新代码
RUN git checkout feature/blog
RUN go mod tidy

# 构建Golang运行程序到 /app 目录下
RUN go build -ldflags="-s -w" -o /app/main main.go

# 使用scratch作为基础镜像,scratch用于构建轻量级的、无操作系统依赖的容器镜像

FROM scratch

# 设置语言环境为en_US.utf8
#ENV LANG=en_US.utf8
# 设置时区环境变量
#ENV TZ Asia/Shanghai

# 复制启动文件到新容器中
COPY --from=builder /app/main /app/main

# 设置新新容器的工作目录
WORKDIR /app

# 设置启动命令(使用config.yaml)
#ENTRYPOINT ./main api --config=config.yaml

# 设置启动命令(使用nacos作为配置中心)
ENTRYPOINT ./main api --n-namespace=test
