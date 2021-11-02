package main

import (
	"fmt"

	fieldhelper "github.com/seigi0714/go-sql-helper/pkg/field_helper"
	sqlmodel "github.com/seigi0714/go-sql-helper/pkg/sql_model"
)

func main() {
	u := sqlmodel.NewUser()
	fields := []string{"id", "name"}
	fs := fieldhelper.AddFields(fields, u)
	fmt.Println("sql :: ", fs)
}
