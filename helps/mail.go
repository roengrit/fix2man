package helps

import (
	"net/smtp"
)

//Send Mail
func SendMail(email, body string) string {

	from := "badcode.th@gmail.com"
	pwd := "MuchP@ssw0rd"
	to := "logon.firstclass@hotmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Your new password : " + body

	auth := smtp.PlainAuth("", from, pwd, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))

	if err != nil {
		return err.Error()
	}
	return ""
}
