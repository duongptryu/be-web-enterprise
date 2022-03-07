package replybiz

import (
	"context"
	"web/common"
	"web/modules/replycomment/replymodel"
	"web/modules/replycomment/replystore"
	"web/pubsub"
)

type deleteCommentBiz struct {
	replyStore replystore.ReplyStore
	pubSub     pubsub.PubSub
}

func NewDeleteReplyBiz(replyStore replystore.ReplyStore, pubSub pubsub.PubSub) *deleteCommentBiz {
	return &deleteCommentBiz{
		replyStore: replyStore,
		pubSub:     pubSub,
	}
}

func (biz *deleteCommentBiz) DeleteReplyBiz(ctx context.Context, replyId, userId int) error {
	replyExist, err := biz.replyStore.FindReply(ctx, map[string]interface{}{"id": replyId, "user_id": userId})
	if err != nil {
		return err
	}
	if replyExist.Id == 0 {
		return common.ErrDataNotFound(replymodel.EntityName)
	}

	if err := biz.replyStore.SoftDeleteReply(ctx, replyId); err != nil {
		return common.ErrCannotDeleteEntity(replymodel.EntityName, err)
	}

	go biz.pubSub.Publish(ctx, common.TopicDecreaseReplyCountComment, pubsub.NewMessage(replyExist.CommentId))
	go biz.pubSub.Publish(ctx, common.TopicDecreaseCommentCountIdea, pubsub.NewMessage(replyExist.IdeaId))

	return nil
}
