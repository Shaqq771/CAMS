package repository

import (
	"backend-nabati/domain/request/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
	"fmt"
	"strings"
)

func (lr requestRepository) RejectResponseRepository(ctx context.Context, id int, update *model.UpdatedApprovalRequest) (err error) {

	buildQuery := []string{}
	if update.Description != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("description = '%s'", update.Description))
	}
	if update.Status != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("status ='%s'", update.Status))
	}

	updateQuery := strings.Join(buildQuery, ",")
	query := fmt.Sprintf("UPDATE approval SET %s , modified_at = now() WHERE id = $1", updateQuery)

	tx := lr.Database.DB.MustBegin()
	_, err = tx.QueryContext(ctx, query, &id)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = tx.Rollback()
			if err != nil {
				err = Error.New(constant.ErrDatabase, constant.ErrWhenRollBackDataToDB, err)
				return
			}
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		err = tx.Rollback()
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenRollBackDataToDB, err)
			return
		}
		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		if err == context.DeadlineExceeded {
			err = tx.Rollback()
			if err != nil {
				err = Error.New(constant.ErrDatabase, constant.ErrWhenRollBackDataToDB, err)
				return
			}
			err = Error.New(constant.ErrTimeout, constant.ErrWhenCommitDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
		return
	}

	return
}
