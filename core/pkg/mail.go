package pkg

import (
	"cloud/core/define"
	"crypto/tls"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "Get <ZHL19174683866@163.com>"
	e.To = []string{"1735950045@qq.com"}
	e.Subject = "验证码测试发送"
	e.HTML = []byte("验证码为：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "ZHL19174683866@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}
