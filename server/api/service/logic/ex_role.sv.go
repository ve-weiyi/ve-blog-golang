package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/sqlx"
)

// 分页获取Role记录
func (s *RoleService) FindRoleDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.RoleInfo, total int64, err error) {

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

		r := response.RoleInfo{
			Role:           *role,
			MenuIdList:     menuIds,
			ResourceIdList: apiIds,
		}
		list = append(list, &r)
	}

	return list, total, nil
}

// 设置角色菜单
func (s *RoleService) UpdateRoleMenus(reqCtx *request.Context, req *request.UpdateRoleMenus) (data interface{}, err error) {
	// 重置角色菜单权限
	_, _, err = s.svcCtx.RoleRepository.UpdateRoleResources(reqCtx, req.RoleId, req.MenuIds)
	if err != nil {
		return nil, err
	}

	return nil, err
}

// 设置角色菜单
func (s *RoleService) UpdateRoleResources(reqCtx *request.Context, req *request.UpdateRoleResources) (data interface{}, err error) {
	// 重置角色接口权限
	role, _, err := s.svcCtx.RoleRepository.UpdateRoleResources(reqCtx, req.RoleId, req.ResourceIds)
	if err != nil {
		return nil, err
	}

	// 查询资源列表
	page := &request.PageQuery{Conditions: []*sqlx.Condition{
		{
			Flag:  "and",
			Field: "api_id",
			Rule:  "in",
			Value: req.ResourceIds,
		},
	},
	}

	resources, err := s.svcCtx.ApiRepository.FindApiList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, err
	}
	// 添加角色规则
	rbac := s.svcCtx.RBAC
	rbac.DeleteRolePolicy(role.RoleName, role.RoleDomain)
	rbac.AddRolePolicy(role.RoleName, role.RoleDomain, resources)

	return nil, err
}
