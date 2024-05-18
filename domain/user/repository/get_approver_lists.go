package repository

import (
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/domain/shared/query"
	"backend-nabati/domain/user/model"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

func (lr userRepository) GetApproverListsRepository(ctx context.Context, limit, offset int, sortby, search string) (approvers []model.Approver, err error) {

	if sortby == "" {
		sortby = "id asc"
	}

	if search != "" {
		search = query.SearchQueryBuilder(search)
	}

	query := fmt.Sprintf("SELECT * FROM Approver ORDER BY %s LIMIT $1 OFFSET $2", search, sortby)
	logger.LogInfo(constant.QUERY, query)

	err = lr.Database.DB.SelectContext(ctx, &approvers, query, limit, offset)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return approvers, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}

func (lr userRepository) GetApproverListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (approvers []model.Approver, err error) {

	query, err := query.SelectStatementBuilder(model.Approver{}, "approver", filter)
	if err != nil {
		err = Error.New(constant.ErrDatabase, "error when create select statements", err)
		return
	}

	logger.LogInfo(constant.QUERY, query)
	if len(filter.Filters) > 0 {
		err = lr.Database.DB.SelectContext(ctx, &approvers, query, offset)
	} else {
		err = lr.Database.DB.SelectContext(ctx, &approvers, query)
	}

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return approvers, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}