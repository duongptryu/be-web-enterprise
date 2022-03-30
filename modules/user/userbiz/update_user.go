package userbiz

import (
	"context"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"web/common"
	"web/modules/department/departmentmodel"
	"web/modules/department/departmentstore"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

type updateUserBiz struct {
	store           userstore.UserStore
	departmentStore departmentstore.DepartmentStore
}

func NewUpdateUserBiz(store userstore.UserStore, departmentStore departmentstore.DepartmentStore) *updateUserBiz {
	return &updateUserBiz{
		store:           store,
		departmentStore: departmentStore,
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

	if data.Password != "" {
		newPass, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		data.Password = string(newPass)
	}

	//create user in db
	if err := biz.store.UpdateUser(ctx, userId, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	go biz.updateTagUser(ctx, userId)

	return nil
}

func (biz *updateUserBiz) UpdateUserSelfBiz(ctx context.Context, userId int, data *usermodel.UserUpdateSelf) error {
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

	if data.Password != "" {
		newPass, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		data.Password = string(newPass)
	}

	if err := biz.store.UpdateUserSelf(ctx, userId, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	go biz.updateTagUser(ctx, userId)

	return nil
}

func (biz *updateUserBiz) updateTagUser(ctx context.Context, userId int) {
	data, err := biz.store.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		log.Error(err)
		return
	}
	userUpdate := usermodel.UserUpdate{
		Tags: data.SetTags(),
	}
	if err := biz.store.UpdateUser(ctx, userId, &userUpdate); err != nil {
		log.Error(err)
		return
	}
	return
}
