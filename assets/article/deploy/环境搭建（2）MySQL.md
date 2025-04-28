# MySQL 环境搭建指南

## 1. MySQL 简介

### 1.1 什么是 MySQL？

MySQL 是一个开源的关系型数据库管理系统（RDBMS），具有以下特点：

- 支持标准 SQL 语言
- 高性能、高可靠性
- 支持多种存储引擎
- 跨平台支持
- 丰富的功能特性

### 1.2 MySQL 的主要用途

- 用户数据存储
- 日志数据管理
- 配置信息存储
- 业务数据存储
- 数据分析支持

## 2. 安装指南

### 2.1 使用 Docker 安装

创建 `docker-compose.yml` 文件：

```yaml
version: "3.9"

# 数据持久化配置
volumes:
  mysql_data:  # 数据卷名称

services:
  mysql:
    image: mysql:8.0.34
    container_name: mysql-server
    restart: always
    ports:
      - "3306:3306"  # 主机端口:容器端口
    environment:
      MYSQL_ROOT_PASSWORD: 'mysql7914'  # root 用户密码
      MYSQL_DATABASE: 'blog-veweiyi'    # 初始数据库名
      MYSQL_USER: 'veweiyi'             # 初始用户名
      MYSQL_PASSWORD: 'mysql7914'       # 初始用户密码
      TZ: 'Asia/Shanghai'               # 时区设置
    volumes:
      - mysql_data:/var/lib/mysql       # 数据持久化
    command:
      --character-set-server=utf8mb4    # 字符集设置
      --collation-server=utf8mb4_general_ci
```

### 2.2 启动 MySQL

```bash
# 启动服务
docker-compose -f docker-compose.yaml up -d

# 查看运行状态
docker-compose ps

# 查看日志
docker-compose logs -f mysql
```

## 3. 基本使用

### 3.1 连接数据库

```bash
# 进入容器
docker exec -it mysql-server mysql -u root -p

# 使用初始用户连接
docker exec -it mysql-server mysql -u veweiyi -p
```

### 3.2 常用命令

```sql
-- 查看所有数据库
SHOW DATABASES;

-- 使用数据库
USE blog-veweiyi;

-- 查看所有表
SHOW TABLES;

-- 创建新用户
CREATE USER 'newuser'@'%' IDENTIFIED BY 'password';

-- 授予权限
GRANT ALL PRIVILEGES ON blog-veweiyi.* TO 'newuser'@'%';
```

## 4. 数据备份与恢复

### 4.1 备份数据

```bash
# 备份整个数据库
docker exec mysql-server mysqldump -u root -p blog-veweiyi > backup.sql

# 备份特定表
docker exec mysql-server mysqldump -u root -p blog-veweiyi table1 table2 > backup.sql
```

### 4.2 恢复数据

```bash
# 恢复数据库
docker exec -i mysql-server mysql -u root -p blog-veweiyi < backup.sql
```

## 5. 性能优化

### 5.1 配置优化

```yaml
# 在 docker-compose.yml 中添加以下配置
services:
  mysql:
    # ... 其他配置 ...
    command:
      --character-set-server=utf8mb4
      --collation-server=utf8mb4_general_ci
      --max_connections=1000
      --innodb_buffer_pool_size=1G
```

### 5.2 监控建议

- 使用 `SHOW STATUS` 监控性能
- 定期检查慢查询日志
- 监控连接数和资源使用

## 6. 安全建议

### 6.1 基础安全

- 修改默认 root 密码
- 限制远程访问
- 使用强密码策略
- 定期更新密码

### 6.2 数据安全

- 定期备份数据
- 使用数据加密
- 实施访问控制
- 监控异常访问

## 7. 常见问题

### Q: 连接被拒绝怎么办？

A:

1. 检查端口映射是否正确
2. 确认用户权限设置
3. 检查防火墙设置
4. 验证密码是否正确

### Q: 数据丢失怎么办？

A:

1. 使用备份恢复数据
2. 检查数据卷挂载
3. 查看错误日志
4. 联系技术支持

## 8. 参考资源

- [MySQL 官方文档](https://dev.mysql.com/doc/)
- [MySQL 中文社区](https://www.mysql.com/cn/)
- [Docker MySQL 镜像](https://hub.docker.com/_/mysql)
