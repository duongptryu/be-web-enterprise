package uploadstore

import (
	"context"
	"web/common"
)

func (s *sqlStore) FindFile(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*common.File, error) {
	var result common.File

	db := s.db

	db = db.Table(common.File{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
