package feature

import (
	"backend-nabati/domain/user/model"
	repository "backend-nabati/domain/user/repository"
	"context"
)

type UserFeature interface {
	GetListOfApproverFeature(ctx context.Context) (response model.ApproverListNoFilter, err error)
}

type userFeature struct {
	userRepo repository.UserRepository
}

func NewUserFeature(userRepo repository.UserRepository) UserFeature {
	return &userFeature{
		userRepo: userRepo,
	}
}

func (uf userFeature) GetListOfApproverFeature(ctx context.Context) (response model.ApproverListNoFilter, err error) {
	listApprover, err := uf.userRepo.GetListOfApproverRepository(ctx)
	if err != nil {
		return
	}
	response = model.ApproverListNoFilter{
		Approver: listApprover,
	}

	return
}