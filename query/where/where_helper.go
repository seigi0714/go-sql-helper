package where

import (
	"strings"

	query "github.com/seigi0714/go-sql-helper/query/query_builder"
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
	} else {
		o = " OR "
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
		return query.InvalidOperatorError(o)
	}
	return nil
}

func WhereIn(x string, in []string) string {
	inStr := strings.Join(in[:], ",")
	return andWherePrefix + x + " IN(" + inStr + ")"
}

func OrWhereIn(x string, in []string) string {
	inStr := strings.Join(in[:], ",")
	return orWherePrefix + x + " IN(" + inStr + ")"
}
