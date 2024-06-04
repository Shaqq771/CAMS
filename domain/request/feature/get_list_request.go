package feature

import (
	"backend-nabati/domain/request/model"
	"context"
)

func (rf requestFeature) GetListOfRequestFeature(ctx context.Context) (response model.RequestListNoFilter, err error) {

	listRequest, err := rf.requestRepo.GetListOfRequestRepository(ctx)
	if err != nil {
		return
	}
	response = model.RequestListNoFilter{
		Request: listRequest,
	}

	return
}