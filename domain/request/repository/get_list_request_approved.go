package repository

import (
	"backend-nabati/domain/request/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

func (rr requestRepository) GetRequestListsApprovedRepository(ctx context.Context) (requests []model.Request, err error) {

	query := fmt.Sprintf("SELECT * FROM request WHERE status = 'Approved'")
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