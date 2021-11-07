package sortby

import (
	"strings"
)

func GetSortSql(v []string) string {
	sql := ""
	for _, sort := range v {
		addSortSql(&sql, sort)
	}
	return sql
}

func addSortSql(sortSql *string, v string) {
	sortType := sortType(&v)
	if *sortSql == "" {
		*sortSql = " ORDER BY " + v + sortType
	} else {
		*sortSql = *sortSql + "," + v + sortType
	}
}

//先頭文字を見て降順、昇順を取得
func sortType(sortField *string) (sortType string) {
	if strings.HasPrefix(*sortField, "-") {
		sortType = " DESC"
		f := *sortField
		*sortField = f[1:]
	} else {
		sortType = " ASC"
	}
	return
}
