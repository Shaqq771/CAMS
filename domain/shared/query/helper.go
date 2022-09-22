package query

import (
	"reflect"
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
