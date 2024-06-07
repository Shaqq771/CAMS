package feature

import (
	"backend-nabati/domain/module/constant"
	"backend-nabati/domain/module/model"
	Error "backend-nabati/domain/shared/error"
	"context"
	"strconv"
)

func (mf moduleFeature) GetModuleFeature(ctx context.Context, id string) (response model.ModuleListNoFilter, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := mf.moduleRepo.GetModuleByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	// if result.Id == 0 {
	// 	err = Error.New(constant.ErrGeneral, constant.ErrApproverIdNotFound, errors.New(strconv.Itoa(result.Id)))
	// 	return
	// }

	var moduleIds []int
	for _, module := range result {
		moduleIds = append(moduleIds, module.Id)
	}

	response = model.ModuleListNoFilter{
		Module: result,
	}
	return
}