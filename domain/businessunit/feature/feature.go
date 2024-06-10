package feature

import (
	"backend-nabati/domain/businessunit/model"
	repository "backend-nabati/domain/businessunit/repository"
	"context"
)

type BusinessFeature interface {
	GetListOfBusinessFeature(ctx context.Context) (response model.BusinessListNoFilter, err error)
	GetBusinessFeature(ctx context.Context, id string) (response model.BusinessListNoFilter, err error)
}

type businessFeature struct {
	businessRepo repository.BusinessRepository
}

func NewBusinessFeature(businessRepo repository.BusinessRepository) BusinessFeature {
	return &businessFeature{
		businessRepo: businessRepo,
	}
}

