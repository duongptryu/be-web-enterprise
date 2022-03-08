package userbiz

import (
	"context"
	"web/common"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

type listUserBiz struct {
	store userstore.UserStore
}

func NewListUserBiz(store userstore.UserStore) *listUserBiz {
	return &listUserBiz{
		store: store,
	}
}

func (biz *listUserBiz) ListUserBiz(ctx context.Context, paging *common.Paging, filter *usermodel.Filter) ([]usermodel.User, error) {
	result, err := biz.store.ListUser(ctx, nil, filter, paging ,"Department")
	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}

	return result, nil
}
