# 使用centos:7作为基础镜像
FROM centos:7

# 创建 /home 目录，如果它不存在
# RUN mkdir -p /home

# 设置工作目录为/opt
#WORKDIR /home/app

# 设置语言环境为en_US.utf8
ENV LANG=en_US.utf8

# 安装必要的依赖
#RUN yum -y install epel-release && \
#    yum -y install git go npm && \
#    npm install -g pnpm && \
#    yum clean all

# 使用 sed 命令替换默认的软件源为阿里云镜像源
RUN sed -i 's|^mirrorlist=https://mirrorlist.centos.org|mirrorlist=https://mirrors.aliyun.com|g' /etc/yum.repos.d/CentOS-Base.repo \
    && sed -i 's|^#baseurl=https://mirror.centos.org|baseurl=https://mirrors.aliyun.com|g' /etc/yum.repos.d/CentOS-Base.repo

# 安装必要的依赖
RUN yum -y install epel-release && \
    yum -y install git go

# 暴露应用程序的端口
EXPOSE 9999

# 复制entrypoint.sh脚本到容器中
COPY entrypoint.sh .

# 赋予entrypoint.sh脚本可执行权限
RUN chmod +x entrypoint.sh

# 设置入口点命令
ENTRYPOINT ["./entrypoint.sh"]

# 在目录下运行 sudo docker build -t ve-blog-server . 构建镜像
