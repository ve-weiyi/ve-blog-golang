package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
)

// 分页获取Api记录
func (s *ApiService) GetAllApiList(reqCtx *request.Context, page *request.PageInfo) (list []*response.ApiTree, total int64, err error) {
	// 查询api信息
	apis, _, err := s.svcCtx.ApiRepository.FindApiList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	// to tree
	var tree response.ApiTree
	tree.Children = s.getApiChildren(tree, apis)

	list = tree.Children
	return list, int64(len(list)), nil
}

func (s *ApiService) GetUserApis(reqCtx *request.Context, req *request.EmptyRequest) (data []*response.ApiTree, err error) {
	//查询用户信息
	account, err := s.svcCtx.UserAccountRepository.GetUserAccount(reqCtx, reqCtx.UID)
	if err != nil {
		return nil, err
	}

	//查询用户角色
	roles, err := s.svcCtx.RoleRepository.FindUserRoles(account.ID)
	if err != nil {
		return nil, err
	}

	//查询角色权限,取交集
	var menuMaps map[int]*entity.Api
	for _, item := range roles {
		menus, err := s.svcCtx.RoleRepository.FindRoleApis(item.ID)
		if err != nil {
			return nil, err
		}

		for _, m := range menus {
			if _, ok := menuMaps[m.ID]; !ok {
				menuMaps[m.ID] = m
			}
		}
	}

	var list []*entity.Api
	for _, v := range menuMaps {
		list = append(list, v)
	}

	var out response.ApiTree
	out.Children = s.getApiChildren(out, list)

	return out.Children, err
}

func (s *ApiService) getApiChildren(root response.ApiTree, list []*entity.Api) (leafs []*response.ApiTree) {
	for _, item := range list {
		if item.ParentID == root.ID {
			leaf := response.ApiTree{
				Api:      *item,
				Children: nil,
			}
			leaf.Children = s.getApiChildren(leaf, list)
			leafs = append(leafs, &leaf)
		}
	}
	return leafs
}
