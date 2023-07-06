package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
)

type PageController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewPageController(svcCtx *svc.ControllerContext) *PageController {
	return &PageController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// CreatePage 创建页面
// @Tags	 Page
// @Summary  创建页面
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data  body 	 entity.Page		true  "创建页面"
// @Success  200   {object}  response.Response{data=entity.Page}  	"返回信息"
// @Router /page/create [post]
func (s *PageController) CreatePage(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page entity.Page
	err = s.ShouldBindJSON(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PageService.CreatePage(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// DeletePage 删除页面
// @Tags 	Page
// @Summary  删除页面
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data 	body	 	entity.Page 		true "删除页面"
// @Success  200  	{object}  	response.Response{}  	"返回信息"
// @Router /page/delete [delete]
func (s *PageController) DeletePage(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page entity.Page
	err = s.ShouldBindJSON(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PageService.DeletePage(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// UpdatePage 更新页面
// @Tags 	Page
// @Summary  更新页面
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data 	body 		entity.Page 		true "更新页面"
// @Success  200  	{object}  	response.Response{data=entity.Page}  	"返回信息"
// @Router /page/update [put]
func (s *PageController) UpdatePage(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page entity.Page
	err = s.ShouldBindJSON(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PageService.UpdatePage(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// FindPage 用id查询页面
// @Tags 	Page
// @Summary 用id查询页面
// @Security ApiKeyAuth
// @accept	 application/json
// @Produce  application/json
// @Param 	 data 	body 		entity.Page 		true "用id查询页面"
// @Success  200  	{object}  	response.Response{data=entity.Page}  	"返回信息"
// @Router /page/find [get]
func (s *PageController) FindPage(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page entity.Page
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PageService.FindPage(reqCtx, page.ID)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// DeletePageByIds 批量删除页面
// @Tags 	Page
// @Summary  批量删除页面
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data 	body 		[]int 					true "批量删除页面"
// @Success  200  	{object}  	response.Response{}  	"返回信息"
// @Router /page/deleteByIds [delete]
func (s *PageController) DeletePageByIds(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var IDS []int
	err = s.ShouldBindJSON(c, &IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PageService.DeletePageByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// GetPageList 分页获取页面列表
// @Tags 	Page
// @Summary  分页获取页面列表
// @Security ApiKeyAuth
// @accept	 application/json
// @Produce  application/json
// @Param 	 page 	body 		request.PageInfo 	true "分页获取页面列表"
// @Success  200  	{object}  	response.Response{data=response.PageResult{list=[]entity.Page}}  	"返回信息"
// @Router /page/list [get]
func (s *PageController) GetPageList(c *gin.Context) {
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

	list, total, err := s.svcCtx.PageService.GetPageList(reqCtx, &page)
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
