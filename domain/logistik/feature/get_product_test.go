package feature_test

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"backend-nabati/domain/logistik/constant"
	"backend-nabati/domain/logistik/feature"
	"backend-nabati/domain/logistik/model"
	mock_repository "backend-nabati/domain/logistik/repository/mocks"
	Error "backend-nabati/domain/shared/error"
	mock_rabbitmq "backend-nabati/infrastructure/broker/rabbitmq/mocks"
)

func Test_GetProductFeature(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()

	mockRepository := mock_repository.NewMockLogistikRepository(ctl)
	mockRabbitMQ := mock_rabbitmq.NewMockRabbitMQ(ctl)
	w := feature.NewLogistikFeature(mockRepository, mockRabbitMQ)

	t.Run("Error invalid id from request", func(t *testing.T) {
		errId := "x"
		expectResponse := model.Product{}

		resp, err := w.GetProductFeature(ctx, errId)
		assert.Equal(t, expectResponse, resp)
		assert.NotNil(t, err)
	})

	t.Run("Error when get product: product not found", func(t *testing.T) {
		id, err := strconv.Atoi(requestGetIdProduct)
		assert.Equal(t, nil, err)

		mockResultDB := model.Product{}

		expectResponse := model.Product{}
		expectError := Error.New(constant.ErrGeneral, constant.ErrProductIdNotFound, errors.New(strconv.Itoa(expectResponse.Id)))

		mockRepository.EXPECT().GetProductByIdRepository(ctx, id).Return(mockResultDB, nil)

		resp, err := w.GetProductFeature(ctx, requestGetIdProduct)
		assert.Equal(t, expectError, err)
		assert.Equal(t, expectResponse, resp)
	})

	t.Run("Get Product Success", func(t *testing.T) {
		id, err := strconv.Atoi(requestGetIdProduct)
		assert.Equal(t, nil, err)

		mockResultDB := model.Product{
			Id:        1,
			Name:      "Wafer 100gram",
			SKU:       "WFR1001GR",
			Price:     1000,
			UOM:       "pcs",
			Stock:     1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: sql.NullTime{},
		}

		expectResponse := model.Product{
			Id:        1,
			Name:      "Wafer 100gram",
			SKU:       "WFR1001GR",
			Price:     1000,
			UOM:       "pcs",
			Stock:     1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockRepository.EXPECT().GetProductByIdRepository(ctx, id).Return(mockResultDB, nil)

		resp, err := w.GetProductFeature(ctx, requestGetIdProduct)
		assert.Nil(t, err)
		assert.Equal(t, expectResponse, resp)
	})
}
