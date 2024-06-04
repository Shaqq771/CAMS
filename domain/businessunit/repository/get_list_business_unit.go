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

func (br businessRepository) GetListOfBusinessRepository(ctx context.Context) (business []model.Business, err error) {

	query := fmt.Sprintf("SELECT * FROM business_unit")
	logger.LogInfo(constant.QUERY, query)
	fmt.Println(query, "query")
	err = br.Database.DB.SelectContext(ctx, &business, query)
	fmt.Println(err, "err")

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