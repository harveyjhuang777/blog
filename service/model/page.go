package model

type PageInfo struct {
	PageIndex int    `form:"page-index" json:"pageIndex"`
	PageSize  int    `form:"page-size" json:"pageSize"`
	PageOrder string `form:"page-order" json:"pageOrder"`
}

func (page *PageInfo) Validate() {
	if page.PageIndex == 0 {
		page.PageIndex = 1
	}

	//prevent web attack
	if page.PageSize > 1000 {
		page.PageSize = 1000
	}

	if page.PageOrder == "" {
		page.PageOrder = "id"
	}
}
