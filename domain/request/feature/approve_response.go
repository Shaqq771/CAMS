package feature

import (
	"backend-nabati/domain/request/constant"
	"backend-nabati/domain/request/model"
	Error "backend-nabati/domain/shared/error"
	"context"
	"errors"
	"strconv"
)

func (lf requestFeature) ApproveResponseFeature(ctx context.Context, id string, request *model.UpdatedApprovalRequest) (response model.Approval, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	// Check Product Id
	exist, err := lf.requestRepo.CheckApprovalIdRepository(ctx, idInt) 
	if err != nil {
		return
	} else if !exist {
		err = Error.New(constant.ErrGeneral, constant.ErrApprovalIdNotFound, errors.New(id))
		return
	}

	// Check Product SKU
	// if strings.TrimSpace(request.SKU) != "" {
	// 	exist, err = lf.logistikRepo.CheckProductSKURepository(ctx, request.SKU)
	// 	if err != nil {
	// 		return
	// 	} else if exist {
	// 		err = Error.New(constant.ErrGeneral, constant.ErrSKUAlreadyExist, errors.New(request.SKU))
	// 		return
	// 	}
	// }

	// Update Product
	err = lf.requestRepo.ApproveResponseRepository(ctx, idInt, request)
	if err != nil {
		return
	}

	// Get New Product
	result, err := lf.requestRepo.GetApprovalByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	if result.Id == 0 {
		err = Error.New(constant.ErrGeneral, constant.ErrApprovalIdNotFound, errors.New(strconv.Itoa(result.Id)))
		return
	}

	response = result
	return

}
