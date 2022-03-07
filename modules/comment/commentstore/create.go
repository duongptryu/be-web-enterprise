package commentstore

import (
	"context"
	"web/common"
	"web/modules/comment/commentmodel"
)

func (s *sqlStore) CreateComment(ctx context.Context, data *commentmodel.CommentCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
