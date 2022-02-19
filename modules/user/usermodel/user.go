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
	FullName    string    `json:"full_name" gorm:"full_name"`
	Email       string    `json:"email" gorm:"email"`
	Password    string    `json:"-" gorm:"password"`
	Gender      string    `json:"gender" gorm:"gender"`
	Department  string    `json:"department" gorm:"department"`
	DateOfBirth time.Time `json:"date_of_birth" gorm:"date_of_birth"`
	Role        string    `json:"role" gorm:"role"`
	Status      bool      `json:"status" gorm:"status"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModelCreate
	FullName    string    `json:"full_name" gorm:"full_name" binding:"require"`
	Email       string    `json:"email" gorm:"email" binding:"require"`
	Password    string    `json:"password" gorm:"password" binding:"require"`
	Gender      string    `json:"gender" gorm:"gender" binding:"require"`
	Department  string    `json:"department" gorm:"department" binding:"require"`
	DateOfBirth time.Time `json:"date_of_birth" gorm:"date_of_birth" binding:"require"`
	Role        string    `json:"role" gorm:"role" binding:"require"`
	Status      bool      `json:"status" gorm:"status"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (data *UserCreate) Validate() error {
	return nil
}

type UserUpdate struct {
	common.SQLModelUpdate
	FullName    string    `json:"full_name" gorm:"full_name"`
	Email       string    `json:"email" gorm:"email"`
	Password    string    `json:"password" gorm:"password"`
	Gender      string    `json:"gender" gorm:"gender"`
	Department  string    `json:"department" gorm:"department"`
	DateOfBirth time.Time `json:"date_of_birth" gorm:"date_of_birth"`
	Role        string    `json:"role" gorm:"role"`
	Status      bool      `json:"status" gorm:"status"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

func (data *UserUpdate) Validate() error {
	return nil
}
