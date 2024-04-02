package entity

import "time"

// TableNameOperationLog return the table name of <operation_log>
const TableNameOperationLog = "operation_log"

// OperationLog mapped from table <operation_log>
type OperationLog struct {
	Id             int64     `gorm:"column:id" json:"id" `                           // 主键id
	UserId         int64     `gorm:"column:user_id" json:"user_id" `                 // 用户id
	Nickname       string    `gorm:"column:nickname" json:"nickname" `               // 用户昵称
	IpAddress      string    `gorm:"column:ip_address" json:"ip_address" `           // 操作ip
	IpSource       string    `gorm:"column:ip_source" json:"ip_source" `             // 操作地址
	OptModule      string    `gorm:"column:opt_module" json:"opt_module" `           // 操作模块
	OptDesc        string    `gorm:"column:opt_desc" json:"opt_desc" `               // 操作描述
	RequestUrl     string    `gorm:"column:request_url" json:"request_url" `         // 请求地址
	RequestMethod  string    `gorm:"column:request_method" json:"request_method" `   // 请求方式
	RequestHeader  string    `gorm:"column:request_header" json:"request_header" `   // 请求头参数
	RequestData    string    `gorm:"column:request_data" json:"request_data" `       // 请求参数
	ResponseData   string    `gorm:"column:response_data" json:"response_data" `     // 返回数据
	ResponseStatus int64     `gorm:"column:response_status" json:"response_status" ` // 响应状态码
	Cost           string    `gorm:"column:cost" json:"cost" `                       // 耗时（ms）
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at" `           // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at" `           // 更新时间
}

// TableName OperationLog 's table name
func (*OperationLog) TableName() string {
	return TableNameOperationLog
}
