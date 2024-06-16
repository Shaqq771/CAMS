package repository

import (
	"backend-nabati/domain/approver/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
)

func (ar *approverRepository) InsertApproverRepository(ctx context.Context, approver model.Approver) (id int64, err error) {

	tx := ar.Database.DB.MustBegin()
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO approver (approver_user_id, name, email, role, job_title, department, location, description, business_unit, business_unit_id, delegation_status, flag_skip_status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenPrepareStatementDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenPrepareStatementDB, err)
		return
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, &approver.ApproverUserId, &approver.Name, &approver.Email, &approver.Role, &approver.JobTitle, &approver.Department, &approver.Location, &approver.Description, &approver.BusinessUnit, &approver.BusinessUnitId, &approver.DelegateStatus, &approver.FlagSkipStatus)
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

	id, err = res.LastInsertId()
	if err != nil {
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
