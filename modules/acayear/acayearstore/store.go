package acayearstore

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearmodel"

	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type AcademicYearStore interface {
	CreateAcaYear(ctx context.Context, data *acayearmodel.AcademicYearCreate) error
	DeleteAcaYear(ctx context.Context, id int) error
	ListAcaYear(ctx context.Context,
		condition map[string]interface{},
		filter *acayearmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]acayearmodel.AcademicYear, error)
	UpdateAcaYear(ctx context.Context, id int, data *acayearmodel.AcademicYearUpdate) error
	FindAcaYear(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*acayearmodel.AcademicYear, error)
	ListAcaYearWithoutPaging(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) ([]acayearmodel.AcademicYear, error)
}
