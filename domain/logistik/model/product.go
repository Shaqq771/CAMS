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
