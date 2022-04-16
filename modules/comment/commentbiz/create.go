package commentbiz

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"
	"web/common"
	"web/components/mailprovider"
	"web/modules/acayear/acayearstore"
	"web/modules/comment/commentmodel"
	"web/modules/comment/commentstore"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
	"web/modules/notification/notificationmodel"
	"web/modules/notification/notificationstore"
	"web/modules/user/userstore"
	"web/pubsub"
)

type createCommentBiz struct {
	ideaStore         ideastore.IdeaStore
	commentStore      commentstore.CommentStore
	acaYearStore      acayearstore.AcademicYearStore
	mailProvider      mailprovider.MailProvider
	userStore         userstore.UserStore
	pubSub            pubsub.PubSub
	notificationStore notificationstore.NotificationStore
}

func NewCreateCommentBiz(store ideastore.IdeaStore, commentStore commentstore.CommentStore, acaYearStore acayearstore.AcademicYearStore, pubSub pubsub.PubSub, mailProvider mailprovider.MailProvider, userStore userstore.UserStore, notificationStore notificationstore.NotificationStore) *createCommentBiz {
	return &createCommentBiz{
		ideaStore:         store,
		commentStore:      commentStore,
		acaYearStore:      acaYearStore,
		pubSub:            pubSub,
		mailProvider:      mailProvider,
		userStore:         userStore,
		notificationStore: notificationStore,
	}
}

func (biz *createCommentBiz) CreateCommentBiz(ctx context.Context, data *commentmodel.CommentCreate) error {
	ideaExist, err := biz.ideaStore.FindIdea(ctx, map[string]interface{}{"id": data.IdeaId})
	if err != nil {
		return err
	}
	if ideaExist.Status == false {
		return ideamodel.ErrIdeaAlreadyRemoved
	}
	if ideaExist.Id == 0 {
		return common.ErrDataNotFound(ideamodel.EntityName)
	}

	acaExist, err := biz.acaYearStore.FindAcaYear(ctx, map[string]interface{}{"status": true})
	if err != nil {
		return err
	}
	if acaExist.Id == 0 {
		return ideamodel.ErrAcademicYearNotReady
	}

	timeNow := time.Now()
	if timeNow.After(acaExist.FinalClosureDate) {
		return ideamodel.ErrFinalClosureDateExpired
	}

	data.Status = true
	if err := biz.commentStore.CreateComment(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(commentmodel.EntityName, err)
	}

	go func(b *createCommentBiz, ownerId int, ideaId int, content string, commenterId int) {
		owner, err := b.userStore.FindUser(ctx, map[string]interface{}{"id": ownerId})
		if err != nil {
			log.Error(err)
		}
		commenter, err := b.userStore.FindUser(ctx, map[string]interface{}{"id": commenterId})
		if err != nil {
			log.Error(err)
		}
		if data.IsAnonymous {
			commenter.FullName = "Anonymous"
		}
		go b.mailProvider.SendMailNotifyNewComment(ctx, &mailprovider.MailDataForComment{CommentContent: content, Email: owner.Email, Name: owner.FullName, CommentBy: commenter.FullName, CreatedAt: time.Now()})

		if ownerId == commenterId {
			return
		}

		newNoti := notificationmodel.NotificationIdeaCreate{
			TypeNoti: common.NewCommentNotification,
			OwnerId:  ownerId,
			IdeaId:   ideaId,
			UserId:   commenterId,
		}
		if err := b.notificationStore.CreateNotification(ctx, &newNoti); err != nil {
			log.Error(err)
		}
	}(biz, ideaExist.UserId, ideaExist.Id, data.Content, data.UserId)

	go biz.pubSub.Publish(ctx, common.TopicIncreaseCommentCountIdea, pubsub.NewMessage(data.IdeaId))

	return nil
}
