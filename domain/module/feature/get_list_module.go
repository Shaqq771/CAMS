package feature

import (
	"backend-nabati/domain/module/model"
	"context"
)

func (mf moduleFeature) GetListOfModuleFeature(ctx context.Context) (response model.ModuleListNoFilter, err error) {

	listModule, err := mf.moduleRepo.GetListOfModuleRepository(ctx)
	if err != nil {
		return
	}
	response = model.ModuleListNoFilter{
		Module: listModule,
	}

	return
}