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
	//var ddd GETVERSIONINPARAMS
	//marshal, err := cstruct.Marshal(&ddd)
	//if !mycheck.Error(err) {
	//	return
	//}
	//mylog.HexDump("marshal", marshal)
	//mylog.Info("unsafe.Sizeof(ddd)", unsafe.Sizeof(ddd))
	//mylog.Info("len(marshal)", len(marshal))

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
	//mylog.HexDump("outBuffer", outBuffer) //struct__SENDCMDOUTPARAMS == struct__IDINFO
	mylog.MarshalJson("getVersionInParam", *getVersionInParam)

	if getVersionInParam.BIDEDeviceMap == 1 { //?
	}

	var BytesReturned uint32
	sendcmdinparams := struct__SENDCMDINPARAMS{ //mem align problem
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

	//mylog.MarshalJson("sendcmdinparams", &sendcmdinparams)

	input := *(*[]byte)(unsafe.Pointer(&sendcmdinparams))
	mylog.HexDump("input", input)
	mylog.Hex("ioControlCode", windef.SMART_RCV_DRIVE_DATA)

	iii := make([]byte, 32)
	iii[10] = ID_CMD
	mylog.HexDump("iii", iii)

	//00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 ec 00  |................|
	//0x0000  00 00 00 00 00 00 00 00 00 00 ec 00 00 00 00 00
	//reason │The parameter is incorrect.
	marshal = marshal[:32] //ID_CMD offset was error
	mylog.HexDump("marshal", marshal)
	if !mycheck.Error(syscall.DeviceIoControl(
		handle,
		windef.SMART_RCV_DRIVE_DATA,
		//(*byte)(unsafe.Pointer(&sendcmdinparams)), //shouild 32,but in 44 ,so arg error
		(*byte)(unsafe.Pointer(&marshal)), //shouild 32,but in 44 ,so arg error
		//&marshal[0],
		32,
		&outBuffer[0],
		528,
		&BytesReturned,
		nil,
	)) {
		return
	}
	outParams_ := (*struct__SENDCMDOUTPARAMS)(unsafe.Pointer(&outBuffer[0]))
	//mylog.MarshalJson("SENDCMDOUTPARAMS", *outParams_) //BDriverError and BIDEError must return 0
	b := outParams_.BBuffer[:] // index 0 address
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

//
//type (
//	idSector struct {
//		wGenConfig                 uint16
//		wNumCyls                   uint16
//		wReserved                  uint16
//		wNumHeads                  uint16
//		wBytesPerTrack             uint16
//		wBytesPerSector            uint16
//		wSectorsPerTrack           uint16
//		wVendorUnique              [3]uint16
//		sSerialNumber              [20]byte
//		wBufferType                uint16
//		wBufferSize                uint16
//		wECCSize                   uint16
//		sFirmwareRev               [8]byte
//		sModelNumber               [40]byte
//		wMoreVendorUnique          uint16
//		wReserved48                uint16
//		wCapabilities              uint16
//		wReserved1                 uint16
//		wPIOTiming                 uint16
//		wDMATiming                 uint16
//		wFieldValidity             uint16
//		wNumCurrentCyls            uint16
//		wNumCurrentHeads           uint16
//		wNumCurrentSectorsPerTrack uint16
//		wCurSectorsLow             uint16
//		wCurSectorsHigh            uint16
//		wMultSectorStuff           uint16
//		ulTotalAddressableSectors  uint32
//		wSingleWordDMA             uint16
//		wMultiWordDMA              uint16
//		wPIOCapacity               uint16
//		wMinMultiWordDMACycle      uint16
//		wRecMultiWordDMACycle      uint16
//		wMinPIONoFlowCycle         uint16
//		wMinPOIFlowCycle           uint16
//		wReserved69                [11]uint16
//		wMajorVersion              uint16
//		wMinorVersion              uint16
//		wReserved82                [6]uint16
//		wUltraDMA                  uint16
//		bReserved                  [167]byte
//	}
//	getVersionInParams struct {
//		bVersion      byte
//		bRevision     byte
//		bReserved     byte
//		bIDEDeviceMap byte
//		fCapabilities uint32
//		dwReserved    [4]byte
//	}
//	ideRegS struct {
//		bFeaturesReg     byte
//		bSectorCountReg  byte
//		bSectorNumberReg byte
//		bCylLowReg       byte
//		bCylHighReg      byte
//		bDriveHeadReg    byte
//		bCommandReg      byte
//		bReserved        byte
//	}
//	sendCmdInParams struct {
//		cBufferSize  uint32
//		regs         ideRegS
//		bDriveNumber byte
//		bReserved    [3]byte
//		dwReserved   [4]byte
//		bBuffer      [1]byte
//	}
//	sendCmdOutParams struct {
//		cBufferSize  uint32
//		DriverStatus uint32
//		bBuffer      [512]byte
//	}
//)
//
//type GetVersionInParam struct {
//	bVersion      byte
//	bRevision     byte
//	bReserved     byte
//	bIDEDeviceMap byte
//	fCapabilities uint64
//	dwReserved    [4]uint64
//}
