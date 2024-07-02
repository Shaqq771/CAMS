package feature

import (
	"backend-nabati/domain/request/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
	"errors"
	"strconv"
)

func (rf requestFeature) HandleApprovalRequestFeature(ctx context.Context, id string, status string, reason string) (err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	// Check Request Id
	exist, err := rf.requestRepo.CheckRequestByIdRepository(ctx, idInt)
	if err != nil {
		return
	} else if !exist {
		err = Error.New(constant.ErrGeneral, constant.ErrApprovalIdNotFound, errors.New(id))
		return
	}

	// Update logic depends on approval status
	switch status {
	case constant.StatusApproved:
		err = rf.requestRepo.ApproveRequestRepository(ctx, idInt) // Implement approval logic here
	case constant.StatusRejected:
		err = rf.requestRepo.RejectRequestRepository(ctx, idInt, reason) // Implement rejection logic here
	case constant.StatusRevised:
		err = rf.requestRepo.ReviseRequestRepository(ctx, idInt) // Implement revision logic here
	default:
		err = Error.New(constant.ErrGeneral, constant.ErrInvalidApprovalStatus, errors.New(status))
		return
	}

	if err != nil {
		return
	}

	// Optional: Perform further actions after approval (e.g., send notifications)
	// ...

	return
}