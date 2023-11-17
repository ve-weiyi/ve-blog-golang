package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// @Tags		Comment
// @Summary		分页获取评论列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string																false	"token"
// @Param		uid		header		string																false	"uid"
// @Param		page	body		request.PageQuery													true	"分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]response.CommentDTO}}	"返回信息"
// @Router		/comment/details_list [post]
func (s *CommentController) FindCommentDetailsList(c *gin.Context) {
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

	list, total, err := s.svcCtx.CommentService.FindCommentDetailsList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.PageSize,
	})
}

// @Tags		Comment
// @Summary		获取用户评论列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		page	body		request.PageQuery			true	"分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]response.CommentBackDTO}}	"返回信息"
// @Router		/comment/list/back [post]
func (s *CommentController) FindCommentBackList(c *gin.Context) {
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

	s.Log.JsonIndent(page)
	list, total, err := s.svcCtx.CommentService.FindCommentBackList(reqCtx, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.PageSize,
	})
}

// @Tags		Comment
// @Summary		查询评论回复列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		id		path		int										true	"id"
// @Param		page	body		request.PageQuery						true	"请求body"
// @Success		200		{object}	response.Response{data=response.ReplyDTO}	"返回信息"
// @Router		/comment/{id}/reply_list [post]
func (s *CommentController) FindCommentReplyList(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	id := cast.ToInt(c.Param("id"))

	var page request.PageQuery
	err = s.ShouldBind(c, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	list, total, err := s.svcCtx.CommentService.FindCommentReplyList(reqCtx, id, &page)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Page,
		PageSize: page.PageSize,
	})
}

// @Tags		Comment
// @Summary		点赞评论
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		id		path		int										true	"id"
// @Success		200		{object}	response.Response{data=any}	"返回信息"
// @Router		/comment/{id}/like [post]
func (s *CommentController) LikeComment(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	id := cast.ToInt(c.Param("id"))

	data, err := s.svcCtx.CommentService.LikeComment(reqCtx, id)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}
