package model

const (
	pagingDefaultIndex = 1
	pagingDefaultSize  = 25
	pagingMaxSize      = 1000
)

type Paging struct {
	Index int    `form:"pi" json:"pi"`
	Size  int    `form:"ps" json:"ps"`
	Order string `form:"po" json:"po"`
}

func (p *Paging) GetIndex() int {
	if p.Index < 1 {
		return pagingDefaultIndex
	}
	return p.Index
}

func (p *Paging) GetSize() int {
	if p.Size < 1 {
		return pagingDefaultSize
	}
	if p.Size > pagingMaxSize {
		return pagingMaxSize
	}
	return p.Size
}

func (p *Paging) GetOffset() int {
	return p.GetSize() * (p.GetIndex() - 1)
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
