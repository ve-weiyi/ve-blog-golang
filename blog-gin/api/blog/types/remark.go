package types

type NewRemarkReq struct {
	MessageContent string `json:"message_content"` // 留言内容
}

type QueryRemarkReq struct {
	PageQuery
}

type Remark struct {
	Id             int64       `json:"id,optional"`     // 主键id
	UserId         string      `json:"user_id"`         // 用户id
	TerminalId     string      `json:"terminal_id"`     // 终端id
	MessageContent string      `json:"message_content"` // 留言内容
	IpAddress      string      `json:"ip_address"`      // IP地址
	IpSource       string      `json:"ip_source"`       // IP归属地
	IsReview       int64       `json:"is_review"`       // 是否审核
	CreatedAt      int64       `json:"created_at"`      // 发布时间
	UpdatedAt      int64       `json:"updated_at"`      // 更新时间
	User           *UserInfoVO `json:"user"`            // 用户信息
}
