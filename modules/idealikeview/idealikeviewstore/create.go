package idealikeviewstore

import (
	"context"
	"web/common"
	"web/modules/idealikeview/idealikeviewmodel"
)

func (s *sqlStore) CreateUserLikeIdea(ctx context.Context, data *idealikeviewmodel.UserLikeIdea) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) CreateUserDislikeIdea(ctx context.Context, data *idealikeviewmodel.UserDislikeIdea) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) CreateUserViewIdea(ctx context.Context, data *idealikeviewmodel.UserViewIdea) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
