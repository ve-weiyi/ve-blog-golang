package entity

import "time"

// TableNameOperationLog return the table name of <operation_log>
const TableNameOperationLog = "operation_log"

// OperationLog mapped from table <operation_log>
type OperationLog struct {
	ID            int       `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:主键id" json:"id"`            // 主键id
	OptModule     string    `gorm:"column:opt_module;type:varchar(20);not null;comment:操作模块" json:"opt_module"`         // 操作模块
	OptType       string    `gorm:"column:opt_type;type:varchar(20);not null;comment:操作类型" json:"opt_type"`             // 操作类型
	OptMethod     string    `gorm:"column:opt_method;type:varchar(255);not null;comment:操作方法" json:"opt_method"`        // 操作方法
	OptDesc       string    `gorm:"column:opt_desc;type:varchar(255);not null;comment:操作描述" json:"opt_desc"`            // 操作描述
	Cost          string    `gorm:"column:cost;type:varchar(32);comment:耗时（ms）" json:"cost"`                            // 耗时（ms）
	Status        int       `gorm:"column:status;type:int;comment:响应状态码" json:"status"`                                 // 响应状态码
	RequestUrl    string    `gorm:"column:request_url;type:varchar(255);not null;comment:操作url" json:"request_url"`     // 操作url
	RequestMethod string    `gorm:"column:request_method;type:varchar(20);not null;comment:请求方式" json:"request_method"` // 请求方式
	RequestHeader string    `gorm:"column:request_header;type:varchar(1024);comment:请求头" json:"request_header"`         // 请求头
	RequestParam  string    `gorm:"column:request_param;type:text;not null;comment:请求参数" json:"request_param"`          // 请求参数
	ResponseData  string    `gorm:"column:response_data;type:text;not null;comment:返回数据" json:"response_data"`          // 返回数据
	UserID        int       `gorm:"column:user_id;type:int;not null;comment:用户id" json:"user_id"`                       // 用户id
	Nickname      string    `gorm:"column:nickname;type:varchar(50);not null;comment:用户昵称" json:"nickname"`             // 用户昵称
	IpAddress     string    `gorm:"column:ip_address;type:varchar(255);not null;comment:操作ip" json:"ip_address"`        // 操作ip
	IpSource      string    `gorm:"column:ip_source;type:varchar(255);not null;comment:操作地址" json:"ip_source"`          // 操作地址
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at"`            // 创建时间
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`                     // 更新时间
}

// TableName OperationLog's table name
func (*OperationLog) TableName() string {
	return TableNameOperationLog
}
