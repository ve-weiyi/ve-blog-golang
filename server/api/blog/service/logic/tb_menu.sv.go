package logic

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
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

// 删除Menu记录
func (s *MenuService) DeleteMenu(reqCtx *request.Context, menu *entity.Menu) (rows int64, err error) {
	return s.svcCtx.MenuRepository.DeleteMenu(reqCtx, menu)
}

// 更新Menu记录
func (s *MenuService) UpdateMenu(reqCtx *request.Context, menu *entity.Menu) (data *entity.Menu, err error) {
	return s.svcCtx.MenuRepository.UpdateMenu(reqCtx, menu)
}

// 查询Menu记录
func (s *MenuService) GetMenu(reqCtx *request.Context, menu *entity.Menu) (data *entity.Menu, err error) {
	return s.svcCtx.MenuRepository.GetMenu(reqCtx, menu.ID)
}

// 批量删除Menu记录
func (s *MenuService) DeleteMenuByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.MenuRepository.DeleteMenuByIds(reqCtx, ids)
}

// 分页获取Menu记录
func (s *MenuService) FindMenuList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.Menu, total int64, err error) {
	return s.svcCtx.MenuRepository.FindMenuList(reqCtx, page)
}
