package categorystore

import (
	"context"
	"web/common"
	"web/modules/category/categorymodel"
)

func (s *sqlStore) FindCategory(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*categorymodel.Category, error) {
	var result categorymodel.Category

	db := s.db

	db = db.Table(categorymodel.Category{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
