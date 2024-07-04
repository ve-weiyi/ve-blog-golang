package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type PageService struct {
	svcCtx *svc.ServiceContext
}

func NewPageService(svcCtx *svc.ServiceContext) *PageService {
	return &PageService{
		svcCtx: svcCtx,
	}
}

// 创建Page记录
func (l *PageService) CreatePage(reqCtx *request.Context, page *entity.Page) (data *entity.Page, err error) {
	return l.svcCtx.PageRepository.Create(reqCtx, page)
}

// 更新Page记录
func (l *PageService) UpdatePage(reqCtx *request.Context, page *entity.Page) (data *entity.Page, err error) {
	return l.svcCtx.PageRepository.Update(reqCtx, page)
}

// 删除Page记录
func (l *PageService) DeletePage(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.PageRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Page记录
func (l *PageService) FindPage(reqCtx *request.Context, req *request.IdReq) (data *entity.Page, err error) {
	return l.svcCtx.PageRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Page记录
func (l *PageService) DeletePageList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.PageRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Page记录
func (l *PageService) FindPageList(reqCtx *request.Context, page *dto.PageQuery) (list []*entity.Page, total int64, err error) {
	p, s := page.PageClause()
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.PageRepository.FindList(reqCtx, p, s, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.PageRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
