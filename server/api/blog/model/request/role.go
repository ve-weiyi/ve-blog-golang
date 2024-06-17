package request

type UpdateUserRolesReq struct {
	UserId  int64   `json:"user_id"`
	RoleIds []int64 `json:"role_ids"`
}

type UpdateRoleMenusReq struct {
	RoleId  int64   `json:"role_id"`
	MenuIds []int64 `json:"menu_ids"`
}

type UpdateRoleApisReq struct {
	RoleId      int64   `json:"role_id"`
	ResourceIds []int64 `json:"resource_ids"`
}
