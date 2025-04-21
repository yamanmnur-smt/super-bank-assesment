package pkg_requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type PageRequest struct {
	Search        string `query:"search"`
	PageNumber    uint   `query:"page_number"`
	PageSize      uint   `query:"page_size"`
	SortBy        string `query:"sort_by"`
	SortDirection string `query:"sort_direction"`
}

func (r PageRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.PageSize, validation.Required),
	)
}
