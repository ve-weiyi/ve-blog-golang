package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/jordan-wright/email"

	"github.com/ve-weiyi/ve-blog-golang/server/config/properties"
)

type EmailSender struct {
	Host     string   // 服务器地址
	Port     int      // 端口
	Username string   // 发件人
	Password string   // 密钥
	Nickname string   // 发件人昵称
	Deliver  []string // 抄送邮箱:多个以英文逗号分隔
	IsSSL    bool     // 是否使用 SSL/TLS
}

func NewEmailSender(cfg *properties.Email) *EmailSender {

	return &EmailSender{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.Username,
		Password: cfg.Password,
		Nickname: cfg.Nickname,
		Deliver:  strings.Split(cfg.Deliver, ","),
		IsSSL:    cfg.IsSSL,
	}
}
func (s *EmailSender) SendEmailMessage(message EmailMessage) error {
	return s.send(message.To, message.Subject, message.Content, message.Type == 1)
}

// 发送邮件
// ReplyTo: 回复邮件时，邮件的回复地址。可以是多个地址。
// From: 发件人的邮箱地址。
// BindExchange: 收件人的邮箱地址。可以是多个地址。
// Bcc: 密送（暗送）的邮箱地址。收件人不可见。可以是多个地址。
// Cc: 抄送的邮箱地址。收件人可见。可以是多个地址。
// Subject: 邮件的主题。
// Text: 邮件的纯文本内容，可选字段。
// HTML: 邮件的 HTML 内容，可选字段。
// Sender: 覆盖发件人的 SMTP 信封发送者，可选字段。
// Headers: 邮件的附加头部信息，使用 textproto.MIMEHeader 类型存储。
// Attachments: 邮件的附件列表，每个附件是一个 Attachment 结构体的实例。
// ReadReceipt: 邮件的回执邮箱地址，表示邮件阅读回执的接收地址。可以是多个地址。
func (s *EmailSender) send(to []string, subject string, body string, needCc bool) error {
	host := s.Host
	port := s.Port
	isSSL := s.IsSSL
	from := s.Username
	secret := s.Password
	nickname := s.Nickname

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	// 抄送
	if needCc {
		e.Cc = s.Deliver
	}
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}

func main() {
	// 配置 SMTP 服务器地址和端口
	smtpHost := "smtp.qq.com"
	smtpPort := 587

	// 邮箱账号和密码
	email := "your_email@qq.com"
	password := "your_password"

	// 发件人和收件人
	from := email
	to := "recipient@example.com"

	// 构建邮件内容
	subject := "Hello from Golang"
	body := "This is a test email sent using Golang."

	// 组装邮件内容
	message := "From: " + from + "\n" +
		"BindExchange: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// 发送邮件
	err := sendMail(smtpHost, smtpPort, email, password, from, to, []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("To sent successfully")
}

// sendMail 发送邮件
func sendMail(smtpHost string, smtpPort int, email, password, from, to string, message []byte) error {
	auth := smtp.PlainAuth("", email, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+strconv.Itoa(smtpPort), auth, from, []string{to}, message)
	if err != nil {
		return err
	}

	return nil
}
