package feature

import (
	"backend-nabati/domain/approver/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
	"strconv"
)

func (af approverFeature) UpdateSkipStatusFeature(ctx context.Context, id string, isSkip bool) (err error) {
	// 1. Convert ID to integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	// 2. Update delegate status in repository
	err = af.approverRepository.UpdateSkipStatusRepository(ctx, idInt, isSkip)
	if err != nil {
		return
	}

	return nil
}