package container

import (
	"backend-nabati/config"
	health_feature "backend-nabati/domain/health/feature"
	logistik_feature "backend-nabati/domain/logistik/feature"

	// module_feature "backend-nabati/domain/module/feature"
	request_feature "backend-nabati/domain/request/feature"
	request_repository "backend-nabati/domain/request/repository"
	user_repository "backend-nabati/domain/user/repository"

	business_feature "backend-nabati/domain/businessunit/feature"
	business_repository "backend-nabati/domain/businessunit/repository"
	sales_feature "backend-nabati/domain/sales/feature"
	sales_repository "backend-nabati/domain/sales/repository"
	user_feature "backend-nabati/domain/user/feature"

	// "backend-nabati/infrastructure/broker/rabbitmq"
	"backend-nabati/infrastructure/database"
	"backend-nabati/infrastructure/logger"
	"backend-nabati/infrastructure/service/queue"
	"backend-nabati/infrastructure/shared/constant"
	"fmt"
	"log"
)

type Container struct {
	EnvironmentConfig config.EnvironmentConfig
	HealthFeature     health_feature.HealthFeature
	LogistikFeature   logistik_feature.LogistikFeature
	SalesFeature      sales_feature.SalesFeature
	QueueServices     queue.QueueService
	RequestFeature    request_feature.RequestFeature
	UserFeature       user_feature.UserFeature
	BusinessFeature       business_feature.BusinessFeature
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

	fmt.Println("Loading repository's...")
	salesRepository := sales_repository.NewSalesRepository(db)
	requestRepository := request_repository.NewRequestRepository(db)
	userRepository := user_repository.NewUserRepository(db)
	businessRepository := business_repository.NewBusinessRepository(db)


	salesFeature := sales_feature.NewSalesFeature(salesRepository)
	requestFeature := request_feature.NewRequestFeature(requestRepository)
	userFeature := user_feature.NewUserFeature(userRepository)
	businessFeature := business_feature.NewBusinessFeature(businessRepository)

	return Container{
		EnvironmentConfig: config,
		SalesFeature:      salesFeature,
		RequestFeature:    requestFeature,
		UserFeature:    userFeature,
		BusinessFeature:    businessFeature,
	}
}
