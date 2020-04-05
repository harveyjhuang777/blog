package model

type Paging struct {
	Index int    `form:"pi" json:"pi"`
	Size  int    `form:"ps" json:"ps"`
	Order string `form:"po" json:"po"`
}

func (page *Paging) Validate() {
	if page.Index == 0 {
		page.Index = 1
	}

	if page.Size == 0 {
		page.Size = 20
	}

	//prevent web attack
	if page.Size > 1000 {
		page.Size = 1000
	}

	if page.Order == "" {
		page.Order = "id"
	}
}

func NewPagingResult(paging *Paging, count int) *PagingResult {
	totalPage := count / paging.Size
	if count%paging.Size > 0 {
		totalPage++
	}

	return &PagingResult{
		Index:     paging.Index,
		Size:      paging.Size,
		TotalPage: totalPage,
		TotalRow:  count,
	}
}

type PagingResult struct {
	Index     int `json:"pi"`         // 頁碼
	Size      int `json:"ps"`         // 比數
	TotalPage int `json:"total_page"` // 總頁數
	TotalRow  int `json:"total_row"`  // 總筆數
}
