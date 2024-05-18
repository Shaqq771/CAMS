package feature

import (
	"backend-nabati/domain/request/model"
	repository "backend-nabati/domain/request/repository"
	"backend-nabati/domain/sales/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
	"errors"
	"fmt"
)

type RequestFeature interface {
	GetListOfRequestFeature(ctx context.Context) (requestList model.RequestListNoFilter, err error)
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
	fmt.Println(listRequest)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrUserProductIdNotFound, errors.New(""))
		return
	}

	return
}
