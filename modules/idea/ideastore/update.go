package ideastore

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
)

func (s *sqlStore) UpdateIdea(ctx context.Context, id int, data *ideamodel.IdeaUpdate) error {
	db := s.db

	if err := db.Table(ideamodel.IdeaUpdate{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
