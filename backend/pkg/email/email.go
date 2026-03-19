package email

import (
	"fmt"

	"tbox-backend/config"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	cfg *config.EmailConfig
}

func NewEmailSender(cfg *config.EmailConfig) *EmailSender {
	return &EmailSender{cfg: cfg}
}

func (s *EmailSender) SendVerificationCode(to, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.cfg.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "TBox - 邮箱验证码")
	m.SetBody("text/html", fmt.Sprintf(`
		<div style="padding: 20px; font-family: Arial, sans-serif;">
			<h2>TBox 软件商城</h2>
			<p>您好，您的验证码是：</p>
			<h1 style="color: #4A90D9; letter-spacing: 5px;">%s</h1>
			<p>验证码有效期为 5 分钟，请勿泄露给他人。</p>
			<p>如果您没有请求此验证码，请忽略此邮件。</p>
		</div>
	`, code))

	d := gomail.NewDialer(s.cfg.Host, s.cfg.Port, s.cfg.Username, s.cfg.Password)
	return d.DialAndSend(m)
}
