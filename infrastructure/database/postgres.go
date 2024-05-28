package database

import (
	"backend-nabati/infrastructure/shared/constant"
	"fmt"
	"os"

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
	// 	config.Name) "root:root@(localhost:3306)/CAMS"
	db, err := sqlx.Connect(config.Dialect,
		os.Getenv("DB_USERNAME")+":"+
			os.Getenv("DB_PASSWORD")+"@"+
			"("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")"+"/"+os.Getenv("DB_NAME"))
	if err != nil {
		err = fmt.Errorf(constant.ErrConnectToDB, err)
		return
	}

	database = &Database{
		db,
	}

	return
}
