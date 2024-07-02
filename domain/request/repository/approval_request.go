package repository

import (
	"backend-nabati/domain/request/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/infrastructure/logger"
	"context"
	"errors"
	"fmt"
)

func (rr requestRepository) ApproveRequestRepository(ctx context.Context, id int) (err error) {
	// Update request status to "approved"
	query := fmt.Sprintf("UPDATE request SET status = '%s' WHERE id = %d", constant.StatusApproved, id)
	logger.LogInfo(constant.QUERY, query)

	result, err := rr.Database.ExecContext(ctx, query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
		return
	}

	if rowsAffected == 0 {
		return Error.New(constant.ErrGeneral, constant.ErrRequestNotFound, errors.New("Request not found for approval"))
	}

	return nil
}

func (rr requestRepository) RejectRequestRepository(ctx context.Context, id int, reason string) (err error) {
	// Update request status to "rejected" and store reason (optional)
	query := fmt.Sprintf("UPDATE request SET status = '%s', rejection_reason = '%s' WHERE id = %d", constant.StatusRejected, reason, id)
	logger.LogInfo(constant.QUERY, query)

	result, err := rr.Database.ExecContext(ctx, query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
		return
	}

	if rowsAffected == 0 {
		return Error.New(constant.ErrGeneral, constant.ErrRequestNotFound, errors.New("Request not found for rejection"))
	}

	return nil
}

func (rr requestRepository) ReviseRequestRepository(ctx context.Context, id int) (err error) {
	// Update request status to "revised" (specific logic might be needed)
	query := fmt.Sprintf("UPDATE request SET status = '%s' WHERE id = %d", constant.StatusRevised, id)
	logger.LogInfo(constant.QUERY, query)

	result, err := rr.Database.ExecContext(ctx, query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
		return
	}

	if rowsAffected == 0 {
		return Error.New(constant.ErrGeneral, constant.ErrRequestNotFound, errors.New("Request not found for revision"))
	}

	// Additional logic for revision might be required here (e.g., update specific fields)
	// ...

	return nil
}