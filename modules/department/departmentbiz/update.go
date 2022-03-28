package departmentbiz

import (
	"context"
	log "github.com/sirupsen/logrus"
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

func (biz *updateDepartmentBiz) UpdateIdeaBiz(ctx context.Context, departmentId int, data *departmentmodel.DepartmentUpdate) error {
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

	go biz.updateTagDepartment(ctx, departmentId)

	return nil
}

func (biz *updateDepartmentBiz) updateTagDepartment(ctx context.Context, id int) {
	data, err := biz.store.FindDepartment(ctx, map[string]interface{}{"id": id}, "Leader")
	if err != nil {
		log.Error(err)
		return
	}
	updateDepartment := departmentmodel.DepartmentUpdate{
		Tags: data.SetTags(),
	}
	if err := biz.store.UpdateDepartment(ctx, id, &updateDepartment); err != nil {
		log.Error(err)
		return
	}
	return
}
