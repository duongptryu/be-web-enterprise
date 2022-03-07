package common

type SimpleUser struct {
	Id       int    `json:"id" gorm:"column:id"`
	FullName string `json:"full_name" gorm:"column:full_name"`
	Role     string `json:"role" gorm:"column:role;"`
}

func (SimpleUser) TableName() string {
	return "users"
}
