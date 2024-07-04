package entity

import "time"

// TableNameMenu return the table name of <menu>
const TableNameMenu = "menu"

// Menu mapped from table <menu>
type Menu struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // 主键
	ParentId  int64     `gorm:"column:parent_id" json:"parent_id" `   // 父id
	Title     string    `gorm:"column:title" json:"title" `           // 菜单标题
	Path      string    `gorm:"column:path" json:"path" `             // 路由路径
	Name      string    `gorm:"column:name" json:"name" `             // 路由名称
	Component string    `gorm:"column:component" json:"component" `   // 路由组件
	Redirect  string    `gorm:"column:redirect" json:"redirect" `     // 路由重定向
	Type      int64     `gorm:"column:type" json:"type" `             // 菜单类型
	Rank      int64     `gorm:"column:rank" json:"rank" `             // 排序
	Extra     string    `gorm:"column:extra" json:"extra" `           // 菜单元数据
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName Menu 's table name
func (*Menu) TableName() string {
	return TableNameMenu
}
