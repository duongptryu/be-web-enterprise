package idealikeviewbiz

import (
	"context"
	"web/common"
	"web/modules/idealikeview/idealikeviewmodel"
	"web/modules/idealikeview/idealikeviewstore"
)

type listUserLikeViewIdeaBiz struct {
	store idealikeviewstore.UserLikeViewIdeaStore
}

func NewListUserLikeViewIdeaBiz(store idealikeviewstore.UserLikeViewIdeaStore) *listUserLikeViewIdeaBiz {
	return &listUserLikeViewIdeaBiz{
		store: store,
	}
}

func (biz *listUserLikeViewIdeaBiz) ListUserLikeIdea(ctx context.Context, paging *common.Paging, filter *idealikeviewmodel.Filter) ([]idealikeviewmodel.UserLikeIdea, error) {
	result, err := biz.store.ListUserLikeIdea(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(idealikeviewmodel.EntityName, err)
	}

	return result, nil
}

func (biz *listUserLikeViewIdeaBiz) ListUserDislikeIdea(ctx context.Context, paging *common.Paging, filter *idealikeviewmodel.Filter) ([]idealikeviewmodel.UserDislikeIdea, error) {
	result, err := biz.store.ListUserDislikeIdea(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(idealikeviewmodel.EntityName, err)
	}

	return result, nil
}

func (biz *listUserLikeViewIdeaBiz) ListUserViewIdea(ctx context.Context, paging *common.Paging, filter *idealikeviewmodel.Filter) ([]idealikeviewmodel.UserViewIdea, error) {
	result, err := biz.store.ListUserViewIdea(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(idealikeviewmodel.EntityName, err)
	}

	return result, nil
}
