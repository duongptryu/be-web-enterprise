package departmentbiz

import (
	"context"
	"web/common"
	"web/modules/department/departmentmodel"
	"web/modules/department/departmentstore"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

type createDepartmentBiz struct {
	departmentStore departmentstore.DepartmentStore
	userStore       userstore.UserStore
}

func NewCreateDepartmentBiz(departmentStore departmentstore.DepartmentStore, userStore userstore.UserStore) *createDepartmentBiz {
	return &createDepartmentBiz{
		departmentStore: departmentStore,
		userStore:       userStore,
	}
}

func (biz *createDepartmentBiz) CreateDepartment(ctx context.Context, data *departmentmodel.DepartmentCreate) error {
	userDB, err := biz.userStore.FindUser(ctx, map[string]interface{}{"id": data.LeaderId})
	if err != nil {
		return err
	}
	if userDB.Id == 0 {
		return common.ErrDataNotFound(usermodel.EntityName)
	}

	if userDB.Role != common.RoleQACoordinator {
		return departmentmodel.ErrInvalidLeaderId
	}

	data.Status = true
	data.Tags = data.SetTags()
	if err := biz.departmentStore.CreateDepartment(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(departmentmodel.EntityName, err)
	}

	return nil
}
