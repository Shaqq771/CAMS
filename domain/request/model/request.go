package model

type AddProductRequest struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price int    `json:"price"`
	UOM   string `json:"uom"`
}

type UpdatedApprovalRequest struct {
	Name        string `json:"name"`
	Id          int    `json:"id" db:"id"`
	Description string `json:"description" db:"description"`
	Status      string `json:"status" db:"status"`
}
