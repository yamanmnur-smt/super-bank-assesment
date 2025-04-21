package pkg_data

type PageData struct {
	Limit      int    `json:"limit,omitempty"`
	Page       int    `json:"page,omitempty"`
	Sort       string `json:"sort,omitempty"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}

func (p *PageData) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *PageData) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *PageData) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PageData) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

type PaginateResponse[T any] struct {
	Data     []T      `json:"data"`
	PageData PageData `json:"page_data"`
}

type PaginateData[T any] struct {
	Data     []T      `json:"data"`
	PageData PageData `json:"page_data"`
}

type InvalidReqPayloadError struct {
	Message string
}

func (e InvalidReqPayloadError) Error() string {
	return e.Message
}
