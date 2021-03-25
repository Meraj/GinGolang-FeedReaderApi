package Classes

import (
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
)
type EmailHelper struct {
	
}

func (EmailHelper) SendEmail(from string,to string,title string,content string)  {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", title)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", content)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "from@gmail.com", "<email_password>")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	//if err := d.DialAndSend(m); err != nil {
//		fmt.Println(err)
//		panic(err)
//	}
}
