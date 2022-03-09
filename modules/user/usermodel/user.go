package usermodel

import (
	"time"
	"web/common"
)

const (
	EntityName = "User"
)

type User struct {
	common.SQLModel
	FullName     string                   `json:"full_name" gorm:"full_name"`
	Email        string                   `json:"email" gorm:"email"`
	Password     string                   `json:"-" gorm:"password"`
	Gender       string                   `json:"gender" gorm:"gender"`
	DepartmentId int                      `json:"department_id" gorm:"department_id"`
	Department   *common.SimpleDepartment `json:"department" gorm:"preload:false"`
	DateOfBirth  time.Time                `json:"date_of_birth" gorm:"date_of_birth"`
	Role         string                   `json:"role" gorm:"role"`
	Status       bool                     `json:"status" gorm:"status"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModelCreate
	FullName     string    `json:"full_name" gorm:"full_name" binding:"required"`
	Email        string    `json:"email" gorm:"email" binding:"required"`
	Password     string    `json:"password" gorm:"password" binding:"required"`
	Gender       string    `json:"gender" gorm:"gender" binding:"required"`
	DepartmentId int       `json:"department_id" gorm:"department_id"`
	DateOfBirth  time.Time `json:"date_of_birth" gorm:"date_of_birth" binding:"required"`
	Role         string    `json:"role" gorm:"role" binding:"required"`
	Status       bool      `json:"status" gorm:"status"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (data *UserCreate) Validate() error {
	return nil
}

type UserUpdate struct {
	common.SQLModelUpdate
	FullName     string    `json:"full_name" gorm:"full_name"`
	Email        string    `json:"email" gorm:"email"`
	Gender       string    `json:"gender" gorm:"gender"`
	DepartmentId int       `json:"department_id" gorm:"department_id"`
	DateOfBirth  time.Time `json:"date_of_birth" gorm:"date_of_birth"`
	Role         string    `json:"role" gorm:"role"`
	Status       *bool     `json:"status" gorm:"status"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

func (data *UserUpdate) Validate() error {
	return nil
}
