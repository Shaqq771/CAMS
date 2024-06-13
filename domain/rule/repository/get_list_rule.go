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

func (rr ruleRepository) GetListOfRuleRepository(ctx context.Context) (rules []model.Rule, err error) {

	query := fmt.Sprintf("SELECT * FROM rule")
	logger.LogInfo(constant.QUERY, query)
	fmt.Println(query, "query")
	err = rr.Database.DB.SelectContext(ctx, &rules, query)
	fmt.Println(err, "err")

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return rules, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}