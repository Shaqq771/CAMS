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

func (rr requestRepository) GetRequestByIdRepository(ctx context.Context, id int) (request []model.Request, err error) {

	query := fmt.Sprintf("SELECT * FROM request where id = %d", id)
	logger.LogInfo(constant.QUERY, query)
	err = rr.Database.DB.SelectContext(ctx, &request, query)

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return request, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}

func (rr requestRepository) GetTotalRequestRepository(ctx context.Context) (count int, err error) {

	query := "SELECT COUNT(*) FROM request WHERE created_by IS NOT NULL"
	rows, err := rr.Database.Query(query)
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

func (rr requestRepository) GetTotalRequestWithConditionsRepository(ctx context.Context, conditions string) (count int, err error) {

	if conditions != "" {
		conditions = query.SearchQueryBuilder(conditions)
	}
	query := fmt.Sprintf("SELECT COUNT(*) FROM request WHERE created_by IS NOT NULL %s", conditions)
	logger.LogInfo(constant.QUERY, query)

	rows, err := rr.Database.Query(query)
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

func (rr requestRepository) GetTotalRequestWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error) {

	var (
		conditions string
	)

	if filter != nil {
		conditions = query.ConditionsBuilder(filter)
	}
	fmt.Println(conditions, "conditions")
	query := "SELECT COUNT(*) FROM request WHERE created_by IS NOT NULL"
	if len(filter.Filters) > 0 {
		query = fmt.Sprintf("SELECT COUNT(*) FROM request WHERE created_by IS NOT NULL AND %s", conditions)
	}

	logger.LogInfo(constant.QUERY, query)
	rows, err := rr.Database.Query(query)
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

func (rr requestRepository) CheckRequestIdRepository(ctx context.Context, id int) (exist bool, err error) {
	rows, err := rr.Database.QueryContext(ctx, "SELECT COUNT(*) FROM request WHERE id = %s", &id)
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

func (rr requestRepository) GetTotalRequestByStatusRepository(ctx context.Context, status string) (count int, err error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM request WHERE status = '%s'", status)
	logger.LogInfo(constant.QUERY, query)
  
	rows, err := rr.Database.QueryContext(ctx, query)
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
