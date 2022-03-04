package ideastore

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
)

func (s *sqlStore) FindIdea(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*ideamodel.Idea, error) {
	var result ideamodel.Idea

	db := s.db

	db = db.Table(ideamodel.Idea{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
