
<div align=center>
<img src="https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.20-blue"/>
<img src="https://img.shields.io/badge/gin-1.9.0-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-1.24.7-red"/>
<img src="https://img.shields.io/badge/redis-9.0.2-brightgreen"/>
<img src="https://img.shields.io/badge/swagger-v1.5.3-green"/>

</div>

联系方式：qq 791422171

简体中文

## 1. 基本介绍

### 1.1 项目介绍

ve-blog 是一个前后端分离的博客系统，项目采用了前后端分离的架构进行开发。前后端通过RESTful API进行数据交互。  
博客前台展示页面使用 vite4+vue3+vuetify 开发，  
博客后台管理系统使用 vite4+vue3+element-plus 开发，  
博客后端服务使用golang语言+ gin框架开发。  


ve-blog-golang项目是博客项目的后台服务支撑。

[博客前台项目地址](https://github.com/ve-weiyi/ve-blog-vite)
 
### 预览页面
![img.png](images%2Fimg.png)

![img_1.png](images%2Fimg_1.png)
## 项目启动

### 1.拉取submodule
```shell
git submodule update
```

### 2.拉取golang依赖库
```shell
cd server
go mod tidy
```

### 3.项目本地启动

复制 [config.default.yaml](server%2Fconfig.default.yaml) 文件，命名为config.yaml  
修改config.yaml配置文件中的配置信息，然后执行以下命令启动项目

```shell
cd server
go run main.go run --config=./conf/config.yaml 
```

### 4.更新swagger文档

格式化代码

```shell
go fmt ./...
```

格式化导入包

```shell
go get golang.org/x/tools/cmd/goimports
go install golang.org/x/tools/cmd/goimports
goimports -w .
./scripts/gofmt.sh
```

格式化swagger注解

```shell
swag fmt
```

```shell
swag init
```
