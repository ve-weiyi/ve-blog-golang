package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
)

// @Tags		PhotoAlbum
// @Summary		获取相册详情列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		page	body		request.PageQuery						true	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.PhotoAlbumDetailsDTO}}	"返回信息"
// @Router		/photo_album/find_photo_album_details_list [post]
func (s *PhotoAlbumController) FindPhotoAlbumDetailsList(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var page dto.PageQuery
	err = request.ShouldBind(c, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	list, total, err := service.NewPhotoAlbumService(s.svcCtx).FindPhotoAlbumDetailsList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     0,
		PageSize: int64(len(list)),
	})
}

// @Tags		PhotoAlbum
// @Summary		获取相册详情
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param 	 	request		body		request.IdReq										true	"PhotoAlbum.id"
// @Success		200		{object}	response.Body{data=dto.PhotoAlbumDetailsDTO}	"返回信息"
// @Router		/photo_album/find_photo_album_details [get]
func (s *PhotoAlbumController) FindPhotoAlbumDetails(c *gin.Context) {
	reqCtx, err := request.GetRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req request.IdReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewPhotoAlbumService(s.svcCtx).FindPhotoAlbumDetails(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
