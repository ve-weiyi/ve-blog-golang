package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
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
func (s *MenuService) CreateMenu(reqCtx *request.Context, menu *entity.Menu) (data *entity.Menu, err error) {
	return s.svcCtx.MenuRepository.CreateMenu(reqCtx, menu)
}

// 更新Menu记录
func (s *MenuService) UpdateMenu(reqCtx *request.Context, menu *entity.Menu) (data *entity.Menu, err error) {
	return s.svcCtx.MenuRepository.UpdateMenu(reqCtx, menu)
}

// 删除Menu记录
func (s *MenuService) DeleteMenu(reqCtx *request.Context, id int) (rows int, err error) {
	return s.svcCtx.MenuRepository.DeleteMenuById(reqCtx, id)
}

// 查询Menu记录
func (s *MenuService) FindMenu(reqCtx *request.Context, id int) (data *entity.Menu, err error) {
	return s.svcCtx.MenuRepository.FindMenuById(reqCtx, id)
}

// 批量删除Menu记录
func (s *MenuService) DeleteMenuByIds(reqCtx *request.Context, ids []int) (rows int, err error) {
	return s.svcCtx.MenuRepository.DeleteMenuByIds(reqCtx, ids)
}

// 分页获取Menu记录
func (s *MenuService) FindMenuList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Menu, total int64, err error) {
	list, err = s.svcCtx.MenuRepository.FindMenuList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.MenuRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
