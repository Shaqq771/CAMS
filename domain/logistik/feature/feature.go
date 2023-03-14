package feature

import (
	"backend-nabati/domain/logistik/model"
	repository "backend-nabati/domain/logistik/repository"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/infrastructure/service/queue"
	"context"
)

type LogistikFeature interface {
	AddProductFeature(ctx context.Context, request *model.AddProductRequest) (response model.AddedProductResponse, err error)
	GetProductFeature(ctx context.Context, id string) (response model.Product, err error)
	GetProductListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (productList model.ProductLists, err error)
	DeleteProductFeature(ctx context.Context, id string) (response model.DeletedProductResponse, err error)
	UpdateProductFeature(ctx context.Context, id string, request *model.UpdateProductRequest) (response model.Product, err error)
	BulkCounterFeature(ctx context.Context) (err error)
	GetListsProductWithFilters(ctx context.Context, filter *shared_model.Filter) (productList model.ProductListsByFilter, err error)
}

type logistikFeature struct {
	logistikRepo repository.LogistikRepository
	queueService queue.QueueService
}

func NewLogistikFeature(logistikRepo repository.LogistikRepository, queueService queue.QueueService) LogistikFeature {
	return &logistikFeature{
		logistikRepo: logistikRepo,
		queueService: queueService,
	}
}
