package repository

import (
	"backend-nabati/domain/sales/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/infrastructure/database"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

type SalesRepository interface {
	InsertUserProductRepository(ctx context.Context, userId, counter int) (id int, err error)
	ReadSalesProductRepository(ctx context.Context, userId int) (userProduct model.UserProduct, err error)
	UpdateSalesProductRepository(ctx context.Context, userId int, count int) (err error)
}

type salesRepository struct {
	Database *database.Database
}

func NewSalesRepository(db *database.Database) SalesRepository {
	return &salesRepository{
		Database: db,
	}
}

func (sf salesRepository) InsertUserProductRepository(ctx context.Context, userId, counter int) (id int, err error) {

	tx := sf.Database.DB.MustBegin()

	query := "INSERT INTO user_product (userId, product_count) VALUES ($1, $2) RETURNING id"
	logger.LogInfo(constant.QUERY, query)
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenPrepareStatementDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenPrepareStatementDB, err)
		return
	}

	err = stmt.QueryRowContext(ctx, &userId, &counter).Scan(&id)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		tx.Rollback()
		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			tx.Rollback()
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
		return
	}

	return
}

func (sf salesRepository) ReadSalesProductRepository(ctx context.Context, userId int) (userProduct model.UserProduct, err error) {

	query := "SELECT * FROM user_product where userId = $1 AND deleted_at IS NULL LIMIT 1"
	logger.LogInfo(constant.QUERY, query)

	rows, err := sf.Database.QueryxContext(ctx, query, &userId)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return userProduct, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err := rows.StructScan(&userProduct)
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}

func (sf salesRepository) UpdateSalesProductRepository(ctx context.Context, userId int, count int) (err error) {

	query := fmt.Sprintf("UPDATE user_product SET product_count = $1, updated_at = now() WHERE userId = $2")
	logger.LogInfo(constant.QUERY, query)

	tx := sf.Database.DB.MustBegin()
	_, err = tx.QueryContext(ctx, query, &count, &userId)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			tx.Rollback()
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenCommitDB, err)
			tx.Rollback()
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
		return
	}

	return
}
