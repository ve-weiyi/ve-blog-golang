package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type PhotoController struct {
	svcCtx *svctx.ServiceContext
}

func NewPhotoController(svcCtx *svctx.ServiceContext) *PhotoController {
	return &PhotoController{
		svcCtx: svcCtx,
	}
}

// @Tags		Photo
// @Summary		"创建照片"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.PhotoNewReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PhotoBackVO}	"返回信息"
// @Router		/admin-api/v1/photo/add_photo [POST]
func (s *PhotoController) AddPhoto(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.PhotoNewReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewPhotoService(s.svcCtx).AddPhoto(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Photo
// @Summary		"删除照片"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.IdsReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.BatchResp}	"返回信息"
// @Router		/admin-api/v1/photo/deletes_photo [DELETE]
func (s *PhotoController) DeletesPhoto(c *gin.Context) {
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

	data, err := service.NewPhotoService(s.svcCtx).DeletesPhoto(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Photo
// @Summary		"分页获取照片列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.PhotoQuery		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/admin-api/v1/photo/find_photo_list [POST]
func (s *PhotoController) FindPhotoList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.PhotoQuery
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewPhotoService(s.svcCtx).FindPhotoList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Photo
// @Summary		"预删除照片"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.PreDeletePhotoReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.BatchResp}	"返回信息"
// @Router		/admin-api/v1/photo/pre_delete_photo [PUT]
func (s *PhotoController) PreDeletePhoto(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.PreDeletePhotoReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewPhotoService(s.svcCtx).PreDeletePhoto(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Photo
// @Summary		"更新照片"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.PhotoNewReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PhotoBackVO}	"返回信息"
// @Router		/admin-api/v1/photo/update_photo [PUT]
func (s *PhotoController) UpdatePhoto(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.PhotoNewReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewPhotoService(s.svcCtx).UpdatePhoto(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
