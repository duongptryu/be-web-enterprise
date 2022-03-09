package smtp

import (
	"context"
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
	"web/components/mailprovider"
)

type smtp struct {
	Email    string
	Password string
}

func NewSmtpProvider(from, password string) *smtp {
	return &smtp{
		Email:    from,
		Password: password,
	}
}

func (p *smtp) SendMailNotifyNewComment(ctx context.Context, data *mailprovider.MailDataForComment) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", p.Email)

	// Set E-Mail receivers
	m.SetHeader("To", data.Email)

	// Set E-Mail subject
	m.SetHeader("Subject", "You got a new comment in your post idea!!")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", data.Content)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, p.Email, p.Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Send email success")

	return
}
