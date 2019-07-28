package monkey

import (
	"reflect"
	"syscall"
	"unsafe"
)

var winProc = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")

func writeCode(p uintptr, code []byte) error {
	iLen, upLen := len(code), uintptr(len(code))
	sl := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{p, iLen, iLen}))
	var perms uint32
	pp := uintptr(unsafe.Pointer(&perms))
	if ret, _, _ := winProc.Call(p, upLen, 0x40, pp); 0 == ret {
		return syscall.GetLastError()
	}
	copy(sl, code)
	winProc.Call(p, upLen, uintptr(perms), pp)
	return nil
}
