package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type RoleService struct {
	svcCtx *svctx.ServiceContext
}

func NewRoleService(svcCtx *svctx.ServiceContext) *RoleService {
	return &RoleService{
		svcCtx: svcCtx,
	}
}

// 创建角色
func (s *RoleService) AddRole(reqCtx *request.Context, in *dto.RoleNewReq) (out *dto.RoleBackVO, err error) {
	// todo

	return
}

// 删除角色
func (s *RoleService) DeletesRole(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取角色列表
func (s *RoleService) FindRoleList(reqCtx *request.Context, in *dto.RoleQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 获取角色资源列表
func (s *RoleService) FindRoleResources(reqCtx *request.Context, in *dto.IdReq) (out *dto.RoleResourcesResp, err error) {
	// todo

	return
}

// 更新角色
func (s *RoleService) UpdateRole(reqCtx *request.Context, in *dto.RoleNewReq) (out *dto.RoleBackVO, err error) {
	// todo

	return
}

// 更新角色接口权限
func (s *RoleService) UpdateRoleApis(reqCtx *request.Context, in *dto.UpdateRoleApisReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 更新角色菜单权限
func (s *RoleService) UpdateRoleMenus(reqCtx *request.Context, in *dto.UpdateRoleMenusReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}
