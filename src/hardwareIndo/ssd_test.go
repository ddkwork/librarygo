package hardwareIndo

import (
	"github.com/ddkwork/librarygo/src/windef"
	"testing"
)

func TestName(t *testing.T) {
	s := `
#pragma pack (1).
typedef struct _IDINFO {
  USHORT wGenConfig;
  USHORT wNumCyls;
  USHORT wReserved2;
  USHORT wNumHeads;
  USHORT wReserved4;
  USHORT wReserved5;
  USHORT wNumSectorsPerTrack;
  USHORT wVendorUnique[3];
  CHAR sSerialNumber[20];
  USHORT wBufferType;
  USHORT wBufferSize;
  USHORT wECCSize;
  CHAR sFirmwareRev[8];
  CHAR sModelNumber[40];
  USHORT wMoreVendorUnique;
  USHORT wReserved48;
  struct {
    USHORT reserved1 : 8;
    USHORT DMA : 1;
    USHORT LBA : 1;
    USHORT DisIORDY : 1;
    USHORT IORDY : 1;
    USHORT SoftReset : 1;
    USHORT Overlap : 1;
    USHORT Queue : 1;
    USHORT InlDMA : 1;
  } wCapabilities;
  USHORT wReserved1;
  USHORT wPIOTiming;
  USHORT wDMATiming;
  struct {
    USHORT CHSNumber : 1;
    USHORT CycleNumber : 1;
    USHORT UnltraDMA : 1;
    USHORT reserved : 13;
  } wFieldValidity;
  USHORT wNumCurCyls;
  USHORT wNumCurHeads;
  USHORT wNumCurSectorsPerTrack;
  USHORT wCurSectorsLow;
  USHORT wCurSectorsHigh;
  struct {
    USHORT CurNumber : 8;
    USHORT Multi : 1;
    USHORT reserved1 : 7;
  } wMultSectorStuff;
  ULONG dwTotalSectors;
  USHORT wSingleWordDMA;
  struct {
    USHORT Mode0 : 1;
    USHORT Mode1 : 1;
    USHORT Mode2 : 1;
    USHORT Reserved1 : 5;
    USHORT Mode0Sel : 1;
    USHORT Mode1Sel : 1;
    USHORT Mode2Sel : 1;
    USHORT Reserved2 : 5;
  } wMultiWordDMA;
  struct {
    USHORT AdvPOIModes : 8;
    USHORT reserved : 8;
  } wPIOCapacity;
  USHORT wMinMultiWordDMACycle;
  USHORT wRecMultiWordDMACycle;
  USHORT wMinPIONoFlowCycle;
  USHORT wMinPOIFlowCycle;
  USHORT wReserved69[11];
  struct {
    USHORT Reserved1 : 1;
    USHORT ATA1 : 1;
    USHORT ATA2 : 1;
    USHORT ATA3 : 1;
    USHORT ATA4 : 1;
    USHORT ATA5 : 1;
    USHORT ATA6 : 1;
    USHORT ATA7 : 1;
    USHORT ATA8 : 1;
    USHORT ATA9 : 1;
    USHORT ATA10 : 1;
    USHORT ATA11 : 1;
    USHORT ATA12 : 1;
    USHORT ATA13 : 1;
    USHORT ATA14 : 1;
    USHORT Reserved2 : 1;
  } wMajorVersion;
  USHORT wMinorVersion;
  USHORT wReserved82[6];
  struct {
    USHORT Mode0 : 1;
    USHORT Mode1 : 1;
    USHORT Mode2 : 1;
    USHORT Mode3 : 1;
    USHORT Mode4 : 1;
    USHORT Mode5 : 1;
    USHORT Mode6 : 1;
    USHORT Mode7 : 1;
    USHORT Mode0Sel : 1;
    USHORT Mode1Sel : 1;
    USHORT Mode2Sel : 1;
    USHORT Mode3Sel : 1;
    USHORT Mode4Sel : 1;
    USHORT Mode5Sel : 1;
    USHORT Mode6Sel : 1;
    USHORT Mode7Sel : 1;
  } wUltraDMA;
  USHORT wReserved89[167];
} IDINFO, *PIDINFO;
typedef struct _DRIVERSTATUS {
  BYTE bDriverError;
  BYTE bIDEError;
  BYTE bReserved[2];
  DWORD dwReserved[2];
} DRIVERSTATUS, *PDRIVERSTATUS, *LPDRIVERSTATUS;
typedef struct _SENDCMDOUTPARAMS {
  DWORD cBufferSize;
  DRIVERSTATUS DriverStatus;
  BYTE bBuffer[1];
} SENDCMDOUTPARAMS, *PSENDCMDOUTPARAMS, *LPSENDCMDOUTPARAMS;
typedef struct _GETVERSIONINPARAMS {
  BYTE bVersion;
  BYTE bRevision;
  BYTE bReserved;
  BYTE bIDEDeviceMap;
  DWORD fCapabilities;
  DWORD dwReserved[4];
} GETVERSIONINPARAMS, *PGETVERSIONINPARAMS, *LPGETVERSIONINPARAMS;
typedef struct _IDEREGS {
  BYTE bFeaturesReg;
  BYTE bSectorCountReg;
  BYTE bSectorNumberReg;
  BYTE bCylLowReg;
  BYTE bCylHighReg;
  BYTE bDriveHeadReg;
  BYTE bCommandReg;
  BYTE bReserved;
} IDEREGS, *PIDEREGS, *LPIDEREGS;
typedef struct _SENDCMDINPARAMS {
        DWORD cBufferSize;
        IDEREGS irDriveRegs;
        BYTE bDriveNumber;
        BYTE bReserved[3];
        DWORD dwReserved[4];
        BYTE bBuffer[1];
} SENDCMDINPARAMS, *PSENDCMDINPARAMS, *LPSENDCMDINPARAMS;
`
	windef.Creat(s)
}
