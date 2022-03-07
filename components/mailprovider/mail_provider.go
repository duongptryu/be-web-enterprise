package mailprovider

import "context"

type MailData struct {
	Email   string
	Name    string
	Content string
}

type MailProvider interface {
	SendMailNotifyNewComment(ctx context.Context, data *MailData)
}
