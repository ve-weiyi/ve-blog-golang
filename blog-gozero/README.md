# 项目介绍

## 1.项目结构

+ [.goctl](.goctl) goctl工具生成所需的模板，可以根据自己的项目进行定制
+ [internal](common) 相当于common，存放公共的代码
+ [service](service) 服务代码
    + [api](service/api) api服务
        + [admin](service/api/admin) 管理后台服务
        + [blog](service/api/blog) 博客服务
    + [model](service/model) 数据库操作层
    + [rpc](service/rpc) rpc服务
        + [blog](service/rpc/blog) 博客rpc服务

## 2.启动项目

### 准备工作：

1. 确认已经安装好golang环境。

2. 确保运行了 mysql、redis、rabbitmq 等服务。    
   如果没有运行过，可以在[data](../deploy/docker-compose/data)目录下输入命令`docker-compose up -d`快速启动这些服务。     
   Mysql需要初始化数据，在数据库执行[blog-veweiyi-init.sql](../blog-veweiyi-init.sql)、[blog-veweiyi-data.sql](../blog-veweiyi-data.sql)

3. 修改配置文件。在[blog-api](service/api/blog/etc)、[admin-api](service/api/admin/etc)、[blog-rpc](service/rpc/blog/etc)
   目录下

4. 【可选】确认启动etcd（使用直连rpc模式时不需要）

```sh
etcd
```

### 使用nacos配置启动

**该模式可用于生产模式**

1. 启动rpc服务

```sh
  go run service/rpc/blog/blog.go
```

2. 启动api服务

```sh
  go run service/api/blog/blog.go
```

```sh
  go run service/api/admin/admin.go
```

### 使用本地配置启动

**该模式适用于本地开发**
本地配置是 etc/xx.yaml 文件，启动时指定配置文件即可。

1. 启动rpc服务

```sh
  go run service/rpc/blog/blog.go -f service/rpc/blog/etc/blog-rpc.yaml
```

2. 启动api服务

```sh
  go run service/api/blog/blog.go -f service/api/blog/etc/blog-api.yaml
```

```sh
  go run service/api/admin/admin.go -f service/api/admin/etc/admin-api.yaml
```

## 3.部署服务

[docker-compose.yaml](docker-compose.yaml)

运行docker-compose.yaml文件

```sh
  docker-compose up -d -f docker-compose.yaml
```
