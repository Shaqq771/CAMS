package model

type AddApproverRequest struct {
	ApproverUserId int    `json:"approver_user_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	JobTitle       string `json:"job_title"`
	Department     string `json:"department"`
	Location       string `json:"location"`
	BusinessUnit   string `json:"business_unit"`
	Description    string `json:"description"`
	DelegateStatus bool   `json:"delegate_status" db:"delegate_status"`
	FlagSkipStatus bool   `json:"flag_skip_status" db:"flag_skip_status"`
}

type UpdateApproverRequest struct {
	ApproverUserId int    `json:"approver_user_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	JobTitle       string `json:"job_title"`
	Department     string `json:"department"`
	Location       string `json:"location"`
	BusinessUnit   string `json:"business_unit"`
	Description    string `json:"description"`
	DelegateStatus bool   `json:"delegate_status" db:"delegate_status"`
	FlagSkipStatus bool   `json:"flag_skip_status" db:"flag_skip_status"`
}
