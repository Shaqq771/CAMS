package repository

import (
	"backend-nabati/domain/rule/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

func (rr ruleRepository) GetRuleByIdRepository(ctx context.Context, id int) (rule []model.Rule, err error) {

	query := fmt.Sprintf("SELECT * FROM rule where id = %d", id)
	logger.LogInfo(constant.QUERY, query)
	err = rr.Database.DB.SelectContext(ctx, &rule, query)

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return rule, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}
