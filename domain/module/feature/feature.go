package feature

import (
	"backend-nabati/domain/module/model"
	"backend-nabati/domain/module/repository"
	"context"
)

type ModuleFeature interface {
	GetListOfModuleFeature(ctx context.Context) (response model.ModuleListNoFilter, err error)
	GetModuleFeature(ctx context.Context, id string) (response model.ModuleListNoFilter, err error)
	// GetRequestFilterFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (requestList model.RequestLists, err error)
}

type moduleFeature struct {
	moduleRepo repository.ModuleRepository
}

func NewModuleFeature(moduleRepo repository.ModuleRepository) ModuleFeature {
	return &moduleFeature{
		moduleRepo: moduleRepo,
	}
}