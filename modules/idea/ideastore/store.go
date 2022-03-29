package ideastore

import (
	"context"
	"gorm.io/gorm"
	"web/common"
	"web/modules/idea/ideamodel"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type IdeaStore interface {
	CreateIdea(ctx context.Context, data *ideamodel.IdeaCreate) error
	DeleteIdea(ctx context.Context, id int) error
	FindIdea(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*ideamodel.Idea, error)
	ListIdea(ctx context.Context,
		condition map[string]interface{},
		filter *ideamodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]ideamodel.Idea, error)
	UpdateIdea(ctx context.Context, id int, data *ideamodel.IdeaUpdate) error
	ListALlIdea(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) ([]ideamodel.Idea, error)
	CountIdeaByCategoryId(ctx context.Context,
		categoryIds []int,
		moreKey ...string,
	) (map[int]int, error)
	CountUserPostIdea(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (map[int]int, error)
}
