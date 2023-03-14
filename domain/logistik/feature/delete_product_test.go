package feature_test

import (
	"backend-nabati/domain/logistik/feature"
	"backend-nabati/domain/logistik/model"
	mock_repository "backend-nabati/domain/logistik/repository/mocks"
	"context"
	"strconv"
	"testing"

	mock_queue "backend-nabati/infrastructure/service/queue/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_DeleteProductFeature(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()

	mockRepository := mock_repository.NewMockLogistikRepository(ctl)
	mockQueueService := mock_queue.NewMockQueueService(ctl)
	w := feature.NewLogistikFeature(mockRepository, mockQueueService)

	t.Run("Error invalid id from request", func(t *testing.T) {
		errId := "x"
		expectResponse := model.DeletedProductResponse{}

		resp, err := w.DeleteProductFeature(ctx, errId)
		assert.Equal(t, expectResponse, resp)
		assert.NotNil(t, err)
	})

	t.Run("Deleted Product Success", func(t *testing.T) {
		id, err := strconv.Atoi(requestGetIdProduct)
		assert.Equal(t, nil, err)

		expectResponse := model.DeletedProductResponse{
			Id: id,
		}

		mockRepository.EXPECT().DeleteProductRepository(ctx, id).Return(nil) // error

		resp, err := w.DeleteProductFeature(ctx, requestGetIdProduct)
		assert.Nil(t, err)
		assert.Equal(t, expectResponse, resp)
	})
}
