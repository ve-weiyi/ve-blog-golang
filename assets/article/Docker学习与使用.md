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

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .
RUN go mod tidy
RUN go build -o /app/rpc/blog/blog service/rpc/blog/blog.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app/rpc/blog
COPY --from=builder /app/rpc/blog/blog /app/rpc/blog/blog

CMD ["./blog"]

# 启动命令，只能有一条命令。
#ENTRYPOINT ./blog -f etc/blog.yaml

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

### docker-compose.yaml 结构

1. network

```yaml
networks:
  my_bridge_network:
    driver: bridge
  my_overlay_network:
    driver: overlay
  my_host_network:
    driver: host

```

| 网络模式       | 命令指定方式            | 描述                                                  | 理解        |
|------------|-------------------|-----------------------------------------------------|-----------|
| bridge	    | –network bridge	  | 为每一个容器分配、设置 ip ，并将容器连接到 docker0 虚拟网桥上，虚拟网桥，默认为该网络模式 | 	一人一个     |
| host	      | –network host	    | 容器不会创建自己的网卡，配置 ip 等，而是使用宿主机的 ip 和端口	                | 多人一个,用宿主的 |
| container	 | –network 容器名称或id	 | 新创建的容器不会创建自己的网卡和配置自己的ip，而是和一个指定的容器共享ip、端口范围等	       | 自己没有，用别人的 |
| none	      | –network none	    | 容器有独立的Network namespace，但并没有对其进行任何网络设置，如网桥，ip等	     | 有，但是空     |

2. volumes

```yaml
services:
  app:
    image: my-app
    volumes:
      - data_volume:/var/lib/data_volume
      - /var/log/app_logs
      - /var/lib/host_data:/var/lib/container_data

volumes:
  data_volume1:                # 本地自动创建的命名卷
  data_volume2:
    driver: local              # 使用本地驱动创建的命名卷。和 data_volume1 的效果类似，但显式声明了存储驱动
  data_volume3:
    external: true             # 预先存在的外部卷
```

| 类型                     | 定义方式        | 存储位置           | 适用场景                           |
|------------------------|-------------|----------------|--------------------------------|
| (挂载卷)Host Volumes      | /主机路径:/容器路径 | 主机指定路径         | 开发环境、需要实时修改主机文件或配置文件。          |
| (命名卷)Named Volumes     | 卷名:/容器路径    | Docker 管理路径    | 生产环境持久化存储，数据与容器解耦，适合长期存储和数据共享。 |
| (匿名卷)Anonymous Volumes | /容器路径       | Docker 自动创建匿名卷 | 临时数据存储，生命周期短的数据。               |

## 其他

1.

2.

Q：在docker容器内使用127.0.0.1 和 0.0.0.0有什么区别？
A：在容器中，使用127.0.0.1启动的rpc，无法被宿主机的api连接到。

在 Docker 容器内使用 127.0.0.1 和 0.0.0.0 之间有着重要的区别，这与容器的网络环境和网络配置有关。

127.0.0.1：
在 Docker 容器内，127.0.0.1 仍然表示本地主机回环地址（loopback address），即容器自身。
当容器内的应用程序尝试连接到 127.0.0.1 时，它实际上是在尝试与容器内部的另一个进程进行通信，而不是宿主机的进程或其他容器的进程。

0.0.0.0：
在 Docker 容器内，0.0.0.0 表示容器的所有网络接口，即容器的整个网络命名空间。
当容器内的应用程序将服务绑定到 0.0.0.0 上时，它将在容器的所有网络接口上监听连接，包括容器自身和宿主机的网络接口。
如果你在容器内的应用程序中绑定服务到 0.0.0.0，则可以从容器外部访问该服务，前提是 Docker 主机已将容器端口映射到宿主机的端口。

因此，要根据具体的需求选择使用 127.0.0.1 还是 0.0.0.0。如果你希望容器内的服务只能在容器内部访问，应使用
127.0.0.1。如果你希望从容器外部访问容器内的服务，则应将服务绑定到 0.0.0.0。

## 总结

在本文中，全面探讨了Docker的核心概念，涵盖了镜像、容器、仓库、卷以及Dockerfile的最佳实践。通过详实的示例代码，演示了从创建自定义镜像到构建多容器应用和管理Docker网络的方方面面。强调了Docker的灵活性和便携性，使得应用程序能够轻松、一致地在不同环境中运行。

介绍了Docker Compose的高级用法，包括服务扩展、网络设置和环境变量定义，大家能够更灵活地管理复杂应用的部署。此外，还提供了关于Docker安全性的最佳实践，强调了使用官方镜像、最小化镜像层数、定期更新镜像等关键策略，以确保容器化环境的安全性。

总体而言，通过对Docker基础概念的深入剖析和实用示例的呈现，大家可以建立起对Docker技术栈全貌的清晰认识。希望这篇文章能够成为初学者的入门指南，同时为有经验的开发者提供更深层次的实践经验，使其能够更好地应用Docker来实现高效、可靠和安全的容器化部署。

## Docker的常用命令
```shell
# 删除所有未使用的镜像、容器、卷和网络
docker system prune

# 清理容器
docker container prune

# 清理镜像
docker image prune

# 清除构建缓存
docker builder prune

# 搜索镜像
docker search nginx

# 拉取镜像
docker pull nginx:latest

# 使用Dockerfile构建镜像
docker build -t my-image .
```

## 参考文档
[Docker基础（三）镜像制作](https://blog.csdn.net/du_zhe_/article/details/132357258)
