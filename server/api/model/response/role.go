package response

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

type ApiDetails struct {
	entity.Api
	Children []*ApiDetails `json:"children"`
}

type MenuDetails struct {
	entity.Menu
	Children []*MenuDetails `json:"children"`
}

type RoleInfo struct {
	entity.Role
	MenuIdList     []int `json:"menu_id_list"`
	ResourceIdList []int `json:"resource_id_list"`
}

type RoleDTO struct {
	RoleName    string `json:"role_name"`
	RoleComment string `json:"role_comment"`
}
