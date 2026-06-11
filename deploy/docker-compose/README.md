# Docker Compose 服务配置

## 服务列表

- `mysql/mysql.yaml` - MySQL 数据库
- `postgres/postgres.yaml` - PostgreSQL 数据库
- `redis/redis.yaml` - Redis 缓存
- `rabbitmq/rabbitmq.yaml` - RabbitMQ 消息队列
- `kafka/kafka-single.yaml` - Kafka 单节点
- `kafka/kafka-mutil.yaml` - Kafka 三节点集群

## 使用方法

### 启动单个服务
```bash
# MySQL
docker-compose -f mysql/mysql.yaml up -d

# PostgreSQL
docker-compose -f postgres/postgres.yaml up -d

# Redis
docker-compose -f redis/redis.yaml up -d

# RabbitMQ
docker-compose -f rabbitmq/rabbitmq.yaml up -d

# Kafka 单节点
docker-compose -f kafka/kafka-single.yaml up -d

# Kafka 三节点集群
docker-compose -f kafka/kafka-mutil.yaml up -d
```

### 启动多个服务
```bash
docker-compose -f mysql/mysql.yaml -f redis/redis.yaml up -d
```

### 停止服务
```bash
docker-compose -f mysql/mysql.yaml down
```

## 服务连接信息

### MySQL
- Host: localhost
- Port: 3306
- Username: veweiyi
- Password: mysql7914

### PostgreSQL
- Host: localhost
- Port: 5432
- Username: postgres
- Password: postgres

### Redis
- Host: localhost
- Port: 6379
- Password: redis7914

### RabbitMQ
- AMQP Port: 5672
- Management UI: http://localhost:15672
- Username: veweiyi
- Password: rabbitmq7914

### Kafka 单节点
- Bootstrap: localhost:9092
- Kafka UI: http://localhost:9080

### Kafka 三节点集群
- Bootstrap: localhost:19094, localhost:29094, localhost:39094
- Schema Registry: http://localhost:8085
- Kafka UI: http://localhost:8099
