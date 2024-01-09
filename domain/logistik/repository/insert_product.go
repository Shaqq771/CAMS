package repository

import (
	"backend-nabati/domain/logistik/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
)

func (lr *logistikRepository) InsertProductRepository(ctx context.Context, product model.Product) (id int, err error) {

	tx := lr.Database.DB.MustBegin()
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO product (name, sku, price, uom) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenPrepareStatementDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenPrepareStatementDB, err)
		return
	}

	err = stmt.QueryRowContext(ctx, &product.Name, &product.SKU, &product.Price, &product.UOM).Scan(&id)
	if err != nil {
		if err == context.DeadlineExceeded {
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
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
		return
	}

	return
}
