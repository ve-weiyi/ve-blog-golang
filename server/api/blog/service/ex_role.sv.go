package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
)

// 分页获取Role记录
func (l *RoleService) FindRoleDetailsList(reqCtx *request.Context, page *dto.PageQuery) (list []*dto.RoleDetailsDTO, total int64, err error) {

	// 查找角色列表
	roles, total, err := l.FindRoleList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}
	var roleIds []int64
	var menuMap = make(map[int64][]int64)
	var apiMap = make(map[int64][]int64)
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
	}

	// 查找角色菜单
	menus, err := l.svcCtx.RoleMenuRepository.FindALL(reqCtx, "role_id in (?)", roleIds)
	if err != nil {
		return nil, 0, err
	}

	for _, menu := range menus {
		menuMap[menu.RoleId] = append(menuMap[menu.RoleId], menu.MenuId)
	}

	// 查找角色资源
	apis, err := l.svcCtx.RoleApiRepository.FindALL(reqCtx, "role_id in (?)", roleIds)
	if err != nil {
		return nil, 0, err
	}

	for _, api := range apis {
		apiMap[api.RoleId] = append(apiMap[api.RoleId], api.ApiId)
	}

	// 拼装数据
	for _, role := range roles {
		menuIds := menuMap[role.Id]
		apiIds := apiMap[role.Id]

		r := dto.RoleDetailsDTO{
			Role:           *role,
			MenuIdList:     menuIds,
			ResourceIdList: apiIds,
		}
		list = append(list, &r)
	}

	return list, total, nil
}

// 设置角色菜单
func (l *RoleService) UpdateRoleMenus(reqCtx *request.Context, req *dto.UpdateRoleMenusReq) (data interface{}, err error) {
	// 重置角色菜单权限
	menu, _, err := l.svcCtx.RoleRepository.UpdateRoleMenus(reqCtx, req.RoleId, req.MenuIds)
	if err != nil {
		return nil, err
	}

	return menu, err
}

// 设置角色菜单
func (l *RoleService) UpdateRoleResources(reqCtx *request.Context, req *dto.UpdateRoleApisReq) (data interface{}, err error) {
	// 重置角色接口权限
	role, _, err := l.svcCtx.RoleRepository.UpdateRoleResources(reqCtx, req.RoleId, req.ResourceIds)
	if err != nil {
		return nil, err
	}

	//// 查询资源列表
	//resources, err := l.svcCtx.ApiRepository.FindALL(reqCtx, "api_id in (?)", req.ResourceIds)
	//if err != nil {
	//	return nil, err
	//}
	//// 添加角色规则
	//rbac := l.svcCtx.RBAC
	//rbac.DeleteRolePolicy(role.RoleName, role.RoleDomain)
	//rbac.AddRolePolicy(role.RoleName, role.RoleDomain, resources)

	return role, err
}
