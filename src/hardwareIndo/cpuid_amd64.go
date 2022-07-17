//go:build amd64
// +build amd64

package hardwareIndo

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/hardwareIndo/cpuid"
	"github.com/ddkwork/librarygo/src/mybinary"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream"
	"strings"
)

func cpuid_low(arg1, arg2 uint32) (eax, ebx, ecx, edx uint32) // implemented in cpuidlow_amd64.s
func xgetbv_low(arg1 uint32) (eax, edx uint32)                // implemented in cpuidlow_amd64.s

type (
	Reg struct {
		eax, ebx, ecx, edx uint32
	}
	cpuInfo struct {
		Cpu0                 Reg
		Cpu1                 Reg
		Vendor               string
		ProcessorBrandString string
	}
)

func (c *cpuInfo) FormatCpu0() []string {
	return []string{
		fmt.Sprintf("%08X", c.Cpu0.eax),
		fmt.Sprintf("%08X", c.Cpu0.ebx),
		fmt.Sprintf("%08X", c.Cpu0.ecx),
		fmt.Sprintf("%08X", c.Cpu0.edx),
	}
}
func (c *cpuInfo) FormatCpu1() []string {
	return []string{
		fmt.Sprintf("%08X", c.Cpu1.eax),
		fmt.Sprintf("%08X", c.Cpu1.ebx),
		fmt.Sprintf("%08X", c.Cpu1.ecx),
		fmt.Sprintf("%08X", c.Cpu1.edx),
	}
}
func (c *cpuInfo) Get() (ok bool) {
	eax, ebx, ecx, edx := cpuid_low(0, 0)
	c.Cpu0 = Reg{
		eax: eax,
		ebx: ebx,
		ecx: ecx,
		edx: edx,
	}
	b := stream.New()
	b.Write(mybinary.LittleEndian.PutUint32(ebx))
	b.Write(mybinary.LittleEndian.PutUint32(edx))
	b.Write(mybinary.LittleEndian.PutUint32(ecx))
	c.Vendor = b.String()
	mylog.Info("cpu vendor", b.String())

	eax, ebx, ecx, edx = cpuid_low(1, 0)
	c.Cpu1 = Reg{
		eax: eax,
		ebx: ebx,
		ecx: ecx,
		edx: edx,
	}
	mylog.Hex("eax", eax)
	mylog.Hex("ebx", ebx)
	mylog.Hex("ecx", ecx)
	mylog.Hex("edx", edx)
	mylog.Info("ProcessorBrandString:", strings.TrimSpace(cpuid.ProcessorBrandString))
	c.ProcessorBrandString = cpuid.ProcessorBrandString
	return true
}
