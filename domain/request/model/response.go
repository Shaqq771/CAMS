package model

import "backend-nabati/domain/shared/model"

type RequestLists struct {
	Pagination model.Pagination `json:"pagination"`
	Approval   []Approval       `json:"approvals"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

//please create request model
type RequestListNoFilter struct {
	Approval []Approval `json:"approvals"`
}

type ApprovalListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	Approval   []Approval       `json:"approvals"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}
