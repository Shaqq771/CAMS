package config

import (
	"backend-nabati/infrastructure/broker/rabbitmq"
	"backend-nabati/infrastructure/database"
	"backend-nabati/infrastructure/shared/constant"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Env      string
	App      App
	Database database.DatabaseConfig
	RabbitMq rabbitmq.RabbitmqConfig
}

type App struct {
	Name    string
	Version string
	Port    int
}

func LoadENVConfig() (config EnvironmentConfig, err error) {
	err = godotenv.Load()
	if err != nil {
		err = fmt.Errorf(constant.ErrLoadENV, err)
		return
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		err = fmt.Errorf(constant.ErrConvertStringToInt, err)
		return
	}

	config = EnvironmentConfig{
		Env: os.Getenv("ENV"),
		App: App{
			Name:    os.Getenv("APP_NAME"),
			Version: os.Getenv("APP_VERSION"),
			Port:    port,
		},
		Database: database.DatabaseConfig{
			Dialect:  os.Getenv("DB_DIALECT"),
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		RabbitMq: rabbitmq.RabbitmqConfig{
			Host:         "localhost",
			Username:     "guest",
			Password:     "guest",
			Port:         5672,
			ConsumerName: os.Getenv("RABBITMQ_CONSUMER_NAME"),
		},
	}

	return
}
