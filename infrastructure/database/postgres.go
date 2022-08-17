package database

import (
	"backend-nabati/infrastructure/shared/constant"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	datasource := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
		config.Dialect,
		config.Username,
		config.Password,
		config.Host,
		config.Name)
	db, err := sqlx.Connect(config.Dialect, datasource)
	if err != nil {
		err = fmt.Errorf(constant.ErrConnectToDB, err)
		return
	}

	database = &Database{
		db,
	}

	return
}
