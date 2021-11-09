package main

import (
	"fmt"

	query "github.com/seigi0714/go-sql-helper/query/query_builder"
	"github.com/seigi0714/go-sql-helper/query/sqlmodel"
	"github.com/seigi0714/go-sql-helper/query/where"
)

func main() {
	u := sqlmodel.NewUser()
	q := query.New(&u)

	q.AddFields([]string{})
	q.Sort([]string{"id"})
	scopeId, _ := where.Where("id", "=", "1")
	q.AddWhere(scopeId)
	sql, err := q.ToSql()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(sql)
}
