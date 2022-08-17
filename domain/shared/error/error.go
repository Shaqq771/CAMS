package error

import (
	"backend-nabati/domain/shared/constant"
	"backend-nabati/infrastructure/logger"
	"errors"
	"fmt"
	"strings"
)

func New(tipe string, msg string, err error) error {
	logger.LogError(constant.Err, tipe, fmt.Sprintf(msg+err.Error()))
	return fmt.Errorf("%s | %s: %w", tipe, msg, err)
}

func TrimMesssage(err error) (tipe string, newErr error) {
	errs := strings.Split(err.Error(), "|")
	tipe = strings.TrimSpace(errs[0])

	newErr = errors.New(strings.TrimSpace(errs[1]))
	return
}
