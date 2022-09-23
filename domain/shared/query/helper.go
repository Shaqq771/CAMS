package query

import (
	"backend-nabati/domain/shared/model"
	"reflect"
	"sort"
	"strings"
)

func GetFieldModel(data interface{}) (fields string) {
	var (
		model    = reflect.ValueOf(data)
		mapField []string
	)

	for i := 0; i < model.Type().NumField(); i++ {
		if model.Type().Field(i).Tag.Get("db") == "" {
			continue
		}

		mapField = append(mapField, model.Type().Field(i).Tag.Get("db"))
	}

	return strings.Join(mapField, ",")
}

func FindCounterField(filters []model.Fields) (duplicated map[string]int) {
	makeMap := make(map[string]int)

	for _, item := range filters {
		_, exist := makeMap[item.FieldName]
		if exist {
			makeMap[item.FieldName] += 1
		} else {
			makeMap[item.FieldName] = 1
		}
	}

	return makeMap
}

func SortByFieldName(fields []model.Fields) []model.Fields {
	sort.SliceStable(fields, func(i, j int) bool {
		return fields[i].FieldName < fields[j].FieldName
	})

	return fields
}
