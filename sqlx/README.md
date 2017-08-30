Name
====

sqlx

A light-weighted sql client wrapper, containing a sql builder and a result set parser.

Synopsis
========

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ldeng7/go-x/sqlx"
)

type TestModel struct {
	// Setting field tag to specify the column name
	Id int `sql:"id"`

	// Using default column name: vc_str
	// Field type can also be string, but sql.NullString can indicate null value
	VcStr sql.NullString
}

func main() {
	db, _ := sql.Open("mysql", "user:password@/db")
	defer db.Close()
	{
		// Directly passing a sql, and returning a slice of map
		res, _ := sqlx.Query(db, "SELECT id, vc_str AS str FROM tests LIMIT 5;")

		for _, m := range res {
			for k, v := range m {
				fmt.Printf("%s: %s\n", k, v)
			}
		}
	}

	{
		// Returning a slice of struct
        models := make([]*TestModel, 0)
		_ := sqlx.QueryObj(db, &models, "SELECT id, vc_str FROM tests LIMIT 5;")

		for _, model := range models {
			fmt.Printf("id: %d\n", model.Id)
			str, _ := model.VcStr.Value()
			fmt.Printf("str: %s\n", str)
		}
	}

	// Returning a string like:
	// UPDATE tests SET col1 = '1', col2 = '2\n2' WHERE col1 = 3 AND col2 = '4\n4' AND col3 IN ('5', '6\n6');
	sqlx.Sql("UPDATE tests SET #{pairs} WHERE col1 = #{v1} AND col2 = #{v2} AND col3 IN (#{vs});",
		map[string]interface{}{
			"pairs": map[string]sqlx.Quote{"col1": "1", "col2": "2\n2"},
			"v1":    "3",
			"v2":    sqlx.Quote("4\n4"),
			"vs":    []sqlx.Quote{"5", "6\n6"},
		})

	// Returning a string like:
	// INSERT INTO tests (col1, col2) VALUES ('1', '2'), ('3', '4\n4');
	sqlx.Sql("INSERT INTO tests (#{columns}) VALUES #{values};",
		map[string]interface{}{
			"columns": []string{"col1, col2"},
			"values":  [][]sqlx.Quote{[]sqlx.Quote{"1", "2"}, []sqlx.Quote{"3", "4\n4"}},
		})
}
```
