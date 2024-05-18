package feature

import (
	"backend-nabati/domain/request/model"
	repository "backend-nabati/domain/request/repository"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/infrastructure/service/queue"
	"context"
)

type RequestFeature interface {
	GetApprovalListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (approvalList model.ApprovalLists, err error)
	GetListsApprovalWithFilters(ctx context.Context, filter *shared_model.Filter) (approvalList model.ApprovalListsByFilter, err error)
	GetApprovalFeature(ctx context.Context, id string) (response model.Approval, err error)
	
}

type requestFeature struct {
	requestRepo repository.RequestRepository
	queueService queue.QueueService
}

func NewRequestFeature(requestRepo repository.RequestRepository, queueService queue.QueueService) RequestFeature {
	return &requestFeature{
		requestRepo: requestRepo,
		queueService: queueService,
	}
}
