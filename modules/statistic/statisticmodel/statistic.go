package statisticmodel

type StatisticReq struct {
	DepartmentId int `json:"department_id" form:"department_id"`
	AcaYearId    int `json:"aca_year_id" form:"aca_year_id"`
}

type StatisticUser struct {
	DepartmentId int `json:"department_id" form:"department_id"`
}

type StatisticRespIdea struct {
	Title        []string `json:"title"`
	Id           []int    `json:"id"`
	LikeCount    []int    `json:"like_count"`
	DislikeCount []int    `json:"dislike_count"`
	ViewCount    []int    `json:"view_count"`
	CommentCount []int    `json:"comment_count"`
}

type StatisticRespIdeaByDay struct {
	CountIdea []int    `json:"count_idea"`
	Days      []string `json:"days"`
}

type StatisticRespTotal struct {
	TotalUser        int `json:"total_user"`
	TotalIdea        int `json:"total_idea"`
	TotalInteractive int `json:"total_interactive"`
}

type StatisticRespUser struct {
	UsersName        []string `json:"users_name"`
	UsersId          []int    `json:"users_id"`
	UsersInteractive []int    `json:"users_interactive"`
}

type StatisticCountIdeaCategory struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	NumberIdea   int    `json:"number_idea"`
}

type StatisticEachUser struct {
	LikeCount     int `json:"like_count"`
	DislikeCount  int `json:"dislike_count"`
	CommentCount  int `json:"comment_count"`
	PostIdeaCount int `json:"post_idea_count"`
	ViewCount     int `json:"view_count"`
}
