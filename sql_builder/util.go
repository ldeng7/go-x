package sql_builder

import (
	"database/sql"
	"reflect"
	"strconv"
	"unicode"
)

var f2s map[string]string = make(map[string]string)

func sqlName(f *reflect.StructField) string {
	out := f.Tag.Get("sql")
	if len(out) > 0 {
		return out
	}

	fn := f.Name
	out, found := f2s[fn]
	if found {
		return out
	}

	runes := []rune(fn)
	runesOut := make([]rune, len(runes)*2)
	runesOut[0] = unicode.ToLower(runes[0])
	runes = runes[1:]
	i := 1
	for _, r := range runes {
		if unicode.IsUpper(r) {
			runesOut[i] = '_'
			runesOut[i+1] = unicode.ToLower(r)
			i += 2
		} else {
			runesOut[i] = r
			i++
		}
	}

	out = string(runesOut[:i])
	f2s[fn] = out
	return out
}

func setIntField(field *reflect.Value, v interface{}) {
	switch vv := v.(type) {
	case int:
		field.SetInt(int64(vv))
	case uint:
		field.SetInt(int64(vv))
	case int8:
		field.SetInt(int64(vv))
	case uint8:
		field.SetInt(int64(vv))
	case int16:
		field.SetInt(int64(vv))
	case uint16:
		field.SetInt(int64(vv))
	case int32:
		field.SetInt(int64(vv))
	case uint32:
		field.SetInt(int64(vv))
	case int64:
		field.SetInt(vv)
	case uint64:
		field.SetInt(int64(vv))
	case float32:
		field.SetInt(int64(vv))
	case float64:
		field.SetInt(int64(vv))
	case []byte:
		i, _ := strconv.ParseInt(string(vv), 10, 64)
		field.SetInt(i)
	case string:
		i, _ := strconv.ParseInt(vv, 10, 64)
		field.SetInt(i)
	}
}

func setUintField(field *reflect.Value, v interface{}) {
	switch vv := v.(type) {
	case int:
		field.SetUint(uint64(vv))
	case uint:
		field.SetUint(uint64(vv))
	case int8:
		field.SetUint(uint64(vv))
	case uint8:
		field.SetUint(uint64(vv))
	case int16:
		field.SetUint(uint64(vv))
	case uint16:
		field.SetUint(uint64(vv))
	case int32:
		field.SetUint(uint64(vv))
	case uint32:
		field.SetUint(uint64(vv))
	case int64:
		field.SetUint(uint64(vv))
	case uint64:
		field.SetUint(vv)
	case float32:
		field.SetUint(uint64(vv))
	case float64:
		field.SetUint(uint64(vv))
	case []byte:
		i, _ := strconv.ParseUint(string(vv), 10, 64)
		field.SetUint(i)
	case string:
		i, _ := strconv.ParseUint(vv, 10, 64)
		field.SetUint(i)
	}
}

func setFloatField(field *reflect.Value, v interface{}) {
	switch vv := v.(type) {
	case int:
		field.SetFloat(float64(vv))
	case uint:
		field.SetFloat(float64(vv))
	case int8:
		field.SetFloat(float64(vv))
	case uint8:
		field.SetFloat(float64(vv))
	case int16:
		field.SetFloat(float64(vv))
	case uint16:
		field.SetFloat(float64(vv))
	case int32:
		field.SetFloat(float64(vv))
	case uint32:
		field.SetFloat(float64(vv))
	case int64:
		field.SetFloat(float64(vv))
	case uint64:
		field.SetFloat(float64(vv))
	case float32:
		field.SetFloat(float64(vv))
	case float64:
		field.SetFloat(vv)
	case []byte:
		f, _ := strconv.ParseFloat(string(vv), 64)
		field.SetFloat(f)
	case string:
		f, _ := strconv.ParseFloat(vv, 64)
		field.SetFloat(f)
	}
}

func setStringField(field *reflect.Value, v interface{}) {
	switch vv := v.(type) {
	case int:
		field.SetString(strconv.FormatInt(int64(vv), 10))
	case uint:
		field.SetString(strconv.FormatUint(uint64(vv), 10))
	case int8:
		field.SetString(strconv.FormatInt(int64(vv), 10))
	case uint8:
		field.SetString(strconv.FormatUint(uint64(vv), 10))
	case int16:
		field.SetString(strconv.FormatInt(int64(vv), 10))
	case uint16:
		field.SetString(strconv.FormatUint(uint64(vv), 10))
	case int32:
		field.SetString(strconv.FormatInt(int64(vv), 10))
	case uint32:
		field.SetString(strconv.FormatUint(uint64(vv), 10))
	case int64:
		field.SetString(strconv.FormatInt(int64(vv), 10))
	case uint64:
		field.SetString(strconv.FormatUint(vv, 10))
	case float32:
		field.SetString(strconv.FormatFloat(float64(vv), 'f', 20, 64))
	case float64:
		field.SetString(strconv.FormatFloat(vv, 'f', 20, 64))
	case []byte:
		field.SetString(string(vv))
	case string:
		field.SetString(vv)
	}
}
func setField(field *reflect.Value, v interface{}) {
	switch field.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		setIntField(field, v)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		setUintField(field, v)

	case reflect.Float32, reflect.Float64:
		setFloatField(field, v)

	case reflect.String:
		setStringField(field, v)

	case reflect.Struct:
		vField, eField := field.Field(0), field.Field(1)
		switch reflect.Zero(field.Type()).Interface().(type) {

		case sql.NullInt64:
			eField.SetBool(true)
			setIntField(&vField, v)

		case sql.NullFloat64:
			eField.SetBool(true)
			setFloatField(&vField, v)

		case sql.NullString:
			eField.SetBool(true)
			setStringField(&vField, v)
		}
	}
}
