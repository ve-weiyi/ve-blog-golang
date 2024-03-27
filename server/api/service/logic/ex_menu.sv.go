package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

// 分页获取Menu记录
func (s *MenuService) FindMenuDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.MenuDetailsDTO, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()
	// 创建db
	menuList, err := s.svcCtx.MenuRepository.FindList(reqCtx, 0, 0, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	// to tree
	var tree response.MenuDetailsDTO
	tree.Children = s.getMenuChildren(tree, menuList)

	list = tree.Children
	return list, int64(len(list)), nil
}

func (s *MenuService) SyncMenuList(reqCtx *request.Context, req *request.SyncMenuRequest) (data int64, err error) {

	for _, item := range req.Menus {
		// 已存在则跳过
		exist, _ := s.svcCtx.MenuRepository.First(reqCtx, "path = ?", item.Path)
		if exist == nil {

			// 插入数据
			exist = &entity.Menu{
				Title:     item.Meta.Title,
				Path:      item.Path,
				Name:      item.Name,
				Component: jsonconv.ObjectToJson(item.Component),
				Redirect:  item.Redirect,
				Type:      item.Type,
				Meta:      jsonconv.ObjectToJson(item.Meta),
			}
			_, err = s.svcCtx.MenuRepository.Create(reqCtx, exist)
			if err != nil {
				return data, err
			}

			data++
		}

		for i, child := range item.Children {
			// 已存在则跳过
			menu, _ := s.svcCtx.MenuRepository.First(reqCtx, "path = ?", child.Path)
			if menu == nil {
				if child.Meta.Rank == 0 {
					child.Meta.Rank = i
				}

				// 插入数据
				menu = &entity.Menu{
					ParentID:  exist.ID,
					Title:     child.Meta.Title,
					Path:      child.Path,
					Name:      child.Name,
					Component: jsonconv.ObjectToJson(child.Component),
					Redirect:  child.Redirect,
					Type:      child.Type,
					Meta:      jsonconv.ObjectToJson(child.Meta),
				}
				_, err = s.svcCtx.MenuRepository.Create(reqCtx, menu)
				if err != nil {
					return data, err
				}

				data++
			}
		}

	}

	return data, err
}

func (s *MenuService) GetUserMenus(reqCtx *request.Context, req interface{}) (data []*response.MenuDetailsDTO, err error) {
	//查询用户信息
	account, err := s.svcCtx.UserAccountRepository.First(reqCtx, "id = ?", reqCtx.UID)
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
			leaf := convertMenu(item)
			leaf.Children = s.getMenuChildren(*leaf, list)
			leafs = append(leafs, leaf)
		}
	}
	return leafs
}
