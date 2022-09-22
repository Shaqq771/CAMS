package repository

import (
	"backend-nabati/domain/logistik/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/domain/shared/query"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

func (lr logistikRepository) GetProductBySKURepository(ctx context.Context, sku string) (product model.Product, err error) {

	query := "SELECT * FROM product where SKU = $1 AND deleted_at IS NULL LIMIT 1"
	logger.LogInfo(constant.QUERY, query)

	rows, err := lr.Database.Queryx(query, &sku)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return product, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err := rows.StructScan(&product)
		if err != nil {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}

func (lr logistikRepository) GetProductByIdRepository(ctx context.Context, id int) (product model.Product, err error) {

	query := "SELECT * FROM product where id = $1 AND deleted_at IS NULL LIMIT 1"
	logger.LogInfo(constant.QUERY, query)

	rows, err := lr.Database.Queryx(query, &id)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return product, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err := rows.StructScan(&product)
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}

func (lr logistikRepository) GetTotalProductRepository(ctx context.Context) (count int, err error) {

	query := "SELECT COUNT(*) FROM product WHERE deleted_at IS NULL"
	rows, err := lr.Database.Query(query)
	logger.LogInfo(constant.QUERY, query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return 0, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}

func (lr logistikRepository) GetTotalProductWithConditionsRepository(ctx context.Context, conditions string) (count int, err error) {

	if conditions != "" {
		conditions = query.SearchQueryBuilder(conditions)
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM product WHERE deleted_at IS NULL %s", conditions)
	logger.LogInfo(constant.QUERY, query)

	rows, err := lr.Database.Query(query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return 0, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}

func (lr logistikRepository) GetTotalProductWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error) {

	var (
		conditions string
	)

	if filter != nil {
		conditions = query.ConditionsBuilder(filter)
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM product WHERE deleted_at IS NULL AND %s", conditions)
	logger.LogInfo(constant.QUERY, query)

	rows, err := lr.Database.Query(query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return 0, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}
