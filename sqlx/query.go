package sqlx

import (
	"database/sql"
	"errors"
	"reflect"
)

func ParseRows(rows *sql.Rows) ([]map[string]interface{}, error) {
	cols, err := rows.Columns()
	if nil != err {
		return nil, err
	}
	nCol := len(cols)
	valsCol := make([]interface{}, nCol)
	ptrsCol := make([]interface{}, nCol)
	for i := 0; i < nCol; i++ {
		ptrsCol[i] = &valsCol[i]
	}

	out := make([]map[string]interface{}, 0)
	for rows.Next() {
		e := make(map[string]interface{})
		rows.Scan(ptrsCol...)
		for i := 0; i < nCol; i++ {
			if v := valsCol[i]; nil != v {
				e[cols[i]] = v
			}
		}
		out = append(out, e)
	}
	return out, nil
}

func ParseRowsObj(rows *sql.Rows, models interface{}) error {
	errTyp := errors.New("models must be a pointer to a slice of pointer of struct")
	typPtr := reflect.TypeOf(models)
	if reflect.Ptr != typPtr.Kind() {
		return errTyp
	}
	typSl := typPtr.Elem()
	if reflect.Slice != typSl.Kind() {
		return errTyp
	}
	typElem := typSl.Elem()
	if reflect.Ptr != typElem.Kind() {
		return errTyp
	}
	typModel := typElem.Elem()
	if reflect.Struct != typModel.Kind() {
		return errTyp
	}
	valPtr := reflect.ValueOf(models)
	if 0 == valPtr.Pointer() {
		return errors.New("models is an empty pointer")
	}

	nField := typModel.NumField()
	fieldMap := make(map[string]int)
	for i := 0; i < nField; i++ {
		field := typModel.Field(i)
		fieldMap[sqlName(&field)] = i
	}

	cols, err := rows.Columns()
	if nil != err {
		return err
	}
	nCol := len(cols)
	valsCol := make([]interface{}, nCol)
	ptrsCol := make([]interface{}, nCol)
	for i := 0; i < nCol; i++ {
		ptrsCol[i] = &valsCol[i]
	}

	valSl := reflect.MakeSlice(reflect.SliceOf(typElem), 0, 8)
	for rows.Next() {
		valElem := reflect.New(typModel)
		valModel := reflect.Indirect(valElem)
		rows.Scan(ptrsCol...)
		for i := 0; i < nCol; i++ {
			iField, need := fieldMap[cols[i]]
			if v := valsCol[i]; need && (nil != v) {
				field := valModel.Field(iField)
				setField(&field, v)
			}
		}
		valSl = reflect.Append(valSl, valElem)
	}
	valPtr.Elem().Set(valSl)
	return nil
}

func Query(db *sql.DB, sql string) ([]map[string]interface{}, error) {
	rows, err := db.Query(sql)
	if nil != err {
		return nil, err
	}
	defer rows.Close()
	return ParseRows(rows)
}

func QueryObj(db *sql.DB, models interface{}, sql string) error {
	rows, err := db.Query(sql)
	if nil != err {
		return err
	}
	defer rows.Close()
	return ParseRowsObj(rows, models)
}
