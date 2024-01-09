package query

import (
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/shared/model"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

func BulkInsert(ctx context.Context, db *sqlx.DB, query string, lastCounter, limit int) (err error) {
	for i := lastCounter; i <= limit; i++ {
		go func(db *sqlx.DB, number int) {
			var (
				id   int
				data = fmt.Sprintf("%09d", number)
			)
			tx := db.MustBegin()
			stmt, errPrepare := tx.PrepareContext(ctx, query)
			if errPrepare != nil {
				if errPrepare == context.DeadlineExceeded {
					err = Error.New(constant.ErrTimeout, constant.ErrWhenPrepareStatementDB, errPrepare)
					return
				}

				err = Error.New(constant.ErrDatabase, constant.ErrWhenPrepareStatementDB, errPrepare)
				return
			}

			err = stmt.QueryRowContext(ctx, &data).Scan(&id)
			if err != nil {
				if err == context.DeadlineExceeded {
					err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
					return
				}

				err = tx.Rollback()
				if err != nil {
					err = Error.New(constant.ErrTimeout, constant.ErrWhenRollBackDataToDB, err)
					return
				}
				err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
				return
			}

			err = tx.Commit()
			if err != nil {
				if err == context.DeadlineExceeded {
					err = tx.Rollback()
					if err != nil {
						err = Error.New(constant.ErrTimeout, constant.ErrWhenRollBackDataToDB, err)
						return
					}
					err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
					return
				}

				err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
				return
			}

			fmt.Println(fmt.Printf("number %d created: %s", number, data))
		}(db, i)
	}

	return
}

func SelectStatementBuilder(data interface{}, tableName string, filter *model.Filter) (query string, err error) {

	var (
		condition string
		fields    string
	)

	if filter != nil {
		condition = ConditionsBuilder(filter)
	}

	fields = GetFieldModel(data)
	if strings.TrimSpace(fields) == "" {
		err = errors.New("no tag 'db' in table model")
		return
	}

	if len(filter.Filters) == 0 {
		query = fmt.Sprintf("SELECT %s FROM %s WHERE deleted_at IS NULL ", fields, tableName)
	} else {
		if filter.Limit != 0 {
			query = fmt.Sprintf("SELECT %s FROM %s WHERE deleted_at IS NULL AND %s LIMIT %d OFFSET $1", fields, tableName, condition, filter.Limit)
			return
		}

		query = fmt.Sprintf("SELECT %s FROM %s WHERE deleted_at IS NULL AND %s", fields, tableName, condition)
	}

	return
}
