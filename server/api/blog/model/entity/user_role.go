package entity

// TableNameUserRole return the table name of <user_role>
const TableNameUserRole = "user_role"

// UserRole mapped from table <user_role>
type UserRole struct {
	Id     int64 `gorm:"column:id" json:"id" `           // 主键id
	UserId int64 `gorm:"column:user_id" json:"user_id" ` // 用户id
	RoleId int64 `gorm:"column:role_id" json:"role_id" ` // 角色id
}

// TableName UserRole 's table name
func (*UserRole) TableName() string {
	return TableNameUserRole
}
