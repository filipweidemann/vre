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

func (irc *IRClient) Init() error {
	fdptr, err := syscall.UTF16PtrFromString(irc.MmapFileLocation)
	if err != nil {
		panic(err)
	}

	accessMode := uint32(windows.PAGE_READONLY)

	fileHandle, err := syscall.CreateFileMapping(syscall.InvalidHandle, nil, accessMode, 0, uint32(vre.MMAP_FILESIZE), fdptr)
	defer syscall.CloseHandle(fileHandle)

	if err != nil {
		return err
	}

	addr, err := syscall.MapViewOfFile(fileHandle, syscall.FILE_MAP_READ, 0, 0, uintptr(vre.MMAP_FILESIZE))
	irc.MmapAddr = addr
	if err != nil {
		return err
	}

	irc.MmapPtr = unsafe.Pointer(addr)
	return nil
}

func (irc *IRClient) GetHeaderBytes() []byte {
	b := *(*[40]byte)(irc.MmapPtr)
	return b[:]
}

func (irc *IRClient) GetHeaderData() []uint32 {
	headerBytes := irc.GetHeaderBytes()
	headerFields := make([]uint32, 10)

	for i := 0; i < len(headerBytes); i += 4 {
		val := binary.LittleEndian.Uint32(headerBytes[i : i+4])
		headerFields = append(headerFields, val)
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
