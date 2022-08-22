package feature

import (
	"backend-nabati/domain/logistik/model"
	repository "backend-nabati/domain/logistik/repository"
	"backend-nabati/infrastructure/broker/rabbitmq"
	"context"
)

type LogistikFeature interface {
	AddProductFeature(ctx context.Context, request *model.AddProductRequest) (response model.AddedProductResponse, err error)
	GetProductFeature(ctx context.Context, id string) (response model.Product, err error)
	GetProductListsFeature(ctx context.Context, queryRequest model.QueryRequest) (productList model.ProductLists, err error)
	DeleteProductFeature(ctx context.Context, id string) (response model.DeletedProductResponse, err error)
	UpdateProductFeature(ctx context.Context, id string, request *model.UpdateProductRequest) (response model.Product, err error)
}

type logistikFeature struct {
	logistikRepo repository.LogistikRepository
	rabbitmq     rabbitmq.RabbitMQ
}

func NewLogistikFeature(logistikRepo repository.LogistikRepository, rabbitmq rabbitmq.RabbitMQ) LogistikFeature {
	return &logistikFeature{
		logistikRepo: logistikRepo,
		rabbitmq:     rabbitmq,
	}
}
