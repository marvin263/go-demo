package mail

import (
	"log"
	"net/smtp"
)

const (
	username = "abc@qq.com"
	password = ""
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
	user := "abc@qq.com"
	to := []string{"abc@qq.com"}
	msg := []byte("Subject: Hello there\r\nFrom: abc <abc@qq.com>\r\nTo: abc@qq.com\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=utf-8\r\n\r\n" + body)

	err := smtp.SendMail(addr, auth, user, to, msg)
	if err != nil {
		log.Printf("Error when send email:%s", err.Error())
	}
}
