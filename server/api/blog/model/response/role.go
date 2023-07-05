package response

import "github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"

type ApiTree struct {
	entity.Api
	Children []*ApiTree
}

type MenuTree struct {
	entity.Menu
	Children []*MenuTree
}

type RoleInfo struct {
	entity.Role
	MenuIdList     []int
	ResourceIdList []int
}
