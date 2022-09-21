package query

import (
	"backend-nabati/domain/shared/constant"
	"context"
	"fmt"

	Error "backend-nabati/domain/shared/error"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func BulkInsert(ctx context.Context, db *sqlx.DB, query string, lastCounter, limit int) (err error) {
	for i := lastCounter; i <= limit; i++ {
		go func(db *sqlx.DB, number int) {
			var (
				id   int
				data = fmt.Sprintf("%09d", number)
			)
			tx := db.MustBegin()
			stmt, err := tx.PrepareContext(ctx, query)
			if err != nil {
				if err == context.DeadlineExceeded {
					err = Error.New(constant.ErrTimeout, constant.ErrWhenPrepareStatementDB, err)
					return
				}

				err = Error.New(constant.ErrDatabase, constant.ErrWhenPrepareStatementDB, err)
				return
			}

			err = stmt.QueryRowContext(ctx, &data).Scan(&id)
			if err != nil {
				if err == context.DeadlineExceeded {
					err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
					return
				}

				tx.Rollback()
				err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
				return
			}

			err = tx.Commit()
			if err != nil {
				if err == context.DeadlineExceeded {
					err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
					tx.Rollback()
					return
				}

				err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
				return
			}

			fmt.Println(fmt.Sprintf("number %d created: %s", number, data))
		}(db, i)
	}

	return
}
