package feature

import (
	"backend-nabati/domain/request/constant"
	"backend-nabati/domain/request/model"
	repository "backend-nabati/domain/request/repository"
	Error "backend-nabati/domain/shared/error"
	"context"
	"strconv"
)

type RequestFeature interface {
	GetListOfRequestFeature(ctx context.Context) (response model.RequestListNoFilter, err error)
	GetRequestFeature(ctx context.Context, id string) (response model.RequestListNoFilter, err error)
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

func (uf requestFeature) GetRequestFeature(ctx context.Context, id string) (response model.RequestListNoFilter, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := uf.requestRepo.GetRequestByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	// if result.Id == 0 {
	// 	err = Error.New(constant.ErrGeneral, constant.ErrApproverIdNotFound, errors.New(strconv.Itoa(result.Id)))
	// 	return
	// }

	var requestIds []int
  		for _, request := range result {
    	requestIds = append(requestIds, request.Id)
 	 }

	response = model.RequestListNoFilter{
		Request: result,
	}
	return
}
