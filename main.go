package main

import "gopkg.in/gomail.v2"

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "fde@inbox.ru")
	m.SetHeader("To", "filippov@doc-lvv.ru", "fde@inbox.ru")
	//m.SetAddressHeader("Cc", "filippov@doc-lvv.ru", "Dan")
	m.SetHeader("Subject", "Место на харде!!")
	m.SetBody("text/html", "Ну давай - спать охота")
	//m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.mail.ru", 465, "fde@inbox.ru", "9dx4ebfaSedVFeeXcEZ7")

	// Отправить электронное письмо Bob, Cora и Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
