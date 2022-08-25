package query

import "strings"

func SearchQuery(conditions string) string {
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

func Query(cmd string) (query string) {

	return
}
