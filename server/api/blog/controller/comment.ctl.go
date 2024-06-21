package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/response"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type CommentController struct {
	svcCtx *svc.ServiceContext
}

func NewCommentController(svcCtx *svc.ServiceContext) *CommentController {
	return &CommentController{
		svcCtx: svcCtx,
	}
}

// @Tags		Comment
// @Summary		创建评论
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		data	body		entity.Comment		true	"请求参数"
// @Success		200		{object}	response.Body{data=entity.Comment}	"返回信息"
// @Router		/comment/create_comment [post]
func (s *CommentController) CreateComment(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var comment entity.Comment
	err = request.ShouldBind(c, &comment)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCommentService(s.svcCtx).CreateComment(reqCtx, &comment)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags 	 	Comment
// @Summary		更新评论
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	data	body 	 	entity.Comment		true	"请求参数"
// @Success		200		{object}	response.Body{data=entity.Comment}	"返回信息"
// @Router 		/comment/update_comment [put]
func (s *CommentController) UpdateComment(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var comment entity.Comment
	err = request.ShouldBind(c, &comment)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCommentService(s.svcCtx).UpdateComment(reqCtx, &comment)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags		Comment
// @Summary		删除评论
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	request		body		request.IdReq							true	"Comment.id"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}			"返回信息"
// @Router		/comment/delete_comment [delete]
func (s *CommentController) DeleteComment(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
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

	data, err := service.NewCommentService(s.svcCtx).DeleteComment(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags 	 	Comment
// @Summary		查询评论
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	request		body		request.IdReq							true	"Comment.id"
// @Success		200		{object}	response.Body{data=entity.Comment}	"返回信息"
// @Router 		/comment/find_comment [get]
func (s *CommentController) FindComment(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
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

	data, err := service.NewCommentService(s.svcCtx).FindComment(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}

// @Tags 	 	Comment
// @Summary		批量删除评论
// @Accept 	 	application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	req		body		request.IdsReq				true	"删除id列表"
// @Success		200		{object}	response.Body{data=dto.BatchResult}	"返回信息"
// @Router		/comment/delete_comment_list [delete]
func (s *CommentController) DeleteCommentList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var req request.IdsReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCommentService(s.svcCtx).DeleteCommentList(reqCtx, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.BatchResult{
		SuccessCount: data,
	})
}

// @Tags 	 	Comment
// @Summary		分页获取评论列表
// @Accept 		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param 	 	page 	body		request.CommentQueryReq 			true 	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]entity.Comment}}	"返回信息"
// @Router		/comment/find_comment_list [post]
func (s *CommentController) FindCommentList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	var page dto.CommentQueryReq
	err = request.ShouldBind(c, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	list, total, err := service.NewCommentService(s.svcCtx).FindCommentList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
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
// @Param		req		body		request.IdReq										true	"id"
// @Param		page	body		request.PageQuery						true	"请求body"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.ReplyDTO}}	"返回信息"
// @Router		/comment/find_comment_reply_list [post]
func (s *CommentController) FindCommentReplyList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	id := cast.ToInt64(c.Param("id"))

	var page dto.PageQuery
	err = request.ShouldBind(c, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	list, total, err := service.NewCommentService(s.svcCtx).FindCommentReplyList(reqCtx, id, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Limit.Page,
		PageSize: page.Limit.PageSize,
	})
}

// @Tags		Comment
// @Summary		获取用户评论列表
// @Accept		application/json
// @Produce		application/json
// @Param		token	header		string						false	"token"
// @Param		uid		header		string						false	"uid"
// @Param		page	body		request.PageQuery			true	"分页参数"
// @Success		200		{object}	response.Body{data=dto.PageResult{list=[]dto.CommentBackDTO}}	"返回信息"
// @Router		/comment/find_comment_back_list [post]
func (s *CommentController) FindCommentBackList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
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

	list, total, err := service.NewCommentService(s.svcCtx).FindCommentBackList(reqCtx, &page)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, response.PageResult{
		List:     list,
		Total:    total,
		Page:     page.Limit.Page,
		PageSize: page.Limit.PageSize,
	})
}

// @Tags		Comment
// @Summary		点赞评论
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		token	header		string									false	"token"
// @Param		uid		header		string									false	"uid"
// @Param		req		body		request.IdReq										true	"id"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/comment/like_comment [post]
func (s *CommentController) LikeComment(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	id := cast.ToInt64(c.Param("id"))

	data, err := service.NewCommentService(s.svcCtx).LikeComment(reqCtx, id)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	response.ResponseOk(c, data)
}
