package feature_test

import (
	"backend-nabati/domain/logistik/constant"
	"backend-nabati/domain/logistik/feature"
	"backend-nabati/domain/logistik/model"
	mock_repository "backend-nabati/domain/logistik/repository/mocks"
	"context"
	"testing"

	mock_queue "backend-nabati/infrastructure/service/queue/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_AddProductFeature(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()

	mockRepository := mock_repository.NewMockLogistikRepository(ctl)
	mockQueueService := mock_queue.NewMockQueueService(ctl)
	w := feature.NewLogistikFeature(mockRepository, mockQueueService)

	request := model.AddProductRequest{
		Name:  "Wafer 100gram",
		SKU:   "WFR1001GR",
		Price: 1000,
		UOM:   "pcs",
	}

	t.Run("Error SKU Already Exist", func(t *testing.T) {
		expectResponse := model.AddedProductResponse{}

		mockRepository.EXPECT().CheckProductSKURepository(ctx, request.SKU).Return(true, nil) // exist, error

		resp, err := w.AddProductFeature(ctx, &request)
		assert.NotNil(t, err)
		assert.Equal(t, expectResponse, resp)
	})

	t.Run("Added Product Success", func(t *testing.T) {

		mockInsertProduct := model.Product{
			Name:  "Wafer 100gram",
			SKU:   "WFR1001GR",
			Price: 1000,
			UOM:   "pcs",
		}

		mockInsertId := 1

		expectResponse := model.AddedProductResponse{
			Id:   mockInsertId,
			Name: request.Name,
		}

		mockRepository.EXPECT().CheckProductSKURepository(ctx, request.SKU).Return(false, nil) // exist, error

		mockRepository.EXPECT().InsertProductRepository(ctx, mockInsertProduct).Return(mockInsertId, nil) // id, error

		mockQueueService.EXPECT().PublishData(ctx, constant.CONSUMER_PRODUCT_INSERT_RABBITMQ, 1)

		resp, err := w.AddProductFeature(ctx, &request)
		assert.Nil(t, err)
		assert.Equal(t, expectResponse, resp)
	})
}
