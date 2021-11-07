package main

import (
	"fmt"

	query "github.com/seigi0714/go-sql-helper/query/query_builder"
	"github.com/seigi0714/go-sql-helper/query/sqlmodel"
)

func main() {
	u := sqlmodel.NewUser()
	q := query.New(&u)

	q.AddFields([]string{})
	q.Sort([]string{"id"})
	sql, err := q.ToSql()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(sql)
}
