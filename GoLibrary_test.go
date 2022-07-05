package librarygo

//import (
//	"go.uber.org/goleak"
//	"testing"
//)
//
//func leak() {
//	f := make(chan struct{})
//	go func() {
//		f <- struct{}{}
//	}()
//}
//
//func TestName(t *testing.T) {
//	defer goleak.VerifyNone(t)
//	leak()
//}
