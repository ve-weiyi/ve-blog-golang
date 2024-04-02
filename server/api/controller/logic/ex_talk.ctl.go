package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// @Tags		Talk
// @Summary		分页获取说说详情列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string																false	"token"
// @Param		uid		header		string																false	"uid"
// @Param		page	body		request.PageQuery													true	"分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]response.TalkDetailsDTO}}	"返回信息"
// @Router		/talk/find_talk_details_list [post]
func (s *TalkController) FindTalkDetailsList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var page request.PageQuery
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.TalkService.FindTalkDetailsList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Limit.Page,
		PageSize: page.Limit.PageSize,
	})
}

// @Tags		Talk
// @Summary		分页获取说说详情列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		req		body		request.IdReq										true	"id"
// @Success		200		{object}	response.Response{data=response.TalkDetailsDTO}	"返回信息"
// @Router		/talk/find_talk_details [get]
func (s *TalkController) FindTalkDetail(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.IdReq
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.TalkService.FindTalkDetailsDTO(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Talk
// @Summary		点赞说说
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		req		body		request.IdReq										true	"id"
// @Success		200		{object}	response.Response{data=response.EmptyResp}	"返回信息"
// @Router		/talk/like_talk [put]
func (s *TalkController) LikeTalk(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var req request.IdReq
	err = s.ShouldBind(c, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.TalkService.LikeTalk(reqCtx, &req)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
