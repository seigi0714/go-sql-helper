package where

import (
	"strings"

	query_error "github.com/seigi0714/go-sql-helper/query/myerror"
	"github.com/thoas/go-funk"
)

var comparisonOperator = []string{
	Equal,
	NotEqual,
	Greater,
	Small,
	GreaterEqual,
	SmallEqual,
}

const Equal = "="
const NotEqual = "!="
const Greater = ">"
const Small = "<"
const GreaterEqual = ">="
const SmallEqual = "<="

const andWherePrefix = "&&"
const orWherePrefix = "||"

func Get(wss []string) string {
	var whereSql string
	for _, ws := range wss {
		o := getOption(&ws)
		if whereSql == "" {
			whereSql = " WHERE " + ws
		} else {
			whereSql = whereSql + o + ws
		}
	}
	return whereSql
}

// 先頭二文字を見てANDかORを判別する
func getOption(ws *string) (o string) {
	s := *ws
	prefix := s[:2]
	where := s[2:]
	if prefix == andWherePrefix {
		o = " AND "
	} else if prefix == orWherePrefix {
		o = " OR "
	} else {
		// 指定しない場合,ANDで返却.
		// WHERE区のPrefixを排除しない.
		o = " AND "
		return
	}
	*ws = where
	return
}

func Where(x, o, y string) (string, error) {
	err := checkCO(o)
	if err != nil {
		return "", err
	}
	return andWherePrefix + x + o + y, nil
}

func OrWhere(x, o, y string) (string, error) {
	err := checkCO(o)
	if err != nil {
		return "", err
	}
	return orWherePrefix + x + "=" + y, nil
}

func checkCO(o string) error {
	if !funk.Contains(comparisonOperator, o) {
		return query_error.InvalidOperatorError(o)
	}
	return nil
}

func WhereNull(col string) string {
	return andWherePrefix + col + " IS NULL"
}

func WhereNot(col string) string {
	return andWherePrefix + col + " IS NOT NULL"
}

func OrWhereNull(col string) string {
	return orWherePrefix + col + " IS NULL"
}

func OrWhereNotNull(col string) string {
	return orWherePrefix + col + " IS NOT NULL"
}

func WhereIn(x string, in []string) string {
	inStr := strings.Join(in[:], ",")
	return andWherePrefix + x + " IN(" + inStr + ")"
}

func OrWhereIn(x string, in []string) string {
	inStr := strings.Join(in[:], ",")
	return orWherePrefix + x + " IN(" + inStr + ")"
}
