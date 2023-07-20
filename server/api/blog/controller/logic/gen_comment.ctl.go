package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/controller/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/controller"
)

type CommentController struct {
	controller.BaseController
	svcCtx *svc.ControllerContext
}

func NewCommentController(svcCtx *svc.ControllerContext) *CommentController {
	return &CommentController{
		svcCtx:         svcCtx,
		BaseController: controller.NewBaseController(svcCtx),
	}
}

// @Tags		Comment
// @Summary		创建评论
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body		entity.Comment							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.Comment}	"返回信息"
// @Router		/comment/create [post]
func (s *CommentController) CreateComment(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var comment entity.Comment
	err = s.ShouldBind(c, &comment)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.CommentService.CreateComment(reqCtx, &comment)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags		Comment
// @Summary		删除评论
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param		data	body	 	entity.Comment 		true "请求body"
// @Success		200		{object}	response.Response{}		"返回信息"
// @Router		/comment/delete [delete]
func (s *CommentController) DeleteComment(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var comment entity.Comment
	err = s.ShouldBind(c, &comment)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.CommentService.DeleteComment(reqCtx, &comment)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Comment
// @Summary		更新评论
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data	body 	 	entity.Comment							true		"请求参数"
// @Success		200		{object}	response.Response{data=entity.Comment}	"返回信息"
// @Router 		/comment/update [put]
func (s *CommentController) UpdateComment(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var comment entity.Comment
	err = s.ShouldBind(c, &comment)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.CommentService.UpdateComment(reqCtx, &comment)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Comment
// @Summary		查询评论
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	data		query		entity.Comment							true		"请求参数"
// @Success		200			{object}	response.Response{data=entity.Comment}	"返回信息"
// @Router 		/comment/find [get]
func (s *CommentController) FindComment(c *gin.Context) {
	reqCtx, err := s.GetRequestContext(c)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	var comment entity.Comment
	err = s.ShouldBind(c, &comment)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	data, err := s.svcCtx.CommentService.FindComment(reqCtx, &comment)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Comment
// @Summary		批量删除评论
// @Security 	ApiKeyAuth
// @accept 	 	application/json
// @Produce		application/json
// @Param		data 	body		[]int 				true "删除id列表"
// @Success		200		{object}	response.Response{}	"返回信息"
// @Router		/comment/deleteByIds [delete]
func (s *CommentController) DeleteCommentByIds(c *gin.Context) {
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

	data, err := s.svcCtx.CommentService.DeleteCommentByIds(reqCtx, IDS)
	if err != nil {
		s.ResponseError(c, err)
		return
	}

	s.ResponseOk(c, data)
}

// @Tags 	 	Comment
// @Summary		分页获取评论列表
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce		application/json
// @Param 	 	page 	body		request.PageInfo 	true "分页参数"
// @Success		200		{object}	response.Response{data=response.PageResult{list=[]entity.Comment}}	"返回信息"
// @Router		/comment/list [post]
func (s *CommentController) FindCommentList(c *gin.Context) {
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

	list, total, err := s.svcCtx.CommentService.FindCommentList(reqCtx, &page)
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