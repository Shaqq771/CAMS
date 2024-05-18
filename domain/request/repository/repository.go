package repository

import (
	"backend-nabati/domain/request/model"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/infrastructure/database"
	"context"
)

type RequestRepository interface {
	GetTotalApprovalWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error)
	GetTotalApprovalRepository(ctx context.Context) (count int, err error)
	GetApprovalByIdRepository(ctx context.Context, id int) (approval model.Approval, err error)
	GetTotalApprovalWithConditionsRepository(ctx context.Context, conditions string) (count int, err error)
	GetApprovalListsRepository(ctx context.Context, limit, offset int, sortby, search string) (approvals []model.Approval, err error)
	GetApprovalListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (approvals []model.Approval, err error)
	ApproveResponseRepository(ctx context.Context, id int, update *model.UpdatedApprovalRequest) (err error)
	RejectResponseRepository(ctx context.Context, id int, update *model.UpdatedApprovalRequest) (err error)
	ReviseResponseRepository(ctx context.Context, id int, update *model.UpdatedApprovalRequest) (err error)
	CheckApprovalIdRepository(ctx context.Context, id int) (exist bool, err error)
}

type requestRepository struct {
	Database *database.Database
}

func NewRequestRepository(db *database.Database) RequestRepository {
	return &requestRepository{
		Database: db,
	}
}
