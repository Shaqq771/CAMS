package repository

import (
	"backend-nabati/domain/approver/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
)

func (ar *approverRepository) InsertApproverRepository(ctx context.Context, approver model.Approver) (id int, err error) {

	tx := ar.Database.DB.MustBegin()
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO approver (approver_user_id, name, email, role, job_title, department, location, business_unit, description, business_unit_id, delegate_status, flag_skip_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id")
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenPrepareStatementDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenPrepareStatementDB, err)
		return
	}

	err = stmt.QueryRowContext(ctx, &approver.ApproverUserId, &approver.Name, &approver.Email, &approver.Role, &approver.JobTitle, &approver.Department, &approver.Location, &approver.BusinessUnit, &approver.Description, &approver.BusinessUnitId, &approver.DelegateStatus, &approver.FlagSkipStatus).Scan(&id)
	if err != nil {
		if err == context.DeadlineExceeded {
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
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
		return
	}

	return
}
