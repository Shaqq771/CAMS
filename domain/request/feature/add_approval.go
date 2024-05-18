package feature

// import (
// 	"backend-nabati/domain/request/constant"
// 	"backend-nabati/domain/request/model"
// 	Error "backend-nabati/domain/shared/error"
// 	shared_constant "backend-nabati/infrastructure/shared/constant"
// 	"context"
// 	"errors"
// )

// func (lf requestFeature) AddApprovalFeature(ctx context.Context, request *model.AddApprovalRequest) (resp model.AddedApprovalResponse, err error) {

// 	// Added Bussiness logic here
// 	exist, err := lf.requestRepo.CheckProductSKURepository(ctx, request.SKU)
// 	if err != nil {
// 		return
// 	} else if exist {
// 		err = Error.New(constant.ErrGeneral, constant.ErrSKUAlreadyExist, errors.New(request.SKU))
// 		return
// 	}

// 	approval := model.Approval{
// 		Name:  request.Name,
// 		SKU:   request.SKU,
// 		UOM:   request.UOM,
// 		Price: request.Price,
// 	}

// 	id, err := lf.requestRepo.InsertApprovalRepository(ctx, approval)
// 	if err != nil {
// 		return
// 	}

// 	resp = model.AddedApprovalResponse{
// 		Id:   id,
// 		Name: product.Name,
// 	}

// 	userId := 1
// 	// Check Health sales
// 	err = lf.queueService.PublishData(ctx, constant.CONSUMER_PRODUCT_INSERT_RABBITMQ, userId)
// 	if err != nil {
// 		err = Error.New(constant.ErrGeneral, shared_constant.ErrPublishQueueToBroker, err)
// 		return
// 	}

// 	return
// }
