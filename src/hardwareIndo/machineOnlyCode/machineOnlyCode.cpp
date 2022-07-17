#include <cstdio>
#include <windows.h>

typedef struct _IDINFO {
	USHORT wGenConfig; // WORD 0: ������Ϣ��
	USHORT wNumCyls; // WORD 1: ������
	USHORT wReserved2; // WORD 2: ����
	USHORT wNumHeads; // WORD 3: ��ͷ��
	USHORT wReserved4; // WORD 4: ����
	USHORT wReserved5; // WORD 5: ����
	USHORT wNumSectorsPerTrack; // WORD 6: ÿ�ŵ�������
	USHORT wVendorUnique[3]; // WORD 7-9: �����趨ֵ
	CHAR sSerialNumber[20]; // WORD 10-19:���к�
	USHORT wBufferType; // WORD 20: ��������
	USHORT wBufferSize; // WORD 21: �����С
	USHORT wECCSize; // WORD 22: ECCУ���С
	CHAR sFirmwareRev[8]; // WORD 23-26: �̼��汾
	CHAR sModelNumber[40]; // WORD 27-46: �ڲ��ͺ�
	USHORT wMoreVendorUnique; // WORD 47: �����趨ֵ
	USHORT wReserved48; // WORD 48: ����
	struct {
		USHORT reserved1 : 8; USHORT DMA : 1; // 1=֧��DMA
		USHORT LBA : 1; // 1=֧��LBA
		USHORT DisIORDY : 1; // 1=�ɲ�ʹ��IORDY
		USHORT IORDY : 1; // 1=֧��IORDY
		USHORT SoftReset : 1; // 1=��ҪATA������
		USHORT Overlap : 1; // 1=֧���ص�����
		USHORT Queue : 1; // 1=֧���������
		USHORT InlDMA : 1; // 1=֧�ֽ����ȡDMA
	} wCapabilities;
	// WORD 49: һ������
	USHORT wReserved1; // WORD 50: ����
	USHORT wPIOTiming; // WORD 51: PIOʱ��
	USHORT wDMATiming; // WORD 52: DMAʱ��
	struct {
		USHORT CHSNumber : 1; // 1=WORD 54-58��Ч
		USHORT CycleNumber : 1; // 1=WORD 64-70��Ч
		USHORT UnltraDMA : 1; // 1=WORD 88��Ч
		USHORT reserved : 13;
	} wFieldValidity; // WORD 53: �����ֶ���Ч�Ա�־
	USHORT wNumCurCyls; // WORD 54: CHS��Ѱַ��������
	USHORT wNumCurHeads; // WORD 55: CHS��Ѱַ�Ĵ�ͷ��
	USHORT wNumCurSectorsPerTrack; // WORD 56: CHS��Ѱַÿ�ŵ�������
	USHORT wCurSectorsLow; // WORD 57: CHS��Ѱַ����������λ��
	USHORT wCurSectorsHigh; // WORD 58: CHS��Ѱַ����������λ��
	struct {
		USHORT CurNumber : 8; // ��ǰһ���Կɶ�д������
		USHORT Multi : 1; // 1=��ѡ���������д
		USHORT reserved1 : 7;
	} wMultSectorStuff;

	// WORD 59: ��������д�趨
	ULONG dwTotalSectors; // WORD 60-61: LBA��Ѱַ��������
	USHORT wSingleWordDMA; // WORD 62: ���ֽ�DMA֧������

	struct {
		USHORT Mode0 : 1; // 1=֧��ģʽ0 (4.17Mb/s)
		USHORT Mode1 : 1; // 1=֧��ģʽ1 (13.3Mb/s)
		USHORT Mode2 : 1; // 1=֧��ģʽ2 (16.7Mb/s)
		USHORT Reserved1 : 5; USHORT Mode0Sel : 1; // 1=��ѡ��ģʽ0
		USHORT Mode1Sel : 1; // 1=��ѡ��ģʽ1
		USHORT Mode2Sel : 1; // 1=��ѡ��ģʽ2
		USHORT Reserved2 : 5;
	} wMultiWordDMA; // WORD 63: ���ֽ�DMA֧������

	struct {
		USHORT AdvPOIModes : 8; // ֧�ָ߼�POIģʽ��
		USHORT reserved : 8;
	} wPIOCapacity; // WORD 64: �߼�PIO֧������

	USHORT wMinMultiWordDMACycle; // WORD 65: ���ֽ�DMA�������ڵ���Сֵ

	USHORT wRecMultiWordDMACycle; // WORD 66: ���ֽ�DMA�������ڵĽ���ֵ
	USHORT wMinPIONoFlowCycle; // WORD 67: ��������ʱPIO�������ڵ���Сֵ
	USHORT wMinPOIFlowCycle; // WORD 68: ��������ʱPIO�������ڵ���Сֵ
	USHORT wReserved69[11]; // WORD 69-79: ����

	struct {
		USHORT Reserved1 : 1;
		USHORT ATA1 : 1; // 1=֧��ATA-1
		USHORT ATA2 : 1; // 1=֧��ATA-2
		USHORT ATA3 : 1; // 1=֧��ATA-3
		USHORT ATA4 : 1; // 1=֧��ATA/ATAPI-4
		USHORT ATA5 : 1; // 1=֧��ATA/ATAPI-5
		USHORT ATA6 : 1; // 1=֧��ATA/ATAPI-6
		USHORT ATA7 : 1; // 1=֧��ATA/ATAPI-7
		USHORT ATA8 : 1; // 1=֧��ATA/ATAPI-8
		USHORT ATA9 : 1; // 1=֧��ATA/ATAPI-9
		USHORT ATA10 : 1; // 1=֧��ATA/ATAPI-10
		USHORT ATA11 : 1; // 1=֧��ATA/ATAPI-11
		USHORT ATA12 : 1; // 1=֧��ATA/ATAPI-12
		USHORT ATA13 : 1; // 1=֧��ATA/ATAPI-13
		USHORT ATA14 : 1; // 1=֧��ATA/ATAPI-14
		USHORT Reserved2 : 1;
	} wMajorVersion; // WORD 80: ���汾

	USHORT wMinorVersion; // WORD 81: ���汾
	USHORT wReserved82[6]; // WORD 82-87: ����

	struct {
		USHORT Mode0 : 1; // 1=֧��ģʽ0 (16.7Mb/s)
		USHORT Mode1 : 1; // 1=֧��ģʽ1 (25Mb/s)
		USHORT Mode2 : 1; // 1=֧��ģʽ2 (33Mb/s)
		USHORT Mode3 : 1; // 1=֧��ģʽ3 (44Mb/s)
		USHORT Mode4 : 1; // 1=֧��ģʽ4 (66Mb/s)
		USHORT Mode5 : 1; // 1=֧��ģʽ5 (100Mb/s)
		USHORT Mode6 : 1; // 1=֧��ģʽ6 (133Mb/s)
		USHORT Mode7 : 1; // 1=֧��ģʽ7 (166Mb/s) ???
		USHORT Mode0Sel : 1; // 1=��ѡ��ģʽ0
		USHORT Mode1Sel : 1; // 1=��ѡ��ģʽ1
		USHORT Mode2Sel : 1; // 1=��ѡ��ģʽ2
		USHORT Mode3Sel : 1; // 1=��ѡ��ģʽ3
		USHORT Mode4Sel : 1; // 1=��ѡ��ģʽ4
		USHORT Mode5Sel : 1; // 1=��ѡ��ģʽ5
		USHORT Mode6Sel : 1; // 1=��ѡ��ģʽ6
		USHORT Mode7Sel : 1; // 1=��ѡ��ģʽ7
	} wUltraDMA;

	// WORD 88: Ultra DMA֧������
	USHORT wReserved89[167]; // WORD 89-255
} IDINFO, * PIDINFO;

static void hexdump(const char* title, const void* pdata, int len) {
	printf("%s\n", title);
	int i, j, k, l;
	const char* data = (const char*)pdata;
	char buf[256], str[64], t[] = "0123456789ABCDEF";
	for (i = j = k = 0; i < len; i++) {
		if (0 == i % 16)
			j += sprintf(buf + j, "%08X  ", i);
		buf[j++] = t[0x0f & (data[i] >> 4)];
		buf[j++] = t[0x0f & data[i]];
		buf[j++] = ' ';
		str[k++] = isprint(data[i]) ? data[i] : '.';
		if (0 == (i + 1) % 16) {
			str[k] = 0;
			j += sprintf(buf + j, " |%s|\n", str);
			printf("%s", buf);
			j = k = buf[0] = str[0] = 0;
		}
	}
	str[k] = 0;
	if (k) {
		for (l = 0; l < 3 * (16 - k); l++)
			buf[j++] = ' ';
		j += sprintf(buf + j, " |%s|\n", str);
	}
	if (buf[0]) printf("%s\n", buf);
	printf("\n");
}

void  exchange_char(char* in, char* out, size_t strlen_in) {
	for (size_t i = 0; i < (strlen_in); i += 2) {
		out[i] = in[i + 1];
		out[i + 1] = in[i];
	}
}

int main()
{
	auto hDevice = CreateFileW(L"\\\\.\\PhysicalDrive0", GENERIC_READ | GENERIC_WRITE,
		FILE_SHARE_READ | FILE_SHARE_WRITE, NULL, OPEN_EXISTING, 0, NULL);
	if (hDevice == INVALID_HANDLE_VALUE)
	{
		//MessageBoxW(NULL, L"���豸ʧ����, ��!", L"����", 0);
	}
	char IdentifyResult[sizeof(SENDCMDOUTPARAMS) + IDENTIFY_BUFFER_SIZE - 1];
	GETVERSIONINPARAMS get_version;
	SENDCMDINPARAMS send_cmd = { 0 };
	DWORD cbBytesReturned = 0;
	DeviceIoControl(hDevice, SMART_GET_VERSION, NULL, 0, &get_version,
		sizeof(get_version), &cbBytesReturned, NULL);
	send_cmd.irDriveRegs.bCommandReg = (get_version.bIDEDeviceMap & 0x10) ? ATAPI_ID_CMD : ID_CMD;

	DeviceIoControl(hDevice, SMART_RCV_DRIVE_DATA, &send_cmd,
		sizeof(SENDCMDINPARAMS) - 1, IdentifyResult, sizeof(IdentifyResult), &cbBytesReturned, NULL);
	auto out = reinterpret_cast<PSENDCMDOUTPARAMS>(IdentifyResult);
	out->cBufferSize = sizeof(IDINFO);
	auto hd = reinterpret_cast<PIDINFO>(out->bBuffer);

	char data[512] = {};
	memcpy(data, out->bBuffer, 512);
	hexdump("data", data, 512);

	char disk_id[512] = {};
	char disk_model[512] = {};
	char sSerialNumber[512] = {};
	char sModelNumber[512] = {};
	exchange_char((hd->sSerialNumber), disk_id, sizeof(hd->sSerialNumber));
	exchange_char((hd->sModelNumber), disk_model, sizeof(hd->sModelNumber));

	exchange_char((hd->sSerialNumber), sSerialNumber, sizeof(hd->sSerialNumber));
	exchange_char((hd->sModelNumber), sModelNumber, sizeof(hd->sModelNumber));
	CloseHandle(hDevice);
	return 0;
}



