package repository

import (
	"backend-nabati/domain/request/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
	"fmt"
	"strings"
)

func (rr requestRepository) UpdateRequestRepository(ctx context.Context, id int, update *model.UpdateApprovalRequest) (err error) {

	buildQuery := []string{}
	if update.Status != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("status = '%d'", update.Status))
	}

	updateQuery := strings.Join(buildQuery, ",")
	query := fmt.Sprintf("UPDATE request SET %s , updated_at = now() WHERE id = $1", updateQuery)

	tx := rr.Database.DB.MustBegin()
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