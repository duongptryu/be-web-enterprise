package commentbiz

import (
	"context"
	"web/common"
	"web/modules/comment/commentmodel"
	"web/modules/comment/commentstore"
)

type listCommentBiz struct {
	store commentstore.CommentStore
}

func NewListComment(store commentstore.CommentStore) *listCommentBiz {
	return &listCommentBiz{
		store: store,
	}
}

func (biz *listCommentBiz) ListComment(ctx context.Context, ideaId int, paging *common.Paging, filter *commentmodel.Filter) ([]commentmodel.Comment, error) {
	result, err := biz.store.ListComment(ctx, map[string]interface{}{"idea_id": ideaId}, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(commentmodel.EntityName, err)
	}

	return result, nil
}

func (biz *listCommentBiz) ListCommentForStaff(ctx context.Context, ideaId int, paging *common.Paging, filter *commentmodel.Filter) ([]commentmodel.Comment, error) {
	result, err := biz.store.ListComment(ctx, map[string]interface{}{"idea_id": ideaId, "status": true}, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(commentmodel.EntityName, err)
	}

	return result, nil
}
