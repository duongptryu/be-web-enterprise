package acayearmodel

import (
	"time"
	"web/common"
)

const EntityName = "Academic_Year"

type AcademicYear struct {
	common.SQLModel
	Title            string    `json:"title" gorm:"title"`
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
	Title            string    `json:"title" gorm:"title" binding:"required"`
	StartDate        time.Time `json:"start_date" gorm:"start_date" binding:"required"`
	EndDate          time.Time `json:"end_date" gorm:"end_date" binding:"required"`
	FirstClosureDate time.Time `json:"first_closure_date" gorm:"first_closure_date" binding:"required"`
	FinalClosureDate time.Time `json:"final_closure_date" gorm:"final_closure_date" binding:"required"`
	Status           bool      `json:"status" gorm:"status"`
}

func (AcademicYearCreate) TableName() string {
	return AcademicYear{}.TableName()
}

func (data *AcademicYearCreate) Validate() error {
	if data.StartDate.After(data.EndDate) || data.StartDate.After(data.FirstClosureDate) || data.StartDate.After(data.FinalClosureDate) {
		return ErrTimeOverLap
	}

	if data.FirstClosureDate.After(data.FinalClosureDate) || data.FirstClosureDate.After(data.EndDate) {
		return ErrTimeOverLap
	}

	if data.FinalClosureDate.After(data.EndDate) {
		return ErrTimeOverLap
	}

	return nil
}

type AcademicYearUpdate struct {
	common.SQLModelUpdate
	Title            string    `json:"title" gorm:"title"`
	StartDate        time.Time `json:"start_date" gorm:"start_date"`
	EndDate          time.Time `json:"end_date" gorm:"end_date"`
	FirstClosureDate time.Time `json:"first_closure_date" gorm:"first_closure_date"`
	FinalClosureDate time.Time `json:"final_closure_date" gorm:"final_closure_date"`
	Status           *bool     `json:"status" gorm:"status"`
}

func (AcademicYearUpdate) TableName() string {
	return AcademicYear{}.TableName()
}

func (data *AcademicYearUpdate) Validate() error {
	if data.StartDate.After(data.EndDate) || data.StartDate.After(data.FirstClosureDate) || data.StartDate.After(data.FinalClosureDate) {
		return ErrTimeOverLap
	}

	if data.FirstClosureDate.After(data.FinalClosureDate) || data.FirstClosureDate.After(data.EndDate) {
		return ErrTimeOverLap
	}

	if data.FinalClosureDate.After(data.EndDate) {
		return ErrTimeOverLap
	}

	return nil
}

type CheckStatusPostIdeaComment struct {
	CanPostComment bool `json:"can_post_comment"`
	CanPostIdea    bool `json:"can_post_idea"`
}

var ErrTimeOverLap = common.NewCustomError(nil, "Time overlap, please check again", "ErrTimeOverLap")
var ErrOverlapAcaYear = common.NewFullErrorResponse(409, nil, "There is currently an active academic year, please deactivate that academic year or change the status of the academic year you created to false to create", "There is currently an active academic year, please deactivate that academic year or change the status of the academic year you created to false to create", "ErrOverlapAcaYear")
