package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
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

// 更新Comment记录
func (s *CommentService) UpdateComment(reqCtx *request.Context, comment *entity.Comment) (data *entity.Comment, err error) {
	return s.svcCtx.CommentRepository.UpdateComment(reqCtx, comment)
}

// 删除Comment记录
func (s *CommentService) DeleteComment(reqCtx *request.Context, id int) (rows int, err error) {
	return s.svcCtx.CommentRepository.DeleteCommentById(reqCtx, id)
}

// 查询Comment记录
func (s *CommentService) FindComment(reqCtx *request.Context, id int) (data *entity.Comment, err error) {
	return s.svcCtx.CommentRepository.FindCommentById(reqCtx, id)
}

// 批量删除Comment记录
func (s *CommentService) DeleteCommentByIds(reqCtx *request.Context, ids []int) (rows int, err error) {
	return s.svcCtx.CommentRepository.DeleteCommentByIds(reqCtx, ids)
}

// 分页获取Comment记录
func (s *CommentService) FindCommentList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Comment, total int64, err error) {
	list, err = s.svcCtx.CommentRepository.FindCommentList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.CommentRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
