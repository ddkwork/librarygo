//go:build windows
// +build windows

package hardwareIndo

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"strconv"
	"syscall"
	"unsafe"
)

type ssdInfo struct {
	SerialNumber string
	ModelNumber  string
	Version      string
}

func (s *ssdInfo) Get() (ok bool) {
	path := fmt.Sprintf("\\\\.\\PhysicalDrive%d", 0)
	fromString, err := syscall.UTF16PtrFromString(path)
	if !mycheck.Error(err) {
		return
	}
	fd, err := syscall.CreateFile(
		fromString,
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if !mycheck.Error(err) {
		return
	}

	out := make([]byte, IDENTIFY_BUFFER_SIZE)
	var bytesReturned uint32
	if !mycheck.Error(syscall.DeviceIoControl(
		fd,
		SMART_GET_VERSION,
		nil,
		0,
		&out[0],
		IDENTIFY_BUFFER_SIZE,
		&bytesReturned,
		nil,
	)) {
		return
	}
	//getVersionInParam := (*getVersionInParams)(unsafe.Pointer(&out[0]))

	var BytesReturned uint32
	out = make([]byte, IDENTIFY_BUFFER_SIZE)
	inParams := &sendCmdInParams{
		cBufferSize: 0,
		regs: ideRegS{
			bFeaturesReg:     0,
			bSectorCountReg:  1,
			bSectorNumberReg: 1,
			bCylLowReg:       0,
			bCylHighReg:      0,
			bDriveHeadReg:    SMART_CMD,
			bCommandReg:      ID_CMD,
			bReserved:        0,
		},
		bDriveNumber: 0, //todo
		bReserved:    [3]byte{},
		dwReserved:   [4]byte{},
		bBuffer:      [1]byte{},
	}
	if !mycheck.Error(syscall.DeviceIoControl(
		fd,
		SMART_RCV_DRIVE_DATA,
		(*byte)(unsafe.Pointer(inParams)),
		IDENTIFY_BUFFER_SIZE*2,
		&out[0],
		IDENTIFY_BUFFER_SIZE*2,
		&BytesReturned,
		nil,
	)) {
		return
	}
	outParams := (*sendCmdOutParams)(unsafe.Pointer(&out[0]))
	b := outParams.bBuffer[8:][:] //todo bug
	info := (*idSector)(unsafe.Pointer(&b[0]))

	sSerialNumber := stream.NewBytes(info.sSerialNumber[:])
	serialNumber := tool.New().Swap().SerialNumber(sSerialNumber.String())

	sModelNumber := stream.NewBytes(info.sModelNumber[:])
	ModelNumber := tool.New().Swap().SerialNumber(sModelNumber.String())

	sFirmwareRev := stream.NewBytes(info.sFirmwareRev[:])
	FirmwareRev := tool.New().Swap().SerialNumber(sFirmwareRev.String())

	mylog.Info("serialNumber", strconv.Quote(serialNumber))
	mylog.Info("ModelNumber", strconv.Quote(ModelNumber))
	mylog.Info("FirmwareRev", strconv.Quote(FirmwareRev))
	*s = ssdInfo{
		SerialNumber: serialNumber,
		ModelNumber:  ModelNumber,
		Version:      FirmwareRev,
	}
	return true
}

const (
	IDENTIFY_BUFFER_SIZE = 512
	ID_CMD               = 0xEC
	ATAPI_ID_CMD         = 0xA1
	SMART_CMD            = 0xB0

	DFP_GET_VERSION        = 0x00074080
	DFP_SEND_DRIVE_COMMAND = 0x0007c084
	DFP_RECEIVE_DRIVE_DATA = 0x0007c088
)

type (
	idSector struct {
		wGenConfig                 uint16
		wNumCyls                   uint16
		wReserved                  uint16
		wNumHeads                  uint16
		wBytesPerTrack             uint16
		wBytesPerSector            uint16
		wSectorsPerTrack           uint16
		wVendorUnique              [3]uint16
		sSerialNumber              [20]byte
		wBufferType                uint16
		wBufferSize                uint16
		wECCSize                   uint16
		sFirmwareRev               [8]byte
		sModelNumber               [40]byte
		wMoreVendorUnique          uint16
		wReserved48                uint16
		wCapabilities              uint16
		wReserved1                 uint16
		wPIOTiming                 uint16
		wDMATiming                 uint16
		wFieldValidity             uint16
		wNumCurrentCyls            uint16
		wNumCurrentHeads           uint16
		wNumCurrentSectorsPerTrack uint16
		wCurSectorsLow             uint16
		wCurSectorsHigh            uint16
		wMultSectorStuff           uint16
		ulTotalAddressableSectors  uint32
		wSingleWordDMA             uint16
		wMultiWordDMA              uint16
		wPIOCapacity               uint16
		wMinMultiWordDMACycle      uint16
		wRecMultiWordDMACycle      uint16
		wMinPIONoFlowCycle         uint16
		wMinPOIFlowCycle           uint16
		wReserved69                [11]uint16
		wMajorVersion              uint16
		wMinorVersion              uint16
		wReserved82                [6]uint16
		wUltraDMA                  uint16
		bReserved                  [167]byte
	}
	getVersionInParams struct {
		bVersion      byte
		bRevision     byte
		bReserved     byte
		bIDEDeviceMap byte
		fCapabilities uint32
		dwReserved    [4]byte
	}
	ideRegS struct {
		bFeaturesReg     byte
		bSectorCountReg  byte
		bSectorNumberReg byte
		bCylLowReg       byte
		bCylHighReg      byte
		bDriveHeadReg    byte
		bCommandReg      byte
		bReserved        byte
	}
	sendCmdInParams struct {
		cBufferSize  uint32
		regs         ideRegS
		bDriveNumber byte
		bReserved    [3]byte
		dwReserved   [4]byte
		bBuffer      [1]byte
	}
	sendCmdOutParams struct {
		cBufferSize  uint32
		DriverStatus uint32
		bBuffer      [512]byte
	}
)

const GENERIC_READ uint32 = 0x80000000
const OPEN_EXISTING uint32 = 0x3
const FILE_FLAG_OPEN_REPARSE_POINT uint32 = 0x00200000
const FILE_FLAG_BACKUP_SEMANTICS uint32 = 0x02000000

const (
	FILE_DEVICE_BEEP                = 0x00000001
	FILE_DEVICE_CD_ROM              = 0x00000002
	FILE_DEVICE_CD_ROM_FILE_SYSTEM  = 0x00000003
	FILE_DEVICE_CONTROLLER          = 0x00000004
	FILE_DEVICE_DATALINK            = 0x00000005
	FILE_DEVICE_DFS                 = 0x00000006
	FILE_DEVICE_DISK                = 0x00000007
	FILE_DEVICE_DISK_FILE_SYSTEM    = 0x00000008
	FILE_DEVICE_FILE_SYSTEM         = 0x00000009
	FILE_DEVICE_INPORT_PORT         = 0x0000000a
	FILE_DEVICE_KEYBOARD            = 0x0000000b
	FILE_DEVICE_MAILSLOT            = 0x0000000c
	FILE_DEVICE_MIDI_IN             = 0x0000000d
	FILE_DEVICE_MIDI_OUT            = 0x0000000e
	FILE_DEVICE_MOUSE               = 0x0000000f
	FILE_DEVICE_MULTI_UNC_PROVIDER  = 0x00000010
	FILE_DEVICE_NAMED_PIPE          = 0x00000011
	FILE_DEVICE_NETWORK             = 0x00000012
	FILE_DEVICE_NETWORK_BROWSER     = 0x00000013
	FILE_DEVICE_NETWORK_FILE_SYSTEM = 0x00000014
	FILE_DEVICE_NULL                = 0x00000015
	FILE_DEVICE_PARALLEL_PORT       = 0x00000016
	FILE_DEVICE_PHYSICAL_NETCARD    = 0x00000017
	FILE_DEVICE_PRINTER             = 0x00000018
	FILE_DEVICE_SCANNER             = 0x00000019
	FILE_DEVICE_SERIAL_MOUSE_PORT   = 0x0000001a
	FILE_DEVICE_SERIAL_PORT         = 0x0000001b
	FILE_DEVICE_SCREEN              = 0x0000001c
	FILE_DEVICE_SOUND               = 0x0000001d
	FILE_DEVICE_STREAMS             = 0x0000001e
	FILE_DEVICE_TAPE                = 0x0000001f
	FILE_DEVICE_TAPE_FILE_SYSTEM    = 0x00000020
	FILE_DEVICE_TRANSPORT           = 0x00000021
	FILE_DEVICE_UNKNOWN             = 0x00000022
	FILE_DEVICE_VIDEO               = 0x00000023
	FILE_DEVICE_VIRTUAL_DISK        = 0x00000024
	FILE_DEVICE_WAVE_IN             = 0x00000025
	FILE_DEVICE_WAVE_OUT            = 0x00000026
	FILE_DEVICE_8042_PORT           = 0x00000027
	FILE_DEVICE_NETWORK_REDIRECTOR  = 0x00000028
	FILE_DEVICE_BATTERY             = 0x00000029
	FILE_DEVICE_BUS_EXTENDER        = 0x0000002a
	FILE_DEVICE_MODEM               = 0x0000002b
	FILE_DEVICE_VDM                 = 0x0000002c
	FILE_DEVICE_MASS_STORAGE        = 0x0000002d
	FILE_DEVICE_SMB                 = 0x0000002e
	FILE_DEVICE_KS                  = 0x0000002f
	FILE_DEVICE_CHANGER             = 0x00000030
	FILE_DEVICE_SMARTCARD           = 0x00000031
	FILE_DEVICE_ACPI                = 0x00000032
	FILE_DEVICE_DVD                 = 0x00000033
	FILE_DEVICE_FULLSCREEN_VIDEO    = 0x00000034
	FILE_DEVICE_DFS_FILE_SYSTEM     = 0x00000035
	FILE_DEVICE_DFS_VOLUME          = 0x00000036
	FILE_DEVICE_SERENUM             = 0x00000037
	FILE_DEVICE_TERMSRV             = 0x00000038
	FILE_DEVICE_KSEC                = 0x00000039
	FILE_DEVICE_FIPS                = 0x0000003A
	FILE_DEVICE_INFINIBAND          = 0x0000003B
	IOCTL_DISK_BASE                 = FILE_DEVICE_DISK
	FILE_ANY_ACCESS                 = 0
	FILE_READ_ACCESS                = 0x0001
	FILE_WRITE_ACCESS               = 0x0002
	FILE_SPECIAL_ACCESS             = FILE_ANY_ACCESS
	METHOD_BUFFERED                 = 0
	METHOD_IN_DIRECT                = 1
	METHOD_OUT_DIRECT               = 2
	METHOD_NEITHER                  = 3
)

var (
	SMART_GET_VERSION    uint32 = CTL_CODE(IOCTL_DISK_BASE, 0x0020, METHOD_BUFFERED, FILE_READ_ACCESS)
	SMART_RCV_DRIVE_DATA uint32 = CTL_CODE(IOCTL_DISK_BASE, 0x0022, METHOD_BUFFERED, FILE_READ_ACCESS|FILE_WRITE_ACCESS)
)

var (
	kernel32, _             = syscall.LoadLibrary("kernel32.dll")
	globalMemoryStatusEx, _ = syscall.GetProcAddress(kernel32, "GlobalMemoryStatusEx")
)

type GetVersionInParam struct {
	bVersion      byte
	bRevision     byte
	bReserved     byte
	bIDEDeviceMap byte
	fCapabilities uint64
	dwReserved    [4]uint64
}

func CTL_CODE(deviceType, function, method, access uint32) uint32 {
	return ((deviceType) << 16) | ((access) << 14) | ((function) << 2) | (method)
}
