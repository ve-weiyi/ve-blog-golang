# gin框架运行的服务


### 1.安装项目运行环境
- 确保安装了 mysql、redis、rabbitmq 等服务。

如果你不想使用这些服务，可以在[server.go](server%2Fcmd%2Fserver.go) OnInitialize 函数中，注释掉相关初始化代码。

如果您不知道如何安装这些服务，可以参考[docker-compose](..%2Fdeploy%2Fdocker-compose)目录下的docker-compose.yaml文件，使用docker-compose启动这些服务。

如果您不会使用docker，可以参考[README.md](..%2Fdeploy%2Fdocker-compose%2FREADME.md)文档，或者其他相关博客。

- 初始化数据库,创建表和数据
```shell
go run main.go migrate --action=migrate \
--file=./blog-veweiyi.sql \
--host=127.0.0.1 \
--port=3307 \
--username=root \
--password=mysql7914 \
--name='blog-veweiyi'  
```

### 2.修改配置文件
在server目录下新增config.yaml文件，修改config.yaml配置文件中的配置信息，文件内容参考[config.default.yaml](server%2Fconfig.default.yaml)


### 3.项目本地启动

使用本地配置文件 config.yaml 启动项目
```shell
go run main.go api --config=./config.yaml 
```

使用nacos的配置文件启动项目
```shell
go run main.go api --use-nacos=true --n-ns=test
```

### 4.其他

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
