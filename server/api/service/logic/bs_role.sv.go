package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type RoleService struct {
	svcCtx *svc.ServiceContext
}

func NewRoleService(svcCtx *svc.ServiceContext) *RoleService {
	return &RoleService{
		svcCtx: svcCtx,
	}
}

// 创建Role记录
func (s *RoleService) CreateRole(reqCtx *request.Context, role *entity.Role) (data *entity.Role, err error) {
	return s.svcCtx.RoleRepository.CreateRole(reqCtx, role)
}

// 更新Role记录
func (s *RoleService) UpdateRole(reqCtx *request.Context, role *entity.Role) (data *entity.Role, err error) {
	return s.svcCtx.RoleRepository.UpdateRole(reqCtx, role)
}

// 删除Role记录
func (s *RoleService) DeleteRole(reqCtx *request.Context, id int) (rows int64, err error) {
	return s.svcCtx.RoleRepository.DeleteRole(reqCtx, id)
}

// 查询Role记录
func (s *RoleService) FindRole(reqCtx *request.Context, id int) (data *entity.Role, err error) {
	return s.svcCtx.RoleRepository.FindRole(reqCtx, id)
}

// 批量删除Role记录
func (s *RoleService) DeleteRoleByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.RoleRepository.DeleteRoleByIds(reqCtx, ids)
}

// 分页获取Role记录
func (s *RoleService) FindRoleList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Role, total int64, err error) {
	return s.svcCtx.RoleRepository.FindRoleList(reqCtx, page)
}
