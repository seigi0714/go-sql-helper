package where

import (
	"testing"

	query_error "github.com/seigi0714/go-sql-helper/query/myerror"
	"github.com/stretchr/testify/assert"
)

func TestAddWhere(t *testing.T) {
	where1, err := Where("id", Equal, "1")
	where2, err2 := OrWhere("name", Equal, "test")
	wss := []string{where1, where2}
	ws := Get(wss)

	assert.Equal(t, " WHERE id='1' OR name='test'", ws)
	assert.Nil(t, err)
	assert.Nil(t, err2)
}

func TestInvalidOperator(t *testing.T) {
	where1, err := Where("id", "==", "'1'")
	assert.Equal(t, where1, "")
	assert.Equal(t, err, query_error.InvalidOperatorError("=="))
}

func TestAddWhereIn(t *testing.T) {
	wss := []string{WhereIn("id", []string{"Go言語", "PHP"})}
	ws := Get(wss)
	assert.Equal(t, " WHERE id IN('Go言語','PHP')", ws)
}

func TestNotSpecified(t *testing.T) {
	notSpecified := []string{}
	ws := Get(notSpecified)
	assert.Equal(t, "", ws)
}
