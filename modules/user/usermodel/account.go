package usermodel

import (
	"errors"
	"web/common"
	"web/components/tokenprovider"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (data *UserLogin) Validate() error {
	return nil
}

type Account struct {
	AccessToken  string               `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"-"`
	Profile      *User                `json:"profile"`
}

func NewAccount(at string, rt *tokenprovider.Token, profile *User) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
		Profile:      profile,
	}
}

var (
	ErrLengthPassword            = common.NewCustomError(nil, "Length of password must be greater than 8 character", "ErrLengthPassword")
	ErrLengthFirstName           = common.NewCustomError(nil, "Length of first name must be greater than 3 character", "ErrLengthFirstName")
	ErrLengthLastName            = common.NewCustomError(nil, "Length of last name must be greater than 3 character", "ErrLengthLastName")
	ErrUsernameOrPasswordInvalid = common.NewFullErrorResponse(401,
		errors.New("username or password invalid"),
		"username or password invalid",
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
	ErrAccountIsInactive = common.NewFullErrorResponse(409,
		nil,
		"This account is inactive",
		"This account is inactive",
		"ErrAccountIsInactive",
	)
	ErrEmailAlreadyExist = common.NewFullErrorResponse(409,
		nil,
		"This email already exist",
		"This email already exist",
		"ErrEmailAlreadyExist",
	)
)
