package entity

// TableNameRoleApi return the table name of <role_api>
const TableNameRoleApi = "role_api"

// RoleApi mapped from table <role_api>
type RoleApi struct {
	Id     int64 `gorm:"column:id" json:"id" `           // 主键id
	RoleId int64 `gorm:"column:role_id" json:"role_id" ` // 角色id
	ApiId  int64 `gorm:"column:api_id" json:"api_id" `   // 接口id
}

// TableName RoleApi 's table name
func (*RoleApi) TableName() string {
	return TableNameRoleApi
}
