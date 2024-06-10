package feature

import (
	"backend-nabati/domain/user/model"
	repository "backend-nabati/domain/user/repository"
	"context"
)

type UserFeature interface {
	GetListOfApproverFeature(ctx context.Context) (response model.ApproverListNoFilter, err error)
	GetApproverFeature(ctx context.Context, id string) (response model.ApproverListNoFilter, err error)
}

type userFeature struct {
	userRepo repository.UserRepository
}

func NewUserFeature(userRepo repository.UserRepository) UserFeature {
	return &userFeature{
		userRepo: userRepo,
	}
}
