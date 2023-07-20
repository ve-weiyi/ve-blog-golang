package entity

import "time"

// TableNameMenu return the table name of <menu>
const TableNameMenu = "menu"

// Menu mapped from table <menu>
type Menu struct {
	ID        int       `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:主键" json:"id"`                             // 主键
	Name      string    `gorm:"column:name;type:varchar(32);not null;comment:菜单名" json:"name"`                                     // 菜单名
	Path      string    `gorm:"column:path;type:varchar(64);not null;comment:菜单路径" json:"path"`                                    // 菜单路径
	Component string    `gorm:"column:component;type:varchar(64);not null;comment:组件" json:"component"`                            // 组件
	Icon      string    `gorm:"column:icon;type:varchar(64);not null;comment:菜单icon" json:"icon"`                                  // 菜单icon
	Rank      int       `gorm:"column:rank;type:tinyint;not null;comment:排序" json:"rank"`                                          // 排序
	ParentID  int       `gorm:"column:parent_id;type:int;comment:父id" json:"parent_id"`                                            // 父id
	IsHidden  bool      `gorm:"column:is_hidden;type:tinyint(1);not null;comment:是否隐藏  0否1是" json:"is_hidden"`                     // 是否隐藏  0否1是
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
}

// TableName Menu's table name
func (*Menu) TableName() string {
	return TableNameMenu
}
