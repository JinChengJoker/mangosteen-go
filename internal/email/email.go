package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "jinchengjoker@foxmail.com")
	m.SetHeader("To", "jinchengjoker@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>金成</b>!")

	d := gomail.NewDialer("smtp.qq.com", 587, "jinchengjoker@foxmail.com", viper.GetString("email.password"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
