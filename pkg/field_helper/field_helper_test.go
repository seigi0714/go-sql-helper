package fieldhelper

import (
	"testing"

	sqlmodel "github.com/seigi0714/go-sql-helper/pkg/sql_model"
	"github.com/stretchr/testify/assert"
)

var fields = []string{"id", "name"}
var notSelect = []string{}
var failedSelect = []string{"存在しないカラム"}

const defaultSql = "SELECT user.id as id,user.name as name,user.age as age"
const priSql = "SELECT user.id as id"

func TestAddField(t *testing.T) {
	u := sqlmodel.NewUser()
	fs := AddFields(fields, u)

	expectedSql := "SELECT user.id as id,user.name as name"
	assert.Equal(t, expectedSql, fs)

	defaultFs := AddFields(notSelect, u)
	assert.Equal(t, defaultSql, defaultFs)

	priFs := AddFields(failedSelect, u)
	assert.Equal(t, priSql, priFs)
}
