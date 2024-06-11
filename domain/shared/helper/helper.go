package helper

import (
	"backend-nabati/infrastructure/shared/constant"
	"encoding/json"
	"strconv"
	"strings"

	Error "backend-nabati/domain/shared/error"
)

func StringToArrString(str string) (arrString []string, err error) {
	if err = json.Unmarshal([]byte(str), &arrString); err != nil {
		return
	}

	return
}

func ArrStringToString(arrString []string, sep string) (str string) {
	str = strings.Join(arrString, sep)
	str = strings.Trim(str, " ")
	return
}

func SortBy(sortby string) (sort string, sortList []string, err error) {
	sortList, err = StringToArrString(sortby)
	if err != nil {
		return
	}

	sort = ArrStringToString(sortList, ",")
	return
}

func FilterBy(filterBy string) (filter string, filterList []string, err error) {
	filterList, err = StringToArrString(filterBy)
	if err != nil {
		return
	}
	filter = ArrStringToString(filterList, "|")
	return
}

func LastDocNumber(number, start, end, skip string) (lastNumber int, err error) {
	var (
		intNumber, intStart, intEnd, intSkip int
	)

	intStart, err = strconv.Atoi(start)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrConvertStringToInt, err)
		return
	}

	intEnd, err = strconv.Atoi(end)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrConvertStringToInt, err)
		return
	}

	intSkip, err = strconv.Atoi(skip)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrConvertStringToInt, err)
		return
	}

	if strings.TrimSpace(number) == "" || number == "0" {
		intNumber = intStart
		return intNumber, nil
	} else {
		intNumber, err = strconv.Atoi(number)
		if err != nil {
			err = Error.New(constant.ErrGeneral, constant.ErrConvertStringToInt, err)
			return
		}
	}
	if intNumber < intStart {
		intNumber = intStart
		return intNumber, nil
	}

	if intNumber > intEnd {
		return
	}

	lastNumber = intNumber + 1 + intSkip

	return
}
