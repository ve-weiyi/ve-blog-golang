
# Nginx 实现 WSS 协议  
文章封面:  https://veport.oss-cn-beijing.aliyuncs.com/articles/f9fa18da262910eb13f802b003147915.jpg   
文章类型: 1   
文章分类: 技术   
文章标签: [学习 nginx]   
创建时间: 2022-02-09 23:52:33 +0800 CST   

文章内容:
# [WebSocket 结合 Nginx 实现域名及 WSS 协议访问](https://www.cnblogs.com/mafly/p/websocket.html)

## 简单了解一下 WebSocket

现在，很多网站为了实现推送技术，所用的技术都是轮询。轮询是在特定的的时间间隔（如每1秒），由浏览器对服务器发出HTTP请求，然后由服务器返回最新的数据给客户端的浏览器。这种传统的模式带来很明显的缺点，即浏览器需要不断的向服务器发出请求，然而HTTP请求可能包含较长的头部，其中真正有效的数据可能只是很小的一部分，显然这样会浪费很多的带宽等资源。
在这种情况下，HTML5定义了WebSocket协议，能更好的节省服务器资源和带宽，并且能够更实时地进行通讯。
WebSocket一种在单个 TCP 连接上进行全双工通讯的协议。使得客户端和服务器之间的数据交换变得更加简单，允许服务端主动向客户端推送数据。在 WebSocket API 中，浏览器和服务器只需要完成一次握手，两者之间就直接可以创建持久性的连接，并进行双向数据传输。

> 以上信息摘自维基百科（https://zh.wikipedia.org/wiki/WebSocket）

简单点说，WebSocket 就是减小客户端与服务器端建立连接的次数，减小系统资源开销，只需要一次 HTTP 握手，整个通讯过程是建立在一次连接/状态中，也就避免了HTTP的非状态性，服务端会一直与客户端保持连接，直到你关闭请求，同时由原本的客户端主动询问，转换为服务器有信息的时候推送。当然，它还能做实时通信、更好的二进制支持、支持扩展、更好的压缩效果等这些优点。

推荐一个知乎上叫 Ovear 的网友关于 WebSocket 原理的回答，嘻哈风格科普文，简直不要更赞了！地址：https://www.zhihu.com/question/20215561/answer/40316953

## ws 和 wss 又是什么鬼？

Websocket使用 `ws` 或 `wss` 的统一资源标志符，类似于 `HTTP` 或 `HTTPS`，其中 `wss` 表示在 TLS 之上的 Websocket ，相当于 HTTPS 了。如：

```bash
ws://example.com/chat
wss://example.com/chat
```

默认情况下，Websocket 的 ws 协议使用 80 端口；运行在TLS之上时，wss 协议默认使用 443 端口。其实说白了，wss 就是 ws 基于 SSL 的安全传输，与 HTTPS 一样样的道理。

如果你的网站是 HTTPS 协议的，那你就不能使用 `ws://` 了，浏览器会 block 掉连接，和 HTTPS 下不允许 HTTP 请求一样，如下图：
![image.png]( https://veport.oss-cn-beijing.aliyuncs.com/articles/374bca19e5fc1b45850781e4131d3d9c.png)
```rust
Mixed Content: The page at 'https://domain.com/' was loaded over HTTPS, but attempted to connect to the insecure WebSocket endpoint 'ws://x.x.x.x:xxxx/'. This request has been blocked; this endpoint must be available over WSS.
```

这种情况，毫无疑问我们就需要使用 `wss:\\` 安全协议了，我们是不是简单的把 `ws:\\` 改为 `wss:\\` 就行了？那试试呗。

改好了，报错啦！！！
![image.png]( https://veport.oss-cn-beijing.aliyuncs.com/articles/449049fb73844e5210bd5f20cfec9582.png)

```vbnet
VM512:35 WebSocket connection to 'wss://IP地址:端口号/websocket' failed: Error in connection establishment: net::ERR_SSL_PROTOCOL_ERROR
```

很明显 SSL 协议错误，说明就是证书问题了。记着，这时候我们一直拿的是 `IP地址 + 端口号` 这种方式连接 WebSocket 的，这根本就没有证书存在好么，况且生成环境你也要用 `IP地址 + 端口号` 这种方式连接 WebSocket 吗？肯定不行阿，要用域名方式连接 WebSocket 阿。

## Nginx 配置域名支持 WSS

不用废话，直接在配置 HTTPS 域名位置加入如下配置：

```bash
location /websocket {
    proxy_pass http://backend;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
}
```

接着拿域名再次连接试一下，不出意外会看 101 状态码：
![upgrade_101](https://images2015.cnblogs.com/blog/539095/201706/539095-20170622132017570-2009453161.png)

这样就完成了，在 HTTPPS 下以域名方式连接 WebSocket ，可以开心的玩耍了。

**稍微解释一下 Nginx 配置**
Nginx 自从 1.3 版本就开始支持 WebSocket 了，并且可以为 WebSocket 应用程序做反向代理和负载均衡。
WebSocket 和 HTTP 协议不同，但是 WebSocket 中的握手和 HTTP 中的握手兼容，它使用 HTTP 中的 Upgrade 协议头将连接从 HTTP 升级到 WebSocket，当客户端发过来一个 `Connection: Upgrade`请求头时，Nginx 是不知道的，所以，当 Nginx 代理服务器拦截到一个客户端发来的 `Upgrade` 请求时，需要显式来设置`Connection` 、`Upgrade` 头信息，并使用 101（交换协议）返回响应，在客户端和代理服务器、后端服务器之间建立隧道来支持 WebSocket。

当然，还需要注意一下，WebSockets 仍然受到 Nginx 缺省为60秒的 proxy_read_timeout 的影响。这意味着，如果你有一个程序使用了 WebSockets，但又可能超过60秒不发送任何数据的话，那你要么需要增加超时时间，要么实现一个 ping 的消息以保持联系。使用 ping 的解决方法有额外的好处，可以发现连接是否被意外关闭。

更具体文档详见 Nginx 官方文档：http://nginx.org/en/docs/http/websocket.html

## [总结一下](http://blog.mayongfa.cn/291.html)

这一篇文章主要了解一下 WebSocket 基本原理和一些使用用途，并解决在实际开发使用过程中遇到的坑，HTTPS 下使用 wss 协议的问题，以及配合 Nginx 使用域名方式建立连接，不使用 `IP地址 + 端口号` 连接 WebSocket，因为这种方式不够优雅。

原文链接:https://www.cnblogs.com/mafly/p/websocket.html
