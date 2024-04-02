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
	return s.svcCtx.CommentRepository.Create(reqCtx, comment)
}

// 更新Comment记录
func (s *CommentService) UpdateComment(reqCtx *request.Context, comment *entity.Comment) (data *entity.Comment, err error) {
	return s.svcCtx.CommentRepository.Update(reqCtx, comment)
}

// 删除Comment记录
func (s *CommentService) DeleteComment(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return s.svcCtx.CommentRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Comment记录
func (s *CommentService) FindComment(reqCtx *request.Context, req *request.IdReq) (data *entity.Comment, err error) {
	return s.svcCtx.CommentRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Comment记录
func (s *CommentService) DeleteCommentList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return s.svcCtx.CommentRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Comment记录
func (s *CommentService) FindCommentList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Comment, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = s.svcCtx.CommentRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.CommentRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
