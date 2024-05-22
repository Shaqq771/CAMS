package repository

import (
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/user/model"
	"backend-nabati/infrastructure/database"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
)

type UserRepository interface {
// 	GetTotalApproverWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error)
// 	GetTotalApproverRepository(ctx context.Context) (count int, err error)
// 	GetApproverByIdRepository(ctx context.Context, id int) (approver model.Approver, err error)
// 	GetTotalApproverWithConditionsRepository(ctx context.Context, conditions string) (count int, err error)
	GetListOfApproverRepository(ctx context.Context) (approvers []model.Approver, err error)
	GetApproverByIdRepository(ctx context.Context, id int) (approver []model.Approver, err error)
	//GetApproverListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (approvers []model.Approver, err error)

}

type userRepository struct {
	Database *database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{
		Database: db,
	}
}

func (ur userRepository) GetListOfApproverRepository(ctx context.Context) (approvers []model.Approver, err error) {

	query := fmt.Sprintf("SELECT * FROM approver")
	logger.LogInfo(constant.QUERY, query)
	fmt.Println(query, "query")
	err = ur.Database.DB.SelectContext(ctx, &approvers, query)
	fmt.Println(err, "err")

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return approvers, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}

func (ur userRepository) GetApproverByIdRepository(ctx context.Context, id int) (approver []model.Approver, err error) {

	query := fmt.Sprintf("SELECT * FROM approver where id = %d", id)
	logger.LogInfo(constant.QUERY, query)
	fmt.Println(query, "query")
	err = ur.Database.DB.SelectContext(ctx, &approver, query)
	fmt.Println(err, "err")

	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
		}

		if err == sql.ErrNoRows {
			return approver, nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	return
}