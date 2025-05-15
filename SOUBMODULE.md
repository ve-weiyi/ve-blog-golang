
### 添加远程监控分支

git remote add upstream https://github.com/ve-weiyi/ve-blog-golang.git      
git fetch upstream

### submodule 
1. init 为项目添加submodule

```shell
git submodule init
git submodule add https://github.com/ve-weiyi/stomp.git  submodule/stomp
```

2. submodule update 克隆项目后拉取submodule
```shell
git submodule update
```

