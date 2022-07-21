package windef

import (
	os "os"
	unsafe "unsafe"
)

type DWORD = uint64
type BOOL = int32
type BYTE = uint8
type WORD = uint16
type FLOAT = float32
type PFLOAT = *float32
type INT = int32
type UINT = uint32
type PUINT = *uint32
type PBOOL = *int32
type LPBOOL = *int32
type PBYTE = *uint8
type LPBYTE = *uint8
type PINT = *int32
type LPINT = *int32
type PWORD = *uint16
type LPWORD = *uint16
type LPLONG = *int64
type PDWORD = *uint64  //bug
type LPDWORD = *uint64 //bug
type LPVOID = unsafe.Pointer
type LPCVOID = unsafe.Pointer
type ULONG = uint64
type PULONG = *uint64
type USHORT = uint16
type PUSHORT = *uint16
type UCHAR = uint8
type PUCHAR = *uint8
type CHAR = int8
type SHORT = int16
type LONG = int64
type struct__SENDCMDINPARAMS struct {
	cBufferSize  uint64 //bug
	bDriveNumber uint8
	bReserved    [3]uint8
	dwReserved   [4]uint64 //bug
	bBuffer      [1]uint8
}
type SENDCMDINPARAMS = struct__SENDCMDINPARAMS
type PSENDCMDINPARAMS = *struct__SENDCMDINPARAMS
type LPSENDCMDINPARAMS = *struct__SENDCMDINPARAMS

func _cgo_main() int32 {
	return int32(0)
}
func main() {
	os.Exit(int(_cgo_main()))
}
