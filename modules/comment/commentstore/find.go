package commentstore

import (
	"context"
	"web/common"
	"web/modules/comment/commentmodel"
)

func (s *sqlStore) FindComment(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*commentmodel.Comment, error) {
	var result commentmodel.Comment

	db := s.db

	db = db.Table(commentmodel.Comment{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
