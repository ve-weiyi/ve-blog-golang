
# Tomcat-使用https访问  
文章封面:  https://veport.oss-cn-beijing.aliyuncs.com/articles/07afc2963f27e63239e50bc65bed6a6f.jpg   
文章类型: 1   
文章分类: 技术   
文章标签: [网站 tomcat]   
创建时间: 2022-01-22 13:29:27 +0800 CST   

文章内容:

# 前言
tomcat配置好了以后默认是使用8080端口访问的，也就是需要在使用"域名.com:8080"才能访问。这篇总结一下如何修改tomcat配置，使可以用"http://域名.com"或"https://域名.com" 访问。
## 前期准备
环境配置：
 1. 腾讯云轻量应用服务器: CentOS 8.2 64bit
 2. 远程访问推荐使用图形化界面(Mac 建议Royal TSX,Windows建议Mobaxterm)
 3. Tomcat 10.0.4 ；
 4. Java 1.8 ；
 
 前提条件:

 配置访问80端口即"http://域名.com"不需要证书
 配置访问443端口即"https://域名.com" 需要SSL证书，证书可以从你购买服务器的运营商那里获取

## 具体操作步骤
话不多说，直接进入正题
编辑在 /usr/tomcat/*/conf 目录(这个目录是你安装tomcat的目录)下的 server.xml 文件。添加如下内容：
```xml
// An highlighted block
<Connector port="443" protocol="HTTP/1.1" SSLEnabled="true"
  maxThreads="150" scheme="https" secure="true"
#证书保存的路径
  keystoreFile="/usr/*/conf/域名.com.jks" 
#密钥库密码
  keystorePass="******"
  clientAuth="false"/>
```
详细 server.xml 文件和一些参数解释如下(可以直接复制过去)：
```xml
<?xml version="1.0" encoding="UTF-8"?>
<!--
Server 根元素，创建⼀个Server实例，⼦标签有 Listener、GlobalNamingResources、Service
port：关闭服务器的监听端⼝
shutdown：关闭服务器的指令字符串
-->
<Server port="8005" shutdown="SHUTDOWN">

    <!-- 创建 5 个监听器  start -->
    <!-- 以⽇志形式输出服务器 、操作系统、JVM的版本信息 -->
    <Listener className="org.apache.catalina.startup.VersionLoggerListener"/>
    <!-- 加载（服务器启动） 和 销毁 （服务器停⽌） APR。 如果找不到APR库， 则会输出⽇志， 并不影响 Tomcat启动 -->
    <Listener className="org.apache.catalina.core.AprLifecycleListener" SSLEngine="on"/>
    <!-- 避免JRE内存泄漏问题 -->
    <Listener className="org.apache.catalina.core.JreMemoryLeakPreventionListener"/>
    <!-- 加载（服务器启动） 和 销毁（服务器停⽌） 全局命名服务 -->
    <Listener className="org.apache.catalina.mbeans.GlobalResourcesLifecycleListener"/>
    <!-- 在Context停⽌时重建 Executor 池中的线程， 以避免ThreadLocal 相关的内存泄漏 -->
    <Listener className="org.apache.catalina.core.ThreadLocalLeakPreventionListener"/>
    <!-- 创建 5 个监听器  end -->


    <!--
         定义服务器全局的JNDI 资源 命名服务
    -->
    <GlobalNamingResources>
        <Resource name="UserDatabase" auth="Container"
                  type="org.apache.catalina.UserDatabase"
                  description="User database that can be updated and saved"
                  factory="org.apache.catalina.users.MemoryUserDatabaseFactory"
                  pathname="conf/tomcat-users.xml"/>
    </GlobalNamingResources>

    <!--
            该标签⽤于创建 Service 实例，默认使⽤ org.apache.catalina.core.StandardService。
       默认情况下，Tomcat 仅指定了Service 的名称， 值为 "Catalina"。
       Service ⼦标签为 ： Listener、Executor、Connector、Engine，
       其中：
       Listener ⽤于为Service添加⽣命周期监听器，
       Executor ⽤于配置Service 共享线程池，(可以给多个 Connector连接器使用)
       Connector ⽤于配置Service 包含的链接器，
       Engine ⽤于配置Service中链接器对应的Servlet 容器引擎
     -->
    <Service name="Catalina">

        <!-- 默认情况下，Service 并未添加共享线程池配置。 如果我们想添加⼀个线程池， 可以在<Executor> 下添加如下配置：
              name：线程池名称，⽤于 Connector中指定
              namePrefix：所创建的每个线程的名称前缀，⼀个单独的线程名称为：namePrefix+线程编号
              maxThreads：池中最⼤线程数
              minSpareThreads：活跃线程数，也就是核⼼池线程数，这些线程不会被销毁，会⼀直存在
              maxIdleTime：线程空闲时间，超过该时间后，空闲线程会被销毁，默认值为6000（1分钟），单位毫秒
              maxQueueSize：在被执⾏前最⼤线程排队数⽬，默认为Int的最⼤值，也就是⼴义的⽆限。除⾮特殊情况，这个值 不需要更改，否则会有请求不会被处理的情况发⽣
              prestartminSpareThreads：启动线程池时是否启动 minSpareThreads部分线程。默认值为false，即不启动
              threadPriority：线程池中线程优先级，默认值为5，值从1到10
              className：线程池实现类，未指定情况下，默认实现类为
              org.apache.catalina.core.StandardThreadExecutor。
              如果想使⽤⾃定义线程池⾸先需要实现org.apache.catalina.Executor接⼝-->
        <Executor name="tomcatThreadPool"
                  namePrefix="catalina-exec-"
                  maxThreads="200"
                  minSpareThreads="100"
                  maxIdleTime="60000"
                  maxQueueSize="Integer.MAX_VALUE"
                  prestartminSpareThreads="true"
                  threadPriority="5"
                  className="org.apache.catalina.core.StandardThreadExecutor"/>

        <!--
           Connector 标签⽤于创建链接器实例，默认情况下，server.xml 配置了两个链接器，⼀个⽀持HTTP协议，⼀个⽀持AJP协议
           ⼤多数情况下，我们并不需要新增链接器配置，只是根据需要对已有链接器进⾏优化
                port：
                     端⼝号，Connector ⽤于创建服务端Socket 并进⾏监听， 以等待客户端请求链接。如果该属性设置为0， Tomcat将会随机选择⼀个可⽤的端⼝号给当前Connector 使⽤
                protocol：
                     当前Connector ⽀持的访问协议。 默认为 HTTP/1.1 ， 并采⽤⾃动切换机制选择⼀个基于 JAVA NIO 的链接器或者基于本地APR的链接器（根据本地是否含有Tomcat的本地库判定）
                connectionTimeOut:
                     Connector 接收链接后的等待超时时间， 单位为 毫秒。 -1 表示不超时。
                redirectPort：
                     如果当前接收的是一个 https 的请求，那么tomcat 会将请求转发到 redirectPort指定的端口。
                     比如现在设定的：8443 端口当前Connector 不⽀持SSL请求， 接收到了⼀个请求， 并且也符合security-constraint 约束，需要SSL传输，Catalina⾃动将请求重定向到指定的端⼝。
                executor：
                     指定共享线程池的名称， 也可以通过maxThreads、minSpareThreads 等属性配置内部线程池。
                URIEncoding:
                     ⽤于指定编码URI的字符编码， Tomcat8.x版本默认的编码为 UTF-8 , Tomcat7.x版本默认为ISO8859-1
 -->
        <!--org.apache.coyote.http11.Http11NioProtocol， ⾮阻塞式 Java NIO 链接器，tomcat8配置nio会报错，可能是已经集成了nio的原因-->
        <Connector port="80"
                   protocol="HTTP/1.1"
                   connectionTimeout="20000"
                   redirectPort="443"
                   executor="tomcatThreadPool"
                   URIEncoding="utf-8"/>


        <!-- certificateKeystoreFile 用于指定证书所在的目录 ；
                        certificateKeystorePassword 用于指定证书的密码;type是使用的加密算法-->
        <Connector port="443" protocol="org.apache.coyote.http11.Http11NioProtocol"
                   maxThreads="150" schema="https" secure="true" SSLEnabled="true">
            <SSLHostConfig>
                <Certificate
                        certificateKeystoreFile="conf/你的域名.cn.jks"
                        certificateKeystorePassword="你申请证书时提交密码"
                        type="RSA" />
            </SSLHostConfig>
        </Connector>


        <!-- Define an AJP 1.3 Connector on port 8009 -->

        <Connector protocol="AJP/1.3"
                   address="::1"
                   port="8009"
                   redirectPort="443" />


        <!--name： ⽤于指定Engine 的名称， 默认为Catalina
         defaultHost：默认使⽤的虚拟主机名称， 当客户端请求指向的主机⽆效时， 将交由默认的虚拟主机处
              理， 默认为localhost-->
        <Engine name="Catalina" defaultHost="localhost">
            <Realm className="org.apache.catalina.realm.LockOutRealm">
                <Realm className="org.apache.catalina.realm.UserDatabaseRealm"
                       resourceName="UserDatabase"/>
            </Realm>

            <!--Host 标签⽤于配置⼀个虚拟主机
                      name：该host的名称
                      appBase ：指定 war包放置的路径，可以是绝对路径，也可以是相对路径（相对路径，相对的就是tomcat的安装目录
                      unpackWARs ：是否自动解压 war包
                      autoDeploy：是否自动部署 （有点热部署的效果）-->
            <Host name="localhost" appBase="webapps"
                  unpackWARs="true" autoDeploy="true">

                <!-- 记录当前 host 处理请求的日志 -->
                <Valve className="org.apache.catalina.valves.AccessLogValve" directory="logs"
                       prefix="localhost_access_log" suffix=".txt"
                       pattern="%h %l %u %t &quot;%r&quot; %s %b"/>
            </Host>
        </Engine>
    </Service>
</Server>

```
其中有一个需要注意的地方就是，证书的位置certificateKeystoreFile可以填绝对路径，也可以填相对路径。如果填写的是相对路径，那地址应该是conf的上一层目录(如果你把jks文件放在server.xml的同级目录下，此处应该填"conf/域名.jks")，我因为这个踩过一些坑。
## HTTP 自动跳转 HTTPS 的安全配置（可选）
如果您需要将 HTTP 请求自动重定向到 HTTPS。您可以通过以下操作设置：

编辑 /usr/*/conf 目录下的 web.xml 文件，找到 </welcome-file-list> 标签。
请在结束标签 </welcome-file-list> 后面换行，并添加以下内容：
```xml
	<login-config>
    <!-- Authorization setting for SSL -->
    <auth-method>CLIENT-CERT</auth-method>
    <realm-name>Client Cert Users-only Area</realm-name>
    </login-config>
    
    <security-constraint>
    <!-- Authorization setting for SSL -->
    <web-resource-collection>
    <web-resource-name>SSL</web-resource-name>
    <url-pattern>/*</url-pattern>
    </web-resource-collection>
    <user-data-constraint>
    <transport-guarantee>CONFIDENTIAL</transport-guarantee>
    </user-data-constraint>
    </security-constraint>
```
## 如何检验配置是否成功

修改server.xml文件后，停止tomcat服务,在/usr/tomcat/*/bin目录下输入：
```linux
./shutdown.sh
```
然后以下命令检查配置文件是否有误，如果有报错信息就在网上搜索一下或者自己解决就好了：
```linux
./configtest.sh
```

以上步骤没有问题以后，输入以下命令就可以使用"https://域名.com"访问tomcat了。
```linux
./startup.sh
```

网上的教程大部分都只说./shutdown.sh和./startup.sh两个命令重启tomcat，但是有时候重启时有一些报错信息并不显示，就是访问80端口没有问题，但是443端口配置有错无法访问。

还要注意的是，./configtest.sh命令一定要在./shutdown.sh停止tomcat之后执行，要不然会出现端口已被占用的错误。

## 结语
最近在忙最近的毕业设计，最近做后端接口和最近写网页，然后在腾讯云上面买了一台服务器还有一个域名(总共花了80块钱，租了一年)。花了半个月终于备案完成，于是迫不及待的把自己写的网页传上去。其中配置docker、tomcat、mysql、rabbitmq、redis都遇到一些坑。希望可以帮到大家吧。

 [1]: 腾讯云Tomcat 服务器 SSL 证书安装部署（JKS 格式）https://cloud.tencent.com/document/product/400/35224
 
