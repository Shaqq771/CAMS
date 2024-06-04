package model

type Module struct {
	Id          int    `json:"id" db:"id"`
	ModuleName  string `json:"module_name" db:"module_name"`
	TypeCount   int    `json:"type_count" db:"type_count"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	ModifiedAt  string `json:"modified_at" db:"modified_at"`
	CreatedBy   string `json:"created_by" db:"created_by"`
	ModifiedBy  string `json:"modified_by" db:"modified_by"`
}
