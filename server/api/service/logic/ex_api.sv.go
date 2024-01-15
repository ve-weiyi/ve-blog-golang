package logic

import (
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/tools/apidocs/apiparser"
)

// 分页获取Api记录
func (s *ApiService) FindApiDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.ApiDetailsDTO, total int64, err error) {
	cond, args := page.ConditionClause()
	// 查询api信息
	apis, err := s.svcCtx.ApiRepository.FindALL(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	// to tree
	var tree response.ApiDetailsDTO
	tree.Children = s.getApiChildren(tree, apis)

	list = tree.Children
	return list, int64(len(list)), nil
}

func (s *ApiService) SyncApiList(reqCtx *request.Context, req interface{}) (data int64, err error) {
	ap := apiparser.NewSwaggerParser()
	apis, err := ap.ParseApiDocsByRoots(global.GetRuntimeRoot() + "server/docs")
	if err != nil {
		return 0, err
	}

	for _, api := range apis {
		if api.Router == "" {
			continue
		}

		// 已存在则跳过
		exist, _ := s.svcCtx.ApiRepository.First(reqCtx, "path = ? and method = ?", api.Router, api.Method)
		if exist != nil {
			continue
		}

		// 查找父分类，没有则创建
		parent, _ := s.svcCtx.ApiRepository.First(reqCtx, "name = ? and parent_id = ?", api.Tag, 0)
		if parent == nil {
			parent = &entity.Api{
				Name: api.Tag,
			}
			_, err = s.svcCtx.ApiRepository.Create(reqCtx, parent)
			if err != nil {
				return 0, err
			}
		}

		// 插入数据
		model := &entity.Api{
			Name:     api.Summary,
			Path:     api.Router,
			Method:   strings.ToUpper(api.Method),
			ParentID: parent.ID,
		}
		_, err = s.svcCtx.ApiRepository.Create(reqCtx, model)
		if err != nil {
			return 0, err
		}
		data++
	}

	return data, nil
}

func (s *ApiService) GetUserApis(reqCtx *request.Context, req interface{}) (data []*response.ApiDetailsDTO, err error) {
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
	menuMaps := make(map[int]*entity.Api)
	for _, item := range roles {
		menus, err := s.svcCtx.RoleRepository.FindRoleApis(reqCtx, item.ID)
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

	var list []*entity.Api
	for _, v := range menuMaps {
		list = append(list, v)
	}

	var out response.ApiDetailsDTO
	out.Children = s.getApiChildren(out, list)

	return out.Children, err
}

func (s *ApiService) getApiChildren(root response.ApiDetailsDTO, list []*entity.Api) (leafs []*response.ApiDetailsDTO) {
	for _, item := range list {
		if item.ParentID == root.ID {
			leaf := response.ApiDetailsDTO{
				Api:      *item,
				Children: nil,
			}
			leaf.Children = s.getApiChildren(leaf, list)
			leafs = append(leafs, &leaf)
		}
	}
	return leafs
}
