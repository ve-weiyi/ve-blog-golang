package response

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

type ApiTree struct {
	entity.Api
	Children []*ApiTree `json:"children"`
}

type MenuTree struct {
	entity.Menu
	Children []*MenuTree `json:"children"`
}

type RoleInfo struct {
	entity.Role
	MenuIdList     []int `json:"menu_id_list"`
	ResourceIdList []int `json:"resource_id_list"`
}
