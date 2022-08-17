package model

import (
	"database/sql"
	"time"
)

type Product struct {
	Id        int          `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	SKU       string       `json:"sku" db:"sku"`
	Price     int          `json:"price" db:"price"`
	UOM       string       `json:"uom" db:"uom"`
	Stock     int          `json:"stock" db:"stock"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"-" db:"deleted_at"`
}

type AddProductRequest struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price int    `json:"price"`
	UOM   string `json:"uom"`
}

type AddedProductResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DeletedProductResponse struct {
	Id int `json:"id"`
}

type Pagination struct {
	Limit     int `json:"limit_per_page"`
	Page      int `json:"current_page"`
	TotalPage int `json:"total_page"`
	TotalRows int `json:"total_rows"`
}

type ProductLists struct {
	Pagination Pagination `json:"pagination"`
	Product    []Product  `json:"products"`
	Sort       []string   `json:"sort,omitempty"`
	Filter     []string   `json:"filter,omitempty"`
}

type UpdateProductRequest struct {
	Name  string `json:"name"`
	SKU   string `json:"sku"`
	Price int    `json:"price"`
	UOM   string `json:"uom"`
}

type QueryRequest struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	SortBy string `query:"sort_by"`
	Search string `query:"search"`
}
