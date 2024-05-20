# Docker学习

## Docker基础概念概览
### 镜像（Image）
在Docker中，镜像是一个轻量级、独立的可执行软件包，包含运行应用程序所需的所有内容，
包括代码、运行时、库、环境变量和配置文件。镜像是容器的基石，容器实际上是在镜像的基础上创建的运行实例。

示例代码：创建一个简单的Nginx镜像
```bash
# Dockerfile
FROM nginx:latest
COPY index.html /usr/share/nginx/html/index.html
```
在上述示例中，基于最新版本的Nginx官方镜像创建了一个自定义镜像，并将一个index.html文件复制到Nginx的默认HTML目录。


### 容器（Container）
容器是镜像的运行实例。它包含了应用程序及其所有依赖项，以隔离应用程序及其环境，
确保它在任何环境中都能一致运行。容器是可移植和可部署的，可以在任何支持Docker的系统上运行。

示例代码：运行Nginx容器
```bash
docker run -d -p 8080:80 my-custom-nginx:latest
```
上述命令将在后台运行一个基于我们自定义镜像的Nginx容器，并将容器的80端口映射到主机的8080端口。

### 仓库（Repository）
Docker仓库是用于存储和组织镜像的地方。仓库可以是公共的（如Docker Hub）或私有的，
用户可以通过仓库来分享和获取镜像。

示例代码：推送和拉取镜像
```bash
# 将自定义镜像推送到Docker Hub
docker push your-username/my-custom-nginx:latest

# 从Docker Hub拉取镜像
docker pull your-username/my-custom-nginx:latest

```

在上述示例中，演示了如何将自定义镜像推送到Docker Hub，并从Docker Hub拉取镜像。


- **镜像类似汽车设计图纸，包含了汽车的各种详细参数和配置细节。**
- **容器类似具体的汽车，根据汽车设计图纸生产而来。用户可以在生产后直接运行，也可以修改配置后运行。**
- **仓库类似汽车设计图纸生产商，负责存储和分发设计图纸。**

## Docker安装
### 在 macOS 上使用命令行安装 Docker
```bash
brew install docker
```

### 在 Windows 上使用 Docker Desktop 安装 Docker
```
在 Windows 上，Docker Desktop for Windows 提供了一个可执行文件安装程序，
不能通过命令行直接安装 Docker。您需要从 Docker 官网下载并手动安装 Docker Desktop for Windows。

https://docs.docker.com/desktop/windows/install/
```

### 在 linux 上使用命令行安装 Docker

Ubuntu 或 Debian
```bash
sudo apt-get update
sudo apt-get install docker.io
```

CentOS 或 Fedora
```bash
sudo yum install docker
```

安装完成后，启动 Docker 服务
```bash
sudo systemctl start docker
```

## Docker的常用命令


## Docker的使用
### 如何构建一个镜像？

**Dockerfile**

Dockerfile 是用于构建 Docker 镜像的文本文件。它包含了一系列指令，用于描述如何构建一个 Docker 镜像。
以下是 Dockerfile 的主要用途和功能：
1. 定义基础镜像：基础镜像是构建你的应用程序或服务所需的操作系统和基本软件包的基础。
2. 安装依赖项：安装你的应用程序或服务所需的软件包、库和依赖项。
3. 配置环境：设置环境变量、工作目录、用户权限等配置，以确保你的应用程序在容器内部以期望的方式运行。
4. 复制文件：使用 Dockerfile COPY 或 ADD 指令将本地文件或目录复制到镜像中的指定位置。这使得你可以将应用程序的代码、配置文件、静态资源等打包到镜像中。
5. 运行命令：使用 RUN 指令在构建过程中执行命令。这些命令可以包括安装软件包、下载文件、设置环境等。
6. 暴露端口：使用 EXPOSE 指令，使得在运行容器时可以方便地将容器端口映射到主机端口。
7. 设置入口点：使用 Dockerfile 中的 CMD 或 ENTRYPOINT 指令，定义容器启动时要执行的默认命令或程序。

```bash
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

FROM alpine:latest

#维护者信息
LABEL maintainer="791422171@qq.com"

WORKDIR /app
COPY --from=0 /app/blog /app/blog
COPY --from=0 /app/etc /app/etc

EXPOSE 8888

# 启动命令，只能有一条命令。
ENTRYPOINT ./blog -f etc/blog.yaml

# 如果有多条命令，可以使用 ENTRYPOINT ["./entrypoint.sh"]
```

**entrypoint.sh**

entrypoint.sh是一个常见的脚本文件，用于作为 Docker 容器的入口点。在 Docker 容器启动时，该脚本会被执行。
当容器启动时，运行命令不止一条时，可以一起写在 entrypoint.sh 中。

### 如何启动一个容器？
**docker-compose.yaml**

docker-compose.yaml 是一个 YAML 格式的文件，用于定义和配置 Docker Compose 服务的组成和运行方式。
docker-compose.yaml 文件通常包含以下内容：
1. 服务定义：在 docker-compose.yaml 文件中，你可以定义多个服务（services），每个服务通常对应一个 Docker 容器。服务定义包括容器的镜像、端口映射、环境变量、依赖关系等信息。
2. 网络设置：你可以指定容器应该使用的网络，例如默认的 Docker 网络或自定义网络。
3. 卷挂载：你可以定义卷（volumes）以将主机文件系统上的目录或文件挂载到容器内部，从而实现数据持久化或共享。
4. 环境变量：你可以在 docker-compose.yaml 文件中设置环境变量，这些变量会传递给容器内运行的应用程序。
5. 构建配置：如果你希望构建自定义镜像而不是使用现有的镜像，你可以在 docker-compose.yaml 文件中指定构建上下文和构建参数。
6. 扩展性：docker-compose.yaml 文件支持使用 YAML 的特性，如继承和扩展。这使得你可以重用部分配置，并将其应用于多个服务或环境。


## 总结

在本文中，全面探讨了Docker的核心概念，涵盖了镜像、容器、仓库、卷以及Dockerfile的最佳实践。通过详实的示例代码，演示了从创建自定义镜像到构建多容器应用和管理Docker网络的方方面面。强调了Docker的灵活性和便携性，使得应用程序能够轻松、一致地在不同环境中运行。

介绍了Docker Compose的高级用法，包括服务扩展、网络设置和环境变量定义，大家能够更灵活地管理复杂应用的部署。此外，还提供了关于Docker安全性的最佳实践，强调了使用官方镜像、最小化镜像层数、定期更新镜像等关键策略，以确保容器化环境的安全性。

总体而言，通过对Docker基础概念的深入剖析和实用示例的呈现，大家可以建立起对Docker技术栈全貌的清晰认识。希望这篇文章能够成为初学者的入门指南，同时为有经验的开发者提供更深层次的实践经验，使其能够更好地应用Docker来实现高效、可靠和安全的容器化部署。
