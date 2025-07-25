package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type CommentService struct {
	svcCtx *svctx.ServiceContext
}

func NewCommentService(svcCtx *svctx.ServiceContext) *CommentService {
	return &CommentService{
		svcCtx: svcCtx,
	}
}

// 删除评论
func (s *CommentService) DeletesComment(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 查询评论列表(后台)
func (s *CommentService) FindCommentBackList(reqCtx *request.Context, in *dto.CommentQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 更新评论审核状态
func (s *CommentService) UpdateCommentReview(reqCtx *request.Context, in *dto.CommentReviewReq) (out *dto.BatchResp, err error) {
	// todo

	return
}
