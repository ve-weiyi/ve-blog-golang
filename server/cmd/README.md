### cobra使用

1. 安装cobra

```shell
go get -u github.com/spf13/cobra/cobra
```

2. 初始化项目

```shell
cobra-cli init
cobra-cli init --author "791422171@qq.com"
cobra-cli init --license apache
```

3. 添加命令 migrate （例）

```shell
cobra-cli add version
cobra-cli add migrate
```

4. 运行命令 migrate（例）

```shell
go run main.go migrate -h
go run main.go migrate --help
go run main.go migrate --action=reset
```
