package commentbiz

import (
	"context"
	"web/common"
	"web/modules/comment/commentmodel"
	"web/modules/comment/commentstore"
	"web/pubsub"
)

type deleteCommentBiz struct {
	commentStore commentstore.CommentStore
	pubSub       pubsub.PubSub
}

func NewDeleteCommentBiz(commentStore commentstore.CommentStore, pubSub pubsub.PubSub) *deleteCommentBiz {
	return &deleteCommentBiz{
		commentStore: commentStore,
		pubSub:       pubSub,
	}
}

func (biz *deleteCommentBiz) DeleteCommentBiz(ctx context.Context, commentId, userId int) error {
	commentExist, err := biz.commentStore.FindComment(ctx, map[string]interface{}{"id": commentId, "user_id": userId})
	if err != nil {
		return err
	}
	if commentExist.Id == 0 {
		return common.ErrDataNotFound(commentmodel.EntityName)
	}

	if err := biz.commentStore.SoftDeleteComment(ctx, commentId); err != nil {
		return common.ErrCannotCreateEntity(commentmodel.EntityName, err)
	}

	go biz.pubSub.Publish(ctx, common.TopicDecreaseCommentCountIdea, pubsub.NewMessage(commentExist.IdeaId))

	return nil
}
