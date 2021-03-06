package ideastore

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
)

func (s *sqlStore) ListALlIdea(ctx context.Context,
	condition map[string]interface{},
	filter *ideamodel.Filter,
	moreKey ...string,
) ([]ideamodel.Idea, error) {
	var result []ideamodel.Idea

	db := s.db

	db = db.Table(ideamodel.Idea{}.TableName()).Where(condition)

	if v := filter; v != nil {
		if v.CreatedAtGt != nil {
			db = db.Where("created_at >= ?", v.CreatedAtGt)
		}
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Order("id").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
