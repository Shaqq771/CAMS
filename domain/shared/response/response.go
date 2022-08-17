package response

import (
	"backend-nabati/domain/shared/constant"
	Shared "backend-nabati/domain/shared/context"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/infrastructure/logger"
	"context"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

func ResponseOK(c *fiber.Ctx, msg string, data interface{}) error {

	logger.LogInfoWithData(data, constant.RESPONSE, msg)
	response := Response{
		Status:  constant.SUCCESS,
		Message: msg,
		Data:    data,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func ResponseError(c *fiber.Ctx, msg string, err error, data interface{}) error {

	response := Response{
		Status:  constant.ERROR,
		Message: fmt.Sprintf("%s: %s", msg, err.Error()),
		Data:    data,
	}

	return c.Status(http.StatusBadGateway).JSON(response)
}

func ResponseErrorWithContext(ctx context.Context, err error) error {

	var (
		errType    string
		statusCode = http.StatusBadRequest
	)

	errType, err = Error.TrimMesssage(err)
	// Set Status Code
	if errType == constant.ErrDatabase || errType == constant.ErrTimeout {
		statusCode = http.StatusInternalServerError
	} else if errType == constant.ErrAuth {
		statusCode = http.StatusUnauthorized
	}

	logger.LogError(constant.RESPONSE, errType, err.Error())

	response := Response{
		Status:  constant.ERROR,
		Message: errType,
		Data:    nil,
	}

	c := Shared.GetValueFiberFromContext(ctx)

	return c.Status(statusCode).JSON(response)
}
