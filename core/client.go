package vre

import (
	"encoding/binary"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
	vre "virtualrace.engineer"
)

type IRClient struct {
	MmapFileLocation string
	MmapAddr         uintptr
	MmapPtr          unsafe.Pointer
}

func (irc *IRClient) Init() (func(), error) {
	fdptr, err := syscall.UTF16PtrFromString(irc.MmapFileLocation)
	if err != nil {
		panic(err)
	}

	accessMode := uint32(windows.PAGE_READONLY)

	fileHandle, err := syscall.CreateFileMapping(syscall.InvalidHandle, nil, accessMode, 0, uint32(vre.MMAP_FILESIZE), fdptr)
	defer syscall.CloseHandle(fileHandle)

	if err != nil {
		return nil, err
	}

	addr, err := syscall.MapViewOfFile(fileHandle, syscall.FILE_MAP_READ, 0, 0, uintptr(vre.MMAP_FILESIZE))
	irc.MmapAddr = addr
	if err != nil {
		return nil, err
	}

	irc.MmapPtr = unsafe.Pointer(addr)
	return irc.Close, nil
}

func (irc *IRClient) GetHeaderBytes() []byte {
	b := *(*[40]byte)(irc.MmapPtr)
	return b[:]
}

func (irc *IRClient) GetHeaderData() []int {
	headerBytes := irc.GetHeaderBytes()
	headerFields := make([]int, 10)

	for i := 0; i < 10; i++ {
		val := int(binary.LittleEndian.Uint32(headerBytes[i*4 : i*4+4]))
		headerFields[i%10] = val
	}

	return headerFields
}

func (irc *IRClient) Close() {
	err := syscall.UnmapViewOfFile(irc.MmapAddr)
	if err != nil {
		panic(err)
	}
}

func NewIRClient(mmapFileLocation string) *IRClient {
	return &IRClient{
		MmapFileLocation: mmapFileLocation,
	}
}
