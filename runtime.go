package main

import (
	"syscall"
)

// Runtime runtime in go
type Runtime struct {
	altRuntime   uintptr
	resourceImpl []*ResourceImpl
	module       *Module
}

func (r *Runtime) onTick() uintptr {
	return 0
}

func (r *Runtime) destroyImpl() uintptr {
	return 0
}

func (r *Runtime) register() bool {
	proc := r.module.dll.MustFindProc("alt_ICore_RegisterScriptRuntime")
	resPtr := StrPtr("go")
	ret, _, _ := proc.Call(r.module.altCore, resPtr, r.altRuntime)

	if i := int(ret); i == 1 {
		return true
	}

	return false
}

// NewRuntime creates a new runtime
func NewRuntime(m *Module) *Runtime {
	r := &Runtime{module: m}

	r.resourceImpl = []*ResourceImpl{}

	CreateImpl := syscall.NewCallback(NewResourceImpl(r))
	DestroyImpl := syscall.NewCallback(r.destroyImpl)
	OnTick := syscall.NewCallback(r.onTick)

	proc := m.dll.MustFindProc("alt_CAPIScriptRuntime_Create")
	ret, _, _ := proc.Call(CreateImpl, DestroyImpl, OnTick)

	r.altRuntime = ret
	r.register()

	return r
}
