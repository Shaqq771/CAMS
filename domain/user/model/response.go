package model

import "backend-nabati/domain/shared/model"

type ApproverLists struct {
	Pagination model.Pagination `json:"pagination"`
	Approver    []Approver       `json:"approvers"`
	Sort       []string         `json:"sort,omitempty"`
	Filter     []string         `json:"filter,omitempty"`
}

type ApproverListsByFilter struct {
	Pagination model.Pagination `json:"pagination"`
	Approver    []Approver        `json:"approvers"`
	Filters    []model.Fields   `json:"filters,omitempty"`
}

type ApproverListNoFilter struct {
	Approver []Approver `json:"approvers"`
}
