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
	}
	Object struct {
		Status     uint32
		service    *mgr.Service
		manager    *mgr.Mgr
		driverPath string
		DeviceName string
	}
)

func NewObject() *Object {
	return &Object{
		Status:     0,
		service:    nil,
		manager:    nil,
		driverPath: "",
		DeviceName: "",
	}
}

func New() Interface {
	return NewObject()
}
func (o *Object) Load(sysPath string) (ok bool) {
	stat, err := os.Stat(sysPath)
	if !mycheck.Error(err) {
		return
	}
	name := stat.Name()
	o.driverPath = filepath.Join(os.Getenv("SYSTEMROOT"), "system32", "drivers", name)
	if o.DeviceName == "" {
		before, _, found := strings.Cut(name, filepath.Ext(name))
		if !found {
			return
		}
		o.DeviceName = before
	}
	mylog.Trace("deviceName", o.DeviceName)
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

func (o *Object) SetService() (ok bool) {
	var err error
	o.service, err = o.manager.OpenService(o.DeviceName)
	if err == nil {
		mylog.Trace("Service already exists")
		return true
	}
	config := mgr.Config{
		ServiceType: windows.SERVICE_KERNEL_DRIVER,
		StartType:   mgr.StartManual,
	}
	o.service, err = o.manager.CreateService(o.DeviceName, o.driverPath, config)
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
	o.Status = status.ServiceSpecificExitCode
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
