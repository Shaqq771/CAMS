package feature

import (
	"backend-nabati/domain/request/model"
	"context"
)

func (rf requestFeature) GetRequestListsWaitingFeature(ctx context.Context) (response model.RequestListNoFilter, err error) {

	listRequest, err := rf.requestRepo.GetRequestListsWaitingRepository(ctx)
	if err != nil {
		return
	}
	response = model.RequestListNoFilter{
		Request: listRequest,
	}

	return
}