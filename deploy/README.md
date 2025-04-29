## 服务说明

对于本地运行的服务，基础版功能只需要mysql和redis、rabbitmq服务，生产环境需要使用nacos和etcd服务。

### 基础服务 server.com

| 服务名称  | 用途     | 是否必需 | 说明           | docker-compose                                            |
|-------|--------|------|--------------|-----------------------------------------------------------|
| RPC服务 | 内部服务调用 | 必需   | 提供RPC访问接口    | [Dockerfile-blog-rpc](../blog-gozero/Dockerfile-blog-rpc) |
| API服务 | 外部接口访问 | 必需   | 提供HTTP API接口 | [Dockerfile-blog-api](../blog-gozero/Dockerfile-blog-api) |

### 数据服务 (data-server.com)

| 服务名称          | 用途     | 是否必需 | 说明         | docker-compose                                                       |
|---------------|--------|------|------------|----------------------------------------------------------------------|
| MySQL         | 关系型数据库 | 必需   | 存储核心业务数据   | [docker-compose.yaml](docker-compose/data/docker-compose.yaml)       |
| Redis         | 缓存服务   | 必需   | 存储文章访问量    | [docker-compose.yaml](docker-compose/data/docker-compose.yaml)       |
| RabbitMQ      | 消息队列   | 必需   | 异步处理邮件发送数据 | [docker-compose.yaml](docker-compose/data/docker-compose.yaml)       |
| Kafka         | 消息队列   | 可选   | 异步处理页面访问数据 | [docker-compose.yaml](docker-compose/data/kafka/docker-compose.yaml) |
| Elasticsearch | 搜索引擎   | 可选   | 搜索文章内容(未做) |                                                                      |

### 环境服务 (environment-server.com)

| 服务名称    | 用途         | 是否必需 | 说明          |                                                                             |
|---------|------------|------|-------------|-----------------------------------------------------------------------------|
| Nacos   | 服务发现和配置管理  | 生产必需 | 获取配置文件      | [docker-compose.yaml](docker-compose/environment/nacos/docker-compose.yaml) |
| Etcd    | 分布式键值存储    | 可选   | 获取RPC服务注册地址 | [docker-compose.yaml](docker-compose/environment/etcd/docker-compose.yaml)  |
| GitLab  | 代码仓库管理     | 可选   | 代码版本控制      |                                                                             |
| Jenkins | CI/CD自动化部署 | 可选   | 持续集成和部署     |                                                                             |
| Harbor  | Docker镜像仓库 | 可选   | 容器镜像管理      |                                                                             |

### 网关服务 (nginx-gateway.com) - 进阶

- Nginx 反向代理和负载均衡
- SSL证书管理
- 流量控制

### K8s集群服务 -进阶

- 应用服务部署
- 服务编排
- 容器管理

## 部署要求

1. 所有服务器需要预先安装 Docker 和 Docker Compose
2. 确保服务器之间网络互通
3. 建议使用内网环境部署，提高安全性
4. 所有服务建议配置监控和日志收集

## 运行docker-compose

```bash
# 对应文件目录下
docker-compose -f docker-compose.yaml up -d
```


