package repository

import (
	"backend-nabati/domain/approver/model"
	"backend-nabati/infrastructure/database"
	"context"
)

type ApproverRepository interface {
	GetListOfApproverRepository(ctx context.Context) (approvers []model.Approver, err error)
	GetApproverByIdRepository(ctx context.Context, id int) (approver []model.Approver, err error)
}

type approverRepository struct {
	Database *database.Database
}

func NewApproverRepository(db *database.Database) ApproverRepository {
	return &approverRepository{
		Database: db,
	}
}
