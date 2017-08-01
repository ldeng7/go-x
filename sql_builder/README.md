Name
====

sql_builder

A light-weighted sql client wrapper, containing a sql builder and a result set parser.

Synopsis
========

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ldeng7/go-x/sql_builder"
)

type TestModel struct {
	Id    int    `sql:"id"`
	VcStr string // using default column name: vc_str
}

func main() {
	db, _ := sql.Open("mysql", "root:abcabc@/test")
	defer db.Close()
	{
		// Directly passing a sql, and returning a slice of map
		res, _ := sql_builder.Query(db, "SELECT id, vc_str AS str FROM tests LIMIT 5;")
		for _, m := range res {
			for k, v := range m {
				fmt.Printf("%s: %s\n", k, v)
			}
		}
	}
	{
		// Building a sql, passing a struct as the schema, and returning a slice of struct
		ids := []string{"1", "2"}
		sqlStr := sql_builder.Sql("SELECT id, vc_str FROM tests WHERE id IN (#{ids});",
			map[string]sql_builder.Arg{
				"ids": sql_builder.Arg{Type: sql_builder.ArgTypeStringArray, Value: ids},
			})
		res, _ := sql_builder.QueryObj(db, TestModel{}, sqlStr)
		models := res.([]TestModel)
		for _, model := range models {
			fmt.Printf("id: %d\n", model.Id)
			fmt.Printf("id: %s\n", model.VcStr)
		}
	}
	db.Close()
}
```
