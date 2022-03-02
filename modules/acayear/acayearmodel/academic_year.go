package acayearmodel

import (
	"time"
	"web/common"
)

const EntityName = "Academic_Year"

type AcademicYear struct {
	common.SQLModel
	Name             string    `json:"name" gorm:"name"`
	StartDate        time.Time `json:"start_date" gorm:"start_date"`
	EndDate          time.Time `json:"end_date" gorm:"end_date"`
	FirstClosureDate time.Time `json:"first_closure_date" gorm:"first_closure_date"`
	FinalClosureDate time.Time `json:"final_closure_date" gorm:"final_closure_date"`
	Status           bool      `json:"status" gorm:"status"`
}

func (AcademicYear) TableName() string {
	return "academic_years"
}

type AcademicYearCreate struct {
	common.SQLModelCreate
	Name             string    `json:"name" gorm:"name"`
	StartDate        time.Time `json:"start_date" gorm:"start_date"`
	EndDate          time.Time `json:"end_date" gorm:"end_date"`
	FirstClosureDate time.Time `json:"first_closure_date" gorm:"first_closure_date"`
	FinalClosureDate time.Time `json:"final_closure_date" gorm:"final_closure_date"`
	Status           bool      `json:"status" gorm:"status"`
}

func (AcademicYearCreate) TableName() string {
	return AcademicYear{}.TableName()
}

func (data *AcademicYearCreate) Validate() error {
	return nil
}

type AcademicYearUpdate struct {
	common.SQLModelUpdate
	Name             string    `json:"name" gorm:"name"`
	StartDate        time.Time `json:"start_date" gorm:"start_date"`
	EndDate          time.Time `json:"end_date" gorm:"end_date"`
	FirstClosureDate time.Time `json:"first_closure_date" gorm:"first_closure_date"`
	FinalClosureDate time.Time `json:"final_closure_date" gorm:"final_closure_date"`
	Status           bool      `json:"status" gorm:"status"`
}

func (AcademicYearUpdate) TableName() string {
	return AcademicYear{}.TableName()
}

func (data *AcademicYearUpdate) Validate() error {
	return nil
}
