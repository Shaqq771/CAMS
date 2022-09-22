package repository

import (
	"backend-nabati/domain/logistik/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/domain/shared/query"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

func (lr logistikRepository) GetProductListsRepository(ctx context.Context, limit, offset int, sortby, search string) (products []model.Product, err error) {

	if sortby == "" {
		sortby = "id asc"
	}

	if search != "" {
		search = query.SearchQueryBuilder(search)
	}

	query := fmt.Sprintf("SELECT * FROM Product WHERE deleted_at IS NULL %s  ORDER BY %s LIMIT $1 OFFSET $2", search, sortby)
	logger.LogInfo(constant.QUERY, query)

	err = lr.Database.DB.SelectContext(ctx, &products, query, limit, offset)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return products, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}

func (lr logistikRepository) GetProductListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (products []model.Product, err error) {

	var (
		condition string
	)

	if filter != nil {
		condition = query.ConditionsBuilder(filter)
	}

	query := fmt.Sprintf("SELECT * FROM Product WHERE deleted_at IS NULL AND %s ORDER BY created_at asc LIMIT $1 OFFSET $2", condition)
	logger.LogInfo(constant.QUERY, query)

	err = lr.Database.DB.SelectContext(ctx, &products, query, filter.Limit, offset)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return products, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}
