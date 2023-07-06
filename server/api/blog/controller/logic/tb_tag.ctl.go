package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
)

type TagController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewTagController(svcCtx *svc.ControllerContext) *TagController {
	return &TagController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Tag
// @Summary	创建文章标签
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Tag							true	"创建文章标签"
// @Success	200		{object}	response.Response{data=entity.Tag}	"返回信息"
// @Router		/tag/create [post]
func (s *TagController) CreateTag(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var tag entity.Tag
	err = s.ShouldBindJSON(c, &tag)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.TagService.CreateTag(reqCtx, &tag)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Tag
// @Summary	删除文章标签
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Tag			true	"删除文章标签"
// @Success	200		{object}	response.Response{}	"返回信息"
// @Router		/tag/delete [delete]
func (s *TagController) DeleteTag(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var tag entity.Tag
	err = s.ShouldBindJSON(c, &tag)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.TagService.DeleteTag(reqCtx, &tag)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Tag
// @Summary	更新文章标签
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Tag							true	"更新文章标签"
// @Success	200		{object}	response.Response{data=entity.Tag}	"返回信息"
// @Router		/tag/update [put]
func (s *TagController) UpdateTag(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var tag entity.Tag
	err = s.ShouldBindJSON(c, &tag)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.TagService.UpdateTag(reqCtx, &tag)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Tag
// @Summary	用id查询文章标签
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		entity.Tag							true	"用id查询文章标签"
// @Success	200		{object}	response.Response{data=entity.Tag}	"返回信息"
// @Router		/tag/find [get]
func (s *TagController) FindTag(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var tag entity.Tag
	err = s.ShouldBind(c, &tag)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.TagService.FindTag(reqCtx, tag.ID)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Tag
// @Summary	批量删除文章标签
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		data	body		[]int				true	"批量删除文章标签"
// @Success	200		{object}	response.Response{}	"返回信息"
// @Router		/tag/deleteByIds [delete]
func (s *TagController) DeleteTagByIds(c *gin.Context) {
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

	data, err := s.svcCtx.TagService.DeleteTagByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Tag
// @Summary	分页获取文章标签列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		page	body		request.PageInfo												true	"分页获取文章标签列表"
// @Success	200		{object}	response.Response{data=response.PageResult{list=[]entity.Tag}}	"返回信息"
// @Router		/tag/list [get]
func (s *TagController) GetTagList(c *gin.Context) {
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

	list, total, err := s.svcCtx.TagService.GetTagList(reqCtx, &page)
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
