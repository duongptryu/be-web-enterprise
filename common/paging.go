package common

import (
	"encoding/json"
	"net/url"
)

type Paging struct {
	Page       int   `json:"page" form:"page"`
	Limit      int   `json:"limit" form:"limit"`
	Total      int64 `json:"total" form:"total"`
	FakeCursor int   `json:"cursor" form:"cursor"`
	NextCursor int   `json:"next_cursor"`
}

func (p *Paging) ParsePaging(input string) error {
	if len(input) == 0{
		return nil
	}
	decodedValue, err := url.QueryUnescape(input)
	if err != nil {
		return ErrParseJson(err)
	}

	if err := json.Unmarshal([]byte(decodedValue), p); err != nil {
		return ErrParseJson(err)
	}
	return nil
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 50
	}
}
