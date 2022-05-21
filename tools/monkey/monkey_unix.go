//+build !windows

package monkey

import (
	"reflect"
	"syscall"
	"unsafe"
)

func getTargetPages(p, l uintptr) []byte {
	sz := uintptr(syscall.Getpagesize())
	pb, pe := p-p%sz, p+l-1
	pe += sz - pe%sz
	pl := int(pe - pb)
	sl := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{pb, pl, pl}))
	return sl
}

func writeCode(p uintptr, code []byte) error {
	iLen, upLen := len(code), uintptr(len(code))
	sl := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{p, iLen, iLen}))
	pages := getTargetPages(p, upLen)
	prot := syscall.PROT_READ | syscall.PROT_WRITE | syscall.PROT_EXEC
	if err := syscall.Mprotect(pages, prot); nil != err {
		return err
	}
	copy(sl, code)
	syscall.Mprotect(pages, syscall.PROT_READ|syscall.PROT_EXEC)
	return nil
}
