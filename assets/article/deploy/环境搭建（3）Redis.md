# Redis 环境搭建指南

## 1. Redis 简介

### 1.1 什么是 Redis？

Redis 是一个开源的内存数据结构存储系统，具有以下特点：

- 高性能的键值对数据库
- 支持多种数据结构
- 持久化存储
- 主从复制
- 事务支持

### 1.2 Redis 的主要用途

- 缓存系统
- 会话管理
- 消息队列
- 排行榜系统
- 实时数据分析

## 2. 安装指南

### 2.1 使用 Docker 安装

创建 `docker-compose.yml` 文件：

```yaml
version: "3.9"

# 数据持久化配置
volumes:
  redis_data:  # 数据卷名称

services:
  redis:
    image: redis:7.0
    container_name: redis-server
    restart: always
    ports:
      - "6379:6379"  # 主机端口:容器端口
    environment:
      TZ: 'Asia/Shanghai'  # 时区设置
    volumes:
      - redis_data:/data  # 数据持久化
      - ./redis.conf:/usr/local/etc/redis/redis.conf  # 配置文件
    command: redis-server /usr/local/etc/redis/redis.conf  # 使用自定义配置
```

### 2.2 Redis 配置文件

创建 `redis.conf` 文件：

```conf
# 基本配置
bind 0.0.0.0
protected-mode yes
port 6379
tcp-backlog 511
timeout 0
tcp-keepalive 300

# 持久化配置
save 900 1
save 300 10
save 60 10000
stop-writes-on-bgsave-error yes
rdbcompression yes
rdbchecksum yes
dbfilename dump.rdb
dir /data

# 安全配置
requirepass your_password  # 设置密码

# 内存管理
maxmemory 1gb
maxmemory-policy allkeys-lru
```

### 2.3 启动 Redis

```bash
# 启动服务
docker-compose -f docker-compose.yaml up -d

# 查看运行状态
docker-compose ps

# 查看日志
docker-compose logs -f redis
```

## 3. 基本使用

### 3.1 连接 Redis

```bash
# 进入容器
docker exec -it redis-server redis-cli

# 使用密码连接
docker exec -it redis-server redis-cli -a your_password
```

### 3.2 常用命令

```bash
# 键值操作
SET key value
GET key
DEL key
EXISTS key

# 列表操作
LPUSH list value
RPUSH list value
LPOP list
RPOP list

# 集合操作
SADD set member
SMEMBERS set
SREM set member

# 哈希操作
HSET hash field value
HGET hash field
HDEL hash field
```

## 4. 数据持久化

### 4.1 RDB 持久化

- 自动保存：根据配置的时间间隔保存
- 手动保存：使用 SAVE 或 BGSAVE 命令
- 恢复数据：重启 Redis 时自动加载

### 4.2 AOF 持久化

在 `redis.conf` 中添加：

```conf
appendonly yes
appendfilename "appendonly.aof"
appendfsync everysec
```

## 5. 性能优化

### 5.1 配置优化

```conf
# 内存管理
maxmemory 1gb
maxmemory-policy allkeys-lru

# 连接管理
maxclients 10000
timeout 0

# 持久化优化
save 900 1
save 300 10
save 60 10000
```

### 5.2 监控建议

- 使用 INFO 命令监控状态
- 定期检查内存使用
- 监控连接数
- 检查持久化状态

## 6. 安全建议

### 6.1 基础安全

- 设置访问密码
- 限制网络访问
- 使用防火墙
- 定期更新密码

### 6.2 数据安全

- 启用持久化
- 定期备份数据
- 监控异常访问
- 实施访问控制

## 7. 常见问题

### Q: 连接超时怎么办？

A:

1. 检查防火墙设置
2. 确认端口映射
3. 验证密码是否正确
4. 检查网络连接

### Q: 内存不足怎么办？

A:

1. 调整 maxmemory 配置
2. 选择合适的淘汰策略
3. 优化数据结构使用
4. 考虑集群部署

## 8. 参考资源

- [Redis 官方文档](https://redis.io/documentation)
- [Redis 中文网](http://www.redis.cn/)
- [Docker Redis 镜像](https://hub.docker.com/_/redis)
