package feature

import (
	"context"
)

func (lf logistikFeature) BulkCounterFeature(ctx context.Context) (err error) {
	var (
		size = 1000
	)

	go lf.logistikRepo.BulkInsertCounterRepository(ctx, size)

	return
}
