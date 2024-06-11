package feature

import (
	"backend-nabati/domain/request/model"
	repository "backend-nabati/domain/request/repository"
	shared_model "backend-nabati/domain/shared/model"
	"context"
)

type RequestFeature interface {
	GetListOfRequestFeature(ctx context.Context) (response model.RequestListNoFilter, err error)
	GetRequestFeature(ctx context.Context, id string) (response model.RequestListNoFilter, err error)
	GetRequestFilterFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (requestList model.RequestLists, err error)
	GetRequestListsWaitingFeature(ctx context.Context) (response model.RequestListNoFilter, err error)
	GetRequestListsApprovedFeature(ctx context.Context) (response model.RequestListNoFilter, err error)
	GetRequestListsRejectedFeature(ctx context.Context) (response model.RequestListNoFilter, err error)
	GetRequestListsRevisedFeature(ctx context.Context) (response model.RequestListNoFilter, err error)
	UpdateRequestFeature(ctx context.Context, id string) (response model.RequestListNoFilter, err error)
}

type requestFeature struct {
	requestRepo repository.RequestRepository
}

func NewRequestFeature(requestRepo repository.RequestRepository) RequestFeature {
	return &requestFeature{
		requestRepo: requestRepo,
	}
}