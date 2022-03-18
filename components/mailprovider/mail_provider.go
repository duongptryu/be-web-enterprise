package mailprovider

import (
	"context"
	"time"
)

type MailDataForComment struct {
	Email          string
	Name           string
	CommentContent string
	CommentBy      string
	IdeaId         int
	CreatedAt      time.Time
}

type MailDataForIdea struct {
	Email         string
	Name          string
	NameUserPush  string
	EmailUserPush string
	Title         string
	Content       string
	Id            int
	CreatedAt     *time.Time
}

type MailProvider interface {
	SendMailNotifyNewComment(ctx context.Context, data *MailDataForComment)
	SendMailNotifyNewIdea(ctx context.Context, data *MailDataForIdea)
}
