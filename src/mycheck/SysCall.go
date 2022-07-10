package mycheck

import (
	"fmt"
	"syscall"
)

const errorCodeIoPending = 997

var errorIoPending error = syscall.Errno(errorCodeIoPending)

func (o *object) CheckSysCallError(r1 uintptr, lastErr syscall.Errno) (ok bool) {
	switch r1 {
	case 0:
		switch lastErr {
		case 0:
			return true
		case errorCodeIoPending:
			return o.Error(`errorCodeIoPending:` + fmt.Sprint(errorIoPending))
		default:
			return o.Error(` EINVAL:` + fmt.Sprint(syscall.EINVAL))
		}
	default:
		return true
	}
}
