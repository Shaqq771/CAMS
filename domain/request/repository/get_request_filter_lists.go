package repository

import (
	"backend-nabati/domain/request/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/domain/shared/query"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

func (rr requestRepository) GetRequestListsRepository(ctx context.Context, limit, offset int, sortby, search string) (requests []model.Request, err error) {
	if sortby == "" {
		sortby = "id asc"
	}

	if search != "" {
		search = query.SearchQueryBuilder(search)
	}
	query := fmt.Sprintf("SELECT * FROM request %s ORDER BY %s", search, sortby)

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

func (rr requestRepository) GetRequestListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (requests []model.Request, err error) {

	query, err := query.SelectStatementBuilder(model.Request{}, "request", filter)
	if err != nil {
		err = Error.New(constant.ErrDatabase, "error when create select statements", err)
		return
	}

	logger.LogInfo(constant.QUERY, query)
	if len(filter.Filters) > 0 {
		err = rr.Database.DB.SelectContext(ctx, &requests, query, offset)
	} else {
		err = rr.Database.DB.SelectContext(ctx, &requests, query)
	}

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
