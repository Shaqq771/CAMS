package context

import (
	"backend-nabati/domain/logistik/constant"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateContext() context.Context {
	return context.Background()
}

func CreateContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), constant.DefaultTimeout*time.Second)
}

func SetValueToContext(ctx context.Context, c *fiber.Ctx) context.Context {
	return context.WithValue(ctx, constant.FiberContext, c)
}

func GetValueFiberFromContext(ctx context.Context) *fiber.Ctx {
	return ctx.Value(constant.FiberContext).(*fiber.Ctx)
}
