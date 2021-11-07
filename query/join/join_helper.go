package join

import (
	"github.com/seigi0714/go-sql-helper/query/sqlmodel"
	"github.com/thoas/go-funk"
)

func AddJoinTablesSql(joinTables []string, e sqlmodel.SqlModel) string {
	joinSql := ""

	for _, def := range selectedTables(e.JoinTablesDef(), joinTables) {
		addJoinSql(&joinSql, def)
	}
	return joinSql
}

func selectedTables(joinDef []sqlmodel.JoinDefinition, joinTables []string) []sqlmodel.JoinDefinition {
	return funk.Filter(joinDef, func(t sqlmodel.JoinDefinition) bool {
		return isSelected(&t, joinTables)
	}).([]sqlmodel.JoinDefinition)
}

func addJoinSql(joinSql *string, def sqlmodel.JoinDefinition) {
	*joinSql = " " + *joinSql + def.Sql
}

func isSelected(jd *sqlmodel.JoinDefinition, joinTables []string) bool {
	return funk.Contains(joinTables, jd.Alias)
}
