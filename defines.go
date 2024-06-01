package vre

// IRSDK mmap
const (
	MMAP_FILESIZE       = 1191936
	IRSDK_MMAP_FILENAME = "Local\\IRSDKMemMapFileName"
)

// Header Fields
type IRHeaderField int

// //go:generate stringer -type=IRHeaderField
const (
	Version IRHeaderField = iota
	Status
	TickRate

	SessionInfoUpdate
	SessionInfoLen
	SessionInfoOffset

	NumVars
	HeaderOffset
	NumBuff
	BuffLen
)
