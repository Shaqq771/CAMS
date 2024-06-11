package user

import (
	"backend-nabati/domain/user/feature"
)

type UserHandler interface {
}

type userHandler struct {
	feature feature.UserFeature
}

func NewUserHandler(feature feature.UserFeature) UserHandler {
	return &userHandler{
		feature: feature,
	}
}
