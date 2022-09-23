package repository

import (
	"backend-nabati/domain/logistik/model"
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/shared/helper"
	"backend-nabati/infrastructure/logger"
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
)

func (lr logistikRepository) BulkInsertCounterRepository(ctx context.Context, limit int) (err error) {
	var wg sync.WaitGroup
	for i := 0; i < limit; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, ctx context.Context, db *sqlx.DB) {

			tx, err := db.Begin()
			if err != nil {
				err = Error.New(constant.ErrDatabase, constant.ErrWhenBeginTX, err)
				wg.Done()
				return
			}

			lastNumber, err := lr.GetAndUpdateNumberNextRepository(ctx, tx)
			if err != nil {
				wg.Done()
				return
			}

			_, err = tx.ExecContext(ctx, fmt.Sprintf("UPDATE nds_number_range SET last_number = %d WHERE doc_type = '1001'", lastNumber))
			if err != nil {
				if err == context.DeadlineExceeded {
					err = Error.New(constant.ErrTimeout, constant.ErrWhenPrepareStatementDB, err)
					err = tx.Rollback()
					if err != nil {
						err = Error.New(constant.ErrDatabase, constant.ErrWhenRollBackDataToDB, err)
						return
					}
					return
				}

				err = Error.New(constant.ErrDatabase, constant.ErrWhenPrepareStatementDB, err)
				err = tx.Rollback()
				if err != nil {
					err = Error.New(constant.ErrDatabase, constant.ErrWhenRollBackDataToDB, err)
					return
				}
				return
			}

			_, err = tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO counter (number) VALUES (%d)", lastNumber))
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					err = Error.New(constant.ErrDatabase, constant.ErrWhenRollBackDataToDB, err)
					wg.Done()
					return
				}

				err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
				wg.Done()
				return
			}

			err = tx.Commit()
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					err = Error.New(constant.ErrDatabase, constant.ErrWhenRollBackDataToDB, err)
					wg.Done()
					return
				}

				err = Error.New(constant.ErrDatabase, constant.ErrWhenCommitDB, err)
				wg.Done()
				return
			}

			wg.Done()
		}(&wg, ctx, lr.Database.DB)
		wg.Wait()
	}
	return
}

func (lr logistikRepository) GetLastCounterRepository(ctx context.Context) (number string, err error) {

	query := "SELECT number FROM counter order by number desc limit 1"
	rows, err := lr.Database.Query(query)
	logger.LogInfo(constant.QUERY, query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			return "0", nil
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&number)
		if err != nil {
			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}

func (lr logistikRepository) GetDocNumberRangeRepository(ctx context.Context) (data model.NumberRange, err error) {

	query := "SELECT doc_type, plant_id, from_number, to_number, last_number, skip FROM nds_number_range WHERE doc_type = '1001' limit 1 FOR UPDATE;"
	rows, err := lr.Database.Queryx(query)
	defer rows.Close()
	logger.LogInfo(constant.QUERY, query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			err = nil
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err := rows.StructScan(&data)
		if err != nil {
			if err == context.DeadlineExceeded {
				err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
				break
			}

			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	return
}

func (lr logistikRepository) GetAndUpdateNumberNextRepository(ctx context.Context, tx *sql.Tx) (number int, err error) {
	data := model.NumberRange{}

	query := "SELECT doc_type, plant_id, from_number, to_number, last_number, skip FROM nds_number_range WHERE doc_type = '1001' limit 1 FOR UPDATE;"
	rows, err := tx.QueryContext(ctx, query)
	logger.LogInfo(constant.QUERY, query)
	if err != nil {
		if err == context.DeadlineExceeded {
			err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
			return
		}

		if err == sql.ErrNoRows {
			err = nil
			return
		}

		err = Error.New(constant.ErrDatabase, constant.ErrWhenExecuteQueryDB, err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&data.DocType, &data.PlantId, &data.FromNumber, &data.ToNumber, &data.LastNumber, &data.SkipNumber)
		if err != nil {
			if err == context.DeadlineExceeded {
				err = Error.New(constant.ErrTimeout, constant.ErrWhenExecuteQueryDB, err)
				break
			}

			err = Error.New(constant.ErrDatabase, constant.ErrWhenScanResultDB, err)
			break
		}
	}

	number = helper.LastDocNumber(data.LastNumber, data.FromNumber, data.ToNumber, data.SkipNumber)
	if number == 0 {
		logger.LogInfo(constant.QUERY, "skip transaction: "+data.LastNumber)
		return
	}

	return
}
