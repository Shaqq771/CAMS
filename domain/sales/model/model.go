package model

import (
	"database/sql"
	"time"
)

type UserProduct struct {
	Id           int          `db:"id"`
	UserId       int          `db:"userid"`
	ProductCount int          `db:"product_count"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt    sql.NullTime `json:"-" db:"deleted_at"`
}
