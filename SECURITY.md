# 项目启动

### 分支合并规则

仅允许创建以下方向的合并请求：     
功能分支 feature/*/* -> 发布分支 release/*  
发布分支 release/* -> 主干分支 master   
热修复分支 hotfix/* -> 发布分支 release/*    
热修复分支 hotfix/* -> 主干分支 master

### 添加远程监控分支

git remote add upstream https://github.com/ve-weiyi/ve-blog-golang.git      
git fetch upstream

### submodule init 为项目添加submodule

```shell
cd  submodule
git submodule init
git submodule add https://github.com/ve-weiyi/go-sdk.git  submodule/go-sdk
```

submodule update 克隆项目后拉取submodule
```shell
git submodule update
```

