package feature

import (
	"backend-nabati/config"
	"backend-nabati/domain/health/model"
	"backend-nabati/domain/health/repository"
	"backend-nabati/infrastructure/broker/rabbitmq"
	"context"
)

type HealthFeature interface {
	HealthCheck(ctx context.Context) (resp model.HealthCheck, err error)
}

type healthFeature struct {
	config           config.EnvironmentConfig
	healthRepository repository.HealthRepository
	rabbitmq         *rabbitmq.Connection
}

func NewHealthFeature(config config.EnvironmentConfig, healthRepo repository.HealthRepository, rabbitmq *rabbitmq.Connection) HealthFeature {
	return &healthFeature{
		config:           config,
		healthRepository: healthRepo,
		rabbitmq:         rabbitmq,
	}
}
