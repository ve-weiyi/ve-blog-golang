#!/bin/bash

# 创建 /home/app 目录
mkdir -p /home/app

# 切换到 /home/app 目录
cd /home/app

# 下载项目
git clone https://github.com/ve-weiyi/ve-blog-golang.git

# 切换到目标分支并拉取最新代码
# git checkout feature/1.0.0/blog

# 切换到项目目录
cd /home/app/ve-blog-golang/server

# 更新依赖
go mod tidy


# Run the Go application
go run main.go server

echo "blog ALL start!!!"
tail -f /dev/null

# 生成镜像并部署 sudo docker-compose -f docker-compose.yaml up -d
# 进入镜像查看部署结果 sudo docker exec -it ve-blog-server /bin/bash
