package model

import "backend-nabati/domain/shared/model"

type RequestLists struct {
	Pagination model.Pagination `json:"pagination"`
	Request    []Request        `json:"requests"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

//please create request model
type RequestListNoFilter struct {
	Request []Request `json:"requests"`
}

type RequestListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	Request    []Request        `json:"requests"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}
