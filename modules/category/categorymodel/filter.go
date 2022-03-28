package categorymodel

type Filter struct {
	Name   string `json:"name" form:"name"`
	Search string `json:"search" form:"search"`
}
