package sortby

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sorts = []string{"id", "-name"}
var notSelect = []string{}

func TestAddField(t *testing.T) {
	ss := Get(sorts)

	expectedSql := "ORDER BY id ASC ,name DESC "
	fmt.Println("sort sql :: ", ss)
	assert.Equal(t, expectedSql, ss)

	ss = Get(notSelect)
	fmt.Println("not sort sql :: ", ss)
	assert.Equal(t, "", ss)
}
