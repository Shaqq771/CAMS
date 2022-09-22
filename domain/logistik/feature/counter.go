package feature

import (
	"context"
)

func (lf logistikFeature) BulkCounter(ctx context.Context) (err error) {
	var (
		size = 1000
	)

	go lf.logistikRepo.BulkInsertCounter(ctx, size)

	return
}
