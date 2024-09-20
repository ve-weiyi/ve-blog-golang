
# 如何在go-zero框架中使用swagger接口文档（knife4j）  
文章封面: https://static.veweiyi.cn/blog/3/article/8C60E9E6ECBA38FCB8F6E323C19FA902-20240919141800.png   
文章类型: 1   
文章分类: 学习   
文章标签: [测试标签]   
创建时间: 1970-01-01 08:00:00 +0800 CST   

文章内容:
## 前言
在工作中，我接触到go后端框架有几种，beego、gin、go-zero。其中，beego自带swagger文档，gin框架可以使用gin-swagger快捷的使用swagger，go-zero没有专门的swagger文档实现方式，但是依然可以使用http-swagger实现。

但是，单纯的展示文档似乎并不能满足我的想法，因为官方的swagger文档页面实在丑陋，以至于我们公司都不使用这个swagger文档页面作为各端沟通文档。

那么，能否有一个好看又方便，实现成本又低的接口文档呢。我想起了以前写springboot的时候，有一个叫knife4j的接口文档页面，美观又漂亮。（你可能有疑问，我明明是一个go后端工程师，怎么还会springboot，O(∩_∩)O哈哈~）

于是，我开始了探索使用go-zero实现knife4j接口文档的实现方式。

##
