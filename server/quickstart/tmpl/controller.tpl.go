package tmpl

const AppController = `
package controller

import (

)

type AppController struct {
	svcCtx *svc.ControllerContext //持有的service引用
}

func NewController(svcCtx *svc.ControllerContext) *AppController {
	return &AppController{
		svcCtx: svcCtx,
	}
}
`

const ControllerContext = `
package svc

import (

)

// 注册需要用到的rpc
type ControllerContext struct {
	*service.AppService
}

func NewControllerContext(cfg *config.Config) *ControllerContext {
	ctx := svc.NewServiceContext(cfg)
	sv := service.NewService(ctx)
	if sv == nil {
		panic("sv cannot be null")
	}

	return &ControllerContext{
		AppService: sv,
	}
}

`

const Controller = `
package logic

import (
	"github.com/gin-gonic/gin"
	{{range .ImportPkgPaths}}{{.}} ` + "\n" + `{{end}}
)

type {{.StructName}}Controller struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func New{{.StructName}}Controller(svcCtx *svc.ControllerContext) *{{.StructName}}Controller {
	return &{{.StructName}}Controller{
		svcCtx: svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		{{.StructName}}
// @Summary		创建{{.StructComment}}
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		entity.{{.StructName}}		true	"请求参数"
// @Success		200		{object}	response.Response{data=entity.{{.StructName}}}	"返回信息"
// @Router		/{{.JsonName}} [post]
func (s *{{.StructName}}Controller) Create{{.StructName}}(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var {{.ValueName}} entity.{{.StructName}}
	err = s.ShouldBind(c, &{{.ValueName}})
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.{{.StructName}}Service.Create{{.StructName}}(reqCtx, &{{.ValueName}});
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	{{.StructName}}
// @Summary		更新{{.StructComment}}
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	data	body 	 	entity.{{.StructName}}		true	"请求参数"
// @Success		200		{object}	response.Response{data=entity.{{.StructName}}}	"返回信息"
// @Router 		/{{.JsonName}} [put]
func (s *{{.StructName}}Controller) Update{{.StructName}}(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var {{.ValueName}} entity.{{.StructName}}
	err = s.ShouldBind(c, &{{.ValueName}})
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.{{.StructName}}Service.Update{{.StructName}}(reqCtx, &{{.ValueName}});
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		{{.StructName}}
// @Summary		删除{{.StructComment}}
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	id		path		int							true	"{{.StructName}} id"
// @Success		200		{object}	response.Response{data=any}			"返回信息"
// @Router		/{{.JsonName}}/{id} [delete]
func (s *{{.StructName}}Controller) Delete{{.StructName}}(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.{{.StructName}}Service.Delete{{.StructName}}(reqCtx, id);
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	{{.StructName}}
// @Summary		查询{{.StructComment}}
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	id		path		int							true	"{{.StructName}} id"
// @Success		200		{object}	response.Response{data=entity.{{.StructName}}}	"返回信息"
// @Router 		/{{.JsonName}}/{id} [get]
func (s *{{.StructName}}Controller) Find{{.StructName}}(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.{{.StructName}}Service.Find{{.StructName}}(reqCtx, id);
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	{{.StructName}}
// @Summary		批量删除{{.StructComment}}
// @Accept 	 	application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data 	body		[]int 						true 	"删除id列表"
// @Success		200		{object}	response.Response{data=response.BatchResult}	"返回信息"
// @Router		/{{.JsonName}}/batch_delete [delete]
func (s *{{.StructName}}Controller) Delete{{.StructName}}ByIds(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var ids []int
	err = s.ShouldBind(c, &ids)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.{{.StructName}}Service.Delete{{.StructName}}ByIds(reqCtx, ids);
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.BatchResult{
		TotalCount:   len(ids),
		SuccessCount: data,
		FailCount:    len(ids) - data,
	})
}

// @Tags 	 	{{.StructName}}
// @Summary		分页获取{{.StructComment}}列表
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	page 	body		request.PageQuery 			true 	"分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.{{.StructName}}}}	"返回信息"
// @Router		/{{.JsonName}}/list [post]
func (s *{{.StructName}}Controller) Find{{.StructName}}List(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageQuery
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.{{.StructName}}Service.Find{{.StructName}}List(reqCtx, &page); 
	if err != nil {
		s.ResponseError(c, err)
		return
	}
	
	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.PageSize,
	})
}
`

const CommonController = `
package logic

import (

)

type {{.StructName}}Controller struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func New{{.StructName}}Controller(svcCtx *svc.ControllerContext) *{{.StructName}}Controller {
	return &{{.StructName}}Controller{
		svcCtx: svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}
`
