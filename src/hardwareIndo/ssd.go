//go:build windows
// +build windows

package hardwareIndo

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"github.com/ddkwork/librarygo/src/windef"
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
		windef.SMART_GET_VERSION,
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
		windef.SMART_RCV_DRIVE_DATA,
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
