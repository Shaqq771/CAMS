package feature

import (
	repository "backend-nabati/domain/user/repository"
)

type UserFeature interface {
}

type userFeature struct {
	userRepo repository.UserRepository
}

func NewUserFeature(userRepo repository.UserRepository) UserFeature {
	return &userFeature{
		userRepo: userRepo,
	}
}
