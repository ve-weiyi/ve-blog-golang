# 项目启动

## Reporting a Vulnerability

Please report security issues to qimiaojiangjizhao@gmail.com

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

