# 非常好用的工具包

## [quickstart]() 快速代码构建器

根据数据库模型生成 router、controller、service、repository、model 代码，并自动注册到 gin 框架中。

### 1. 介绍

inject 使用ast在文件中指定位置注入代码
invent 使用template生成代码文件

### 2. 使用

cobra-cli init
cobra-cli add entity
cobra-cli add repository
cobra-cli add service
cobra-cli add controller
cobra-cli add router

1. 生成entity文件

```shell
go run main.go entity -f=./config.yaml -t=./entity.tpl -o=./
```

2. 生成repository文件

```shell
go run main.go quickstart repository -f=./config.yaml -t=./repository.tpl -o=./
```
