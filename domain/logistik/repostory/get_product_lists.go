package repository

import (
	"backend-nabati/domain/logistik/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
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
		search = query.SearchQuery(search)
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
