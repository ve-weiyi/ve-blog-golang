#!/bin/bash

# 输出日志到控制台
echo "starting entrypoint.sh script..."


# 创建 /home 目录
mkdir -p /home/blog

# 切换到 /home 目录
cd /home/blog

# 安装 git
git config --global user.email "791422171@qq.com"
git config --global user.name "veweiyi"
git config --global --unset http.proxy

# 下载项目
git clone git://github.com/ve-weiyi/ve-blog-golang.git

# 切换到项目目录
cd /home/blog/ve-blog-golang/server

# 切换到目标分支并拉取最新代码
git checkout feature/1.0.0/blog
git pull origin feature/1.0.0/blog


# git branch -d feature/1.0.0/blog
git fetch origin
git reset --hard origin/feature/1.0.0/blog

# 更新依赖
go mod tidy

# Run the Go application
# go run main.go server
go build main.go

nohup ./main server -c /config.yaml > output.log 2>&1 &
# 输出日志到控制台
echo "entrypoint.sh script completed."

tail -f /dev/null

# 生成镜像并部署 sudo docker-compose -f docker-compose.yaml up -d
# 进入镜像查看部署结果 sudo docker exec -it ve-blog-server /bin/bash
# netstat -tulpn | grep 9999 查看端口占用情况
