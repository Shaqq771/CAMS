package repository

import (
	"backend-nabati/domain/module/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

func (mr moduleRepository) GetListOfModuleRepository(ctx context.Context) (modules []model.Module, err error) {

	query := fmt.Sprintf("SELECT * FROM module")
	logger.LogInfo(constant.QUERY, query)
	fmt.Println(query, "query")
	err = mr.Database.DB.SelectContext(ctx, &modules, query)
	fmt.Println(err, "err")

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return modules, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}