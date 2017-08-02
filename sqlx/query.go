package sqlx

import (
	"database/sql"
	"errors"
	"reflect"
)

func ParseRows(rows *sql.Rows) ([]map[string]interface{}, error) {
	cols, err := rows.Columns()
	if nil == cols {
		return nil, err
	}
	nCol := len(cols)
	elems := make([]interface{}, nCol)
	elemPtrs := make([]interface{}, nCol)
	for i := 0; i < nCol; i++ {
		elemPtrs[i] = &elems[i]
	}

	out := make([]map[string]interface{}, 0)
	for rows.Next() {
		e := make(map[string]interface{})
		rows.Scan(elemPtrs...)
		for i := 0; i < nCol; i++ {
			if v := elems[i]; nil != v {
				e[cols[i]] = v
			}
		}
		out = append(out, e)
	}
	return out, nil
}

func ParseRowsObj(rows *sql.Rows, model interface{}) (interface{}, error) {
	typModel := reflect.TypeOf(model)
	if reflect.Struct != typModel.Kind() {
		return nil, errors.New("argument \"model\" must be a struct")
	}
	nField := typModel.NumField()
	fieldMap := make(map[string]int)
	for i := 0; i < nField; i++ {
		field := typModel.Field(i)
		fieldMap[sqlName(&field)] = i
	}

	cols, err := rows.Columns()
	if nil == cols {
		return nil, err
	}
	nCol := len(cols)
	elems := make([]interface{}, nCol)
	elemPtrs := make([]interface{}, nCol)
	for i := 0; i < nCol; i++ {
		elemPtrs[i] = &elems[i]
	}

	out := reflect.MakeSlice(reflect.SliceOf(typModel), 0, 8)
	for rows.Next() {
		e := reflect.Indirect(reflect.New(typModel))
		rows.Scan(elemPtrs...)
		for i := 0; i < nCol; i++ {
			iField, need := fieldMap[cols[i]]
			if v := elems[i]; need && (nil != v) {
				field := e.Field(iField)
				setField(&field, v)
			}
		}
		out = reflect.Append(out, e)
	}
	return out.Interface(), nil
}

func Query(db *sql.DB, sql string) ([]map[string]interface{}, error) {
	rows, err := db.Query(sql)
	if nil == rows {
		return nil, err
	}
	defer rows.Close()
	return ParseRows(rows)
}

func QueryObj(db *sql.DB, model interface{}, sql string) (interface{}, error) {
	rows, err := db.Query(sql)
	if nil == rows {
		return nil, err
	}
	defer rows.Close()
	return ParseRowsObj(rows, model)
}
