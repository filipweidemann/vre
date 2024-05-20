package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	fdptr, err := syscall.UTF16PtrFromString("Local\\IRSDKMemMapFileName")
	if err != nil {
		panic(err)
	}

	accessMode := uint32(windows.PAGE_READONLY)
	mmapFileSize := 1164 * 1024

	fileHandle, err := syscall.CreateFileMapping(syscall.InvalidHandle, nil, accessMode, 0, uint32(mmapFileSize), fdptr)
	defer syscall.CloseHandle(fileHandle)

	if err != nil {
		panic(err)
	}

	addr, err := syscall.MapViewOfFile(fileHandle, syscall.FILE_MAP_READ, 0, 0, uintptr(mmapFileSize))
	if err != nil {
		panic(err)
	}

	ptr := unsafe.Pointer(addr)
	data := unsafe.Slice((*byte)(ptr), mmapFileSize)
	stringData := string(data[:])
	fmt.Printf("String Data: %v\n", stringData)
}
