# 使用Ubuntu 20.04作为基础镜像
FROM ubuntu:20.04

# 创建 /home 目录，如果它不存在
# RUN mkdir -p /home

# 设置工作目录为/opt
#WORKDIR /home/app

# 设置语言环境为en_US.utf8
ENV LANG=en_US.utf8

# 更新 apt 软件包列表并安装必要的依赖
RUN apt-get update && \
    apt-get install -y git golang npm && \
    npm install -g pnpm && \
    apt-get clean

# 复制entrypoint.sh脚本到容器中
COPY entrypoint.sh .

# 赋予entrypoint.sh脚本可执行权限
RUN chmod +x entrypoint.sh

# 设置入口点命令
ENTRYPOINT ["./entrypoint.sh"]

# 在目录下运行 sudo docker build -t ve-blog-server . 构建镜像
