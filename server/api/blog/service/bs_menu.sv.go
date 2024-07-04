package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
)

type MenuService struct {
	svcCtx *svc.ServiceContext
}

func NewMenuService(svcCtx *svc.ServiceContext) *MenuService {
	return &MenuService{
		svcCtx: svcCtx,
	}
}

// 创建Menu记录
func (l *MenuService) CreateMenu(reqCtx *request.Context, menu *entity.Menu) (data *entity.Menu, err error) {
	return l.svcCtx.MenuRepository.Create(reqCtx, menu)
}

// 更新Menu记录
func (l *MenuService) UpdateMenu(reqCtx *request.Context, menu *entity.Menu) (data *entity.Menu, err error) {
	return l.svcCtx.MenuRepository.Update(reqCtx, menu)
}

// 删除Menu记录
func (l *MenuService) DeleteMenu(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.MenuRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Menu记录
func (l *MenuService) FindMenu(reqCtx *request.Context, req *request.IdReq) (data *entity.Menu, err error) {
	return l.svcCtx.MenuRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Menu记录
func (l *MenuService) DeleteMenuList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.MenuRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Menu记录
func (l *MenuService) FindMenuList(reqCtx *request.Context, page *dto.PageQuery) (list []*entity.Menu, total int64, err error) {
	p, s := page.PageClause()
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.MenuRepository.FindList(reqCtx, p, s, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.MenuRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
