package repository

import (
	"backend-nabati/domain/approver/model"
	"backend-nabati/infrastructure/database"
	"context"
)

type ApproverRepository interface {
	GetListOfApproverRepository(ctx context.Context) (approvers []model.Approver, err error)
	GetApproverByIdRepository(ctx context.Context, id int) (approver []model.Approver, err error)
	CheckApproverEmailRepository(ctx context.Context, email string) (exist bool, err error)
	InsertApproverRepository(ctx context.Context, approver model.Approver) (id int64, err error)
	UpdateDelegateStatusRepository(ctx context.Context, id int, isDelegate bool) (err error)
	UpdateSkipStatusRepository(ctx context.Context, id int, isSkip bool) (err error)
}

type approverRepository struct {
	Database *database.Database
}

func NewApproverRepository(db *database.Database) ApproverRepository {
	return &approverRepository{
		Database: db,
	}
}
