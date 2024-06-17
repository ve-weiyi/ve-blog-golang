package entity

import "time"

// TableNameApi return the table name of <api>
const TableNameApi = "api"

// Api mapped from table <api>
type Api struct {
	Id        int64     `gorm:"column:id" json:"id" `                 // 主键id
	ParentId  int64     `gorm:"column:parent_id" json:"parent_id" `   // 分组id
	Name      string    `gorm:"column:name" json:"name" `             // api名称
	Path      string    `gorm:"column:path" json:"path" `             // api路径
	Method    string    `gorm:"column:method" json:"method" `         // api请求方法
	Traceable int64     `gorm:"column:traceable" json:"traceable" `   // 是否追溯操作记录 0需要，1是
	Status    int64     `gorm:"column:status" json:"status" `         // 状态 1开，2关
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" ` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" ` // 更新时间
}

// TableName Api 's table name
func (*Api) TableName() string {
	return TableNameApi
}
