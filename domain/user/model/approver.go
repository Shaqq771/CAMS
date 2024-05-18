package model

import "time"

type Approver struct {
	Id             int
	ApproverUserId string `json:"approver_user_id" db:"approver_user_id"`
	Name           string `json:"name" db:"name"`
	Email          string `json:"email" db:"email"`
	Role           string `json:"role" db:"role"`
	JobTitle       string `json:"job_title" db:"job_title"`
	Department     string `json:"department" db:"department"`
	BusinessUnit   string `json:"business_unit" db:"business_unit"`
	Description    string `json:"description" db:"description"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	ModifiedAt     time.Time `json:"modified_at" db:"modified_at"`
	CreatedBy      string `json:"created_by" db:"created_by"`
	ModifiedBy     string `json:"modified_by" db:"modified_by"`
}