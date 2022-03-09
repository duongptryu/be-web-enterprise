package usermodel

type Filter struct {
	Role         string `json:"role" form:"role"`
	DepartmentId int    `json:"department_id" form:"department_id"`
	Email        string `json:"email" form:"email"`
	FullName     string `json:"full_name" form:"full_name"`
	Status       string `json:"status" form:"status"`
}
