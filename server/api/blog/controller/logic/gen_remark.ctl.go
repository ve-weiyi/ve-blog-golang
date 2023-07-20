package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type RemarkController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewRemarkController(svcCtx *svc.ControllerContext) *RemarkController {
	return &RemarkController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Remark
// @Summary		创建留言
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Remark							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.Remark}	"返回信息"
// @Router		/remark/create [post]
func (s *RemarkController) CreateRemark(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var remark entity.Remark
	err = s.ShouldBind(c, &remark)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.RemarkService.CreateRemark(reqCtx, &remark)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Remark
// @Summary		删除留言
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body	 	entity.Remark 		true "请求body"
// @Success		200		{object}	response.Response{}		"返回信息"
// @Router		/remark/delete [delete]
func (s *RemarkController) DeleteRemark(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var remark entity.Remark
	err = s.ShouldBind(c, &remark)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.RemarkService.DeleteRemark(reqCtx, &remark)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Remark
// @Summary		更新留言
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data	body 	 	entity.Remark							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.Remark}	"返回信息"
// @Router 		/remark/update [put]
func (s *RemarkController) UpdateRemark(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var remark entity.Remark
	err = s.ShouldBind(c, &remark)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.RemarkService.UpdateRemark(reqCtx, &remark)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Remark
// @Summary		查询留言
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data		query		entity.Remark							true		"请求参数"
// @Success		200			{object}	response.Response{data=entity.Remark}	"返回信息"
// @Router 		/remark/find [get]
func (s *RemarkController) FindRemark(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var remark entity.Remark
	err = s.ShouldBind(c, &remark)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.RemarkService.FindRemark(reqCtx, &remark)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Remark
// @Summary		批量删除留言
// @Security 	ApiKeyAuth
// @accept 	 	application/json
// @Produce		application/json
// @Param		data 	body		[]int 				true "删除id列表"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/remark/deleteByIds [delete]
func (s *RemarkController) DeleteRemarkByIds(c *gin.Context) {
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

	data, err := s.svcCtx.RemarkService.DeleteRemarkByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Remark
// @Summary		分页获取留言列表
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	page 	body		request.PageInfo 	true "分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.Remark}}	"返回信息"
// @Router		/remark/list [post]
func (s *RemarkController) FindRemarkList(c *gin.Context) {
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

	list, total, err := s.svcCtx.RemarkService.FindRemarkList(reqCtx, &page)
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