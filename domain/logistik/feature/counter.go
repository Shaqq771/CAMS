package feature

import (
	"context"
)

func (lf logistikFeature) BulkCounter(ctx context.Context) (err error) {
	var (
		size = 10000
	)

	go lf.logistikRepo.BulkInsertCounter(ctx, size)

	return
}
