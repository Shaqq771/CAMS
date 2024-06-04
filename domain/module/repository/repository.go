package repository

import (
	"backend-nabati/domain/module/model"
	"backend-nabati/infrastructure/database"
	"context"
)

type ModuleRepository interface {
	GetListOfModuleRepository(ctx context.Context) (modules []model.Module, err error)
	// GetRequestByIdRepository(ctx context.Context, id int) (request []model.Request, err error)
	// GetRequestListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (requests []model.Request, err error)
	// GetTotalRequestWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error)
	// GetTotalRequestWithConditionsRepository(ctx context.Context, conditions string) (count int, err error)
	// GetTotalRequestRepository(ctx context.Context) (count int, err error)
	// GetRequestListsRepository(ctx context.Context, limit, offset int, sortby, search string) (requests []model.Request, err error)
}

type moduleRepository struct {
	Database *database.Database
}

func NewModuleRepository(db *database.Database) ModuleRepository {
	return &moduleRepository{
		Database: db,
	}
}