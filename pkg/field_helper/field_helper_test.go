package fieldhelper

import (
	"testing"

	sqlmodel "github.com/seigi0714/go-sql-helper/pkg/sql_model"
	"github.com/stretchr/testify/assert"
)

var fields = []string{"id", "name", "todo_id", "post_id"}
var notSelect = []string{}
var failedSelect = []string{"存在しないカラム"}

const defaultSql = "SELECT user.id as id,user.name as name,user.age as age"
const priSql = "SELECT user.id as id"

const todoJoinSql = "INNER JOIN todo ON todo.userId = user.id "
const postJoinSql = "INNER JOIN post ON post.userId = user.id "

func TestAddField(t *testing.T) {
	u := sqlmodel.NewUser()
	fs, js := AddFields(fields, u)

	expectedSql := "SELECT user.id as id,user.name as name,todo.id as todo_id,post.id as post_id"
	assert.Equal(t, expectedSql, fs)
	assert.Equal(t, todoJoinSql+postJoinSql, js)

	defaultFs, defaultJs := AddFields(notSelect, u)
	assert.Equal(t, defaultSql, defaultFs)
	assert.Equal(t, "", defaultJs)

	priFs, priJs := AddFields(failedSelect, u)
	assert.Equal(t, priSql, priFs)
	assert.Equal(t, "", priJs)
}
