package ideastore

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
)

func (s *sqlStore) DeleteIdea(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(ideamodel.Idea{}.TableName()).Where("id = ?", id).Update("status", false).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
