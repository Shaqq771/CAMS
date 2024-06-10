package repository

import (
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/user/model"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

func (ur userRepository) GetListOfApproverRepository(ctx context.Context) (approvers []model.Approver, err error) {

	query := fmt.Sprintf("SELECT * FROM approver")
	logger.LogInfo(constant.QUERY, query)
	fmt.Println(query, "query")
	err = ur.Database.DB.SelectContext(ctx, &approvers, query)
	fmt.Println(err, "err")

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