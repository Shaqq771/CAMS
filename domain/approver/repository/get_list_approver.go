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

func (ar approverRepository) GetListOfApproverRepository(ctx context.Context) (approvers []model.Approver, err error) {

	query := fmt.Sprintf("SELECT * FROM approver")
	logger.LogInfo(constant.QUERY, query)
	fmt.Println(query, "query")
	err = ar.Database.DB.SelectContext(ctx, &approvers, query)

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return approvers, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}
