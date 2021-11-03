package sorthepler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sorts = []string{"id", "-name"}
var notSelect = []string{}

func TestAddField(t *testing.T) {
	ss := GetSortSql(sorts)

	expectedSql := "ORDER BY id ASC ,name DESC "
	fmt.Println("sort sql :: ", ss)
	assert.Equal(t, expectedSql, ss)

	ss = GetSortSql(notSelect)
	fmt.Println("not sort sql :: ", ss)
	assert.Equal(t, "", ss)
}
