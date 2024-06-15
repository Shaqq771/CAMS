package feature

import (
	"backend-nabati/domain/approver/model"
	"context"
)

func (af approverFeature) AddApproverFeature(ctx context.Context, request *model.AddApproverRequest) (resp model.AddedApproverResponse, err error) {

	// Added Bussiness logic here
	exist, err := af.approverRepository.CheckApproverEmailRepository(ctx, request.Email)
	if err != nil {
		return
	} else if exist {
		return
	}

	approver := model.Approver{
		ApproverUserId: request.ApproverUserId,
		Name:  request.Name,
		Email:   request.Email,
		Role:   request.Role,
		JobTitle: request.JobTitle,
		Department: request.Department,
		Location: request.Location,
		BusinessUnit: request.BusinessUnit,
		Description: request.Description,
	}

	id, err := af.approverRepository.InsertApproverRepository(ctx, approver)
	if err != nil {
		return
	}

	resp = model.AddedApproverResponse{
		Id:   id,
		Name: approver.Name,
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
