package feature

import (
	"backend-nabati/domain/request/constant"
	"backend-nabati/domain/request/model"
	Error "backend-nabati/domain/shared/error"
	"context"
	"strconv"
)

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
