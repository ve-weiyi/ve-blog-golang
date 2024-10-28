package mail

type EmailMessage struct {
	To      []string `json:"to"`      // 目标邮箱号
	Subject string   `json:"subject"` // 主题
	Content string   `json:"content"` // 内容
	Type    int      `json:"type"`    // 0:普通邮件 1:需要抄送
}

// 投递邮件
type IEmailDeliver interface {
	DeliveryEmail(message *EmailMessage) error
}
