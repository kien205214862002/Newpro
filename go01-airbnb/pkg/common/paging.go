package common

type Paging struct {
	Page  int   `json:"page,omitempty" form:"page"`
	Limit int   `json:"pageSize" form:"pageSize"`
	Total int64 `json:"total" form:"total"`

	// Phân trang bằng cursor
	Cursor     int `json:"-" form:"cursor"`
	NextCursor int `json:"nextCursor"`
}

func (p *Paging) Fulfill() {
	if p.Cursor == 0 && p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}
}
