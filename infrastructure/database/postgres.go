package database

import (
	"backend-nabati/infrastructure/shared/constant"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	Dialect  string
	Host     string
	Name     string
	Username string
	Password string
}

type Database struct {
	*sqlx.DB
}

func LoadDatabase(config DatabaseConfig) (database *Database, err error) {

	// datasource := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
	// 	config.Dialect,
	// 	config.Username,
	// 	config.Password,
	// 	config.Host,
	// 	config.Name)
	db, err := sqlx.Connect(config.Dialect, "root:root@(localhost:3306)/cams")
	if err != nil {
		err = fmt.Errorf(constant.ErrConnectToDB, err)
		return
	}

	database = &Database{
		db,
	}

	return
}
