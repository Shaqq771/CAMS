package model

type Business struct {
	Id               int    `json:"id" db:"id"`
	BusinessUnitName string `json:"business_unit_name" db:"business_unit_name"`
	Description      string `json:"description" db:"description"`
	CreatedAt        string `json:"created_at" db:"created_at"`
	ModifiedAt       string `json:"modified_at" db:"modified_at"`
	CreatedBy        string `json:"created_by" db:"created_by"`
	ModifiedBy       string `json:"modified_by" db:"modified_by"`
}
