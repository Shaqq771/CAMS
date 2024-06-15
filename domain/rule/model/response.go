package model

import "backend-nabati/domain/shared/model"

//please create request model
type RuleListNoFilter struct {
	Rule []Rule `json:"rules"`
}

type RuleLists struct {
	Pagination model.Pagination `json:"pagination"`
	Rule    []Rule        `json:"rules"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

type RuleListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	Rule    []Rule        `json:"rules"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}