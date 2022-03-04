package ideastore

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
)

func (s *sqlStore) CreateIdea(ctx context.Context, data *ideamodel.IdeaCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
