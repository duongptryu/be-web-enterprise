package departmentmodel

import (
	"fmt"
	"web/common"
)

const (
	EntityName = "Department"
)

type Department struct {
	common.SQLModel
	Name     string             `json:"name" gorm:"name"`
	Status   bool               `json:"status" gorm:"status"`
	LeaderId int                `json:"leader_id" gorm:"leader_id"`
	Leader   *common.SimpleUser `json:"leader" gorm:"foreignKey:LeaderId;preload:false"`
}

func (Department) TableName() string {
	return "departments"
}

func (data *Department) SetTags() string {
	return fmt.Sprintf("%v,%v,%v", data.Name, data.Leader.FullName, data.Status)
}

type DepartmentCreate struct {
	common.SQLModelCreate
	Name     string `json:"name" gorm:"name" binding:"required"`
	Status   bool   `json:"status" gorm:"status"`
	LeaderId int    `json:"leader_id" gorm:"leader_id" binding:"required"`
	Tags     string `json:"-" gorm:"column:tags"`
}

func (DepartmentCreate) TableName() string {
	return Department{}.TableName()
}

func (data *DepartmentCreate) SetTags() string {
	return fmt.Sprintf("%v,%v", data.Name, data.Status)
}

type DepartmentUpdate struct {
	common.SQLModel
	Name     string `json:"name" gorm:"name"`
	Status   *bool  `json:"status" gorm:"status"`
	LeaderId int    `json:"leader_id" gorm:"leader_id"`
	Tags     string `json:"-" gorm:"column:tags"`
}

func (DepartmentUpdate) TableName() string {
	return Department{}.TableName()
}

var ErrInvalidLeaderId = common.NewCustomError(nil, "This user is not QAC", "ErrInvalidLeaderId")
