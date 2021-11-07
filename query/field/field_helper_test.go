package field

import (
	"testing"

	"github.com/seigi0714/go-sql-helper/query/sqlmodel"
	"github.com/stretchr/testify/assert"
)

var fields = []string{"id", "name", "todo_id", "post_id"}
var notSelect = []string{}
var failedSelect = []string{"存在しないカラム"}

const defaultSql = "SELECT user.id as id,user.name as name,user.age as age"
const priSql = "SELECT user.id as id"

// const todoJoinSql = "INNER JOIN todo ON todo.userId = user.id "
// const postJoinSql = "INNER JOIN post ON post.userId = user.id "

func TestAddField(t *testing.T) {
	u := sqlmodel.NewUser()
	fs, jts := AddFields(fields, &u)

	expectedSql := "SELECT user.id as id,user.name as name,todo.id as todo_id,post.id as post_id"
	expectedJts := []string{"user", "todo", "post"}
	assert.Equal(t, expectedSql, fs)
	assert.Equal(t, expectedJts, jts)

	defaultFs, jts := AddFields(notSelect, &u)
	assert.Equal(t, defaultSql, defaultFs)
	assert.Equal(t, []string{"user"}, jts)

	priFs, _ := AddFields(failedSelect, &u)
	assert.Equal(t, priSql, priFs)
}
