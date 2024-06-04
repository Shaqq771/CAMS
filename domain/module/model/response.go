package model

import "backend-nabati/domain/shared/model"

//please create request model
type ModuleListNoFilter struct {
	Module []Module `json:"modules"`
}
type ModuleLists struct {
	Pagination model.Pagination `json:"pagination"`
	Module    []Module        `json:"modules"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

type ModuleListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	Module    []Module        `json:"products"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}