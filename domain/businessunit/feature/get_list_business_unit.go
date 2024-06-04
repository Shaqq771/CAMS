package feature

import (
	"backend-nabati/domain/businessunit/model"
	"context"
)

func (bf businessFeature) GetListOfBusinessFeature(ctx context.Context) (response model.BusinessListNoFilter, err error) {

	listBusiness, err := bf.businessRepo.GetListOfBusinessRepository(ctx)
	if err != nil {
		return
	}
	response = model.BusinessListNoFilter{
		Business: listBusiness,
	}

	return
}