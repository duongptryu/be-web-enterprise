package sendgridprovider

import (
	"context"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	log "github.com/sirupsen/logrus"
	"web/components/mailprovider"
)

type sendgridProvider struct {
	client *sendgrid.Client
}

func NewSendGridProvider(secretKey string) *sendgridProvider {
	return &sendgridProvider{
		client: sendgrid.NewSendClient(secretKey),
	}
}

func (p *sendgridProvider) SendMailNotifyNewComment(ctx context.Context, data *mailprovider.MailData) {
	from := mail.NewEmail("Web Enterprise", "duongpt2503@gmail.com")
	subject := "You got a new comment in your post idea!!"
	to := mail.NewEmail(data.Name, data.Email)
	plainTextContent := data.Content
	htmlContent := fmt.Sprintf("<strong>%s</strong>", data.Content)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	response, err := p.client.Send(message)
	if err != nil {
		log.Error(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
