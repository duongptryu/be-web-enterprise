package common

type SimpleIdea struct {
	Id           int    `json:"id" gorm:"column:id"`
	Title        string `json:"title" gorm:"column:title"`
	ThumbnailUrl string `json:"thumbnail_url" gorm:"column:thumbnail_url"`
}

func (SimpleIdea) TableName() string {
	return "ideas"
}
