package userbiz

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"web/common"
	"web/modules/department/departmentmodel"
	"web/modules/department/departmentstore"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

type createUserBiz struct {
	store           userstore.UserStore
	departmentStore departmentstore.DepartmentStore
}

func NewCreateUserBiz(store userstore.UserStore, departmentStore departmentstore.DepartmentStore) *createUserBiz {
	return &createUserBiz{
		store:           store,
		departmentStore: departmentStore,
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

	if data.Role == common.RoleStaff {
		departmentExist, err := biz.departmentStore.FindDepartment(ctx, map[string]interface{}{"id": data.DepartmentId})
		if err != nil {
			return err
		}
		if departmentExist.Id == 0 {
			return common.ErrDataNotFound(departmentmodel.EntityName)
		}
	} else {
		data.DepartmentId = 0
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
