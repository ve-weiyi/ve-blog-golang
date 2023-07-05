package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/infra/base/controller"
)

type TalkController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewTalkController(ctx *svc.ControllerContext) *TalkController {
	return &TalkController{
		svcCtx:         ctx,
		BaseController: controller.NewBaseController(ctx),
	}
}

// @Tags	 Talk
// @Summary  创建说说
// @Security ApiKeyAuth
// @accept 	 application/json
// @Produce  application/json
// @Param 	 data  body 	 entity.Talk		true  "请求body"
// @Success  200   {object}  response.Response{data=entity.Talk}  	"返回信息"
// @Router /talk/create [post]
func (s *TalkController) CreateTalk(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var talk entity.Talk
	err = s.ShouldBind(c, &talk)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.TalkService.CreateTalk(reqCtx, &talk)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	Talk
// @Summary 删除说说
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body	 	entity.Talk 		true "请求body"
// @Success 200  {object}  	response.Response{}  	"返回信息"
// @Router /talk/delete [delete]
func (s *TalkController) DeleteTalk(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var talk entity.Talk
	err = s.ShouldBind(c, &talk)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.TalkService.DeleteTalk(reqCtx, &talk)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	Talk
// @Summary 更新说说
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body 		entity.Talk 		true "请求body"
// @Success 200  {object}  	response.Response{data=entity.Talk}  	"返回信息"
// @Router /talk/update [put]
func (s *TalkController) UpdateTalk(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var talk entity.Talk
	err = s.ShouldBind(c, &talk)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.TalkService.UpdateTalk(reqCtx, &talk)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	Talk
// @Summary 查询说说
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce	application/json
// @Param 	data query 		entity.Talk 		true "请求body"
// @Success 200  {object}  	response.Response{data=entity.Talk}  	"返回信息"
// @Router /talk/query [get]
func (s *TalkController) GetTalk(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var talk entity.Talk
	err = s.ShouldBind(c, &talk)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.TalkService.GetTalk(reqCtx, &talk)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	Talk
// @Summary 批量删除说说
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce application/json
// @Param 	data body 		[]int 					true "删除id列表"
// @Success 200  {object}  	response.Response{}  	"返回信息"
// @Router /talk/deleteByIds [delete]
func (s *TalkController) DeleteTalkByIds(c *gin.Context) {
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

	data, err := s.svcCtx.TalkService.DeleteTalkByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	Talk
// @Summary 分页获取说说列表
// @Security ApiKeyAuth
// @accept 	application/json
// @Produce	application/json
// @Param 	data query 		request.PageInfo 	true "分页参数"
// @Success 200  {object}  	response.Response{data=response.PageResult{list=[]entity.Talk}}  	"返回信息"
// @Router /talk/list [get]
func (s *TalkController) FindTalkList(c *gin.Context) {
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

	list, total, err := s.svcCtx.TalkService.FindTalkList(reqCtx, &page)
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
