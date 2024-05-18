package repository

import (
	"backend-nabati/domain/request/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/infrastructure/database"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

type RequestRepository interface {
	GetListOfRequestRepository(ctx context.Context) (requests []model.RequestListNoFilter, err error)
}

type requestRepository struct {
	Database *database.Database
}

func NewRequestRepository(db *database.Database) RequestRepository {
	return &requestRepository{
		Database: db,
	}
}

func (rr requestRepository) GetListOfRequestRepository(ctx context.Context) (requests []model.RequestListNoFilter, err error) {

	query := fmt.Sprintf("SELECT * FROM Request")
	logger.LogInfo(constant.QUERY, query)

	err = rr.Database.DB.SelectContext(ctx, &requests, query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return requests, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}
