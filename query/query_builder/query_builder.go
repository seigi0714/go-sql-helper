package query

import (
	"github.com/seigi0714/go-sql-helper/query/field"
	"github.com/seigi0714/go-sql-helper/query/join"
	"github.com/seigi0714/go-sql-helper/query/sortby"
	"github.com/seigi0714/go-sql-helper/query/sqlmodel"
	"github.com/thoas/go-funk"
)

type QueryBuilder struct {
	model      *sqlmodel.SqlModel
	fields     []string
	joinTables []string
	sort       []string
	where      []string
	// TODO: ページング実装
	// page       int
	// per_page   int
}

func New(sm *sqlmodel.SqlModel) *QueryBuilder {
	return &QueryBuilder{model: sm}
}

// 取得したいフィールドのエイリアスを引数に指定
// エイリアスはsqlmodelのFieldsDefに定義してあるものでなければならない
func (qb *QueryBuilder) AddFields(fs []string) {
	qb.fields = append(qb.fields, fs...)
}

// 取得したいフィールドのエイリアスのスライスを指定を引数に指定
// エイリアスはsqlmodelのJoinTablesDefに定義してあるものでなければならない
func (qb *QueryBuilder) AddJoinTable(jt string) {
	model := *qb.model
	if model.Table() != jt && !qb.isAlreadyJoin(jt) {
		qb.joinTables = append(qb.joinTables, jt)
	}
}

func (qb *QueryBuilder) isAlreadyJoin(addTable string) bool {
	return funk.Contains(qb.joinTables, addTable)
}

// ソートカラム順を指定
// {field名} : 昇順
// -{field名} : 降順
func (qb *QueryBuilder) Sort(sf []string) {
	qb.sort = append(qb.joinTables, sf...)
}

// クエリービルダーからSQLを返却
func (qb *QueryBuilder) ToSql() (string, error) {

	err := qb.isDefFields(qb.fields)
	if err != nil {
		return "", err
	}

	err = qb.isDefTables(qb.joinTables)
	if err != nil {
		return "", err
	}

	model := *qb.model
	from := " FROM " + model.Table()
	// セレクト句とそれらに関連したJOINテーブルを取得
	fs, jts := field.Get(qb.fields, qb.model)

	// JOIN句を取得
	qb.joinTables = append(qb.joinTables, jts...)
	js := join.Get(qb.joinTables, *qb.model)

	// ORDER BY句を取得
	ss := sortby.Get(qb.sort)
	return fs + from + js + ss, nil
}

// 選択したフィールドが定義されているかチェック
// 返却されるエラーはBuildingError code 101
func (qb *QueryBuilder) isDefFields(fs []string) error {
	for _, f := range fs {
		err := qb.isDefField(f)
		if err != nil {
			return err
		}
	}
	return nil
}

// 選択したフィールドが定義されているかチェック
// 返却されるエラーはBuildingError code 101
func (qb *QueryBuilder) isDefField(f string) (err error) {
	sm := *qb.model
	for _, fd := range sm.FieldsDef() {
		if fd.Alias != f {
			continue
		}
		err = nil
		return
	}
	err = NotDefFieldError(f)
	return
}

// Joinされるテーブルが定義されているかチェック
// 返却されるエラーはBuildingError code 102
func (qb *QueryBuilder) isDefTables(ts []string) error {
	for _, t := range ts {
		err := qb.isDefTable(t)
		if err != nil {
			return err
		}
	}
	return nil
}

// 選択したフィールドが定義されているかチェック
// 返却されるエラーはBuildingError code 101
func (qb *QueryBuilder) isDefTable(t string) (err error) {
	sm := *qb.model
	for _, jtd := range sm.JoinTablesDef() {
		if jtd.Alias != t {
			continue
		}
		err = nil
		return
	}
	err = NotDefJoinTableError(t)
	return
}
