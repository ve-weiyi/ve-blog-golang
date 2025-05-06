# Docker 环境搭建指南

## 1. Docker 简介

### 1.1 什么是 Docker？

Docker 是一个开源的容器化平台，它使开发者能够创建、部署和运行应用程序在隔离的环境中（称为"容器"
）。容器是一种轻量级的虚拟化技术，它共享宿主操作系统的内核，从而更加高效。

### 1.2 Docker 的核心优势

- **轻量高效**：容器只包含运行应用所需的代码和依赖
- **跨平台兼容**：实现"构建一次，随处运行"
- **版本控制**：支持镜像版本管理
- **隔离安全**：容器间相互隔离，防止资源干扰

### 1.3 通俗理解

想象你需要运行一个项目，传统方式需要：

- 安装 MySQL
- 安装 Redis
- 安装 Java 环境
- 配置各种环境变量
- 处理版本兼容问题

使用 Docker 后：

- 只需安装 Docker
- 拉取所需服务的镜像
- 运行容器即可
- 无需关心底层环境配置

**比喻**：Docker 就像一个公司，各种服务（MySQL、Redis等）就是不同的团队。需要时"招聘"团队，不需要时"解散"团队，不会对系统产生任何影响。

## 2. Docker 核心组件

### 2.1 镜像（Image）

- 只读的模板
- 包含应用程序运行所需的所有内容
- 可以基于其他镜像创建

### 2.2 容器（Container）

- 镜像的运行实例
- 独立的进程环境
- 包含应用程序和运行环境

### 2.3 仓库（Repository）

- 存储和分发镜像的地方
- 可以是公共的（如 Docker Hub）
- 也可以是私有的

## 3. 安装指南

### 3.1 CentOS 7 安装

```bash
# 更新系统
sudo yum update

# 安装必要工具
sudo yum install git

# 安装 Docker
sudo yum install docker

# 安装 Docker Compose
sudo yum install docker-compose

# 启动 Docker 服务
sudo systemctl start docker
sudo systemctl enable docker

# 验证安装
docker --version
```

### 3.2 Docker Desktop

Docker Desktop 是官方提供的桌面版，支持：

- Windows
- macOS
- 提供可视化界面
- 简化容器管理

## 4. 基本使用

### 4.1 常用命令

```bash
# 查看版本
docker --version

# 查看运行中的容器
docker ps

# 查看所有容器
docker ps -a

# 查看镜像
docker images

# 拉取镜像
docker pull mysql:latest

# 运行容器
docker run -d -p 3306:3306 mysql:latest
```

### 4.2 服务安装示例

#### MySQL 安装

```bash
# 拉取镜像
docker pull mysql:8.0

# 运行容器
docker run -d \
  --name mysql \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=your_password \
  -v /path/to/data:/var/lib/mysql \
  mysql:8.0
```

#### Redis 安装

```bash
# 拉取镜像
docker pull redis:latest

# 运行容器
docker run -d \
  --name redis \
  -p 6379:6379 \
  -v /path/to/data:/data \
  redis:latest
```

## 5. 最佳实践

### 5.1 数据持久化

- 使用数据卷挂载
- 避免容器内存储重要数据
- 定期备份数据

### 5.2 资源管理

- 限制容器资源使用
- 监控容器性能
- 及时清理无用资源

### 5.3 安全建议

- 使用官方镜像
- 定期更新镜像
- 限制容器权限
- 使用非 root 用户

## 6. 常见问题

### Q: 容器无法启动怎么办？

A:

1. 检查日志：`docker logs <container_id>`
2. 检查端口冲突
3. 检查资源限制
4. 检查配置文件

### Q: 如何备份容器数据？

A:

1. 使用数据卷
2. 定期导出数据
3. 使用备份工具

## 7. 参考资源

- [Docker 官方文档](https://docs.docker.com/)
- [Docker Hub](https://hub.docker.com/)
- [Docker 中文社区](https://www.docker.org.cn/)
