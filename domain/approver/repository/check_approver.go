package repository

import (
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
	"database/sql"
)

func (ar approverRepository) CheckApproverEmailRepository(ctx context.Context, email string) (exist bool, err error) {
	rows, err := ar.Database.QueryContext(ctx, "SELECT COUNT(*) FROM approver WHERE email = %s", email)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return true, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		var count int
		scanErr := rows.Scan(&count)
		if scanErr != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, scanErr)
			break
		}

		if count == 1 {
			exist = true
			break
		} else {
			exist = false
			break
		}
	}

	return
}