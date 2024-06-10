package model

type Request struct {
	Id              int    `json:"id" db:"id"`
	RuleID          int    `json:"rule_id" db:"rule_id"`
	ApproverID      int    `json:"approver_id" db:"approver_id"`
	UserID          int    `json:"user_id" db:"user_id"`
	RequestModuleId int    `json:"request_module_id" db:"request_module_id"`
	Name            string `json:"name" db:"name"`
	Email           string `json:"email" db:"email"`
	Role            string `json:"role" db:"role"`
	JobTitle        string `json:"job_title" db:"job_title"`
	Department      string `json:"department" db:"department"`
	Description     string `json:"description" db:"description"`
	Status          string `json:"status" db:"status"`
	Module          string `json:"module" db:"module"`
	Type            string `json:"type" db:"type"`
	ReviseDecision  string `json:"revise_decision" db:"revise_decision"`
	CreatedAt       string `json:"created_at" db:"created_at"`
	ModifiedAt      string `json:"modified_at" db:"modified_at"`
	CreatedBy       string `json:"created_by" db:"created_by"`
	ModifiedBy      string `json:"modified_by" db:"modified_by"`
}
