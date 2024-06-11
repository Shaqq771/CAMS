package repository

import (
	"backend-nabati/domain/approver/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

func (ar approverRepository) GetApproverByIdRepository(ctx context.Context, id int) (approver []model.Approver, err error) {

	query := fmt.Sprintf("SELECT * FROM approver where id = %d", id)
	logger.LogInfo(constant.QUERY, query)
	err = ar.Database.DB.SelectContext(ctx, &approver, query)

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return approver, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}
