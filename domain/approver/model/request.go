package model

type AddApproverRequest struct {
	Approver_user_id string `json:"approver_user_id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Role             string `json:"role"`
	Job_title        string `json:"job_title"`
	Department       string `json:"department"`
	Location         string `json:"location"`
	Business_unit    string `json:"business_unit"`
	Description      string `json:"description"`
}

type UpdateApproverRequest struct {
	Approver_user_id string `json:"approver_user_id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Role             string `json:"role"`
	Job_title        string `json:"job_title"`
	Department       string `json:"department"`
	Location         string `json:"location"`
	Business_unit    string `json:"business_unit"`
	Description      string `json:"description"`
}
