package common

type SimpleDepartment struct {
	Id       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	LeaderId int    `json:"leader_id" gorm:"column:leader_id"`
}

func (SimpleDepartment) TableName() string {
	return "departments"
}
