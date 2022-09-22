package feature_test

import (
	"backend-nabati/domain/logistik/feature"
	"backend-nabati/domain/logistik/model"
	mock_repository "backend-nabati/domain/logistik/repository/mocks"
	"backend-nabati/domain/shared/helper"
	shared_model "backend-nabati/domain/shared/model"
	mock_rabbitmq "backend-nabati/infrastructure/broker/rabbitmq/mocks"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetProductListsFeature(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRepository := mock_repository.NewMockLogistikRepository(ctl)
	mockRabbitMQ := mock_rabbitmq.NewMockRabbitMQ(ctl)
	w := feature.NewLogistikFeature(mockRepository, mockRabbitMQ)

	t.Run("Get Lists Product Success", func(t *testing.T) {
		ctx := context.Background()

		mockQueryRequest := shared_model.QueryRequest{
			Page:  1,
			Limit: 5,
		}

		mockTotal := 3
		mockSortBy := ""
		mockSearch := ""

		offset, total_page := helper.GetPaginations(mockTotal, mockQueryRequest.Limit, mockQueryRequest.Page)

		mockRepository.EXPECT().GetTotalProductRepository(ctx).Return(mockTotal, nil)

		mockCreatedAt := time.Now()
		mockUpdatedAt := time.Now()

		mockProduct1 := model.Product{
			Id:        1,
			Name:      "Wafer 100gram",
			SKU:       "WFR1001GR",
			Price:     1000,
			UOM:       "pcs",
			Stock:     1,
			CreatedAt: mockCreatedAt,
			UpdatedAt: mockUpdatedAt,
		}

		mockProduct2 := model.Product{
			Id:        2,
			Name:      "Wafer 1000gram",
			SKU:       "WFR10001GR",
			Price:     10000,
			UOM:       "pcs",
			Stock:     1,
			CreatedAt: mockCreatedAt,
			UpdatedAt: mockUpdatedAt,
		}

		mockProduct3 := model.Product{
			Id:        3,
			Name:      "Wafer 10000gram",
			SKU:       "WFR100001GR",
			Price:     100000,
			UOM:       "pcs",
			Stock:     1,
			CreatedAt: mockCreatedAt,
			UpdatedAt: mockUpdatedAt,
		}

		mockResultsProduct := []model.Product{mockProduct1, mockProduct2, mockProduct3}
		mockRepository.EXPECT().GetProductListsRepository(ctx, mockQueryRequest.Limit, offset, mockSortBy, mockSearch).Return(mockResultsProduct, nil)

		expectResponse := model.ProductLists{
			Pagination: shared_model.Pagination{
				Limit:     mockQueryRequest.Limit,
				TotalPage: total_page,
				TotalRows: mockTotal,
				Page:      mockQueryRequest.Page,
			},
			Product: mockResultsProduct,
			Sort:    nil,
			Filter:  nil,
		}

		resp, err := w.GetProductListsFeature(ctx, mockQueryRequest)
		assert.Nil(t, err)
		assert.Equal(t, expectResponse, resp)
	})
}
