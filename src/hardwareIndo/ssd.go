//go:build windows
// +build windows

package hardwareIndo

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/cstruct"
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

var (
	kernel32, _             = syscall.LoadLibrary("kernel32.dll")
	globalMemoryStatusEx, _ = syscall.GetProcAddress(kernel32, "GlobalMemoryStatusEx")
)

const (
	IDENTIFY_BUFFER_SIZE = 512
	ID_CMD               = 0xEC
	ATAPI_ID_CMD         = 0xA1
	SMART_CMD            = 0xB0

	DFP_GET_VERSION        = 0x00074080
	DFP_SEND_DRIVE_COMMAND = 0x0007c084
	DFP_RECEIVE_DRIVE_DATA = 0x0007c088
)

func (s *ssdInfo) Get() (ok bool) {
	path := fmt.Sprintf("\\\\.\\PhysicalDrive%d", 0)
	fromString, err := syscall.UTF16PtrFromString(path)
	if !mycheck.Error(err) {
		return
	}
	handle, err := syscall.CreateFile(
		fromString,
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if !mycheck.Error(err) { //if (hDevice == INVALID_HANDLE_VALUE)
		return
	}
	outBuffer := make([]byte, 528)
	var bytesReturned uint32
	if !mycheck.Error(syscall.DeviceIoControl( //todo check mapid return is 1 ?
		handle,
		windef.SMART_GET_VERSION,
		nil,
		0,
		&outBuffer[0],
		IDENTIFY_BUFFER_SIZE,
		&bytesReturned,
		nil,
	)) {
		return
	}
	getVersionInParam := (*struct__GETVERSIONINPARAMS)(unsafe.Pointer(&outBuffer[0]))
	mylog.MarshalJson("getVersionInParam", *getVersionInParam)
	if getVersionInParam.BIDEDeviceMap == 1 { //?
	}
	var BytesReturned uint32
	sendcmdinparams := struct__SENDCMDINPARAMS{
		CBufferSize: 0, //32
		IrDriveRegs: struct__IDEREGS{
			BFeaturesReg:     0,
			BSectorCountReg:  0,
			BSectorNumberReg: 0,
			BCylLowReg:       0,
			BCylHighReg:      0,
			BDriveHeadReg:    0,
			BCommandReg:      ID_CMD,
			BReserved:        0,
		},
		BDriveNumber: 0,
		BReserved:    [3]uint8{},
		DwReserved:   [4]uint32{},
		BBuffer:      [1]uint8{},
	}
	marshal, err := cstruct.Marshal(&sendcmdinparams)
	if !mycheck.Error(err) {
		return
	}
	mylog.HexDump("sendcmdinparams", marshal)
	mylog.Info("unsafe.Sizeof(sendcmdinparams)", unsafe.Sizeof(sendcmdinparams))
	mylog.Info("len(sendcmdinparams)", len(marshal))
	input := *(*[]byte)(unsafe.Pointer(&sendcmdinparams))
	mylog.HexDump("input", input)
	mylog.Hex("ioControlCode", windef.SMART_RCV_DRIVE_DATA)
	mylog.HexDump("marshal", marshal)
	if !mycheck.Error(syscall.DeviceIoControl(
		handle,
		windef.SMART_RCV_DRIVE_DATA,
		&marshal[0],
		32,
		&outBuffer[0],
		528,
		&BytesReturned,
		nil,
	)) {
		return
	}
	outParams_ := (*struct__SENDCMDOUTPARAMS)(unsafe.Pointer(&outBuffer[0]))
	b := outParams_.BBuffer[:]
	mylog.HexDump("index 0 address", b)
	info := (*struct__IDINFO)(unsafe.Pointer(&b[0]))
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
