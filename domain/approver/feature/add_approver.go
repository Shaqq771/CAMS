package feature

import (
	"backend-nabati/domain/approver/model"
	"context"
)

func (af approverFeature) AddApproverFeature(ctx context.Context, request *model.AddApproverRequest) (resp model.AddedProductResponse, err error) {

	// Added Bussiness logic here
	exist, err := af.approverRepo.CheckApproverEmailRepository(ctx, request.Email)
	if err != nil {
		return
	} else if exist {
		return
	}

	product := model.Approver{
		Name:  request.Name,
		SKU:   request.SKU,
		UOM:   request.UOM,
		Price: request.Price,
	}

	id, err := lf.logistikRepo.InsertProductRepository(ctx, product)
	if err != nil {
		return
	}

	resp = model.AddedProductResponse{
		Id:   id,
		Name: product.Name,
	}

	// userId := 1
	// Check Health sales
	// err = lf.queueService.PublishData(ctx, constant.CONSUMER_PRODUCT_INSERT_RABBITMQ, userId)
	// if err != nil {
	// 	err = Error.New(constant.ErrGeneral, shared_constant.ErrPublishQueueToBroker, err)
	// 	return
	// }

	return
}
