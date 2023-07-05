package logic

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
)

type UserRoleService struct {
	svcCtx *svc.ServiceContext
}

func NewUserRoleService(svcCtx *svc.ServiceContext) *UserRoleService {
	return &UserRoleService{
		svcCtx: svcCtx,
	}
}

// 创建UserRole记录
func (s *UserRoleService) CreateUserRole(reqCtx *request.Context, userRole *entity.UserRole) (data *entity.UserRole, err error) {
	return s.svcCtx.UserRoleRepository.CreateUserRole(userRole)
}

// 删除UserRole记录
func (s *UserRoleService) DeleteUserRole(reqCtx *request.Context, userRole *entity.UserRole) (rows int64, err error) {
	return s.svcCtx.UserRoleRepository.DeleteUserRole(userRole)
}

// 更新UserRole记录
func (s *UserRoleService) UpdateUserRole(reqCtx *request.Context, userRole *entity.UserRole) (data *entity.UserRole, err error) {
	return s.svcCtx.UserRoleRepository.UpdateUserRole(userRole)
}

// 根据id获取UserRole记录
func (s *UserRoleService) FindUserRole(reqCtx *request.Context, id int) (data *entity.UserRole, err error) {
	return s.svcCtx.UserRoleRepository.FindUserRole(id)
}

// 批量删除UserRole记录
func (s *UserRoleService) DeleteUserRoleByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.UserRoleRepository.DeleteUserRoleByIds(ids)
}

// 分页获取UserRole记录
func (s *UserRoleService) GetUserRoleList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.UserRole, total int64, err error) {
	return s.svcCtx.UserRoleRepository.GetUserRoleList(page)
}
