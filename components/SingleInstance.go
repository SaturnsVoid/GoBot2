package components

import (
	"syscall"
	"unsafe"
)

func createMutex(name string) (uintptr, error) {
	ret, _, err := procCreateMutex.Call(
		0,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
	)
	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}

func checkSingleInstance(name string) bool {
	_, err := createMutex(name)
	if err != nil {
		return true
	} else {
		return false
	}
}
