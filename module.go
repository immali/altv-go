package main

import (
	"syscall"
)

type Module struct {
	dll     *syscall.DLL
	core    uintptr
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

func NewModule() *Module {
	m := &Module{}
	m.loadDLL()

	return m
}
