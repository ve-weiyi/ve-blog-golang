package request

type UpdateUserRoles struct {
	UserId  int   `json:"user_id"`
	RoleIds []int `json:"role_ids"`
}

type UpdateRoleMenus struct {
	RoleId  int   `json:"role_id"`
	MenuIds []int `json:"menu_ids"`
}

type UpdateRoleResources struct {
	RoleId      int   `json:"role_id"`
	ResourceIds []int `json:"resource_ids"`
}
