package model

import "backend-nabati/domain/shared/model"

type ApprovalLists struct {
	Pagination model.Pagination `json:"pagination"`
	Approval    []Approval       `json:"approvals"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

type ApprovalListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	Approval    []Approval        `json:"approvals"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}