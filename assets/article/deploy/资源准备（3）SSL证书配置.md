# SSL证书配置

本文将介绍SSL证书的选择、申请流程、安装配置以及常见问题的解决方法。

SSl证书的作用是让你们网站可以被https访问，浏览器上不会出现安全提示。

## 证书类型选择

### 免费证书

1. Let's Encrypt
    - 优势：免费、自动续期
    - 有效期：90天
    - 适合：个人网站

2. 服务商免费证书
    - 阿里云免费证书
    - 腾讯云免费证书
    - 有效期：1年

### 付费证书

1. DigiCert
    - 优势：企业级安全
    - 价格：较高
    - 适合：企业网站

2. GeoTrust
    - 优势：性价比高
    - 价格：中等
    - 适合：商业网站

3. 通配符证书
    - 优势：支持子域名
    - 价格：较高
    - 适合：多子域名网站

## 证书申请流程

### 腾讯云申请免费证书

进入SSL证书控制台
![img_3.png](https://static.veweiyi.cn/article/deploy/img_3.png)

根据自身经济实力选择对应证书类型
![img_4.png](https://static.veweiyi.cn/article/deploy/img_4.png)

填写如下配置，提交申请
![img_5.png](https://static.veweiyi.cn/article/deploy/img_5.png)

下载证书
![img_7.png](https://static.veweiyi.cn/article/deploy/img_7.png)

## 证书安装配置

### 安装证书

复制证书到服务器 /www/ssl 目录下,并且解压
![img_8.png](https://static.veweiyi.cn/article/deploy/img_8.png)

### Nginx配置

```nginx
    server {
        listen 443 ssl;
        server_name blog.veweiyi.cn;

        # 配置 SSL 证书路径
        ssl_certificate     /www/ssl/blog.veweiyi.cn_nginx/blog.veweiyi.cn_bundle.crt;
        ssl_certificate_key /www/ssl/blog.veweiyi.cn_nginx/blog.veweiyi.cn.key;

        # 博客前台接口映射
        location ^~ /api {
            proxy_pass http://localhost:9090;
            proxy_set_header Host $host;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection $connection_upgrade;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }

        # 博客后台网站映射
        location ^~ / {
            proxy_pass http://localhost:9420;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
    }
```

## 常见问题

### 1. 证书过期

- 检查自动续期配置
- 手动更新证书
- 更新证书文件

### 2. 混合内容

- 检查资源引用
- 更新HTTP链接
- 配置CSP策略

### 3. 浏览器警告

- 检查证书链
- 验证域名匹配
- 检查时间同步 
