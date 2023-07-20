package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type CommentService struct {
	svcCtx *svc.ServiceContext
}

func NewCommentService(svcCtx *svc.ServiceContext) *CommentService {
	return &CommentService{
		svcCtx: svcCtx,
	}
}

// 创建Comment记录
func (s *CommentService) CreateComment(reqCtx *request.Context, comment *entity.Comment) (data *entity.Comment, err error) {
	return s.svcCtx.CommentRepository.CreateComment(reqCtx, comment)
}

// 删除Comment记录
func (s *CommentService) DeleteComment(reqCtx *request.Context, comment *entity.Comment) (rows int64, err error) {
	return s.svcCtx.CommentRepository.DeleteComment(reqCtx, comment)
}

// 更新Comment记录
func (s *CommentService) UpdateComment(reqCtx *request.Context, comment *entity.Comment) (data *entity.Comment, err error) {
	return s.svcCtx.CommentRepository.UpdateComment(reqCtx, comment)
}

// 查询Comment记录
func (s *CommentService) FindComment(reqCtx *request.Context, comment *entity.Comment) (data *entity.Comment, err error) {
	return s.svcCtx.CommentRepository.GetComment(reqCtx, comment.ID)
}

// 批量删除Comment记录
func (s *CommentService) DeleteCommentByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.CommentRepository.DeleteCommentByIds(reqCtx, ids)
}

// 分页获取Comment记录
func (s *CommentService) FindCommentList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.Comment, total int64, err error) {
	return s.svcCtx.CommentRepository.FindCommentList(reqCtx, page)
}
