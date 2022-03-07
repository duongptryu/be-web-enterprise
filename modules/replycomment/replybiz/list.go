package replybiz

import (
	"context"
	"web/common"
	"web/modules/replycomment/replymodel"
	"web/modules/replycomment/replystore"
)

type listReplyBiz struct {
	store replystore.ReplyStore
}

func NewListReplyBiz(store replystore.ReplyStore) *listReplyBiz {
	return &listReplyBiz{
		store: store,
	}
}

func (biz *listReplyBiz) ListReplyBiz(ctx context.Context, commentId int, paging *common.Paging, filter *replymodel.Filter) ([]replymodel.Reply, error) {
	result, err := biz.store.ListReply(ctx, map[string]interface{}{"comment_id": commentId}, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(replymodel.EntityName, err)
	}

	return result, nil
}

func (biz *listReplyBiz) ListReplyForStaff(ctx context.Context, commentId int, paging *common.Paging, filter *replymodel.Filter) ([]replymodel.Reply, error) {
	result, err := biz.store.ListReply(ctx, map[string]interface{}{"comment_id": commentId, "status": true}, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(replymodel.EntityName, err)
	}

	return result, nil
}
