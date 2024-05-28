package model

import "backend-nabati/domain/shared/model"

//please create request model
type RequestListNoFilter struct {
	Request []Request `json:"requests"`
}
type RequestLists struct {
	Pagination model.Pagination `json:"pagination"`
	Request    []Request        `json:"products"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

type RequestListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	Request    []Request        `json:"products"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}
