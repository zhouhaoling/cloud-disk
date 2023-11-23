package test

import (
	"cloud/core/define"
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	//发送者的邮箱
	e.From = "Get <xxx>"
	//接收邮箱
	e.To = []string{"目的邮箱"}
	//发送邮件的主题
	e.Subject = "验证码测试发送"
	//有HTML和Text两种发送方式
	//e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("验证码为：<h1>123456</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "邮箱名", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
