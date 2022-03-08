package departmentstore

import (
	"context"
	"web/common"
	"web/modules/department/departmentmodel"
)

func (s *sqlStore) CreateDepartment(ctx context.Context, data *departmentmodel.DepartmentCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
