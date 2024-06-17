package entity

// TableNameRoleMenu return the table name of <role_menu>
const TableNameRoleMenu = "role_menu"

// RoleMenu mapped from table <role_menu>
type RoleMenu struct {
	Id     int64 `gorm:"column:id" json:"id" `           // 主键id
	RoleId int64 `gorm:"column:role_id" json:"role_id" ` // 角色id
	MenuId int64 `gorm:"column:menu_id" json:"menu_id" ` // 菜单id
}

// TableName RoleMenu 's table name
func (*RoleMenu) TableName() string {
	return TableNameRoleMenu
}
