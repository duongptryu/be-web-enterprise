package departmentstore

import (
	"context"
	"web/common"
	"web/modules/department/departmentmodel"
)

func (s *sqlStore) DeleteDepartment(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(departmentmodel.Department{}.TableName()).Where("id = ?", id).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
