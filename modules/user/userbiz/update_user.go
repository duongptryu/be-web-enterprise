package userbiz

import (
	"context"
	"web/common"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

type updateUserBiz struct {
	store userstore.UserStore
}

func NewUpdateUserBiz(store userstore.UserStore) *updateUserBiz {
	return &updateUserBiz{
		store: store,
	}
}

func (biz *updateUserBiz) UpdateUserBiz(ctx context.Context, userId int, data *usermodel.UserUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	userDB, err := biz.store.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	if userDB.Id == 0 {
		return common.ErrDataNotFound(usermodel.EntityName)
	}

	//create user in db
	if err := biz.store.UpdateUser(ctx, userId, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
