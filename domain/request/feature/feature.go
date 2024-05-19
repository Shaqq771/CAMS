package feature

import (
	"backend-nabati/domain/request/model"
	repository "backend-nabati/domain/request/repository"
	"context"
)

type RequestFeature interface {
	GetListOfRequestFeature(ctx context.Context) (response model.RequestListNoFilter, err error)
}

type requestFeature struct {
	requestRepo repository.RequestRepository
}

func NewRequestFeature(requestRepo repository.RequestRepository) RequestFeature {
	return &requestFeature{
		requestRepo: requestRepo,
	}
}

func (rf requestFeature) GetListOfRequestFeature(ctx context.Context) (response model.RequestListNoFilter, err error) {
	listRequest, err := rf.requestRepo.GetListOfRequestRepository(ctx)
	if err != nil {
		return
	}
	response = model.RequestListNoFilter{
		Request: listRequest,
	}

	return
}
