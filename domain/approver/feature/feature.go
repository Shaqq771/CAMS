package feature

import (
	"backend-nabati/domain/approver/model"
	"backend-nabati/domain/approver/repository"
	"context"
)

type ApproverFeature interface {
	GetListOfApproverFeature(ctx context.Context) (response model.ApproverListNoFilter, err error)
	GetApproverFeature(ctx context.Context, id string) (response model.ApproverListNoFilter, err error)
	AddApproverFeature(ctx context.Context, request *model.AddApproverRequest) (response model.AddedApproverResponse, err error)
}

type approverFeature struct {
	approverRepository repository.ApproverRepository
}

func NewApproverFeature(approverRepo repository.ApproverRepository) ApproverFeature {
	return &approverFeature{
		approverRepository: approverRepo,
	}
}
