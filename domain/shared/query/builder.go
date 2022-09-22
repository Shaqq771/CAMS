package query

import (
	"backend-nabati/domain/shared/constant"
	"backend-nabati/domain/shared/model"
	"fmt"
	"strings"
)

func SearchQueryBuilder(conditions string) string {
	var (
		tempConditions []string
	)

	conds := strings.Split(conditions, "|")
	for _, cond := range conds {
		temps := strings.Split(cond, "=")
		for j, temp := range temps {
			key := strings.ReplaceAll(temps[0], " ", "")
			if j == 1 {
				temp := strings.ReplaceAll(temp, " ", "")
				arrs := strings.Split(temp, ",")
				if len(arrs) > 1 {
					tempConditions = append(tempConditions, key+" IN ('"+strings.Join(arrs, "','")+"')")
				} else {
					tempConditions = append(tempConditions, key+" = '"+temp+"'")
				}
			}
		}
	}

	return " AND " + strings.Join(tempConditions, " AND ")
}

func ConditionsBuilder(condition *model.Filter) (query string) {
	conditionLen := len(condition.Filters)
	for i, field := range condition.Filters {
		if field.Option == constant.EQUAL {
			query += FieldValueBuilder(field, "=", i, conditionLen)
		} else if field.Option == constant.BETWEEN {
			query += FieldValueBuilder(field, "BETWEEN", i, conditionLen)
		} else if field.Option == constant.NOT_BETWEEN {
			query += FieldValueBuilder(field, "NOT BETWEEN", i, conditionLen)
		} else if field.Option == constant.LESS_THAN {
			query += FieldValueBuilder(field, "<", i, conditionLen)
		} else if field.Option == constant.GREATER_THAN {
			query += FieldValueBuilder(field, ">", i, conditionLen)
		} else if field.Option == constant.LESS_THAN_EQUAL {
			query += FieldValueBuilder(field, "<=", i, conditionLen)
		} else if field.Option == constant.GREATER_THAN_EQUAL {
			query += FieldValueBuilder(field, ">=", i, conditionLen)
		} else if field.Option == constant.NOT_EQUAL {
			query += FieldValueBuilder(field, "<>", i, conditionLen)
		} else if field.Option == constant.CONTAINS_PATTERN {
			query += FieldValueBuilder(field, "LIKE", i, conditionLen)
		} else if field.Option == constant.CONTAINS_NO_PATTERN {
			query += FieldValueBuilder(field, "NOT LIKE", i, conditionLen)
		}
	}

	return
}

func FieldValueBuilder(field model.Fields, operations string, idx, fieldsLen int) (query string) {
	if field.ToValue != nil {
		if field.DataType == constant.NUMBER {
			query += fmt.Sprintf("(%s %s %s AND %s)", field.FieldName, operations, CastToString(field.FromValue), CastToString(field.ToValue))
		} else if field.DataType == constant.STRING || field.DataType == constant.TIME || field.DataType == constant.DATE || field.DataType == constant.DECIMAL {
			query += fmt.Sprintf("(%s %s %s AND %s)", field.FieldName, operations, ValueStrBuilder(CastToString(field.FromValue)), ValueStrBuilder(CastToString(field.ToValue)))
		}
	} else {
		if field.DataType == constant.NUMBER {
			query += fmt.Sprintf("%s %s %s", field.FieldName, operations, CastToString(field.FromValue))
		} else if field.DataType == constant.STRING || field.DataType == constant.TIME || field.DataType == constant.DATE || field.DataType == constant.DECIMAL {
			query += fmt.Sprintf("%s %s %s", field.FieldName, operations, ValueStrBuilder(CastToString(field.FromValue)))
		}
	}

	if idx < fieldsLen-1 {
		query += " AND "
	} else {
		query += " "
	}

	return
}
