package query

import (
	"testing"

	"github.com/seigi0714/go-sql-helper/query/sqlmodel"
	"github.com/stretchr/testify/assert"
)

const from = " FROM user"

var fields = []string{"id", "name", "todo_id", "post_id"}
var notSelect = []string{}
var notDefField = []string{"display_name"}

var sort = []string{"-id", "name"}

const defaultSql = "SELECT user.id as id,user.name as name,user.age as age"

func TestSQL(t *testing.T) {
	u := sqlmodel.NewUser()
	qb := New(&u)
	qb.AddFields(fields)
	sql, err := qb.ToSql()
	fs := "SELECT user.id as id,user.name as name,todo.id as todo_id,post.id as post_id"
	js := " INNER JOIN todo ON todo.user_id = user.id INNER JOIN post ON post.user_id = user.id"
	assert.Equal(t, fs+from+js, sql)
	assert.Nil(t, err)
}

func TestNotSelectSQL(t *testing.T) {
	u := sqlmodel.NewUser()
	qb := New(&u)
	qb.AddFields(notSelect)
	qb.Sort(sort)
	order_by := " ORDER BY id DESC,name ASC"
	sql, err := qb.ToSql()
	assert.Equal(t, defaultSql+from+order_by, sql)
	assert.Nil(t, err)
}

func TestNotDefFields(t *testing.T) {
	u := sqlmodel.NewUser()
	qb := New(&u)
	qb.AddFields(notDefField)

	notDefFieldErr := NotDefFieldError("display_name")
	sql, err := qb.ToSql()
	assert.Equal(t, "", sql)
	assert.Equal(t, err, notDefFieldErr, "定義していないフィールドを指定していた場合エラー")
}

func TestNotDefTables(t *testing.T) {
	u := sqlmodel.NewUser()
	qb := New(&u)
	qb.AddFields(notSelect)
	qb.AddJoinTable("memo")

	notDefTableErr := NotDefJoinTableError("memo")
	sql, err := qb.ToSql()
	assert.Equal(t, "", sql)
	assert.Equal(t, err, notDefTableErr, "定義していないテーブルをJOINしようとした場合エラー")
}
