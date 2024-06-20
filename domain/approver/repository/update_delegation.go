package repository

import (
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
)

func (ar *approverRepository) UpdateDelegateStatusRepository(ctx context.Context, id int, isDelegate bool) (err error) {
	// 1. Build the update query
	stmt, err := ar.Database.DB.PrepareContext(ctx, "UPDATE approver SET delegate_status = ? WHERE id = ?")
	if err != nil {
		err = Error.New(constant.ErrDatabase, constant.ErrWhenPrepareStatementDB, err)
		return
	}
	defer stmt.Close()

	// 2. Execute the update query
	_, err = stmt.ExecContext(ctx, isDelegate, id)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}
		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return nil
}