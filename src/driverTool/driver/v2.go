package driver

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type (
	helper interface {
		SetStatus(status uint32)
		SetService() (ok bool)
		SetManager() (ok bool)
		StartService() (ok bool)
		StopService() (ok bool)
		DeleteService() (ok bool)
		QueryService() (ok bool)
	}
	Interface interface {
		Load(sysPath string) (ok bool)
		Unload() (ok bool)
		Status() uint32
	}
	Object struct {
		status     uint32
		service    *mgr.Service
		manager    *mgr.Mgr
		driverPath string
		deviceName string
	}
)

func New() Interface {
	return &Object{
		status:     0,
		service:    nil,
		manager:    nil,
		driverPath: "",
		deviceName: "",
	}
}
func (o *Object) Load(sysPath string) (ok bool) {
	stat, err := os.Stat(sysPath)
	if !mycheck.Error(err) {
		return
	}
	name := stat.Name()
	o.driverPath = filepath.Join(os.Getenv("SYSTEMROOT"), "system32", "drivers", name)
	before, _, found := strings.Cut(name, filepath.Ext(name))
	if !found {
		return
	}
	o.deviceName = before
	mylog.Trace("deviceName", o.deviceName)
	mylog.Trace("driverPath", o.driverPath)
	b, err := ioutil.ReadFile(sysPath)
	if !mycheck.Error(err) {
		return
	}
	f, err := os.Create(o.driverPath)
	if !mycheck.Error(err) {
		return
	}
	if !mycheck.Error2(f.Write(b)) {
		return
	}
	if !mycheck.Error(f.Close()) {
		return
	}
	if !o.SetManager() {
		return
	}
	if !o.SetService() {
		return
	}
	if !o.StartService() {
		return
	}
	return o.QueryService()
}
func (o *Object) Unload() (ok bool) {
	if !o.StopService() {
		return
	}
	if !o.DeleteService() {
		return
	}
	if !mycheck.Error(o.manager.Disconnect()) {
		return
	}
	if !mycheck.Error(o.service.Close()) {
		return
	}
	return mycheck.Error(os.Remove(o.driverPath))
}
func (o *Object) Status() uint32 { return o.status }
func (o *Object) SetStatus(status uint32) {
	o.status = status
	mylog.Trace("status", o.status)
}
func (o *Object) SetService() (ok bool) {
	var err error
	o.service, err = o.manager.OpenService(o.deviceName)
	if err == nil {
		mylog.Trace("Service already exists")
		return true
	}
	config := mgr.Config{
		ServiceType: windows.SERVICE_KERNEL_DRIVER,
		StartType:   mgr.StartManual,
	}
	o.service, err = o.manager.CreateService(o.deviceName, o.driverPath, config)
	return mycheck.Error(err)
}
func (o *Object) SetManager() (ok bool) {
	var err error
	o.manager, err = mgr.Connect()
	if !mycheck.Error(err) {
		return
	}
	return true
}
func (o *Object) QueryService() (ok bool) {
	status, err := o.service.Query()
	if !mycheck.Error(err) {
		return
	}
	o.SetStatus(status.ServiceSpecificExitCode)
	return true
}
func (o *Object) StopService() (ok bool) {
	status, err := o.service.Control(svc.Stop)
	if !mycheck.Error(err) {
		return
	}
	timeout := time.Now().Add(10 * time.Second)
	for status.State != svc.Stopped {
		if timeout.Before(time.Now()) {
			return mycheck.Error("Timed out waiting for service to stop")
		}
		time.Sleep(300 * time.Millisecond)
		if !o.QueryService() {
			return
		}
		mylog.Trace("Service stopped")
	}
	return true
}
func (o *Object) DeleteService() (ok bool) {
	if !mycheck.Error(o.service.Delete()) {
		return
	}
	mylog.Trace("Service deleted")
	return o.QueryService()
}
func (o *Object) StartService() (ok bool) { return mycheck.Error(o.service.Start()) }
