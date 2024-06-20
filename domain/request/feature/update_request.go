package feature

import (
	"backend-nabati/domain/request/constant"
	"backend-nabati/domain/request/model"
	Error "backend-nabati/domain/shared/error"
	"context"
	"errors"
	"strconv"
)

func (rf requestFeature) UpdateRequestFeature(ctx context.Context, id string) (response model.RequestListNoFilter, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	// Check Product Id
	exist, err := rf.requestRepo.CheckRequestByIdRepository(ctx, idInt)
	if err != nil {
		return
	} else if !exist {
		err = Error.New(constant.ErrGeneral, constant.ErrApprovalIdNotFound, errors.New(id))
		return
	}

	// Update Product
	err = rf.requestRepo.UpdateRequestRepository(ctx, idInt)
	if err != nil {
		return
	}

	// Get New Product
	result, err := rf.requestRepo.GetRequestByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	var requestIds []int
	for _, request := range result {
		requestIds = append(requestIds, request.Id)
	}

	response = model.RequestListNoFilter{
		Request: result,
	}
	return
}