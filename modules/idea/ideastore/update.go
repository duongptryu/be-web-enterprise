package ideastore

import (
	"context"
	"gorm.io/gorm"
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

func (s *sqlStore) IncreaseLikeCountIdea(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(ideamodel.IdeaUpdate{}.TableName()).Where("id = ?", id).Update("likes_count", gorm.Expr("likes_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) DecreaseLikeCountIdea(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(ideamodel.IdeaUpdate{}.TableName()).Where("id = ?", id).Update("likes_count", gorm.Expr("likes_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) IncreaseViewCountIdea(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(ideamodel.IdeaUpdate{}.TableName()).Where("id = ?", id).Update("views_count", gorm.Expr("views_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) IncreaseCommentCountIdea(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(ideamodel.IdeaUpdate{}.TableName()).Where("id = ?", id).Update("comments_count", gorm.Expr("comments_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
