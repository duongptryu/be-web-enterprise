package departmentstore

import (
	"context"
	"gorm.io/gorm"
	"web/common"
	"web/modules/department/departmentmodel"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type DepartmentStore interface {
	UpdateDepartment(ctx context.Context, id int, data *departmentmodel.DepartmentUpdate) error
	ListDepartment(ctx context.Context,
		condition map[string]interface{},
		filter *departmentmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]departmentmodel.Department, error)
	FindDepartment(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*departmentmodel.Department, error)
	DeleteDepartment(ctx context.Context, id int) error
	CreateDepartment(ctx context.Context, data *departmentmodel.DepartmentCreate) error
	ListDepartmentForStaff(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) ([]departmentmodel.Department, error)
}
