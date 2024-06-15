package model

type Rule struct {
	Id               int    `json:"id" db:"id"`
	ModuleId         int    `json:"module_id" db:"module_id"`
	Module           string `json:"module" db:"module"`
	Type             string `json:"type" db:"type"`
	Description      string `json:"description" db:"description"`
	Stage            int    `json:"stage" db:"stage"`
	Method           string `json:"method" db:"method"`
	CountLevelling   int    `json:"count_levelling" db:"count_levelling"`
	Approver         string `json:"approver" db:"approver"`
	RejectPermission bool   `json:"reject_permission" db:"reject_permission"`
	Req              int    `json:"req" db:"req"`
	EmailApprover    bool   `json:"email_approver" db:"email_approver"`
	DueTime          int    `json:"due_time" db:"due_time"`
	CaseOverdue      string `json:"case_overdue" db:"case_overdue"`
	ReviseMethod     string `json:"revise_method" db:"revise_method"`
	ReviseDecision   string `json:"revise_decision" db:"revise_decision"`
	Delegation       string `json:"delegation" db:"delegation"`
	DelegationTime   int    `json:"delegation_time" db:"delegation_time"`
	CreatedAt        string `json:"created_at" db:"created_at"`
	ModifiedAt       string `json:"modified_at" db:"modified_at"`
	CreatedBy        string `json:"created_by" db:"created_by"`
	ModifiedBy       string `json:"modified_by" db:"modified_by"`
}