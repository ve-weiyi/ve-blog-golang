package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
)

type PhotoAlbumController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewPhotoAlbumController(ctx *svc.ControllerContext) *PhotoAlbumController {
	return &PhotoAlbumController{
		svcCtx:         ctx,
		BaseController: controller.NewBaseController(ctx),
	}
}

// @Tags	 PhotoAlbum
// @Summary  创建相册
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data  body 	 entity.PhotoAlbum		true  "请求body"
// @Success  200   {object}  response.Response{data=entity.PhotoAlbum}  	"返回信息"
// @Router /photoAlbum/create [post]
func (s *PhotoAlbumController) CreatePhotoAlbum(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var photoAlbum entity.PhotoAlbum
	err = s.ShouldBind(c, &photoAlbum)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PhotoAlbumService.CreatePhotoAlbum(reqCtx, &photoAlbum)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	PhotoAlbum
// @Summary 删除相册
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body	 	entity.PhotoAlbum 		true "请求body"
// @Success 200  {object}  	response.Response{}  	"返回信息"
// @Router /photoAlbum/delete [delete]
func (s *PhotoAlbumController) DeletePhotoAlbum(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var photoAlbum entity.PhotoAlbum
	err = s.ShouldBind(c, &photoAlbum)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PhotoAlbumService.DeletePhotoAlbum(reqCtx, &photoAlbum)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	PhotoAlbum
// @Summary 更新相册
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body 		entity.PhotoAlbum 		true "请求body"
// @Success 200  {object}  	response.Response{data=entity.PhotoAlbum}  	"返回信息"
// @Router /photoAlbum/update [put]
func (s *PhotoAlbumController) UpdatePhotoAlbum(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var photoAlbum entity.PhotoAlbum
	err = s.ShouldBind(c, &photoAlbum)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PhotoAlbumService.UpdatePhotoAlbum(reqCtx, &photoAlbum)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	PhotoAlbum
// @Summary 查询相册
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce	application/json
// @Param 	data query 		entity.PhotoAlbum 		true "请求body"
// @Success 200  {object}  	response.Response{data=entity.PhotoAlbum}  	"返回信息"
// @Router /photoAlbum/query [get]
func (s *PhotoAlbumController) GetPhotoAlbum(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var photoAlbum entity.PhotoAlbum
	err = s.ShouldBind(c, &photoAlbum)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PhotoAlbumService.GetPhotoAlbum(reqCtx, &photoAlbum)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	PhotoAlbum
// @Summary 批量删除相册
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body 		[]int 					true "删除id列表"
// @Success 200  {object}  	response.Response{}  	"返回信息"
// @Router /photoAlbum/deleteByIds [delete]
func (s *PhotoAlbumController) DeletePhotoAlbumByIds(c *gin.Context) {
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

	data, err := s.svcCtx.PhotoAlbumService.DeletePhotoAlbumByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	PhotoAlbum
// @Summary 分页获取相册列表
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce	application/json
// @Param 	data query 		request.PageInfo 	true "分页参数"
// @Success 200  {object}  	response.Response{data=response.PageResult{list=[]entity.PhotoAlbum}}  	"返回信息"
// @Router /photoAlbum/list [get]
func (s *PhotoAlbumController) FindPhotoAlbumList(c *gin.Context) {
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

	list, total, err := s.svcCtx.PhotoAlbumService.FindPhotoAlbumList(reqCtx, &page)
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
