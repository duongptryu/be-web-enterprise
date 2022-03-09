package ideamodel

type Filter struct {
	Title        string `json:"title" form:"title"`
	UserId       int    `json:"user_id" form:"user_id"`
	CategoryId   int    `json:"category_id" form:"category_id"`
	AcaYearId    int    `json:"aca_year_id" form:"aca_year_id"`
	DepartmentId int    `json:"department_id" form:"department_id"`
	IsAnonymous  string `json:"is_anonymous" form:"is_anonymous"`
	LikeGt       int    `json:"like_gt" form:"like_gt"`
	DislikeGt    int    `json:"dislike_gt" form:"dislike_gt"`
	ViewGt       int    `json:"view_gt" form:"view_gt"`
	LikeSt       int    `json:"like_st" form:"like_st"`
	DislikeSt    int    `json:"dislike_st" form:"dislike_st"`
	ViewSt       int    `json:"view_st" form:"view_st"`
}
