package feature

import (
	"backend-nabati/domain/sales/constant"
	"backend-nabati/domain/sales/repository"
	Context "backend-nabati/domain/shared/context"
	Error "backend-nabati/domain/shared/error"
	"context"
	"errors"
	"fmt"
	"strconv"
)

type SalesFeature interface {
	UpdateSalesProductFromBroker(data string) (err error)
	UpdateSalesProduct(ctx context.Context, userId int) (err error)
}

type salesFeature struct {
	salesRepository repository.SalesRepository
}

func NewSalesFeature(salesRepo repository.SalesRepository) SalesFeature {
	return &salesFeature{
		salesRepository: salesRepo,
	}
}

func (sf salesFeature) UpdateSalesProductFromBroker(data string) (err error) {
	ctx := Context.CreateContext()
	userId, err := strconv.Atoi(data)
	if err != nil {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, errors.New(fmt.Sprintf("data from broker: %s", data)))
		return
	}

	return sf.UpdateSalesProduct(ctx, userId)
}

func (sf salesFeature) UpdateSalesProduct(ctx context.Context, userId int) (err error) {

	userProduct, err := sf.salesRepository.ReadSalesProductRepository(ctx, userId)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrUserProductIdNotFound, errors.New(strconv.Itoa(userId)))
		return
	}

	if userProduct.UserId == 0 {
		_, errInsert := sf.salesRepository.InsertUserProductRepository(ctx, userId, 1)
		if errInsert != nil {
			err = Error.New(constant.ErrGeneral, constant.ErrFailedInsertData, errInsert)
			return err
		}

		return
	}

	counter := userProduct.ProductCount + 1

	err = sf.salesRepository.UpdateSalesProductRepository(ctx, userId, counter)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedUpdateData, err)
		return
	}

	return
}
