package statisticmodel

type StatisticReq struct {
	DepartmentId int `json:"department_id" form:"department_id"`
	AcaYearId    int `json:"aca_year_id" gorm:"aca_year_id"`
}

type StatisticRespIdea struct {
	Title        []string `json:"title"`
	Id           []int    `json:"id"`
	LikeCount    []int    `json:"like_count"`
	DislikeCount []int    `json:"dislike_count"`
	ViewCount    []int    `json:"view_count"`
	CommentCount []int    `json:"comment_count"`
}

type StatisticRespTotal struct {
	TotalUser        int `json:"total_user"`
	TotalIdea        int `json:"total_idea"`
	TotalComment     int `json:"total_comment"`
	TotalInteractive int `json:"total_interactive"`
}
