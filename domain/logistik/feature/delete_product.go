package feature

import (
	"backend-nabati/domain/logistik/constant"
	"backend-nabati/domain/logistik/model"
	Error "backend-nabati/domain/shared/error"
	"context"
	"strconv"
)

func (lf logistikFeature) DeleteProductFeature(ctx context.Context, id string) (response model.DeletedProductResponse, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	err = lf.logistikRepo.DeleteProductRepository(ctx, idInt)
	if err != nil {
		return
	}

	response = model.DeletedProductResponse{
		Id: idInt,
	}

	return
}
