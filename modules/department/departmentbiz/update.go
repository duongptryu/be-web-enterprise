package departmentbiz

import (
	"context"
	"web/common"
	"web/modules/department/departmentmodel"
	"web/modules/department/departmentstore"
)

type updateDepartmentBiz struct {
	store departmentstore.DepartmentStore
}

func NewUpdateIdeaBiz(store departmentstore.DepartmentStore) *updateDepartmentBiz {
	return &updateDepartmentBiz{
		store: store,
	}
}

func (biz *updateDepartmentBiz) UpdateIdeaBiz(ctx context.Context, departmentId int,  data *departmentmodel.DepartmentUpdate) error {
	departmentDB, err := biz.store.FindDepartment(ctx, map[string]interface{}{"id": departmentId})
	if err != nil {
		return common.ErrCannotUpdateEntity(departmentmodel.EntityName, err)
	}
	if departmentDB.Id == 0 {
		return common.ErrDataNotFound(departmentmodel.EntityName)
	}

	if err := biz.store.UpdateDepartment(ctx, departmentId, data); err != nil {
		return common.ErrCannotUpdateEntity(departmentmodel.EntityName, err)
	}

	return nil
}
