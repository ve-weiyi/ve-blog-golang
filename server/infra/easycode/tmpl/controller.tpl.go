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
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.{{.StructName}}							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.{{.StructName}}}	"返回信息"
// @Router		/{{.ValueName}}/create [post]
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

// @Tags		{{.StructName}}
// @Summary		删除{{.StructComment}}
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body	 	entity.{{.StructName}} 		true "请求body"
// @Success		200		{object}	response.Response{}		"返回信息"
// @Router		/{{.ValueName}}/delete [delete]
func (s *{{.StructName}}Controller) Delete{{.StructName}}(c *gin.Context) {
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

	data, err := s.svcCtx.{{.StructName}}Service.Delete{{.StructName}}(reqCtx, &{{.ValueName}});
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	{{.StructName}}
// @Summary		更新{{.StructComment}}
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data	body 	 	entity.{{.StructName}}							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.{{.StructName}}}	"返回信息"
// @Router 		/{{.ValueName}}/update [put]
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

// @Tags 	 	{{.StructName}}
// @Summary		查询{{.StructComment}}
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data		query		entity.{{.StructName}}							true		"请求参数"
// @Success		200			{object}	response.Response{data=entity.{{.StructName}}}	"返回信息"
// @Router 		/{{.ValueName}}/find [get]
func (s *{{.StructName}}Controller) Find{{.StructName}}(c *gin.Context) {
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

	data, err := s.svcCtx.{{.StructName}}Service.Find{{.StructName}}(reqCtx, &{{.ValueName}});
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	{{.StructName}}
// @Summary		批量删除{{.StructComment}}
// @Security 	ApiKeyAuth
// @accept 	 	application/json
// @Produce		application/json
// @Param		data 	body		[]int 				true "删除id列表"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/{{.ValueName}}/deleteByIds [delete]
func (s *{{.StructName}}Controller) Delete{{.StructName}}ByIds(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var IDS []int
	err = s.ShouldBind(c, &IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.{{.StructName}}Service.Delete{{.StructName}}ByIds(reqCtx, IDS);
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	{{.StructName}}
// @Summary		分页获取{{.StructComment}}列表
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	page 	body		request.PageInfo 	true "分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.{{.StructName}}}}	"返回信息"
// @Router		/{{.ValueName}}/list [post]
func (s *{{.StructName}}Controller) Find{{.StructName}}List(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageInfo
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
		PageSize: page.Limit(),
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
