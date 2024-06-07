package feature

import (
	"backend-nabati/domain/businessunit/constant"
	"backend-nabati/domain/businessunit/model"
	Error "backend-nabati/domain/shared/error"
	"context"
	"strconv"
)

func (bf businessFeature) GetBusinessFeature(ctx context.Context, id string) (response model.BusinessListNoFilter, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := bf.businessRepo.GetBusinessByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	// if result.Id == 0 {
	// 	err = Error.New(constant.ErrGeneral, constant.ErrApproverIdNotFound, errors.New(strconv.Itoa(result.Id)))
	// 	return
	// }

	var businessIds []int
	for _, business := range result {
		businessIds = append(businessIds, business.Id)
	}

	response = model.BusinessListNoFilter{
		Business: result,
	}
	return
}