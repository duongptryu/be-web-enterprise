package exportmodel

type Export struct {
	AcaYearId   int `json:"aca_year_id" binding:"required"`
	NameAcaYear string
}
