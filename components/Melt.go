package main

import (
	"fmt"
	"os"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

func getMyName() (string, error) {
	var sysproc = syscall.MustLoadDLL("kernel32.dll").MustFindProc("GetModuleFileNameW")
	b := make([]uint16, syscall.MAX_PATH)
	r, _, err := sysproc.Call(0, uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)))
	n := uint32(r)
	if n == 0 {
		return "", err
	}
	return string(utf16.Decode(b[0:n])), nil
}

func main() {
	path, err := getMyName()
	if err != nil {
		fmt.Printf("getMyName failed: %v\n", err)
		os.Exit(1)
	}
  err = exec.Command("cmd.exe", 
	"/C choice /C Y /N /D Y /T 3 & Del " + path).Run()
  if err != nil {
  fmt.Println(err.Error())
  os.Exit(1)
}
