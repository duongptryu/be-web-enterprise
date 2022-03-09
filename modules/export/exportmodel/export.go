package exportmodel

type ExportIdea struct {
	AcaYearId   int `json:"aca_year_id" binding:"required"`
	NameAcaYear string
}
