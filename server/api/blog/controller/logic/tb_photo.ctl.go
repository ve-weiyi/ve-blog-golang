package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
)

type PhotoController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewPhotoController(ctx *svc.ControllerContext) *PhotoController {
	return &PhotoController{
		svcCtx:         ctx,
		BaseController: controller.NewBaseController(ctx),
	}
}

//	@Tags		Photo
//	@Summary	创建相片
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		entity.Photo							true	"请求body"
//	@Success	200		{object}	response.Response{data=entity.Photo}	"返回信息"
//	@Router		/photo/create [post]
func (s *PhotoController) CreatePhoto(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var photo entity.Photo
	err = s.ShouldBind(c, &photo)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PhotoService.CreatePhoto(reqCtx, &photo)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

//	@Tags		Photo
//	@Summary	删除相片
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		entity.Photo		true	"请求body"
//	@Success	200		{object}	response.Response{}	"返回信息"
//	@Router		/photo/delete [delete]
func (s *PhotoController) DeletePhoto(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var photo entity.Photo
	err = s.ShouldBind(c, &photo)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PhotoService.DeletePhoto(reqCtx, &photo)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

//	@Tags		Photo
//	@Summary	更新相片
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		entity.Photo							true	"请求body"
//	@Success	200		{object}	response.Response{data=entity.Photo}	"返回信息"
//	@Router		/photo/update [put]
func (s *PhotoController) UpdatePhoto(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var photo entity.Photo
	err = s.ShouldBind(c, &photo)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PhotoService.UpdatePhoto(reqCtx, &photo)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

//	@Tags		Photo
//	@Summary	查询相片
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		entity.Photo							true	"请求body"
//	@Success	200		{object}	response.Response{data=entity.Photo}	"返回信息"
//	@Router		/photo/query [get]
func (s *PhotoController) GetPhoto(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var photo entity.Photo
	err = s.ShouldBind(c, &photo)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.PhotoService.GetPhoto(reqCtx, &photo)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

//	@Tags		Photo
//	@Summary	批量删除相片
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		[]int				true	"删除id列表"
//	@Success	200		{object}	response.Response{}	"返回信息"
//	@Router		/photo/deleteByIds [delete]
func (s *PhotoController) DeletePhotoByIds(c *gin.Context) {
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

	data, err := s.svcCtx.PhotoService.DeletePhotoByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

//	@Tags		Photo
//	@Summary	分页获取相片列表
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		page	body		request.PageInfo													true	"分页参数"
//	@Success	200		{object}	response.Response{data=response.PageResult{list=[]entity.Photo}}	"返回信息"
//	@Router		/photo/list [get]
func (s *PhotoController) FindPhotoList(c *gin.Context) {
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

	list, total, err := s.svcCtx.PhotoService.FindPhotoList(reqCtx, &page)
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
