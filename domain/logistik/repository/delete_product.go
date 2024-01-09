package repository

import (
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
)

func (lr logistikRepository) DeleteProductRepository(ctx context.Context, id int) (err error) {

	tx := lr.Database.DB.MustBegin()
	stmt, err := tx.PrepareContext(ctx, "UPDATE product SET deleted_at = now(), updated_at = now() WHERE id = $1")
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenPrepareStatementDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenPrepareStatementDB, err)
		return
	}

	err = stmt.QueryRowContext(ctx, &id).Err()
	if err != nil {
		if err == context.DeadlineExceeded {
			err = tx.Rollback()
			if err != nil {
				err = Error.New(constant.ErrDatabase, constant.ErrWhenRollBackDataToDB, err)
				return
			}
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		err = tx.Rollback()
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenRollBackDataToDB, err)
			return
		}
		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		if err == context.DeadlineExceeded {
			err = tx.Rollback()
			if err != nil {
				err = Error.New(constant.ErrDatabase, constant.ErrWhenRollBackDataToDB, err)
				return
			}
			err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
		return
	}

	return
}
