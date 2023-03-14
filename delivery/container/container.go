package container

import (
	"backend-nabati/config"
	health_feature "backend-nabati/domain/health/feature"
	health_repository "backend-nabati/domain/health/repository"
	logistik_feature "backend-nabati/domain/logistik/feature"
	logistik_repository "backend-nabati/domain/logistik/repository"
	sales_feature "backend-nabati/domain/sales/feature"
	sales_repository "backend-nabati/domain/sales/repository"
	"backend-nabati/infrastructure/broker/rabbitmq"
	"backend-nabati/infrastructure/database"
	"backend-nabati/infrastructure/logger"
	"backend-nabati/infrastructure/service/queue"
	"backend-nabati/infrastructure/shared/constant"
	"fmt"
	"log"
)

type Container struct {
	EnvironmentConfig config.EnvironmentConfig
	RabbitMQ          rabbitmq.RabbitMQ
	HealthFeature     health_feature.HealthFeature
	LogistikFeature   logistik_feature.LogistikFeature
	SalesFeature      sales_feature.SalesFeature
	QueueServices     queue.QueueService
}

func SetupContainer() Container {
	fmt.Println("Starting new container...")

	fmt.Println("Loading config...")
	config, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	logger.InitializeLogger(constant.LOGRUS) // choose which log, ZAP or LOGRUS. Default: LOGRUS

	fmt.Println("Loading database...")
	db, err := database.LoadDatabase(config.Database)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Loading message broker...")
	rmq := rabbitmq.NewConnection(config.RabbitMq)
	// Connect RabbitMQ
	err = rmq.Connect()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Loading service's...")
	queueService := queue.NewQueueService(rmq, config.RabbitMq)

	fmt.Println("Loading repository's...")
	healthRepository := health_repository.NewHealthFeature(db)
	logistikRepository := logistik_repository.NewLogistikRepository(db)
	salesRepository := sales_repository.NewSalesRepository(db)

	fmt.Println("Loading feature's...")
	healthFeature := health_feature.NewHealthFeature(config, healthRepository, rmq)
	logistikFeature := logistik_feature.NewLogistikFeature(logistikRepository, queueService)
	salesFeature := sales_feature.NewSalesFeature(salesRepository)

	return Container{
		EnvironmentConfig: config,
		RabbitMQ:          rmq,
		HealthFeature:     healthFeature,
		LogistikFeature:   logistikFeature,
		SalesFeature:      salesFeature,
		QueueServices:     queueService,
	}
}
