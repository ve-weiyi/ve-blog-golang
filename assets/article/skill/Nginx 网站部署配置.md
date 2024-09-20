
# Nginx vue打包发布刷新页面空白问题  
文章封面:  https://veport.oss-cn-beijing.aliyuncs.com/articles/f9e9490619d664167657258c21db086a.jpeg   
文章类型: 1   
文章分类: 网站搭建   
文章标签: [vue nginx]   
创建时间: 2022-01-22 23:55:55 +0800 CST   

文章内容:
最近遇到一个问题，网站页面可以通过路由点击进入，但是直接在浏览器url处刷新页面之后会是空白页。

原以为是nginx反向代理出现了问题，然后改了又改。后来又仔细想了想，会不会是vue的history模式出现问题。但是测试了之后发现，vue在本地运行时页面的跳转是正常的，也可以通过url刷新访问到页面。

最后才醒悟，应该是vue打包之后环境问题。在一篇文章发现了这个bug的解决方案。

来看看怎么说：

```JavaScript
因为我们的应用是单页客户端应用，当使用 history 模式时，URL 就像正常的 url，可以直接访问http://www.xxx.com/user/id，但是因为vue-router设置的路径不是真实存在的路径，所以刷新就会返回404错误。

想要history模式正常访问，还需要后台配置支持。要在服务端增加一个覆盖所有情况的候选资源：如果 URL 匹配不到任何静态资源，则应该返回同一个 index.html 页面，这个页面就是你 app 依赖的页面。

也就是在服务端修改404错误页面的配置路径，让其指向到index.html。

**拓展**
部署后，当访问一些页面的时候，报错 Uncaught SyntaxError: Unexpected token ‘＜’。

解决方案：

刚开始publicPath是’./’，需要改成’/’，即在vue.config.js中修改配置

module.exports = {
  ...
  publicPath: '/',
}

```
重要的是这句话”刚开始publicPath是’./’，需要改成’/’，即在vue.config.js中修改配置“,于是我打包时把路径改为‘/’。

但是还是不行，浏览器页面报错找到“https://ve77.cn/static/***.css"文件。

看到这个问题我就想明白了,因为我是放在/blog路径下的,使用此处改成‘/blog’才对。果然，发布之后可以通过url访问了！！！


分享一下自己的nginx配置
```xml
#user  nobody;
worker_processes  2;
#日志位置和日志级别
error_log /usr/local/webserver/nginx/logs/nginx_error.txt;
pid /usr/local/webserver/nginx/nginx.pid;

#pid        logs/nginx.pid;

events {
    worker_connections  1024;
}


http {
    include       mime.types;
    default_type  application/octet-stream;
    sendfile        on;
    keepalive_timeout  65;

    client_max_body_size     50m;
    client_body_buffer_size  10m;
    client_header_timeout    1m;
    client_body_timeout      1m;

    gzip on;
    gzip_min_length  1k;
    gzip_buffers     4 16k;
    gzip_comp_level  4;
    gzip_types text/plain application/javascript application/x-javascript text/css application/xml text/javascript application/x-httpd-php image/jpeg image/gif image/png;
    gzip_vary on;

    access_log  logs/access.log ;

    # http默认端口，转发到https
    # 接口只能http访问，此时被重定向了。
    server {
        listen       80;
        server_name  ve77.com  www.ve77.com  static.ve77.com  www.static.ve77.com;
        rewrite ^(.*) https://$host$request_uri;
        access_log  logs/host80.access.txt ;
    }

    # HTTPS server
    server {
        #SSL 访问端口号为 443
        listen       443 ssl;
        #填写绑定证书的域名
        server_name  ve77.cn;  #填写绑定证书的域名
        #证书文件名称
        ssl_certificate      /jks/ve77.cn_nginx/ve77.cn_bundle.crt;
        #私钥文件名称
        ssl_certificate_key  /jks/ve77.cn_nginx/ve77.cn.key;
        #请按照以下协议配置
        ssl_protocols TLSv1.2 TLSv1.3;
        #请按照以下套件配置，配置加密套件，写法遵循 openssl 标准。
        ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;
        ssl_prefer_server_ciphers  on;

        ssl_session_cache    shared:SSL:1m;
        ssl_session_timeout  5m;
        access_log  logs/host443.access.txt ;

        #末位别加/ 要不然会路径错误
        location  /blog {
            root   /usr/local/vue/ ; #站点目录+/blog
            index  index.html index.htm;
            try_files  $uri $uri/ /blog/index.html;
        }

        location /admin {
            root   /usr/local/vue/;
            index  index.html index.htm;
            try_files $uri $uri/ /admin/index.html;
        }
        #重定向
        location ^~ /api {
            proxy_pass https://ve77.cn:8088;
            proxy_set_header Host $host:8088; #这里是重点,这样配置才不会丢失端口
            proxy_set_header   X-Real-IP        $remote_addr;
            proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
        }

        location ^~ /websocket {
            proxy_pass https://ve77.cn:8088/api;
            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection "Upgrade";
            proxy_set_header Host $host:8088;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }

    }

            #重点 start
            # 这一点写location里面是一样的 但是提在外面就不用重复写了
            # 如果内网nginx监听端口与外网访问的端口不一致 需要配置成这样
            #	proxy_set_header Host $host:$server_port;
            # 	proxy_set_header X-Real-IP $remote_addr;
            #	proxy_set_header REMOTE-HOST $remote_addr;
            # 	proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            #重点 end
            #在 Nginx 根目录下，通过执行以下命令验证配置文件问题。
            #./sbin/nginx -t
}

```

常用的nginx配置命令
```xml
重新加载/重启 nginx服务

/usr/local/webserver/nginx/sbin/nginx -s reload

/usr/local/webserver/nginx/sbin/nginx -s reopen

 

验证nginx配置文件

/usr/local/webserver/nginx/sbin/nginx -t
```
参考文献： 

[vue history模型下的问题](https://www.jb51.net/article/119075.htm)

[Vue路由为history模式的nginx配置](https://blog.csdn.net/kiscon/article/details/115416832)

