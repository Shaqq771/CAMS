package feature

import (
	"backend-nabati/domain/logistik/constant"
	"backend-nabati/domain/logistik/model"
	Error "backend-nabati/domain/shared/error"
	"context"
	"errors"
	"strconv"
	"strings"
)

func (lf logistikFeature) UpdateProductFeature(ctx context.Context, id string, request *model.UpdateProductRequest) (response model.Product, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	// Check Product Id
	exist, err := lf.logistikRepo.CheckProductIdRepository(ctx, idInt)
	if err != nil {
		return
	} else if !exist {
		err = Error.New(constant.ErrGeneral, constant.ErrProductIdNotFound, errors.New(id))
		return
	}

	// Check Product SKU
	if strings.TrimSpace(request.SKU) != "" {
		exist, err = lf.logistikRepo.CheckProductSKURepository(ctx, request.SKU)
		if err != nil {
			return
		} else if exist {
			err = Error.New(constant.ErrGeneral, constant.ErrSKUAlreadyExist, errors.New(request.SKU))
			return
		}
	}

	// Update Product
	err = lf.logistikRepo.UpdateProductRepository(ctx, idInt, request)
	if err != nil {
		return
	}

	// Get New Product
	result, err := lf.logistikRepo.GetProductByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	if result.Id == 0 {
		err = Error.New(constant.ErrGeneral, constant.ErrProductIdNotFound, errors.New(strconv.Itoa(result.Id)))
		return
	}

	response = result
	return

}
