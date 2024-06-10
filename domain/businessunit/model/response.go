package model

import "backend-nabati/domain/shared/model"

//please create request model
type BusinessListNoFilter struct {
	Business []Business `json:"business_unit"`
}
type BusinessLists struct {
	Pagination model.Pagination `json:"pagination"`
	Business    []Business        `json:"business_unit"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

type BusinessListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	Business    []Business        `json:"products"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}
