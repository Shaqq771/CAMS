package model

type AddProductRequest struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price int    `json:"price"`
	UOM   string `json:"uom"`
}

type UpdateProductRequest struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price int    `json:"price"`
	UOM   string `json:"uom"`
}
