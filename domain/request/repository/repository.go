package repository

import (
	"backend-nabati/domain/request/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/infrastructure/database"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

type RequestRepository interface {
	GetListOfRequestRepository(ctx context.Context) (requests []model.Request, err error)
	GetRequestByIdRepository(ctx context.Context, id int) (request []model.Request, err error)
	GetRequestListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (requests []model.Request, err error)
	GetTotalRequestWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error)
	GetTotalRequestWithConditionsRepository(ctx context.Context, conditions string) (count int, err error)
	GetTotalRequestRepository(ctx context.Context) (count int, err error)
	GetRequestListsRepository(ctx context.Context, limit, offset int, sortby, search string) (requests []model.Request, err error)
}

type requestRepository struct {
	Database *database.Database
}

func NewRequestRepository(db *database.Database) RequestRepository {
	return &requestRepository{
		Database: db,
	}
}

func (rr requestRepository) GetListOfRequestRepository(ctx context.Context) (requests []model.Request, err error) {

	query := fmt.Sprintf("SELECT * FROM request")
	logger.LogInfo(constant.QUERY, query)
	fmt.Println(query, "query")
	err = rr.Database.DB.SelectContext(ctx, &requests, query)
	fmt.Println(err, "err")

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

func (rr requestRepository) GetRequestByIdRepository(ctx context.Context, id int) (request []model.Request, err error) {

	query := fmt.Sprintf("SELECT * FROM request where id = %d", id)
	logger.LogInfo(constant.QUERY, query)
	fmt.Println(query, "query")
	err = rr.Database.DB.SelectContext(ctx, &request, query)
	fmt.Println(err, "err")

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return request, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}
