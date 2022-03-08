package departmentmodel

import "web/common"

const (
	EntityName = "Department"
)

type Department struct {
	common.SQLModel
	Name     string             `json:"name" gorm:"name"`
	Status   bool               `json:"status" gorm:"status"`
	LeaderId int                `json:"leader_id" gorm:"leader_id"`
	User     *common.SimpleUser `json:"leader" gorm:"foreignKey:LeaderId;preload:false"`
}

func (Department) TableName() string {
	return "departments"
}

type DepartmentCreate struct {
	common.SQLModelCreate
	Name     string `json:"name" gorm:"name"`
	Status   bool   `json:"status" gorm:"status"`
	LeaderId int    `json:"leader_id" gorm:"leader_id"`
}

func (DepartmentCreate) TableName() string {
	return Department{}.TableName()
}

type DepartmentUpdate struct {
	common.SQLModel
	Name     string `json:"name" gorm:"name"`
	Status   *bool  `json:"status" gorm:"status"`
	LeaderId int    `json:"leader_id" gorm:"leader_id"`
}

func (DepartmentUpdate) TableName() string {
	return Department{}.TableName()
}

var ErrInvalidLeaderId = common.NewCustomError(nil, "This user is not QAC", "ErrInvalidLeaderId")
