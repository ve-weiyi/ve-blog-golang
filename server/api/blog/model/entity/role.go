package entity

import "time"

// TableNameRole return the table name of <role>
const TableNameRole = "role"

// Role mapped from table <role>
type Role struct {
	Id          int64     `gorm:"column:id" json:"id" `                     // 主键id
	ParentId    int64     `gorm:"column:parent_id" json:"parent_id" `       // 父角色id
	RoleDomain  string    `gorm:"column:role_domain" json:"role_domain" `   // 角色域
	RoleName    string    `gorm:"column:role_name" json:"role_name" `       // 角色名
	RoleComment string    `gorm:"column:role_comment" json:"role_comment" ` // 角色备注
	IsDisable   int64     `gorm:"column:is_disable" json:"is_disable" `     // 是否禁用  0否 1是
	IsDefault   int64     `gorm:"column:is_default" json:"is_default" `     // 是否默认角色 0否 1是
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at" `     // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at" `     // 更新时间
}

// TableName Role 's table name
func (*Role) TableName() string {
	return TableNameRole
}
