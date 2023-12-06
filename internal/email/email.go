package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

var d *gomail.Dialer

func New() {
	d = gomail.NewDialer("smtp.qq.com", 587, "jinchengjoker@foxmail.com", viper.GetString("email.password"))
}

func Send(to string) error {
	if d == nil {
		New()
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "Mangosteen <jinchengjoker@foxmail.com>")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>金成</b>!")

	return d.DialAndSend(m)
}
