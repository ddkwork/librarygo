package windef

import (
	"github.com/ddkwork/librarygo/src/stream"
	"github.com/ddkwork/librarygo/src/stream/tool"
)

func Creat(structBody string) (ok bool) {
	s := stream.New()
	s.WriteStringLn(defBody)
	s.WriteStringLn(structBody)
	s.WriteStringLn(`int main() { return 0; }`)
	if !tool.File().WriteTruncate("windef.c", s.Bytes()) {
		return
	}
	return tool.File().WriteTruncate("windef.cmd", "c2go -v windef windef.c")
}

var defBody = `
 //#include <minwindef.h>
#define CONST const
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
typedef CONST void *LPCVOID;
typedef int INT;
typedef unsigned int UINT;
typedef unsigned int *PUINT;
typedef unsigned long ULONG;
typedef ULONG *PULONG;
typedef unsigned short USHORT;
typedef USHORT *PUSHORT;
typedef unsigned char UCHAR;
typedef UCHAR *PUCHAR;
#define VOID void
typedef char CHAR;
typedef short SHORT;
typedef long LONG;
typedef int INT;
`
