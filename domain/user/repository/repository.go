package repository

import (
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/domain/user/model"
	"backend-nabati/infrastructure/database"
	"context"
)

type UserRepository interface {
	GetTotalApproverWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error)
	GetTotalApproverRepository(ctx context.Context) (count int, err error)
	GetApproverByIdRepository(ctx context.Context, id int) (approver model.Approver, err error)
	GetTotalApproverWithConditionsRepository(ctx context.Context, conditions string) (count int, err error)
	GetApproverListsRepository(ctx context.Context, limit, offset int, sortby, search string) (approvers []model.Approver, err error)
	GetApproverListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (approvers []model.Approver, err error)
}

type userRepository struct {
	Database *database.Database
}

func NewRequestRepository(db *database.Database) UserRepository {
	return &userRepository{
		Database: db,
	}
}
