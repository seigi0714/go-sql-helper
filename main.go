package main

import (
	"fmt"

	fieldhelper "github.com/seigi0714/go-sql-helper/pkg/field_helper"
	sqlmodel "github.com/seigi0714/go-sql-helper/pkg/sql_model"
)

func main() {
	u := sqlmodel.NewUser()
	fields := []string{"id", "name", "todo_id"}
	fs, js := fieldhelper.AddFields(fields, u)
	fmt.Println("sql :: ", fs)
	fmt.Println("join sql :: ", js)
}
