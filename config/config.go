package config

import (
	// "backend-nabati/infrastructure/broker/rabbitmq"
	"backend-nabati/infrastructure/database"
	"backend-nabati/infrastructure/shared/constant"
	"fmt"
	"os"
	"strconv"

	// "strings"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Env      string
	App      App
	Database database.DatabaseConfig
	// RabbitMq rabbitmq.RabbitmqConfig
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

	// rmqPort := 0
	// if os.Getenv("RABBITMQ_PORT") != "" {
	// 	rmqPort, err = strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	// 	if err != nil {
	// 		err = fmt.Errorf(constant.ErrConvertStringToInt, err)
	// 		return
	// 	}
	// }

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
		// RabbitMq: rabbitmq.RabbitmqConfig{
		// 	Host:                      strings.TrimPrefix(os.Getenv("RABBITMQ_HOST"), "http://"),
		// 	Username:                  os.Getenv("RABBITMQ_USERNAME"),
		// 	Password:                  os.Getenv("RABBITMQ_PASSWORD"),
		// 	Port:                      rmqPort,
		// 	BillingProducerName:       os.Getenv("RABBITMQ_BILLING_PRODUCER_NAME"),
		// 	BillingConsumerName:       os.Getenv("RABBITMQ_BILLING_CONSUMER_NAME"),
		// 	ProductInsertConsumerName: os.Getenv("RABBITMQ_INSERT_PRODUCT_CONSUMER_NAME"),
		// 	ProductUpdateConsumerName: os.Getenv("RABBITMQ_UPDATE_PRODUCT_CONSUMER_NAME"),
		// },
	}

	return
}
