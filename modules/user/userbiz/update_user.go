package userbiz

import (
	"context"
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

	//create user in db
	if err := biz.store.UpdateUser(ctx, userId, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
