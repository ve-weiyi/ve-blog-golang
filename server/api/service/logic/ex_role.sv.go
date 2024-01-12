package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// 分页获取Role记录
func (s *RoleService) FindRoleDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.RoleDetailsDTO, total int64, err error) {

	roles, total, err := s.FindRoleList(reqCtx, page)

	for _, role := range roles {
		var menuIds []int
		menus, err := s.svcCtx.RoleRepository.FindRoleMenus(reqCtx, role.ID)
		if err != nil {
			return nil, 0, err
		}

		var apiIds []int
		apis, err := s.svcCtx.RoleRepository.FindRoleApis(reqCtx, role.ID)
		if err != nil {
			return nil, 0, err
		}

		for _, menu := range menus {
			menuIds = append(menuIds, menu.ID)
		}
		for _, api := range apis {
			apiIds = append(apiIds, api.ID)
		}

		r := response.RoleDetailsDTO{
			Role:           *role,
			MenuIdList:     menuIds,
			ResourceIdList: apiIds,
		}
		list = append(list, &r)
	}

	return list, total, nil
}

// 设置角色菜单
func (s *RoleService) UpdateRoleMenus(reqCtx *request.Context, req *request.UpdateRoleMenusReq) (data interface{}, err error) {
	// 重置角色菜单权限
	_, _, err = s.svcCtx.RoleRepository.UpdateRoleResources(reqCtx, req.RoleId, req.MenuIds)
	if err != nil {
		return nil, err
	}

	return nil, err
}

// 设置角色菜单
func (s *RoleService) UpdateRoleResources(reqCtx *request.Context, req *request.UpdateRoleApisReq) (data interface{}, err error) {
	// 重置角色接口权限
	role, _, err := s.svcCtx.RoleRepository.UpdateRoleResources(reqCtx, req.RoleId, req.ResourceIds)
	if err != nil {
		return nil, err
	}

	// 查询资源列表
	resources, err := s.svcCtx.ApiRepository.FindALL(reqCtx, "api_id in (?)", req.ResourceIds)
	if err != nil {
		return nil, err
	}
	// 添加角色规则
	rbac := s.svcCtx.RBAC
	rbac.DeleteRolePolicy(role.RoleName, role.RoleDomain)
	rbac.AddRolePolicy(role.RoleName, role.RoleDomain, resources)

	return nil, err
}
