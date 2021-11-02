package fieldhelper

import (
	sqlmodel "github.com/seigi0714/go-sql-helper/pkg/sql_model"

	"github.com/thoas/go-funk"
)

func AddFields(fields []string, e sqlmodel.SqlModel) string {
	var fieldSql = ""

	// 指定フィールドがない場合デフォルトtrueのものを返却
	if len(fields) == 0 {
		for _, fd := range e.FieldsDef() {
			if !(fd.IsDefault) {
				continue
			}
			fieldSqlJoin(&fieldSql, fd)
		}
	} else {
		if !isSelectedPK(e, fields) {
			fields = append([]string{e.PrimaryKey()}, fields...)
		}
		sqlSlice := GetSelectedFieldsSql(e, fields)
		for _, sql := range sqlSlice {
			fieldSqlJoin(&fieldSql, sql)
		}
	}
	return fieldSql
}

func GetSelectedFieldsSql(bs sqlmodel.SqlModel, selectedFields []string) []sqlmodel.FieldDefinition {
	return funk.Filter(bs.FieldsDef(), func(fd sqlmodel.FieldDefinition) bool {
		return isSelected(&fd, selectedFields)
	}).([]sqlmodel.FieldDefinition)
}

func isSelectedPK(e sqlmodel.SqlModel, selectedFields []string) bool {
	return funk.Contains(selectedFields, e.PrimaryKey())
}

func isSelected(fd *sqlmodel.FieldDefinition, selectedFields []string) bool {
	return funk.Contains(selectedFields, fd.Alias)
}

func fieldSqlJoin(joinedFieldSql *string, joinFieldDef sqlmodel.FieldDefinition) {
	if *joinedFieldSql == "" {
		*joinedFieldSql = "SELECT " + joinFieldDef.Sql + " as " + joinFieldDef.Alias
	} else {
		*joinedFieldSql = *joinedFieldSql + "," + joinFieldDef.Sql + " as " + joinFieldDef.Alias
	}
}
