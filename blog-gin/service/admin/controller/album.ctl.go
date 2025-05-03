package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type AlbumController struct {
	svcCtx *svctx.ServiceContext
}

func NewAlbumController(svcCtx *svctx.ServiceContext) *AlbumController {
	return &AlbumController{
		svcCtx: svcCtx,
	}
}

// @Tags		Album
// @Summary		"创建相册"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.AlbumNewReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.AlbumBackVO}	"返回信息"
// @Router		/admin-api/v1/album/add_album [POST]
func (s *AlbumController) AddAlbum(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.AlbumNewReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAlbumService(s.svcCtx).AddAlbum(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Album
// @Summary		"删除相册"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.IdReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.BatchResp}	"返回信息"
// @Router		/admin-api/v1/album/delete_album [DELETE]
func (s *AlbumController) DeleteAlbum(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.IdReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAlbumService(s.svcCtx).DeleteAlbum(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Album
// @Summary		"分页获取相册列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.AlbumQuery		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/admin-api/v1/album/find_album_list [POST]
func (s *AlbumController) FindAlbumList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.AlbumQuery
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAlbumService(s.svcCtx).FindAlbumList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Album
// @Summary		"查询相册"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.IdReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.AlbumBackVO}	"返回信息"
// @Router		/admin-api/v1/album/get_album [POST]
func (s *AlbumController) GetAlbum(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.IdReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAlbumService(s.svcCtx).GetAlbum(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Album
// @Summary		"更新相册"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.AlbumNewReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.AlbumBackVO}	"返回信息"
// @Router		/admin-api/v1/album/update_album [PUT]
func (s *AlbumController) UpdateAlbum(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.AlbumNewReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewAlbumService(s.svcCtx).UpdateAlbum(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
