package mail

import (
	"log"
	"net/smtp"
)

const (
	username = "290500532@qq.com"
	password = "qvwxegtzumjmcbbh"
	host     = "smtp.qq.com"
	addr     = "smtp.qq.com:25"
)

// 发送邮件
func SendEmail() {
	auth := smtp.PlainAuth("", username, password, host)
	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	user := "290500532@qq.com"
	to := []string{"290500532@qq.com"}
	msg := []byte("Subject: Hello there\r\nFrom: libingqq <290500532@qq.com>\r\nTo: 290500532@qq.com\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=utf-8\r\n\r\n" + body)

	err := smtp.SendMail(addr, auth, user, to, msg)
	if err != nil {
		log.Printf("Error when send email:%s", err.Error())
	}
}
