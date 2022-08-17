package helper

import (
	"encoding/json"
	"strings"
)

func GetPaginations(count, limit, page int) (offset, totalPage int) {
	totalPage = ((count - 1) / limit) + 1
	offset = limit * (page - 1)
	return
}

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

func ValueQueryBuilder(str string) (builder string) {
	str = strings.Trim(str, " ")
	builder = `'` + str + `'`
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
