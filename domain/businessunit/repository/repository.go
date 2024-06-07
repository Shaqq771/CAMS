package repository

import (
	"backend-nabati/domain/businessunit/model"
	"backend-nabati/infrastructure/database"
	"context"
)

type BusinessRepository interface {
	GetListOfBusinessRepository(ctx context.Context) (business []model.Business, err error)
	GetBusinessByIdRepository(ctx context.Context, id int) (business []model.Business, err error)
	// GetRequestListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (requests []model.Request, err error)
	// GetTotalRequestWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error)
	// GetTotalRequestWithConditionsRepository(ctx context.Context, conditions string) (count int, err error)
	// GetTotalRequestRepository(ctx context.Context) (count int, err error)
	// GetRequestListsRepository(ctx context.Context, limit, offset int, sortby, search string) (requests []model.Request, err error)
}

type businessRepository struct {
	Database *database.Database
}

func NewBusinessRepository(db *database.Database) BusinessRepository {
	return &businessRepository{
		Database: db,
	}
}

