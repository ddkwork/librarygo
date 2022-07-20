package table_test

import (
	"github.com/ddkwork/librarygo/src/table"
	"testing"
)

func TestName(t *testing.T) {
	table.MakeDoc([]table.GenInterfaceDoc{
		{Api: "RegisterPaApicketHandles(handles []Handle)", Function: "Register Packet Handles as work pool", Note: "", Todo: ""},
		{Api: "SetHandles(handles []Handle)", Function: "RegisterPacketHandles means", Note: "", Todo: ""},
		{Api: "Handles() []Handle ", Function: "work pool", Note: "work center", Todo: ""},
		{Api: "HandlePacket() (ok bool)", Function: "rang work pool packet and post them", Note: "", Todo: ""},
		{Api: "SetEvent(event any)", Function: "set a event for any work", Note: "", Todo: ""},
		{Api: "SetEventsCap(cap int)", Function: "Set Events Cap", Note: "", Todo: ""},
		{Api: "Events() <-chan any", Function: "pop events", Note: "work life", Todo: ""},
		{Api: "HandleEvent()", Function: "rang event and handle them", Note: "need handle by your self", Todo: ""},
		{Api: "HttpClient() httpClient.Interface", Function: "http client", Note: "", Todo: "add udp wss etc"},
		{Api: "", Function: "", Note: "", Todo: "add Worn echo and saveDataBase api"},
	})
}
