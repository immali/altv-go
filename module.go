package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

import "C"

// Module in go
type Module struct {
	dll     *syscall.DLL
	altCore uintptr
	runtime *Runtime
}

func (m *Module) getSDKVersion() int {
	proc, _ := m.dll.FindProc("alt_GetSDKVersion")

	ret, _, _ := proc.Call()

	return int(ret)
}

func (m *Module) loadDLL() {
	if m.dll != nil {
		return
	}

	dll, err := syscall.LoadDLL(".\\modules\\altv-capi-server\\bin\\altv-capi-server.dll")

	if err != nil {
		panic(err)
	}

	m.dll = dll
}

func (m *Module) createRuntime() {
	m.runtime = NewRuntime(m)
}

func (m *Module) logWithLevel(level, msg string) {
	proc := m.dll.MustFindProc(fmt.Sprintf("alt_ICore_Log%s", level))
	cmsg := C.CString(msg)

	proc.Call(m.altCore, (uintptr)(unsafe.Pointer(&cmsg)))
}

func (m *Module) logInfo(msg string) {
	m.logWithLevel("Info", msg)
}

func (m *Module) logDebug(msg string) {
	m.logWithLevel("Debug", msg)
}

func (m *Module) logWarning(msg string) {
	m.logWithLevel("Warning", msg)
}

func (m *Module) logError(msg string) {
	m.logWithLevel("Error", msg)
}

// NewModule initialize module
func NewModule() *Module {
	m := &Module{}
	m.loadDLL()

	return m
}
