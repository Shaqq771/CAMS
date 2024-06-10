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

func (mr moduleRepository) GetModuleByIdRepository(ctx context.Context, id int) (module []model.Module, err error) {

	query := fmt.Sprintf("SELECT * FROM module where id = %d", id)
	logger.LogInfo(constant.QUERY, query)
	err = mr.Database.DB.SelectContext(ctx, &module, query)

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return module, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}