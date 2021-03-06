package categorystore

import (
	"context"
	"web/common"
	"web/modules/category/categorymodel"
)

func (s *sqlStore) DeleteCategory(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(categorymodel.Category{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
