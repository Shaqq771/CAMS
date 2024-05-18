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

func (lr requestRepository) GetApprovalListsRepository(ctx context.Context, limit, offset int, sortby, search string) (approvals []model.Approval, err error) {

	if sortby == "" {
		sortby = "id asc"
	}

	if search != "" {
		search = query.SearchQueryBuilder(search)
	}

	query := fmt.Sprintf("SELECT * FROM Approval ORDER BY %s LIMIT $1 OFFSET $2", search, sortby)
	logger.LogInfo(constant.QUERY, query)

	err = lr.Database.DB.SelectContext(ctx, &approvals, query, limit, offset)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return approvals, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}

func (lr requestRepository) GetApprovalListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (approvals []model.Approval, err error) {

	query, err := query.SelectStatementBuilder(model.Approval{}, "approval", filter)
	if err != nil {
		err = Error.New(constant.ErrDatabase, "error when create select statements", err)
		return
	}

	logger.LogInfo(constant.QUERY, query)
	if len(filter.Filters) > 0 {
		err = lr.Database.DB.SelectContext(ctx, &approvals, query, offset)
	} else {
		err = lr.Database.DB.SelectContext(ctx, &approvals, query)
	}

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return approvals, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}