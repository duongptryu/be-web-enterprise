package userbiz

import (
	"context"
	"web/common"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

type deleteUserBiz struct {
	store userstore.UserStore
}

func NewDeleteUserBiz(store userstore.UserStore) *deleteUserBiz {
	return &deleteUserBiz{
		store: store,
	}
}

func (biz *deleteUserBiz) SoftDeleteUserBiz(ctx context.Context, userId int) error {
	userDB, err := biz.store.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, err)
	}
	if userDB.Id == 0 {
		return common.ErrDataNotFound(usermodel.EntityName)
	}

	//create user in db
	if err := biz.store.SoftDeleteUser(ctx, userId); err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, err)
	}

	return nil
}
