package monkey

import (
	"errors"
	"reflect"
	"unsafe"
)

type Patch struct {
	tar     uintptr
	tarCode []byte
	repCode []byte
}

func NewPatch(tar, rep any) (*Patch, error) {
	tv, rv := reflect.ValueOf(tar), reflect.ValueOf(rep)
	if tv.Kind() != reflect.Func || rv.Kind() != reflect.Func {
		return nil, errors.New("invalid type")
	} else if tv.Type() != rv.Type() {
		return nil, errors.New("unequal type of functions")
	} else if 0 == tv.Pointer() || 0 == rv.Pointer() {
		return nil, errors.New("nil function")
	}

	// fetch rv.ptr, as reflect.unpackEface()
	rp := uintptr(unsafe.Pointer(&rep)) + unsafe.Sizeof(uintptr(0))
	code := getJumpCode(*(*uintptr)(unsafe.Pointer(rp)))
	l := len(code)

	patch := &Patch{tv.Pointer(), make([]byte, l), code}
	tarCode := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{patch.tar, l, l}))
	copy(patch.tarCode, tarCode)
	return patch, nil
}

func (p *Patch) Patch() error {
	return writeCode(p.tar, p.repCode)
}

func (p *Patch) Unpatch() error {
	return writeCode(p.tar, p.tarCode)
}
