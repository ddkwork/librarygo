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
type PDWORD = *uint64
type LPDWORD = *uint64
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
type _cgoa_1_windef struct {
	Xbf_0 uint16
}
type _cgoa_2_windef struct {
	Xbf_0 uint16
}
type _cgoa_3_windef struct {
	Xbf_0 uint16
}
type _cgoa_4_windef struct {
	Xbf_0 uint16
}
type _cgoa_5_windef struct {
	Xbf_0 uint16
}
type _cgoa_6_windef struct {
	Xbf_0 uint16
}
type _cgoa_7_windef struct {
	Xbf_0 uint16
}
type struct__IDINFO struct {
	wGenConfig             uint16
	wNumCyls               uint16
	wReserved2             uint16
	wNumHeads              uint16
	wReserved4             uint16
	wReserved5             uint16
	wNumSectorsPerTrack    uint16
	wVendorUnique          [3]uint16
	sSerialNumber          [20]int8
	wBufferType            uint16
	wBufferSize            uint16
	wECCSize               uint16
	sFirmwareRev           [8]int8
	sModelNumber           [40]int8
	wMoreVendorUnique      uint16
	wReserved48            uint16
	wCapabilities          _cgoa_1_windef
	wReserved1             uint16
	wPIOTiming             uint16
	wDMATiming             uint16
	wFieldValidity         _cgoa_2_windef
	wNumCurCyls            uint16
	wNumCurHeads           uint16
	wNumCurSectorsPerTrack uint16
	wCurSectorsLow         uint16
	wCurSectorsHigh        uint16
	wMultSectorStuff       _cgoa_3_windef
	dwTotalSectors         uint64
	wSingleWordDMA         uint16
	wMultiWordDMA          _cgoa_4_windef
	wPIOCapacity           _cgoa_5_windef
	wMinMultiWordDMACycle  uint16
	wRecMultiWordDMACycle  uint16
	wMinPIONoFlowCycle     uint16
	wMinPOIFlowCycle       uint16
	wReserved69            [11]uint16
	wMajorVersion          _cgoa_6_windef
	wMinorVersion          uint16
	wReserved82            [6]uint16
	wUltraDMA              _cgoa_7_windef
	wReserved89            [167]uint16
}
type IDINFO = struct__IDINFO
type PIDINFO = *struct__IDINFO
type struct__DRIVERSTATUS struct {
	bDriverError uint8
	bIDEError    uint8
	bReserved    [2]uint8
	dwReserved   [2]uint64
}
type DRIVERSTATUS = struct__DRIVERSTATUS
type PDRIVERSTATUS = *struct__DRIVERSTATUS
type LPDRIVERSTATUS = *struct__DRIVERSTATUS
type struct__SENDCMDOUTPARAMS struct {
	cBufferSize  uint64
	DriverStatus struct__DRIVERSTATUS
	bBuffer      [1]uint8
}
type SENDCMDOUTPARAMS = struct__SENDCMDOUTPARAMS
type PSENDCMDOUTPARAMS = *struct__SENDCMDOUTPARAMS
type LPSENDCMDOUTPARAMS = *struct__SENDCMDOUTPARAMS
type struct__GETVERSIONINPARAMS struct {
	bVersion      uint8
	bRevision     uint8
	bReserved     uint8
	bIDEDeviceMap uint8
	fCapabilities uint64
	dwReserved    [4]uint64
}
type GETVERSIONINPARAMS = struct__GETVERSIONINPARAMS
type PGETVERSIONINPARAMS = *struct__GETVERSIONINPARAMS
type LPGETVERSIONINPARAMS = *struct__GETVERSIONINPARAMS
type struct__IDEREGS struct {
	bFeaturesReg     uint8
	bSectorCountReg  uint8
	bSectorNumberReg uint8
	bCylLowReg       uint8
	bCylHighReg      uint8
	bDriveHeadReg    uint8
	bCommandReg      uint8
	bReserved        uint8
}
type IDEREGS = struct__IDEREGS
type PIDEREGS = *struct__IDEREGS
type LPIDEREGS = *struct__IDEREGS
type struct__SENDCMDINPARAMS struct {
	cBufferSize  uint64
	irDriveRegs  struct__IDEREGS
	bDriveNumber uint8
	bReserved    [3]uint8
	dwReserved   [4]uint64
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
