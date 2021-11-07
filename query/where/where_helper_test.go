package where

import (
	"fmt"
	"testing"

	query "github.com/seigi0714/go-sql-helper/query/query_builder"
	"github.com/stretchr/testify/assert"
)

func TestAddWhere(t *testing.T) {
	where1, err := Where("id", Equal, "1")
	where2, err2 := OrWhere("name", Equal, "test")
	wss := []string{where1, where2}
	ws := Get(wss)

	assert.Equal(t, " WHERE id=1 OR name=test", ws)
	assert.Nil(t, err)
	assert.Nil(t, err2)
}

func TestInvalidOperator(t *testing.T) {
	where1, err := Where("id", "==", "1")
	assert.Equal(t, where1, "")
	assert.Equal(t, err, query.InvalidOperatorError("=="))
}

func TestAddWhereIn(t *testing.T) {
	wss := []string{WhereIn("id", []string{"1", "3"})}
	ws := Get(wss)
	assert.Equal(t, " WHERE id IN(1,3)", ws)
	fmt.Println(ws)
}

func TestNotSpecified(t *testing.T) {
	notSpecified := []string{}
	ws := Get(notSpecified)
	assert.Equal(t, "", ws)
	fmt.Println(ws)
}
