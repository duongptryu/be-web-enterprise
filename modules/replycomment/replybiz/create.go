package replybiz

import (
	"context"
	"time"
	"web/common"
	"web/modules/acayear/acayearstore"
	"web/modules/comment/commentmodel"
	"web/modules/comment/commentstore"
	"web/modules/idea/ideamodel"
	"web/modules/replycomment/replymodel"
	"web/modules/replycomment/replystore"
	"web/pubsub"
)

type createReplyBiz struct {
	commentStore commentstore.CommentStore
	replyStore   replystore.ReplyStore
	acaYearStore acayearstore.AcademicYearStore
	pubSub       pubsub.PubSub
}

func NewCreateReplyBiz(replyStore replystore.ReplyStore, commentStore commentstore.CommentStore, acaYearStore acayearstore.AcademicYearStore, pubSub pubsub.PubSub) *createReplyBiz {
	return &createReplyBiz{
		replyStore:   replyStore,
		commentStore: commentStore,
		acaYearStore: acaYearStore,
		pubSub:       pubSub,
	}
}

func (biz *createReplyBiz) CreateReplyBiz(ctx context.Context, data *replymodel.ReplyCreate) error {
	commentExist, err := biz.commentStore.FindComment(ctx, map[string]interface{}{"id": data.CommentId})
	if err != nil {
		return err
	}
	if commentExist.Id == 0 {
		return common.ErrDataNotFound(commentmodel.EntityName)
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
	data.IdeaId = commentExist.IdeaId
	if err := biz.replyStore.CreateReply(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(replymodel.EntityName, err)
	}

	go biz.pubSub.Publish(ctx, common.TopicIncreaseReplyCountComment, pubsub.NewMessage(data.CommentId))
	go biz.pubSub.Publish(ctx, common.TopicIncreaseCommentCountIdea, pubsub.NewMessage(data.IdeaId))

	return nil
}
