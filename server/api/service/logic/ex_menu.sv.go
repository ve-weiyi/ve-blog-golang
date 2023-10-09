package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// 分页获取Menu记录
func (s *MenuService) FindMenuDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.MenuDetailsDTO, total int64, err error) {
	// 创建db
	menuList, err := s.svcCtx.MenuRepository.FindMenuList(reqCtx, nil, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	// to tree
	var tree response.MenuDetailsDTO
	tree.Children = s.getMenuChildren(tree, menuList)

	list = tree.Children
	return list, int64(len(list)), nil
}

func (s *MenuService) GetUserMenus(reqCtx *request.Context, req interface{}) (data []*response.MenuDetailsDTO, err error) {
	//查询用户信息
	account, err := s.svcCtx.UserAccountRepository.FindUserAccountById(reqCtx, reqCtx.UID)
	if err != nil {
		return nil, err
	}

	//查询用户角色
	roles, err := s.svcCtx.RoleRepository.FindUserRoles(reqCtx, account.ID)
	if err != nil {
		return nil, err
	}

	//查询角色权限,取交集
	menuMaps := make(map[int]*entity.Menu)
	for _, item := range roles {
		menus, err := s.svcCtx.RoleRepository.FindRoleMenus(reqCtx, item.ID)
		if err != nil {
			return nil, err
		}
		// 去重
		for _, m := range menus {
			if _, ok := menuMaps[m.ID]; !ok {
				menuMaps[m.ID] = m
			}
		}
	}

	var list []*entity.Menu
	for _, v := range menuMaps {
		list = append(list, v)
	}

	var out response.MenuDetailsDTO
	out.Children = s.getMenuChildren(out, list)

	return out.Children, err
}

func (s *MenuService) getMenuChildren(root response.MenuDetailsDTO, list []*entity.Menu) (leafs []*response.MenuDetailsDTO) {
	for _, item := range list {
		if item.ParentID == root.ID {
			leaf := response.MenuDetailsDTO{
				Menu:     *item,
				Children: nil,
			}
			leaf.Children = s.getMenuChildren(leaf, list)
			leafs = append(leafs, &leaf)
		}
	}
	return leafs
}
