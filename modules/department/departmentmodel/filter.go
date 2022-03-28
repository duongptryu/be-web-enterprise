package departmentmodel

type Filter struct {
	Search     string `json:"search" form:"search"`
	LeaderName string `json:"leader_name" form:"leader_name"`
	Name       string `json:"name" form:"name"`
}
