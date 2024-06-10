package repository

import (
	"backend-nabati/domain/businessunit/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

func (br businessRepository) GetBusinessByIdRepository(ctx context.Context, id int) (business []model.Business, err error) {

	query := fmt.Sprintf("SELECT * FROM business_unit where id = %d", id)
	logger.LogInfo(constant.QUERY, query)
	err = br.Database.DB.SelectContext(ctx, &business, query)

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return business, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}