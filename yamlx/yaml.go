package yamlx

import (
	"errors"
	"reflect"

	yaml "gopkg.in/yaml.v2"
)

var typString = reflect.TypeOf("")

func ReadStructFromYamlMapByKey(bytes []byte, ps interface{}, key string) error {
	errTyp := errors.New("ps must be a pointer to a struct")
	typPointer := reflect.TypeOf(ps)
	if reflect.Ptr != typPointer.Kind() {
		return errTyp
	}
	if reflect.Struct != typPointer.Elem().Kind() {
		return errTyp
	}
	valPtr := reflect.ValueOf(ps)
	if 0 == valPtr.Pointer() {
		return errors.New("ps is an empty pointer")
	}
	typMap := reflect.MapOf(typString, typPointer)
	valMap := reflect.MakeMap(typMap)

	err := yaml.Unmarshal(bytes, valMap.Interface())
	if nil != err {
		return err
	}

	valElem := valMap.MapIndex(reflect.ValueOf(key))
	if reflect.Invalid == valElem.Kind() || 0 == valElem.Pointer() {
		return errors.New("key not found")
	}
	valPtr.Elem().Set(valElem.Elem())
	return nil
}
