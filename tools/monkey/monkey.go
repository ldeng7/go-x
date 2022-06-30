package monkey

import (
	"errors"
	"reflect"
	"unsafe"
)

type Patch struct {
	dst     uintptr
	dstCode []byte
	srcCode []byte
}

func NewPatch(dst, src any) (*Patch, error) {
	dv, sv := reflect.ValueOf(dst), reflect.ValueOf(src)
	if dv.Kind() != reflect.Func || sv.Kind() != reflect.Func {
		return nil, errors.New("invalid type")
	} else if dv.Type() != sv.Type() {
		return nil, errors.New("unequal type of functions")
	} else if dv.Pointer() == 0 || sv.Pointer() == 0 {
		return nil, errors.New("nil function")
	}

	srcCode := getJumpCode(sv.Pointer())
	l := len(srcCode)
	patch := &Patch{dv.Pointer(), make([]byte, l), srcCode}
	dstCode := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{patch.dst, l, l}))
	copy(patch.dstCode, dstCode)
	return patch, nil
}

func (p *Patch) Patch() error {
	return writeCode(p.dst, p.srcCode)
}

func (p *Patch) Unpatch() error {
	return writeCode(p.dst, p.dstCode)
}
