
typedef unsigned long DWORD;
typedef int BOOL;
typedef unsigned char BYTE;
typedef unsigned short WORD;
typedef float FLOAT;
typedef FLOAT *PFLOAT;
typedef int INT;
typedef unsigned int UINT;
typedef unsigned int *PUINT;
typedef unsigned long DWORD;
typedef int BOOL;
typedef unsigned char BYTE;
typedef unsigned short WORD;
typedef float FLOAT;
typedef FLOAT *PFLOAT;
typedef BOOL *PBOOL;
typedef BOOL *LPBOOL;
typedef BYTE *PBYTE;
typedef BYTE *LPBYTE;
typedef int *PINT;
typedef int *LPINT;
typedef WORD *PWORD;
typedef WORD *LPWORD;
typedef long *LPLONG;
typedef DWORD *PDWORD;
typedef DWORD *LPDWORD;
typedef void *LPVOID;
typedef const void *LPCVOID;
typedef int INT;
typedef unsigned int UINT;
typedef unsigned int *PUINT;
typedef unsigned long ULONG;
typedef ULONG *PULONG;
typedef unsigned short USHORT;
typedef USHORT *PUSHORT;
typedef unsigned char UCHAR;
typedef UCHAR *PUCHAR;
typedef char CHAR;
typedef short SHORT;
typedef long LONG;
typedef int INT;

typedef struct _SENDCMDINPARAMS {
        DWORD cBufferSize;
        BYTE bDriveNumber;
        BYTE bReserved[3];
        DWORD dwReserved[4];
        BYTE bBuffer[1];
} SENDCMDINPARAMS, *PSENDCMDINPARAMS, *LPSENDCMDINPARAMS;

int main() { return 0; }

