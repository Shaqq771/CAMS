package feature

import (
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/domain/user/model"
	repository "backend-nabati/domain/user/repository"
	"backend-nabati/infrastructure/service/queue"
	"context"
)

type UserFeature interface {
	GetApproverListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (approvalList model.ApproverLists, err error)
	GetListsApproverWithFilters(ctx context.Context, filter *shared_model.Filter) (approvalList model.ApproverListsByFilter, err error)
}

type userFeature struct {
	userRepo repository.UserRepository
	queueService queue.QueueService
}

func NewUserFeature(userRepo repository.UserRepository, queueService queue.QueueService) UserFeature {
	return &userFeature{
		userRepo: userRepo,
		queueService: queueService,
	}
}
