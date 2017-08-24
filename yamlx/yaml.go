package yamlx

import (
	"errors"
	"reflect"

	yaml "gopkg.in/yaml.v2"
)

func ReadStructFromYamlMapByKey(bytes []byte, ps interface{}, key string) error {
	typErr := errors.New("ps must be a pointer to a struct")
	pt := reflect.TypeOf(ps)
	if reflect.Ptr != pt.Kind() {
		return typErr
	}
	if reflect.Struct != pt.Elem().Kind() {
		return typErr
	}
	pv := reflect.ValueOf(ps)
	if 0 == pv.Pointer() {
		return errors.New("ps is an empty pointer")
	}
	mt := reflect.MapOf(reflect.TypeOf(""), pt)
	mv := reflect.MakeMap(mt)

	err := yaml.Unmarshal(bytes, mv.Interface())
	if nil != err {
		return err
	}

	mev := mv.MapIndex(reflect.ValueOf(key))
	if reflect.Invalid == mev.Kind() || 0 == mev.Pointer() {
		return errors.New("key not found")
	}
	pv.Elem().Set(mev.Elem())
	return nil
}
