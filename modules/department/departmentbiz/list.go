package departmentbiz

import (
	"context"
	"web/common"
	"web/modules/department/departmentmodel"
	"web/modules/department/departmentstore"
)

type listDepartmentBiz struct {
	store departmentstore.DepartmentStore
}

func NewListDepartment(store departmentstore.DepartmentStore) *listDepartmentBiz {
	return &listDepartmentBiz{
		store: store,
	}
}

func (biz *listDepartmentBiz) ListDepartmentBiz(ctx context.Context, paging *common.Paging, filter *departmentmodel.Filter) ([]departmentmodel.Department, error) {
	result, err := biz.store.ListDepartment(ctx, nil, filter, paging, "Leader")
	if err != nil {
		return nil, common.ErrCannotListEntity(departmentmodel.EntityName, err)
	}

	return result, nil
}

func (biz *listDepartmentBiz) ListDepartmentBizForStaff(ctx context.Context) ([]departmentmodel.Department, error) {
	result, err := biz.store.ListDepartmentForStaff(ctx, nil, "Leader")
	if err != nil {
		return nil, common.ErrCannotListEntity(departmentmodel.EntityName, err)
	}

	return result, nil
}
