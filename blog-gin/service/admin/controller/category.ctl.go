package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type CategoryController struct {
	svcCtx *svctx.ServiceContext
}

func NewCategoryController(svcCtx *svctx.ServiceContext) *CategoryController {
	return &CategoryController{
		svcCtx: svcCtx,
	}
}

// @Tags		Category
// @Summary		"创建文章分类"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.CategoryNewReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.CategoryBackVO}	"返回信息"
// @Router		/admin-api/v1/category/add_category [POST]
func (s *CategoryController) AddCategory(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.CategoryNewReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCategoryService(s.svcCtx).AddCategory(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Category
// @Summary		"删除文章分类"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.IdsReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.BatchResp}	"返回信息"
// @Router		/admin-api/v1/category/deletes_category [DELETE]
func (s *CategoryController) DeletesCategory(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.IdsReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCategoryService(s.svcCtx).DeletesCategory(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Category
// @Summary		"分页获取文章分类列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.CategoryQuery		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/admin-api/v1/category/find_category_list [POST]
func (s *CategoryController) FindCategoryList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.CategoryQuery
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCategoryService(s.svcCtx).FindCategoryList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Category
// @Summary		"更新文章分类"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.CategoryNewReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.CategoryBackVO}	"返回信息"
// @Router		/admin-api/v1/category/update_category [PUT]
func (s *CategoryController) UpdateCategory(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.CategoryNewReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCategoryService(s.svcCtx).UpdateCategory(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
