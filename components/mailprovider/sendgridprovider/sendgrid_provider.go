package sendgridprovider

import (
	"context"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
	"web/components/mailprovider"
)

type sendgridProvider struct {
	client    *sendgrid.Client
	secretKey string
	name      string
	email     string
}

func NewSendGridProvider(secretKey string) *sendgridProvider {
	return &sendgridProvider{
		client:    sendgrid.NewSendClient(secretKey),
		secretKey: secretKey,
		name:      "Web Enterprise",
		email:     "duongpt2503@gmail.com",
	}
}

func (s *sendgridProvider) SendMailNotifyNewComment(ctx context.Context, data *mailprovider.MailDataForComment) {
	m := mail.NewV3Mail()
	from := mail.NewEmail(s.name, s.email)
	m.SetFrom(from)

	m.SetTemplateID("d-883583940e884f06b44b6b560e418529")

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail(data.Name, data.Email),
	}
	p.AddTos(tos...)

	p.SetDynamicTemplateData("name", data.CommentBy)
	p.SetDynamicTemplateData("owner", data.Name)
	p.SetDynamicTemplateData("datetime", data.CreatedAt.Format(time.RFC1123))
	p.SetDynamicTemplateData("content", data.CommentContent)
	p.SetDynamicTemplateData("link", fmt.Sprintf("https://groupbar.me/idea/%v", data.IdeaId))

	m.AddPersonalizations(p)

	request := sendgrid.GetRequest(s.secretKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = http.MethodPost
	var Body = mail.GetRequestBody(m)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func (p *sendgridProvider) SendMailNotifyNewIdea(ctx context.Context, data *mailprovider.MailDataForIdea) {
	from := mail.NewEmail("Web Enterprise", "duongpt2503@gmail.com")
	subject := "Staff of your department just push new idea!!"
	to := mail.NewEmail(data.Name, data.Email)
	plainTextContent := data.Title
	htmlContent := fmt.Sprintf("<strong>%s</strong>", data.Title)
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
