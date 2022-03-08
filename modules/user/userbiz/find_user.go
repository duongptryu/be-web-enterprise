package userbiz

import (
	"context"
	"web/common"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

type findUserBiz struct {
	store userstore.UserStore
}

func NewFindUserBiz(store userstore.UserStore) *findUserBiz {
	return &findUserBiz{
		store: store,
	}
}

func (biz *findUserBiz) FindUserBiz(ctx context.Context, userId int) (*usermodel.User, error) {
	result, err := biz.store.FindUser(ctx, map[string]interface{}{"id": userId}, "Department")
	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}

	return result, nil
}
