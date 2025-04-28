# RabbitMQ 环境搭建指南

## 1. RabbitMQ 简介

### 1.1 什么是 RabbitMQ？

RabbitMQ 是一个开源的消息代理软件，具有以下特点：

- 支持多种消息协议
- 高可用性
- 消息持久化
- 集群支持
- 灵活的路由机制

### 1.2 RabbitMQ 的主要用途

- 异步消息处理
- 应用解耦
- 流量削峰
- 日志处理
- 任务队列

## 2. 安装指南

### 2.1 使用 Docker 安装

创建 `docker-compose.yml` 文件：

```yaml
version: "3.9"

# 数据持久化配置
volumes:
  rabbitmq_data:  # 数据卷名称

services:
  rabbitmq:
    image: rabbitmq:3.12-management
    container_name: rabbitmq-server
    restart: always
    ports:
      - "5672:5672"   # AMQP 协议端口
      - "15672:15672" # Web 管理界面端口
    environment:
      RABBITMQ_DEFAULT_USER: admin
      RABBITMQ_DEFAULT_PASS: admin123
      TZ: 'Asia/Shanghai'
    volumes:
      - rabbitmq_data:/var/lib/rabbitmq
```

### 2.2 启动 RabbitMQ

```bash
# 启动服务
docker-compose -f docker-compose.yaml up -d

# 查看运行状态
docker-compose ps

# 查看日志
docker-compose logs -f rabbitmq
```

## 3. 基本使用

### 3.1 访问管理界面

- 地址：http://localhost:15672
- 用户名：admin
- 密码：admin123

### 3.2 常用概念

- **Exchange**：消息交换机，决定消息路由规则
- **Queue**：消息队列，存储消息
- **Binding**：绑定关系，连接交换机和队列
- **Routing Key**：路由键，用于消息路由

### 3.3 消息模式

1. **简单队列**
    - 一个生产者
    - 一个消费者
    - 一个队列

2. **工作队列**
    - 一个生产者
    - 多个消费者
    - 一个队列

3. **发布/订阅**
    - 一个生产者
    - 多个消费者
    - 一个交换机
    - 多个队列

4. **路由模式**
    - 一个生产者
    - 多个消费者
    - 一个交换机
    - 多个队列
    - 使用路由键

## 4. 配置优化

### 4.1 性能优化

```yaml
services:
  rabbitmq:
    # ... 其他配置 ...
    environment:
      RABBITMQ_SERVER_ADDITIONAL_ERL_ARGS: "-rabbitmq_management listener [{port,15672},{ssl,false}]"
      RABBITMQ_ERLANG_COOKIE: "SWQOKODSQALRPCLNMEQG"
      RABBITMQ_DEFAULT_USER: admin
      RABBITMQ_DEFAULT_PASS: admin123
```

### 4.2 集群配置

```yaml
version: "3.9"

services:
  rabbitmq1:
    image: rabbitmq:3.12-management
    hostname: rabbitmq1
    environment:
      RABBITMQ_ERLANG_COOKIE: "SWQOKODSQALRPCLNMEQG"
      RABBITMQ_NODENAME: rabbit@rabbitmq1
    ports:
      - "5672:5672"
      - "15672:15672"

  rabbitmq2:
    image: rabbitmq:3.12-management
    hostname: rabbitmq2
    environment:
      RABBITMQ_ERLANG_COOKIE: "SWQOKODSQALRPCLNMEQG"
      RABBITMQ_NODENAME: rabbit@rabbitmq2
    ports:
      - "5673:5672"
      - "15673:15672"
```

## 5. 监控与维护

### 5.1 监控指标

- 队列长度
- 消息处理速率
- 连接数
- 内存使用
- 磁盘使用

### 5.2 维护建议

- 定期清理过期消息
- 监控队列积压
- 检查节点状态
- 备份配置数据

## 6. 安全建议

### 6.1 基础安全

- 修改默认密码
- 限制网络访问
- 使用 SSL/TLS
- 定期更新密码

### 6.2 数据安全

- 启用消息持久化
- 配置镜像队列
- 实施访问控制
- 监控异常访问

## 7. 常见问题

### Q: 消息丢失怎么办？

A:

1. 启用消息持久化
2. 使用确认机制
3. 配置镜像队列
4. 实施监控告警

### Q: 性能下降怎么办？

A:

1. 优化队列配置
2. 增加节点数量
3. 调整内存设置
4. 检查网络状况

## 8. 参考资源

- [RabbitMQ 官方文档](https://www.rabbitmq.com/documentation.html)
- [RabbitMQ 中文文档](https://rabbitmq.mr-ping.com/)
- [Docker RabbitMQ 镜像](https://hub.docker.com/_/rabbitmq)
