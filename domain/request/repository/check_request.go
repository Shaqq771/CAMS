package repository

import (
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
	"database/sql"
)

func (rr requestRepository) CheckRequestByIdRepository(ctx context.Context, id int) (exist bool, err error) {
	rows, err := rr.Database.QueryContext(ctx, "SELECT COUNT(*) FROM request WHERE id = %s LIMIT 1", id)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return false, nil
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
		}
	}

	return
}