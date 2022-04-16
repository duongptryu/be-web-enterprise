package userbiz

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"web/common"
	"web/components/tokenprovider"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

type userBiz struct {
	storeUser     userstore.UserStore
	tokenProvider tokenprovider.TokenProvider
	expire        int
}

func NewLoginBiz(storeUser userstore.UserStore, tokenProvider tokenprovider.TokenProvider, expiry int) *userBiz {
	return &userBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		expire:        expiry,
	}
}

func (biz *userBiz) UserLogin(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	userDB, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	if userDB.Id == 0 {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	if userDB.Status == false {
		return nil, usermodel.ErrAccountIsInactive
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(data.Password))
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: userDB.Id,
		Role:   userDB.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(&payload, biz.expire)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken.Token, nil, userDB)

	return account, nil
}
