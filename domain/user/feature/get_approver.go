package feature

import (
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/user/constant"
	"backend-nabati/domain/user/model"
	"context"
	"errors"
	"strconv"
)

func (lf userFeature) GetApproverFeature(ctx context.Context, id string) (response model.Approver, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := lf.userRepo.GetApproverByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	if result.Id == 0 {
		err = Error.New(constant.ErrGeneral, constant.ErrApproverIdNotFound, errors.New(strconv.Itoa(result.Id)))
		return
	}

	response = result

	return
}