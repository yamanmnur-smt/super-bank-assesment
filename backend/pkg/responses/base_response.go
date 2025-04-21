package pkg_response

import pkg_data "yamanmnur/simple-dashboard/pkg/data"

type MetaData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

type GenericResponse struct {
	MetaData MetaData    `json:"meta_data"`
	Data     interface{} `json:"data"`
}

type BasicResponse struct {
	MetaData MetaData `json:"meta_data"`
}

type PaginateResponse[T any] struct {
	MetaData MetaData          `json:"meta_data"`
	Data     []T               `json:"data"`
	PageData pkg_data.PageData `json:"page_data"`
}

func NewPaginateResponse[T any](metaData MetaData) *PaginateResponse[T] {
	return &PaginateResponse[T]{
		MetaData: metaData,
	}
}

func (p *PaginateResponse[T]) GetPageData() pkg_data.PageData {
	return p.PageData
}

func (p *PaginateResponse[T]) SetPageData(pageData pkg_data.PageData) {
	p.PageData = pageData
}

func (p *PaginateResponse[T]) GetData() []T {
	return p.Data
}

func (p *PaginateResponse[T]) SetData(data []T) {
	p.Data = data
}
