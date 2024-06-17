package response

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
)

type ApiDetailsDTO struct {
	entity.Api
	Children []*ApiDetailsDTO `json:"children"`
}

type RoleDetailsDTO struct {
	entity.Role
	MenuIdList     []int64 `json:"menu_id_list"`
	ResourceIdList []int64 `json:"resource_id_list"`
}

type RoleDTO struct {
	RoleName    string `json:"role_name"`
	RoleComment string `json:"role_comment"`
}
