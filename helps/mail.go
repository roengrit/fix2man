package helps

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

//SendMail _
func SendMail(email, body string) string {

	m := gomail.NewMessage()
	m.SetHeader("From", "Your mail")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Forget password!")
	m.SetBody("text/html", "Your password has been change to : "+body)
	d := gomail.NewDialer("Mail server", 25, "Your mail", "password") // ส่งผ่าน Server
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	//d.SSL = true
	if err := d.DialAndSend(m); err != nil {
		return err.Error()
	}
	return ""
}
