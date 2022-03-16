package categorymodel

import "web/common"

const EntityName = "Category"

type Category struct {
	common.SQLModel
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryCreate struct {
	common.SQLModelCreate
	Name   string `json:"name" gorm:"name" binding:"required"`
	Status bool   `json:"status" gorm:"status"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

type CategoryUpdate struct {
	common.SQLModelUpdate
	Name   string `json:"name" gorm:"name"`
	Status *bool  `json:"status" gorm:"status"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

var ErrCannotDelCategory = common.NewCustomError(nil, "This category has been used, cannot delete", "ErrCannotDelCategory")
