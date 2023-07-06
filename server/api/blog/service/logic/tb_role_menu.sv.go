package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type RoleMenuService struct {
	svcCtx *svc.ServiceContext
}

func NewRoleMenuService(svcCtx *svc.ServiceContext) *RoleMenuService {
	return &RoleMenuService{
		svcCtx: svcCtx,
	}
}

// 创建RoleMenu记录
func (s *RoleMenuService) CreateRoleMenu(reqCtx *request.Context, roleMenu *entity.RoleMenu) (data *entity.RoleMenu, err error) {
	return s.svcCtx.RoleMenuRepository.CreateRoleMenu(roleMenu)
}

// 删除RoleMenu记录
func (s *RoleMenuService) DeleteRoleMenu(reqCtx *request.Context, roleMenu *entity.RoleMenu) (rows int64, err error) {
	return s.svcCtx.RoleMenuRepository.DeleteRoleMenu(roleMenu)
}

// 更新RoleMenu记录
func (s *RoleMenuService) UpdateRoleMenu(reqCtx *request.Context, roleMenu *entity.RoleMenu) (data *entity.RoleMenu, err error) {
	return s.svcCtx.RoleMenuRepository.UpdateRoleMenu(roleMenu)
}

// 根据id获取RoleMenu记录
func (s *RoleMenuService) FindRoleMenu(reqCtx *request.Context, id int) (data *entity.RoleMenu, err error) {
	return s.svcCtx.RoleMenuRepository.FindRoleMenu(id)
}

// 批量删除RoleMenu记录
func (s *RoleMenuService) DeleteRoleMenuByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.RoleMenuRepository.DeleteRoleMenuByIds(ids)
}

// 分页获取RoleMenu记录
func (s *RoleMenuService) GetRoleMenuList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.RoleMenu, total int64, err error) {
	return s.svcCtx.RoleMenuRepository.GetRoleMenuList(page)
}
