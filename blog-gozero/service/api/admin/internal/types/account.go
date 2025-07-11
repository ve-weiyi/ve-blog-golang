// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4

package types

type AccountQuery struct {
	PageQuery
	Username string   `json:"username,optional"`
	Nickname string   `json:"nickname,optional"`
	Email    string   `json:"email,optional"`
	Phone    string   `json:"phone,optional"`
	Status   int64    `json:"status,optional"`   // 状态: -1删除 0正常 1禁用
	UserIds  []string `json:"user_ids,optional"` // 用户ID
}

type UpdateAccountPasswordReq struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type UpdateAccountRolesReq struct {
	UserId  string  `json:"user_id"`
	RoleIds []int64 `json:"role_ids"`
}

type UpdateAccountStatusReq struct {
	UserId string `json:"user_id"`
	Status int64  `json:"status"` // 状态: -1删除 0正常 1禁用
}
