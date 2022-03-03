package userbiz

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"web/common"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

type createUserBiz struct {
	store userstore.UserStore
}

func NewCreateUserBiz(store userstore.UserStore) *createUserBiz {
	return &createUserBiz{
		store: store,
	}
}

func (biz *createUserBiz) CreateUserBiz(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	userDB, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	if userDB.Id != 0 {
		return usermodel.ErrEmailAlreadyExist
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	data.Password = string(hashedPassword)

	//create user in db
	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
