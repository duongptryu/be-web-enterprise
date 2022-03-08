package departmentstore

import (
	"context"
	"web/common"
	"web/modules/department/departmentmodel"
)

func (s *sqlStore) UpdateDepartment(ctx context.Context, id int, data *departmentmodel.DepartmentUpdate) error {
	db := s.db

	if err := db.Table(departmentmodel.Department{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
