package response

import (
	entity2 "github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

type ApiTree struct {
	entity2.Api
	Children []*ApiTree
}

type MenuTree struct {
	entity2.Menu
	Children []*MenuTree
}

type RoleInfo struct {
	entity2.Role
	MenuIdList     []int
	ResourceIdList []int
}
