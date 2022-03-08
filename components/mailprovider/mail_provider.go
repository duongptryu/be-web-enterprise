package mailprovider

import (
	"context"
	"time"
)

type MailDataForComment struct {
	Email   string
	Name    string
	Content string
}

type MailDataForIdea struct {
	Email         string
	Name          string
	NameUserPush  string
	EmailUserPush string
	Title         string
	CreatedAt     *time.Time
}

type MailProvider interface {
	SendMailNotifyNewComment(ctx context.Context, data *MailDataForComment)
	SendMailNotifyNewIdea(ctx context.Context, data *MailDataForIdea)
}
