
inject 使用ast在文件中指定位置注入代码
--east


plate 模版生成

[provider](plate/provider)
数据库操作帮助

controller --- api.handler
service --- api.logic
repository--- rpc.handler、logic
entity--- model

// @Tags Blog
// @Summary 创建api路由
// @Description 描述,可以有多个。https://www.jianshu.com/p/4bb4283632e4
// @Security ApiKeyUser
// @Param file formData file true    "上传文件"
// @Param id path int true    "id"
// @Param token header string true    "token"
// @Param data body entity.Api true    "创建api路由"
// @Success 200 {object} response.Response{data=entity.Api}    "返回信息"
// @Router /version [get]
