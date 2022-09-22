package query

import (
	"backend-nabati/domain/shared/constant"
	Error "backend-nabati/domain/shared/error"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func ValueStrBuilder(str string) (builder string) {
	str = strings.Trim(str, " ")
	builder = `'` + str + `'`
	return
}

func CastToNumber(value interface{}) (out int) {
	v := fmt.Sprintf("%v", value)
	out, err := strconv.Atoi(v)
	if err != nil {
		Error.New(constant.ErrGeneral, fmt.Sprintf("can't cast value: %v to decimal", value), err)
	}

	return
}

func CastToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func CastToDecimal(value interface{}) (out float64) {
	v := fmt.Sprintf("%v", value)
	out, err := strconv.ParseFloat(v, 64)
	if err != nil {
		Error.New(constant.ErrGeneral, fmt.Sprintf("can't cast value: %v to decimal", value), err)
	}

	return
}
