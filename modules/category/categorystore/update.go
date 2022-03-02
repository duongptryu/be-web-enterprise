package categorystore

import (
	"context"
	"web/common"
	"web/modules/category/categorymodel"
)

func (s *sqlStore) UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	db := s.db

	if err := db.Table(categorymodel.CategoryUpdate{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
