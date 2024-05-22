package feature

import (
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/user/constant"
	"backend-nabati/domain/user/model"
	repository "backend-nabati/domain/user/repository"
	"context"
	"strconv"
)

type UserFeature interface {
	GetListOfApproverFeature(ctx context.Context) (response model.ApproverListNoFilter, err error)
	GetApproverFeature(ctx context.Context, id string) (response model.ApproverListNoFilter, err error)
}

type userFeature struct {
	userRepo repository.UserRepository
}

func NewUserFeature(userRepo repository.UserRepository) UserFeature {
	return &userFeature{
		userRepo: userRepo,
	}
}

func (uf userFeature) GetListOfApproverFeature(ctx context.Context) (response model.ApproverListNoFilter, err error) {
	listApprover, err := uf.userRepo.GetListOfApproverRepository(ctx)
	if err != nil {
		return
	}
	response = model.ApproverListNoFilter{
		Approver: listApprover,
	}

	return
}

func (uf userFeature) GetApproverFeature(ctx context.Context, id string) (response model.ApproverListNoFilter, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := uf.userRepo.GetApproverByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	// if result.Id == 0 {
	// 	err = Error.New(constant.ErrGeneral, constant.ErrApproverIdNotFound, errors.New(strconv.Itoa(result.Id)))
	// 	return
	// }

	var approverIds []int
  		for _, approver := range result {
    	approverIds = append(approverIds, approver.Id)
 	 }

	response = model.ApproverListNoFilter{
		Approver: result,
	}
	return
}

