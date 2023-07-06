package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
)

type CategoryController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewCategoryController(svcCtx *svc.ControllerContext) *CategoryController {
	return &CategoryController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// CreateCategory 创建文章分类
// @Tags	 Category
// @Summary  创建文章分类
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data  body 	 entity.Category		true  "创建文章分类"
// @Success  200   {object}  response.Response{data=entity.Category}  	"返回信息"
// @Router /category/create [post]
func (s *CategoryController) CreateCategory(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var category entity.Category
	err = s.ShouldBindJSON(c, &category)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.CategoryService.CreateCategory(reqCtx, &category)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// DeleteCategory 删除文章分类
// @Tags 	Category
// @Summary  删除文章分类
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data 	body	 	entity.Category 		true "删除文章分类"
// @Success  200  	{object}  	response.Response{}  	"返回信息"
// @Router /category/delete [delete]
func (s *CategoryController) DeleteCategory(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var category entity.Category
	err = s.ShouldBindJSON(c, &category)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.CategoryService.DeleteCategory(reqCtx, &category)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// UpdateCategory 更新文章分类
// @Tags 	Category
// @Summary  更新文章分类
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data 	body 		entity.Category 		true "更新文章分类"
// @Success  200  	{object}  	response.Response{data=entity.Category}  	"返回信息"
// @Router /category/update [put]
func (s *CategoryController) UpdateCategory(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var category entity.Category
	err = s.ShouldBindJSON(c, &category)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.CategoryService.UpdateCategory(reqCtx, &category)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// FindCategory 用id查询文章分类
// @Tags 	Category
// @Summary 用id查询文章分类
// @Security ApiKeyAuth
// @accept	 application/json
// @Produce  application/json
// @Param 	 data 	body 		entity.Category 		true "用id查询文章分类"
// @Success  200  	{object}  	response.Response{data=entity.Category}  	"返回信息"
// @Router /category/find [get]
func (s *CategoryController) FindCategory(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var category entity.Category
	err = s.ShouldBind(c, &category)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.CategoryService.FindCategory(reqCtx, category.ID)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// DeleteCategoryByIds 批量删除文章分类
// @Tags 	Category
// @Summary  批量删除文章分类
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data 	body 		[]int 					true "批量删除文章分类"
// @Success  200  	{object}  	response.Response{}  	"返回信息"
// @Router /category/deleteByIds [delete]
func (s *CategoryController) DeleteCategoryByIds(c *gin.Context) {
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

	data, err := s.svcCtx.CategoryService.DeleteCategoryByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// GetCategoryList 分页获取文章分类列表
// @Tags 	Category
// @Summary  分页获取文章分类列表
// @Security ApiKeyAuth
// @accept	 application/json
// @Produce  application/json
// @Param 	 page 	body 		request.PageInfo 	true "分页获取文章分类列表"
// @Success  200  	{object}  	response.Response{data=response.PageResult{list=[]entity.Category}}  	"返回信息"
// @Router /category/list [get]
func (s *CategoryController) GetCategoryList(c *gin.Context) {
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

	list, total, err := s.svcCtx.CategoryService.GetCategoryList(reqCtx, &page)
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
