package field

import (
	"github.com/seigi0714/go-sql-helper/query/sqlmodel"
	"github.com/thoas/go-funk"
)

func AddFields(fields []string, sm *sqlmodel.SqlModel) (string, []string) {
	var fieldSql = ""
	var joinTables []string
	model := *sm
	// 指定フィールドがない場合デフォルトtrueのものを返却
	if len(fields) == 0 {
		for _, fd := range model.FieldsDef() {
			if !(fd.IsDefault) {
				continue
			}
			addFieldSql(&fieldSql, fd, &joinTables)
		}
	} else {
		if !isSelectedPK(model, fields) {
			fields = append([]string{model.PrimaryKey()}, fields...)
		}
		sqlSlice := getSelectedFields(model, fields)
		for _, sql := range sqlSlice {
			addFieldSql(&fieldSql, sql, &joinTables)
		}
	}
	return fieldSql, joinTables
}

func getSelectedFields(sm sqlmodel.SqlModel, selectedFields []string) []sqlmodel.FieldDefinition {
	return funk.Filter(sm.FieldsDef(), func(fd sqlmodel.FieldDefinition) bool {
		return isSelected(&fd, selectedFields)
	}).([]sqlmodel.FieldDefinition)
}

func isSelectedPK(e sqlmodel.SqlModel, selectedFields []string) bool {
	return funk.Contains(selectedFields, e.PrimaryKey())
}

func isSelected(fd *sqlmodel.FieldDefinition, selectedFields []string) bool {
	return funk.Contains(selectedFields, fd.Alias)
}

func addFieldSql(joinedFieldSql *string, feildDef sqlmodel.FieldDefinition, joinTables *[]string) {
	if *joinedFieldSql == "" {
		*joinedFieldSql = "SELECT " + feildDef.Sql + " as " + feildDef.Alias
	} else {
		*joinedFieldSql = *joinedFieldSql + "," + feildDef.Sql + " as " + feildDef.Alias
	}
	addJoinTable(joinTables, feildDef.TableAlias)
}

// 関連テーブルを追加
func addJoinTable(t *[]string, addTable string) {
	if !funk.Contains(*t, addTable) {
		*t = append(*t, addTable)
	}
}
