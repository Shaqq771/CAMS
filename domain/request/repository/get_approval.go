package repository

import (
	"backend-nabati/domain/request/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/domain/shared/query"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

// func (lr requestRepository) GetProductBySKURepository(ctx context.Context, sku string) (product model.Approval, err error) {

// 	query := "SELECT * FROM product where SKU = $1 AND deleted_at IS NULL LIMIT 1"
// 	logger.LogInfo(constant.QUERY, query)

// 	rows, err := lr.Database.Queryx(query, &sku)
// 	if err != nil {
// 		if err == context.DeadlineExceeded {
// 			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
// 			return
// 		}

// 		if err == sql.ErrNoRows {
// 			return product, nil
// 		}

// 		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
// 		return
// 	}

// 	for rows.Next() {
// 		errScan := rows.StructScan(&product)
// 		if errScan != nil {
// 			err = Error.New(constant.ErrTimeout, constant.ErrWhenScanResultDB, errScan)
// 			break
// 		}
// 	}

// 	return
// }

func (lr requestRepository) GetApprovalByIdRepository(ctx context.Context, id int) (approval model.Approval, err error) {

	query := "SELECT * FROM approval where id = $1 LIMIT 1"
	logger.LogInfo(constant.QUERY, query)

	rows, err := lr.Database.Queryx(query, &id)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return approval, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		errScan := rows.StructScan(&approval)
		if errScan != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, errScan)
			break
		}
	}

	return
}

func (lr requestRepository) GetTotalApprovalRepository(ctx context.Context) (count int, err error) {

	query := "SELECT COUNT(*) FROM approval"
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
		errScan := rows.Scan(&count)
		if errScan != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, errScan)
			break
		}
	}

	return
}

func (lr requestRepository) GetTotalApprovalWithConditionsRepository(ctx context.Context, conditions string) (count int, err error) {

	if conditions != "" {
		conditions = query.SearchQueryBuilder(conditions)
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM approval %s", conditions)
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
		errScan := rows.Scan(&count)
		if errScan != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, errScan)
			break
		}
	}

	return
}

func (lr requestRepository) GetTotalApprovalWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error) {

	var (
		conditions string
	)

	if filter != nil {
		conditions = query.ConditionsBuilder(filter)
	}

	query := "SELECT COUNT(*) FROM approval"
	if len(filter.Filters) > 0 {
		query = fmt.Sprintf("SELECT COUNT(*) FROM approval WHERE %s", conditions)
	}

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
		errScan := rows.Scan(&count)
		if errScan != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, errScan)
			break
		}
	}

	return
}
