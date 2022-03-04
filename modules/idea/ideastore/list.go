package ideastore

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
)

func (s *sqlStore) ListIdea(ctx context.Context,
	condition map[string]interface{},
	filter *ideamodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]ideamodel.Idea, error) {
	var result []ideamodel.Idea

	db := s.db

	db = db.Table(ideamodel.Idea{}.TableName()).Where(condition)

	// if v := filter; v != nil {

	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if paging.FakeCursor > 0 {
		db = db.Where("id < ?", paging.FakeCursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
