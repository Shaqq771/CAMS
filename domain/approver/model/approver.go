package model

type Approver struct {
	Id             int    `json:"id" db:"id"`
	ApproverUserId int    `json:"approver_user_id" db:"approver_user_id"`
	BusinessUnitId int    `json:"business_unit_id" db:"business_unit_id"`
	Name           string `json:"name" db:"name"`
	Email          string `json:"email" db:"email"`
	Role           string `json:"role" db:"role"`
	JobTitle       string `json:"job_title" db:"job_title"`
	Department     string `json:"department" db:"department"`
	Location       string `json:"location" db:"location"`
	BusinessUnit   string `json:"business_unit" db:"business_unit"`
	Description    string `json:"description" db:"description"`
	DelegateStatus bool   `json:"delegate_status" db:"delegate_status"`
	FlagSkipStatus bool   `json:"flag_skip_status" db:"flag_skip_status"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	ModifiedAt     string `json:"modified_at" db:"modified_at"`
	CreatedBy      string `json:"created_by" db:"created_by"`
	ModifiedBy     string `json:"modified_by" db:"modified_by"`
}