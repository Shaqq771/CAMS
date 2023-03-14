package feature_test

import (
	"backend-nabati/domain/logistik/feature"
	"backend-nabati/domain/logistik/model"
	mock_repository "backend-nabati/domain/logistik/repository/mocks"
	"context"
	"database/sql"
	"strconv"
	"testing"
	"time"

	mock_queue "backend-nabati/infrastructure/service/queue/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_UpdateProductFeature(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRepository := mock_repository.NewMockLogistikRepository(ctl)
	mockQueueService := mock_queue.NewMockQueueService(ctl)
	w := feature.NewLogistikFeature(mockRepository, mockQueueService)

	request := model.UpdateProductRequest{
		Name:  "Wafer 100gram",
		SKU:   "WFR100G2",
		Price: 2000,
		UOM:   "bungkus",
	}

	t.Run("Error invalid id from request", func(t *testing.T) {
		ctx := context.Background()

		errId := "x"
		expectResponse := model.Product{}

		resp, err := w.UpdateProductFeature(ctx, errId, &request)
		assert.Equal(t, expectResponse, resp)
		assert.NotNil(t, err)
	})

	t.Run("Product Id does'nt exist", func(t *testing.T) {
		ctx := context.Background()

		id, err := strconv.Atoi(requestGetIdProduct)
		assert.Equal(t, nil, err)

		expectResponse := model.Product{}

		mockRepository.EXPECT().CheckProductIdRepository(ctx, id).Return(false, nil) // exists, error

		resp, err := w.UpdateProductFeature(ctx, requestGetIdProduct, &request)
		assert.Equal(t, expectResponse, resp)
		assert.NotNil(t, err)
	})

	t.Run("SKU request already exist", func(t *testing.T) {
		ctx := context.Background()

		id, err := strconv.Atoi(requestGetIdProduct)
		assert.Equal(t, nil, err)

		expectResponse := model.Product{}

		mockRepository.EXPECT().CheckProductIdRepository(ctx, id).Return(true, nil)           // exists, error
		mockRepository.EXPECT().CheckProductSKURepository(ctx, request.SKU).Return(true, nil) // exists, error

		resp, err := w.UpdateProductFeature(ctx, requestGetIdProduct, &request)
		assert.Equal(t, expectResponse, resp)
		assert.NotNil(t, err)
	})

	t.Run("Error When Get product id for response", func(t *testing.T) {
		ctx := context.Background()

		id, err := strconv.Atoi(requestGetIdProduct)
		assert.Equal(t, nil, err)

		expectResponse := model.Product{}

		mockRepository.EXPECT().CheckProductIdRepository(ctx, id).Return(true, nil)            // exists, error
		mockRepository.EXPECT().CheckProductSKURepository(ctx, request.SKU).Return(false, nil) // exists, error
		mockRepository.EXPECT().UpdateProductRepository(ctx, id, &request).Return(nil)         //error

		mockResultDB := model.Product{}

		mockRepository.EXPECT().GetProductByIdRepository(ctx, id).Return(mockResultDB, nil)

		resp, err := w.UpdateProductFeature(ctx, requestGetIdProduct, &request)
		assert.Equal(t, expectResponse, resp)
		assert.NotNil(t, err)
	})

	t.Run("Update Product Success", func(t *testing.T) {
		ctx := context.Background()

		id, err := strconv.Atoi(requestGetIdProduct)
		assert.Equal(t, nil, err)

		mockRepository.EXPECT().CheckProductIdRepository(ctx, id).Return(true, nil)            // exists, error
		mockRepository.EXPECT().CheckProductSKURepository(ctx, request.SKU).Return(false, nil) // exists, error
		mockRepository.EXPECT().UpdateProductRepository(ctx, id, &request).Return(nil)         //error

		mockCreatedAt := time.Now()
		mockUpdatedAt := time.Now()
		mockResultDB := model.Product{
			Id:        1,
			Name:      "Wafer 100gram",
			SKU:       "WFR100G2",
			Price:     2000,
			UOM:       "bungkus",
			Stock:     1,
			CreatedAt: mockCreatedAt,
			UpdatedAt: mockUpdatedAt,
			DeletedAt: sql.NullTime{},
		}

		mockRepository.EXPECT().GetProductByIdRepository(ctx, id).Return(mockResultDB, nil)

		expectResponse := model.Product{
			Id:        1,
			Name:      "Wafer 100gram",
			SKU:       "WFR100G2",
			Price:     2000,
			UOM:       "bungkus",
			Stock:     1,
			CreatedAt: mockCreatedAt,
			UpdatedAt: mockUpdatedAt,
			DeletedAt: sql.NullTime{},
		}

		resp, err := w.UpdateProductFeature(ctx, requestGetIdProduct, &request)
		assert.Equal(t, expectResponse, resp)
		assert.Nil(t, err)
	})
}
