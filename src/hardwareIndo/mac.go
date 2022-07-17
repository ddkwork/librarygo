//go:build windows
// +build windows

package hardwareIndo

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/spyre-project/spyre/config"
	"golang.org/x/sys/windows"
	"net"
	"strings"
	"unsafe"
)

//func init() { scanner.RegisterSystemScanner(&systemScanner{}) }

type systemScanner struct {
	Buf                  []byte
	Row, Ip, Description string
}

//func (s *systemScanner) FriendlyName() string { return "Diag-Sysinfo" }
//func (s *systemScanner) ShortName() string    { return "sysinfo" }

func (s *systemScanner) Init(*config.ScannerConfig) error { return nil }

func getAdaptersInfo() *windows.IpAdapterInfo {
	var l uint32
	windows.GetAdaptersInfo(nil, &l)
	if l == 0 {
		return nil
	}
	buf := make([]byte, int(l))
	ai := (*windows.IpAdapterInfo)(unsafe.Pointer(&buf[0]))
	if err := windows.GetAdaptersInfo(ai, &l); err != nil {
		return nil
	}
	return ai
}

func (s *systemScanner) Get() (ok bool) {
	for ai := getAdaptersInfo(); ai != nil; ai = ai.Next {
		var row string
		for _, c := range ai.Address[:int(ai.AddressLength)] {
			if len(row) > 0 {
				row += ":"
			}
			row += fmt.Sprintf("%02x", c)
		}
		//bytes.Join()
		//strings.Join()  todo

		//var ipaddr string
		ip := ""
		for ca := &ai.IpAddressList; ca != nil; ca = ca.Next {
			ip = string(ca.IpAddress.String[:])
			//if len(ipaddr) > 0 {
			//	ipaddr += ";"
			//}
			//ipaddr += fmt.Sprintf("%s/%s",
			//	strings.Trim(string(ca.IpAddress.String[:]), " \t\n\000"),
			//	strings.Trim(string(ca.IpMask.String[:]), " \t\n\000"),
			//)
		}
		//AdapterName := string(ai.AdapterName[:])
		Description := string(ai.Description[:])
		join := strings.Join(append([]string{}, row, ip, Description), "  ")
		mylog.Info("", join)
		*s = systemScanner{
			Buf:         ai.Address[:int(ai.AddressLength)],
			Row:         row,
			Ip:          ip,
			Description: Description,
		}
		return true
		//
		//mylog.Info("row", row)
		//mylog.Info("ip", ip)
		//mylog.Info("Description", Description)
		//report.AddStringf("%s: network interface: '%s'(%s): row=%s, ipv4=%s",
		//	s.ShortName(),
		//	strings.Trim(string(ai.Description[:]), " \t\n\000"),
		//	strings.Trim(string(ai.AdapterName[:]), " \t\n\000"),
		//	row, ipaddr,
		//)
	}

	//drives, _ := sys.GetLogicalDriveStrings()
	//for _, d := range drives {
	//	if t, _ := sys.GetDriveType(d); t == sys.DRIVE_FIXED {
	//		var volNameU16 [1024]uint16
	//		var volSerial uint32
	//		if err := sys.GetVolumeInformation(
	//			d+`\`,
	//			&volNameU16[0], uint32(len(volNameU16)),
	//			&volSerial,
	//			nil, nil, nil, 0,
	//		); err != nil {
	//			log.Errorf("Could not determine volume information for %s", d)
	//			continue
	//		}
	//		volName := windows.UTF16ToString(volNameU16[:])
	//		report.AddStringf("%s: %s %08x / \"%s\"", s.ShortName(), d, volSerial, volName)
	//	}
	//}

	return
}

func GetMacAddress() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

func GetIPs() (ips []string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}
