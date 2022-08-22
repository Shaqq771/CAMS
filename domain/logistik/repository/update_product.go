package repository

import (
	"backend-nabati/domain/logistik/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"context"
	"fmt"
	"strings"
)

func (lr logistikRepository) UpdateProductRepository(ctx context.Context, id int, update *model.UpdateProductRequest) (err error) {

	buildQuery := []string{}
	if update.Name != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("name = '%s'", update.Name))
	}
	if update.SKU != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("sku ='%s'", update.SKU))
	}
	if update.Price != 0 {
		buildQuery = append(buildQuery, fmt.Sprintf("price = %d", update.Price))
	}
	if update.UOM != "" {
		buildQuery = append(buildQuery, fmt.Sprintf("uom = '%s'", update.UOM))
	}

	updateQuery := strings.Join(buildQuery, ",")
	query := fmt.Sprintf("UPDATE product SET %s , updated_at = now() WHERE id = $1", updateQuery)

	tx := lr.Database.DB.MustBegin()
	_, err = tx.QueryContext(ctx, query, &id)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			tx.Rollback()
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenCommitDB, err)
			tx.Rollback()
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
		return
	}

	return
}
