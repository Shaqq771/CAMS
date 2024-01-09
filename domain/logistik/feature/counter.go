package feature

import (
	"context"
)

func (lf logistikFeature) BulkCounterFeature(ctx context.Context) (err error) {
	var (
		size = 1000
	)

	go func() {
		// Ignore the error using the underscore (_) if you don't need to handle it explicitly
		_ = lf.logistikRepo.BulkInsertCounterRepository(ctx, size)
	}()

	return
}
