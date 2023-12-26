package main

import (
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

func main() {
	// e := email.NewEmail()
	// e.From = "ZZX <984174877@qq.com>"
	// e.To = []string{"984174877@qq.com"}
	// e.Bcc = []string{"984174877@qq.com"}
	// e.Cc = []string{"984174877@qq.com"}
	// e.Subject = "测试golang email库"
	// e.Text = []byte("本文邮件内容!")
	// e.HTML = []byte("<h1>html 邮件内容!</h1>")
	// e.Send("smtp.qq.com:587", smtp.PlainAuth("", "984174877@qq.com", "mussewjschgnbecd", "smtp.qq.com"))

	// e := &email.Email{
	// 	To:      []string{"test@example.com"},
	// 	From:    "Jordan Wright <test@gmail.com>",
	// 	Subject: "Awesome Subject",
	// 	Text:    []byte("Text Body is, of course, supported!"),
	// 	HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
	// 	Headers: textproto.MIMEHeader{},
	// }

	// e := NewEmail()
	// e.AttachFile("test.txt")

	var ch <-chan *email.Email
	p, _ := email.NewPool(
		"smtp.gmail.com:587",
		4,
		smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"),
	)
	for i := 0; i < 4; i++ {
		go func() {
			for e := range ch {
				p.Send(e, 10*time.Second)
			}
		}()
	}

}
