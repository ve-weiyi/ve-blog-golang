package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/response"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/service"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type CommentController struct {
	svcCtx *svctx.ServiceContext
}

func NewCommentController(svcCtx *svctx.ServiceContext) *CommentController {
	return &CommentController{
		svcCtx: svcCtx,
	}
}

// @Tags		Comment
// @Summary		"查询评论列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.CommentQueryReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/blog-api/v1/comment/find_comment_list [POST]
func (s *CommentController) FindCommentList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.CommentQueryReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCommentService(s.svcCtx).FindCommentList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Comment
// @Summary		"查询最新评论回复列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.CommentQueryReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/blog-api/v1/comment/find_comment_recent_list [POST]
func (s *CommentController) FindCommentRecentList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.CommentQueryReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCommentService(s.svcCtx).FindCommentRecentList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Comment
// @Summary		"查询评论回复列表"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.CommentQueryReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.PageResp}	"返回信息"
// @Router		/blog-api/v1/comment/find_comment_reply_list [POST]
func (s *CommentController) FindCommentReplyList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.CommentQueryReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCommentService(s.svcCtx).FindCommentReplyList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Comment
// @Summary		"创建评论"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.CommentNewReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.Comment}	"返回信息"
// @Router		/blog-api/v1/comment/add_comment [POST]
func (s *CommentController) AddComment(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.CommentNewReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCommentService(s.svcCtx).AddComment(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Comment
// @Summary		"点赞评论"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.IdReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.EmptyResp}	"返回信息"
// @Router		/blog-api/v1/comment/like_comment [POST]
func (s *CommentController) LikeComment(c *gin.Context) {
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

	data, err := service.NewCommentService(s.svcCtx).LikeComment(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}

// @Tags		Comment
// @Summary		"更新评论"
// @accept		application/json
// @Produce		application/json
// @Param		data	body		dto.UpdateCommentReq		true	"请求参数"
// @Success		200		{object}	response.Body{data=dto.Comment}	"返回信息"
// @Router		/blog-api/v1/comment/update_comment [POST]
func (s *CommentController) UpdateComment(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *dto.UpdateCommentReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	data, err := service.NewCommentService(s.svcCtx).UpdateComment(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, data)
}
